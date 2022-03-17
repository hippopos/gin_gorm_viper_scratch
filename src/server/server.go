package server

import (
	"fmt"
	"log"
	"os"
	"sync"

	"scratch/src/config"
	"scratch/src/database"
)

type Server struct {
	Config config.Config
	Logger *log.Logger
	Conn   database.Database
}

func NewServer() *Server {
	var err error
	s := new(Server)
	s.Logger = log.New(os.Stdout, "[bear] ", log.Ldate|log.Lshortfile|log.Ltime)
	s.Config = config.NewConfig() //初始化配置项
	s.Logger.Println(s.Config)
	s.Conn, err = database.NewDatabase(s.Config.PC) //初始化数据库
	if err != nil {
		panic(err.Error())
	}

	s.Logger.Println("bear server is started")
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
