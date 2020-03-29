package szs

import (
	"fmt"
	"net/http"
	"strings"
)

// context.go something

//一个。。。好像还有不少人叫轮子
type Context struct {
	req        *http.Request//http请求
	w          http.ResponseWriter//http响应文本
	queryParam map[string]string//map
	formParam  map[string]string//map
}

type JSONS struct {
	Data interface{}
}

type Renders interface {
	Renders(http.ResponseWriter) error
	WriteContentType(w http.ResponseWriter)
}

func (c *Context) String(s string) {
	_, _ = c.w.Write([]byte(s))//应该是返回前端字符串。。好神奇
}

func NewContext(rw http.ResponseWriter, r *http.Request) (ctx Context) {
	ctx = Context{
		req:       r,
		w:         rw,
		formParam: make(map[string]string),
	}
	ctx.queryParam = parseQuery(r.RequestURI)
	return
}

func (c *Context) Query(key string) string {
	v := c.queryParam[key]
	return v
}

func parseQuery(uri string) (res map[string]string) {
	res = make(map[string]string)
	uris := strings.Split(uri, "?")
	if len(uris) == 1 {
		return
	}
	param := uris[len(uris)-1]
	pair := strings.Split(param, "&")
	for _, kv := range pair {
		kvPair := strings.Split(kv, "=")
		if len(kvPair) != 2 {
			fmt.Println(kvPair)
			panic("request error")
		}
		res[kvPair[0]] = kvPair[1]
	}
	return
}

//不太会写发送json
//JSON
func (c *Context) JSON(code int, obj interface{})  {
	c.Render(code, JSONS{Data: obj})
}

func (c *Context) Render(code int, r Renders){
	c.Status(code)
	if !bodyAllowedForStatus(code){
		r.WriteContentType(c.w)
		c.w.WriteHeaderNow()
		return
	}

	if err := Renders(c.w); err != nil{
		panic(err)
	}
}

func (c *Context) Status(code int)  {
	w.Write.WriteHeader(code)
}
