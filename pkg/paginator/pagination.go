package paginator

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
	"math"
	"mqenergy-go/config"
	"mqenergy-go/pkg/util"
)

type PageBuilder struct {
	DB       *gorm.DB
	Model    interface{} // model struct
	Preloads []string    // 预加载
	Fields   []string    // 查询字段
}

type OnJoins struct {
	LeftTableField, RightTableField JoinTableField // LeftTableField：如：主表.ID  RightTableField：如：关联表.主表ID
}

type JoinTableField struct {
	Table, Field string
}

type SelectTableField struct {
	Model interface{}
	Table string
	Field []string
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

// NewDB 对接原生查询方式
func (pb *PageBuilder) NewDB() *gorm.DB {
	return pb.DB
}

// WithField 查询单表的字段 和 过滤字段 不能与WithFields方法同用
func (pb *PageBuilder) WithField(fields []string) *PageBuilder {
	fieldList := filterFields(pb.Model, fields)
	pb.Fields = fieldList
	pb.DB.Select(pb.Fields)
	return pb
}

// WithFields 单多表字段查询字段（或过滤某些字段不查询 最后一个参数默认为select（不传或者传），如传omit为过滤前面传输的字段）
func (pb *PageBuilder) WithFields(model interface{}, table string, fields []string) *PageBuilder {
	fieldList := filterFields(model, fields)
	for i, _field := range fieldList {
		fieldList[i] = table + "." + _field
	}
	pb.Fields = append(pb.Fields, fieldList...)
	pb.DB.Select(pb.Fields)
	return pb
}

// filterFields 过滤查询字段
func filterFields(model interface{}, fields []string) []string {
	var fieldList []string
	if len(fields) == 1 {
		if fields[0] != "_omit" && fields[0] != "_select" {
			fieldList = fields
		}
	} else {
		switch fields[len(fields)-1] {
		case "_omit":
			fields = fields[:len(fields)-1]
			_fields, _ := util.GetStructColumnName(model, 1)
			fieldList, _ = lo.Difference[string](_fields, fields)
		case "_select":
			fieldList = fields[:len(fields)-1]
		default:
			fieldList = fields[:]
		}
	}
	return fieldList
}

// WithMultiFields 多表多字段查询
func (pb *PageBuilder) WithMultiFields(fields []SelectTableField) *PageBuilder {
	for _, field := range fields {
		pb.WithFields(field.Model, field.Table, field.Field)
	}
	return pb
}

// WithModel 查询的model struct
func (pb *PageBuilder) WithModel(model interface{}) *PageBuilder {
	pb.Model = model
	pb.DB = pb.DB.Model(&model)
	return pb
}

// WithJoins join查询
func (pb *PageBuilder) WithJoins(joinType string, joinFields []OnJoins) *PageBuilder {
	var joins string
	for _, field := range joinFields {
		joins += " " + joinType + " JOIN " + field.RightTableField.Table
		joins += " ON " + field.LeftTableField.Table + "." + field.LeftTableField.Field + " = "
		joins += field.RightTableField.Table + "." + field.RightTableField.Field
	}
	pb.DB.Joins(joins)
	return pb
}

// WithPreloads 多表关联查询主动预加载（无条件）
func (pb *PageBuilder) WithPreloads(querys []string) *PageBuilder {
	pb.Preloads = querys
	return pb
}

// WithPreload 关联查询主动预加载（可传条件）
func (pb *PageBuilder) WithPreload(query string, args ...interface{}) *PageBuilder {
	pb.DB.Preload(query, args)
	return pb
}

// WithCondition 查询条件
func (pb *PageBuilder) WithCondition(query interface{}, args ...interface{}) *PageBuilder {
	pb.DB.Where(query, args...)
	return pb
}

// Pagination 分页查询
func (pb *PageBuilder) Pagination(dst interface{}, currentPage, pageSize int) (Page, error) {
	query := pb.DB
	page := pb.ParsePage(currentPage, pageSize)
	offset := (page.CurrentPage - 1) * page.PerPage

	// 查询总数
	if err := query.Count(&page.Count).Error; err != nil {
		return page, err
	}
	// 预加载
	if len(pb.Preloads) > 0 {
		for _, preload := range pb.Preloads {
			query.Preload(preload)
		}
	}
	// 计算总页数
	if page.Count > int64(page.PerPage) {
		page.LastPage = int(math.Ceil(float64(page.Count) / float64(page.PerPage)))
	}
	// 判断总数跟最后一页的关系
	if page.CurrentPage <= page.LastPage {
		if err := query.Limit(page.PerPage).Offset(offset).Find(dst).Error; err != nil {
			return page, err
		}
	}
	page.List = dst
	return page, nil
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
