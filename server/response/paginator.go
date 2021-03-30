package response

import (
	"math"

	"github.com/jinzhu/gorm"
)

// Param 分页参数
type PageOption struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

type pageMeta struct {
	TotalRecord int `json:"total"`
	TotalPage   int `json:"total_pages"`
	Offset      int `json:"offset"`
	Limit       int `json:"per_page"`
	Page        int `json:"current_page"`
}

// Paginator 分页返回
type Paginator struct {
	Pagination pageMeta    `json:"pagination"`
	Records    interface{} `json:"data"`
}

// Paging 分页
func paging(p *PageOption, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.Records = result
	paginator.Pagination.TotalRecord = count
	paginator.Pagination.Page = p.Page
	paginator.Pagination.Offset = offset
	paginator.Pagination.Limit = p.Limit
	paginator.Pagination.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
