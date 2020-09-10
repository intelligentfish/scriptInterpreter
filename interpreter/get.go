package interpreter

import (
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

type Get struct {
}

func (object *Get) GetCmd() string {
	return "get"
}

func (object *Get) Run(args []string) (err error) {
	if 1 > len(args) {
		err = ErrArgs
		return
	}
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	req.SetRequestURI(args[0])
	req.Header.SetMethod(http.MethodGet)
	if err = fasthttp.Do(req, res); nil != err {
		return
	}
	res.Header.WriteTo(os.Stdout)
	res.BodyWriteTo(os.Stdout)
	return
}
