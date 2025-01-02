package adpter

import (
	"go.uber.org/zap"
	"learning_tools/all_packaged_library/base/tool"
	"learning_tools/ddd-auth2-example/adpter/http"
	"learning_tools/ddd-auth2-example/domain/service"
	"learning_tools/ddd-auth2-example/infrastructure/conf"
	"learning_tools/ddd-auth2-example/infrastructure/pkg/log"
)

type Server struct {
	conf *conf.AppConfig
	log  *log.Logger
	auth service.AuthSrv
}

func NewSrv(c *conf.AppConfig, log *log.Logger, auth service.AuthSrv,
) *Server {
	s := &Server{conf: c, log: log, auth: auth}
	s.Init()
	return s
}
func (s *Server) Init() {
	// hcode.Click()
}

func (s *Server) RunApp() {
	http.NewHttp(s.conf, s.auth)
	tool.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("auth server exit", zap.Any("addr", s.conf.NetConf.ServerAddr))
	})
}

func (s *Server) Close() {

}
