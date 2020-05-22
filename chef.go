package chef 

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

type Chef struct {
	config *Config
	*Router
}

func New() *Chef {
	c := &Chef{}

	return c
}

func (c *Chef) Run() {
	if c.config.UseFastTTP {
		fasthttp.ListenAndServe("", c.Router)
	} else {
		http.ListenAndServe("", c.Router)
	}
}