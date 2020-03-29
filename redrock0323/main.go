package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock0323/szs"

)

func main() {
	fmt.Println("start now:")
    app := szs.Default()
    app.GET("/book", QueryBook)
    app.POST("/cake", QueryCake)
    app.Run(8080)
}

func QueryBook(c *szs.Context) {
	c.PostForm()
	bid := c.Query("id")
	c.JSON(500,{})
	c.String("your book id is " + bid)
}


func QueryCake(c *szs.Context) {
	//id := c.PostForm()
	c.String("Cake 666 ")
	//c.JSON()
}

//JSON解析
func Json(c *szs.Context)  {
	var jsons = `{"bool":true,"string":"hello world"}`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsons),&data)
	fmt.Println(data)
	fmt.Println(data["string"])
}


