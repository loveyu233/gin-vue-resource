package response

import "github.com/gin-gonic/gin"

type R struct {
	State
	Data interface{} `json:"data,omitempty"`
}

func NewR(data interface{}) R {
	return R{Success, data}
}

func ResOk(c *gin.Context, r R) {
	c.JSON(200, r)
}

func ResWithMsg(c *gin.Context, code int, msg interface{}) {
	c.JSON(200, R{
		State: State{code, msg},
	})
}
