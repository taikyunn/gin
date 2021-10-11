package main

import (
	"github.com/gin-gonic/gin"
)

type B struct {
	Id int
}

type C struct {
	Id   int
	Name string
}

func main() {
	// 構造体を使用してiという変数を作成する。
	i := []C{
		{1, "こんにちわ"},
		{2, "こんばんわ"},
	}
	x := B{
		1,
	}
	// エンジンインスタンスの生成
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/index.html")
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/test", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/test.html")
		c.HTML(200, "test.html", gin.H{
			// 左がHTMLでの変数・右はgoの中での変数名
			"a": "abcd",
			"b": []string{"b_todo1", "b_todo2"},
			"c": []C{{1, "mika"}, {2, "risa"}},
			"d": C{3, "d_mayu"},
			"e": true,
			"f": false,
			"h": true,
			"i": i,
			"x": x,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/result.html")

		Id := c.PostForm("id")
		Name := c.PostForm("name")
		c.HTML(200, "result.html", gin.H{
			"Id": Id, "Name": Name})
	})
	// エンジンの実行
	r.Run()
}
