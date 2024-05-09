package helper

import (
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"time"
)

type Header struct {
	ProcessTime float64                `json:"process_time"`
	Status      bool                   `json:"status"`
	StatusCode  int                    `json:"status_code"`
	Reason      string                 `json:"reason"`
	Messages    map[string]interface{} `json:"messages"`
}

type Response struct {
	Header   Header                  `json:"header"`
	Data     interface{}             `json:"data,omitempty"`
	Paginate *model.ResponsePaginate `json:"paginate,omitempty"`
}

func NewResponse(status bool, code int, reason string, message map[string]interface{}, data interface{}, paginate *model.ResponsePaginate, startTime time.Time) Response {
	return Response{
		Header: Header{
			ProcessTime: float64(time.Since(startTime).Seconds()),
			Status:      status,
			StatusCode:  code,
			Reason:      reason,
			Messages:    message,
		},
		Data:     data,
		Paginate: paginate,
	}
}

func ResponseAPI(c *gin.Context, status bool, code int, reason string, message map[string]interface{}, startTime time.Time) {
	response := NewResponse(status, code, reason, message, nil, nil, startTime)

	c.JSON(code, response)
	c.Abort()
}

func ResponseDataAPI(c *gin.Context, status bool, code int, reason string, message map[string]interface{}, data interface{}, startTime time.Time) {
	response := NewResponse(status, code, reason, message, data, nil, startTime)

	c.JSON(code, response)
	c.Abort()
}

func ResponseDataPaginationAPI(c *gin.Context, status bool, code int, reason string, message map[string]interface{}, data interface{}, paginate model.ResponsePaginate, startTime time.Time) {
	response := NewResponse(status, code, reason, message, data, &paginate, startTime)

	c.JSON(code, response)
	c.Abort()
}
