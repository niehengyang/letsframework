package request

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (bool, interface{}) {

	err := c.Bind(form)
	if err != nil {
		return false, nil
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return false, nil
	}
	if !check {

		markErrors := make([]string, 0)
		for _, err := range valid.Errors {
			markErrors = append(markErrors, err.Message)
		}
		return false, markErrors
	}

	return true, nil
}
