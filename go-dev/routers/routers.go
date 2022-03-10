package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Process struct {
	APIServer *http.Server
}

func (p *Process) NewProcess() *Process {
	return &Process{
		APIServer: &http.Server{},
	}
}

func (p *Process) InitRouter() *gin.Engine {
	r := gin.New()
	return r
}

func Start() {
	var p *Process
	p.APIServer.Handler = p.InitRouter()
	p.APIServer.ListenAndServe()

}
