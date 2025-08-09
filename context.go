package goink

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type KV map[string]interface{}

// 兼容Gin的H
type H = KV

type Context struct {
	// 原始请求和响应
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求路径与方法
	Path   string
	Method string
	// 响应码
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:   req.URL.Path,
		Method: req.Method,
		Writer: w,
		Req:    req,
	}
}

// 获取请求内容的基本函数

// 获取key对应的表单值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 获取key对应的查询参数值
// Query是在URL中的参数
// 例如?name=Goink,则key为name,value为Goink
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 构造响应的基本函数
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 响应类型
func (c *Context) String(code int, text string) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(text))
}

func (c *Context) Stringf(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

// Raw格式，具体内容由用户自己定义
func (c *Context) Raw(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
