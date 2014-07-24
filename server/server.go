package main //server   先用main 调试

import (
	"Package"
	"fmt"
	"log"
	"net"
	"uniqID" //这里面MaxClient=20000   这里需要调用
)

var connString = ":1114" //端口号，需要传入端口号才能开一个服务器
//这里先
type SessionTable map[TypeSessionID]*Session
type Server struct {
	ServerID  int           //ServerID   可以调用GetServerID得到，第一期任务默认为0
	listener  net.Listener  //监听器
	sessions  SessionTable  //会话层
	pending   chan net.Conn //等待通道？？？？
	quiting   chan net.Conn //退出？？？？？
	incoming  chan Packet   //输入通道
	outgoing  chan Packet   //输出通道
	maxClient chan byte     //控制一个服务器下只能有20000个客户端
	//通过clientIn()和clientOut() 在maxClient通道取出 放入1byte实现功能
}

/*func (self *Server) Start() {
	//服务器开始先把client 通道放MaxClient个 字节内容
	for i := 0; i < MaxClient; i++ { //注意Maxclient是数量，maxclient实现数量控制的通道
		self.maxClient <- 0
	}
}*/

func (self *Server) Stop() {
	self.listener.Close()
}

func (self *Server) listen() {
	go func() {
		for {
			select {
			case conn := <-self.pending:
				self.join(conn)
			case con := <-self.quiting:
				self.leave(conn)
			}
		}
	}()
}

func (self *Server) join(conn net.Conn) {
	session := CreateSession(conn)
	id := TypeSessionID(uniqID.GetUniq()) //  这里需要改..！！
}

func (self *Server) leave(conn net.Conn) {
	if conn != nil {
		conn.Close()
	}
}

func (self *Server) clientIn() {
	<-self.maxclient
}

func (self *Server) clientOut() {
	self.maxClient <- 0
}

func CreateServer() *Server { //初始化server并开始监听  调用listen()
	server := &Server{ //初始化 通道 会话等，数值型变量不需要初始化
		sessions:  make(SessionTable, MaxClient),
		pending:   make(chan net.Conn), //传输的是 IP等方面的信息，具体是什么？？
		quiting:   make(chan net.Conn),
		incoming:  make(chan Package), // 这里调用了Package包来定义传输Package通道
		outgoing:  make(chan Package),
		maxClient: make(chan byte),
	}
	server.listen()
	////////////////////////////////////////////////////////////
	//下面内容是原来beeim   Start()内容，当CreateServer的时候  先初始化，并写日志，然后运行，可以放在一块
	////////////////////////////////////////////////////////////
	for i := 0; i < Maxclient; i++ {
		server.clientOut()
	}

	for {
		conn, err := server.listener.Accept()

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("A new connection %v is connecting \n", conn)
		server.clientIn()
		server.pending <- conn
	}

}
