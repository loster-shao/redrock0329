package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var jsons = `{"bool":true,"string":"hello world"}`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsons),&data)
	fmt.Println(data)
	fmt.Println(data["string"])

}
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//func main()  {
//	app := gin.Default()
//	app.POST("/s0",mmm)
//
//}
//
//func mmm(c *gin.Context)  {
//c.JSON()
//
//	fmt.Println("szs")
//}
