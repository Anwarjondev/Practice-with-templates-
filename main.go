package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/hello", Hello)
	r.GET("/hellohtml", HellHtml)
	r.GET("/namehtml/:name", NameHtml)
	r.GET("/manydata", Manydata)
	r.GET("/formget", formGet)
	r.POST("/formget", resultHtml)
	err := r.Run(":8080")
	gin.SetMode(gin.ReleaseMode)
	if err != nil {
		log.Fatal(err)

	}
}
func Hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func HellHtml(c *gin.Context) {
	c.HTML(200, "hello.html", nil)
}

func NameHtml(c *gin.Context) {
	name := c.Param("name")
	c.HTML(200, "name.html", name)
}

func Manydata(c *gin.Context) {
	team := []string{"Yamal", "Pedri", "Gavi", "Cubarsi", "Nico", "Aroujo", "De Jong", "Balde"}
	c.HTML(200, "many_data.html", gin.H{
		"name": "Barselona",
		"team": team,
	})
}

func formGet(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}
func resultHtml(c *gin.Context) {
	name := c.PostForm("name")
	food := c.PostForm("food")
	c.HTML(200, "result.html", gin.H{
		"name": name,
		"food": food,
	})
}
