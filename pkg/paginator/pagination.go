package paginator

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
	"math"
	"mqenergy-go/config"
	"mqenergy-go/pkg/util"
)

type PageBuilder struct {
	DB        *gorm.DB
	Query     interface{} // 查询条件
	Args      interface{} // 查询条件映射的参数
	Fields    []string    // 查询的字段
	FieldType string      // 查询类型 _select：代表查询前面字段 _omit：代表过滤前面字段
	Joins     string      // 连表查询拼接的字符串 例如："LEFT JOIN {关联表} ON 主表.ID = 关联表.主表ID"
	Model     interface{} // model struct
}

type OnJoins struct {
	LeftTableField, RightTableField TableField // LeftTableField：如：主表.ID  RightTableField：如：关联表.主表ID
}

type TableField struct {
	Table, Field string
}

type Page struct {
	List        interface{} `json:"list"`         // 查询的列表
	CurrentPage int         `json:"current_page"` // 当前页
	Count       int64       `json:"count"`        // 查询记录总数
	LastPage    int         `json:"last_page"`    // 最后一页
	PerPage     int         `json:"per_page"`     // 每页条数
}

func NewBuilder() *PageBuilder {
	return &PageBuilder{}
}

// WithDB db连接
func (pb *PageBuilder) WithDB(db *gorm.DB) *PageBuilder {
	pb.DB = db
	return pb
}

// WithFields 查询字段（或过滤某些字段不查询 最后一个参数默认为select（不传或者传），如传omit为过滤前面传输的字段）
func (pb *PageBuilder) WithFields(fields []string) *PageBuilder {
	// 返回列表
	if len(fields) == 1 {
		if fields[0] == "_omit" || fields[0] == "_select" {
			fields = []string{"*"}
		}
	} else {
		switch fields[len(fields)-1] {
		case "_omit", "_select":
			pb.FieldType = fields[len(fields)-1]
			fields = fields[:len(fields)-1]
		default:
			pb.FieldType = "_select"
			fields = fields[:]
		}
	}
	pb.Fields = fields
	return pb
}

// WithModel 查询的model struct
func (pb *PageBuilder) WithModel(model interface{}) *PageBuilder {
	pb.Model = model
	return pb
}

// WithJoins join查询
func (pb *PageBuilder) WithJoins(joinType string, joinFields OnJoins) *PageBuilder {
	joins := joinType + " JOIN " + joinFields.RightTableField.Table + " ON "
	joins += joinFields.LeftTableField.Table + "." + joinFields.LeftTableField.Field + "="
	joins += joinFields.RightTableField.Table + "." + joinFields.RightTableField.Field
	pb.Joins = joins
	return pb
}

// WithCondition 查询条件
func (pb *PageBuilder) WithCondition(query interface{}, args interface{}) *PageBuilder {
	pb.Query = query
	pb.Args = args
	return pb
}

// ParsePage 分页超限设置和格式化
func (pb *PageBuilder) ParsePage(currentPage, pageSize int) Page {
	var page Page
	// 返回每页数量
	page.PerPage = pageSize
	// 返回当前页码
	page.CurrentPage = currentPage

	if currentPage < 1 {
		page.CurrentPage = 1
	}
	if pageSize < 1 {
		page.PerPage = config.Conf.Server.DefaultPageSize
	}
	if pageSize > config.Conf.Server.MaxPageSize {
		page.PerPage = config.Conf.Server.MaxPageSize
	}
	if page.LastPage < 1 {
		page.LastPage = 1
	}
	return page
}

// Pagination 分页查询
func (pb *PageBuilder) Pagination(dst interface{}, currentPage, pageSize int) (Page, error) {
	page := pb.ParsePage(currentPage, pageSize)
	offset := (page.CurrentPage - 1) * page.PerPage
	query := pb.DB

	// 查询的model
	if pb.Model != nil {
		query = query.Model(&pb.Model)
	}
	// 查询字段
	if pb.FieldType == "_select" {
		query = query.Select(pb.Fields)
	}
	if pb.FieldType == "_omit" {
		fields, _ := util.GetStructColumnName(pb.Model, 1)
		difference, _ := lo.Difference[string](fields, pb.Fields)
		query = query.Select(difference)
	}
	if pb.Query != nil && pb.Args != nil {
		query = query.Where(pb.Query, pb.Args)
	}
	// join查询
	if pb.Joins != "" {
		query = query.Joins(pb.Joins)
	}
	if err := query.Count(&page.Count).Error; err != nil {
		return page, err
	}
	// 计算总页数
	if page.Count > int64(page.PerPage) {
		page.LastPage = int(math.Ceil(float64(page.Count) / float64(page.PerPage)))
	}
	// 判断总数跟最后一页的关系
	if page.CurrentPage <= page.LastPage {
		if err := query.Limit(page.PerPage).Offset(offset).Find(&dst).Error; err != nil {
			return page, err
		}
	}
	page.List = dst
	return page, nil
}
