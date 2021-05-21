package response

import "github.com/gin-gonic/gin"

type httpResponse struct {
	c *gin.Context
}

type ResponseMap map[string]interface{}

type resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(c *gin.Context) *httpResponse {
	return &httpResponse{c: c}
}

func (hr *httpResponse) Raw(httpCode int, code int, msg string, data interface{}) {
	hr.c.JSON(httpCode, resp{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func (hr *httpResponse) Data(data interface{}) {
	hr.c.JSON(200, data)
}

func (hr *httpResponse) Pagination(option *PageOption, result interface{}, callback func(interface{}) interface{}) {
	paginator := paging(option, result)
	paginator.Records = callback(paginator.Records)
	hr.c.JSON(200, paginator)
}

func (hr *httpResponse) AfterPaginator(result interface{}) interface{} {
	return result
}

func (hr *httpResponse) Success(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "success",
			"code":    200,
		}
	}
	hr.c.JSON(200, message)
}

func (hr *httpResponse) NoContent(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "No Content",
			"code":    204,
		}
	}
	hr.c.JSON(204, message)
}

func (hr *httpResponse) Created(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "Created",
			"code":    201,
		}
	}

	hr.c.JSON(201, message)
}

func (hr *httpResponse) Error401(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "认证失败",
			"code":    401,
		}
	}
	hr.c.JSON(401, message)
}

func (hr *httpResponse) Error409(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "不能重复创建资源",
			"code":    409,
		}
	}
	hr.c.JSON(409, message)
}

func (hr *httpResponse) Error404(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "找不到所需的资源",
			"code":    404,
		}
	}
	hr.c.JSON(404, message)
}

func (hr *httpResponse) Error403(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "无权限操作资源",
			"code":    403,
		}
	}
	hr.c.JSON(403, message)
}

func (hr *httpResponse) Error500(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "服务器错误",
			"code":    500,
		}
	}
	hr.c.JSON(500, message)
}

func (hr *httpResponse) Error400(message interface{}) {
	if message == nil {
		message = map[string]interface{}{
			"message": "请求参数错误",
			"code":    400,
		}
	}
	hr.c.JSON(400, message)
}
