package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线
func (u *User) Online() {
	// 用户上线，将用户加入onlinemap中
	u.server.mapLock.Lock()
	u.server.OnLineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 广播当前上线信息
	u.server.Broadcast(u, "已上线")
}

// 用户下线
func (u *User) Offline() {
	// 用户下线，将用户移除onlinemap
	u.server.mapLock.Lock()
	delete(u.server.OnLineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播当前下线信息
	u.server.Broadcast(u, "已下线")
}

// 给向当前用户对应客户端发消息
func (u *User) SendMessage(msg string) {
	u.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (u *User) DoMessage(msg string) {
	if msg == "@who" {
		// 查询当前在线用户都有哪些
		u.server.mapLock.Lock()
		for _, user := range u.server.OnLineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			u.SendMessage(onlineMsg)
		}
		u.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := u.server.OnLineMap[newName]

		if ok {
			u.SendMessage("当前用户名被使用\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnLineMap, u.Name)
			u.server.OnLineMap[newName] = u
			u.server.mapLock.Unlock()

			u.SendMessage("您已经更新用户名:" + u.Name + "\n")
		}
	} else {
		u.server.Broadcast(u, msg)
	}
}

// 监听当前User channel的方法，一旦有消息，就直接发送给对方客户端
func (s *User) ListenMessage() {
	for {
		msg := <-s.C
		s.conn.Write([]byte(msg + "\n"))
	}
}
