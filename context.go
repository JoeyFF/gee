package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	上下文封装http.ResponseWriter和http.Request
*/

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 参数
func (c *Context) PostForm(key string) string {
	return c.Req.PostFormValue(key)
}

// Query 参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) (err error) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	//_, err = c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	_, err = fmt.Fprintf(c.Writer, format, values...)
	return
}

func (c *Context) JSON(code int, obj interface{}) (err error) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	err = encoder.Encode(obj)
	return
}

func (c *Context) Data(code int, data []byte) (err error) {
	c.Status(code)
	_, err = c.Writer.Write(data)
	return
}

func (c *Context) HTML(code int, html string) (err error) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err = c.Writer.Write([]byte(html))
	return
}
