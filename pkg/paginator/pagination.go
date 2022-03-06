package paginator

import (
	"gorm.io/gorm"
	"math"
)

type PageBuilder struct {
	DB        *gorm.DB
	Query     interface{}
	Args      interface{}
	Fields    []string
	FieldType string
}

type Page struct {
	List        interface{} `json:"list"`
	CurrentPage int         `json:"current_page"`
	Count       int64       `json:"count"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
}

var Builder = PageBuilder{}

// WithDB db连接
func (pb *PageBuilder) WithDB(db *gorm.DB) *PageBuilder {
	pb.DB = db
	return pb
}

// WithFields 查询字段（或过滤某些字段不查询 最后一个参数默认为select（不传或者传），如传omit为过滤前面传输的字段）
func (pb *PageBuilder) WithFields(fields []string) *PageBuilder {
	// 返回列表
	if len(fields) == 1 {
		if fields[0] == "omit" || fields[0] == "select" {
			fields = []string{"*"}
		}
	} else {
		if fields[len(fields)-1] == "omit" {
			pb.FieldType = "omit"
		} else {
			pb.FieldType = "select"
		}
		fields = fields[:len(fields)-1]
	}
	pb.Fields = fields
	return pb
}

// WithCondition 查询条件
func (pb *PageBuilder) WithCondition(query interface{}, args ...interface{}) *PageBuilder {
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
		page.PerPage = 20
	}
	if pageSize > 100 {
		page.PerPage = 100
	}
	if page.LastPage < 1 {
		page.LastPage = 1
	}
	return page
}

// Pagination 分页
func (pb *PageBuilder) Pagination(list interface{}, currentPage, pageSize int) (Page, error) {
	page := pb.ParsePage(currentPage, pageSize)
	offset := (page.CurrentPage - 1) * page.PerPage
	queryBuilder := pb.DB
	if len(pb.Fields) > 0 {
		queryBuilder = pb.DB.Select(pb.Fields)
	}
	if pb.FieldType == "omit" {
		queryBuilder = pb.DB.Omit(pb.Fields...)
	}
	if pb.Query != nil && pb.Args != nil {
		queryBuilder = queryBuilder.Where(pb.Query, pb.Args)
	}
	// 查询总数 total queryBuilder变量随着记录的查询会改变到相应的数据表struct中，所以无需model即可查询到count总数 此处二维struct传给model也可以获取到总数
	if err := queryBuilder.Model(&list).Count(&page.Count).Error; err != nil {
		return page, err
	}
	// 计算总页数
	if page.Count > int64(page.PerPage) {
		page.LastPage = int(math.Ceil(float64(page.Count) / float64(page.PerPage)))
	}
	// 判断总数跟最后一页的关系
	if page.CurrentPage <= page.LastPage {
		if err := queryBuilder.Limit(page.PerPage).Offset(offset).Find(&list).Error; err != nil {
			return page, err
		}
	}
	page.List = list
	return page, nil
}
