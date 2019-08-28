package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	入门基本示例
 */
func main() {
	r := gin.Default()
	data := map[string]interface{}{
		"title":"gin 学习",
		"days":15,
	}
	//1.入门示例
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"pong",
		})
	})

	//2.AsciiJSON：使用AsciiJSON生成只有ascii的JSON，并使用转义的非ascii计时器。这个会输出{"days":15,"title":"gin \u5b66\u4e60"}
	r.GET("/asciiiJson", func(context *gin.Context) {
		context.AsciiJSON(http.StatusOK,data)
	})
	//3.JSON：这个输出会被转换为：{"html":"\u003cb\u003eHello Gin\u003c/b\u003e"}
	r.GET("/normaljson", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"html":"<b>Hello Gin</b>",
		})
	})
	//4.PureJSON：这个会原样输出：{"html":"<b>Hello Gin</b>"}
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK,gin.H{
			"html":"<b>Hello Gin</b>",
		})
	})

	//5.SecureJSON：将输出while(1);["go","java","php"]
	r.GET("/secureJson", func(c *gin.Context) {
		names := []string{"go","java","php"}
		c.SecureJSON(http.StatusOK,names)
	})

	//6.IndentedJSON：这种会带格式的输出，以下会输出以下内容：
	/**
		{
			"days": 15,
			"title": "gin 学习"
		}
	 */
	r.GET("/indentedJson", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK,data)
	})
	//要转换为xml格式，需要特定的对象
	r.GET("/xmldata", func(c *gin.Context) {
		c.XML(http.StatusOK,data)
	})

	r.GET("/yamldata", func(c *gin.Context) {
		c.YAML(http.StatusOK,data)
	})
	//这个需要特殊处理。
	r.GET("/protobuf", func(c *gin.Context) {
		c.ProtoBuf(http.StatusOK,data)
	})

	r.GET("/stringdata", func(c *gin.Context) {
		c.String(http.StatusOK,"%v:%d","age",15)
	})

	r.Run()
}
