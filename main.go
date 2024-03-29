package webook

import "github.com/gin-gonic/gin"

// 程序启动入口
func main() {
	r := gin.Default()

	r.Run(":8080")
}
