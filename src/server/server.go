package server

import (
	"fmt"
	"sync"

	"scratch/src/config"
	"scratch/src/database"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Config config.Config
	Conn   database.Database
}

func NewServer() *Server {
	var err error
	s := new(Server)
	s.Config = config.NewConfig()                   //初始化配置项
	s.Conn, err = database.NewDatabase(s.Config.PC) //初始化数据库
	if err != nil {
		panic(err.Error())
	}
	logrus.Println(s.Config)
	return s
}

func (s *Server) Run() {
	// router := s.makeRouter() // 实例化gin
	// Subscribe to a topic

	wg := sync.WaitGroup{}
	wg.Add(1)
	router := s.makeRouter() // 实例化gin

	go router.Run(fmt.Sprintf(":%d", s.Config.Port)) //监听端口,开启RESTful API
	wg.Wait()
}

func (s *Server) Close() {

}
