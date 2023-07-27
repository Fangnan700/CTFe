package middleware

import (
	"CTFe/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
)

func SecurityMiddleWare(ctx *gin.Context) {
	var (
		err       error
		data      map[string]interface{}
		bodyBytes []byte
		jsonByte  []byte
	)

	bodyBytes, err = io.ReadAll(ctx.Request.Body)
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Println(err)
	}
	// 将请求体放回请求，否则后续解析时出现EOF异常
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// 检查 JSON 数据是否包含特殊字符
	re := regexp.MustCompile(`[~!#$%^&*()+={}\[\]:;<>,?/\\|]`)
	for _, value := range data {
		jsonByte, err = json.Marshal(value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.NewResponse(-1, "服务器异常"))
			ctx.Abort()
			return
		}
		if re.MatchString(string(jsonByte)) {
			ctx.JSON(http.StatusOK, models.NewResponse(-1, "数据包含特殊字符"))
			ctx.Abort()
			return
		}
	}

	ctx.Next()
	return
}
