package chef

import (
	"net/http"
	"net/url"
)

type Context struct {
	request *http.Request 
	writer http.ResponseWriter
	index int8
	path      string
	pnames    []string
	pvalues   []string
	query     url.Values
	params    map[string]string
	handlers []Handler
}

type Handler func(*Context)

func NewContext(req *http.Request, res http.ResponseWriter, maxParam *int) *Context {
	return &Context{}
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}


func (c *Context) SetHandlers(h []Handler) {
	c.handlers = h
}

func (c *Context) GetHandlers() []Handler {
	return c.handlers
}


func (c *Context) SetStatusCode(code int) {
	c.writer.WriteHeader(code)
}

func (c *Context) SetHeader(header, value string) {
	c.writer.Header().Set(header, value)
}

func (c *Context) Host() string {
	return c.request.Host
}

func (c *Context) Write(body []byte) {
	c.writer.Write(body)
}

func (c *Context) WriteString(body string) {
	c.Write([]byte(body))
}

func (c *Context) reset(request *http.Request, writer http.ResponseWriter, cfg *Config) {

}