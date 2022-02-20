package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播channel
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播channel的goroutine
func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message

		// 将msg发送给全部在线User
		s.mapLock.Lock()
		for _, cli := range s.OnLineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息
func (s *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	//fmt.Println("链接建立成功")

	user := NewUser(conn)

	// 用户上线，将用户加入onlinemap中
	s.mapLock.Lock()
	s.OnLineMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播当前上线信息
	s.Broadcast(user, "已上线")

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				s.Broadcast(user, "下线")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn Read error:", err)
				return
			}

			// 提取用户的消息(去除"\n")
			msg := string(buf[:n-1])

			// 将得到的消息进行广播
			s.Broadcast(user, msg)
		}
	}()
	// 当前hanlder阻塞
	select {}
}

// 启动服务器
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go s.ListenMessager()

	// accpet
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpet err:", err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}
}
