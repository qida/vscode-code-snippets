package logs

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/axgle/mahonia"
	"github.com/qida/tcp_server"
)

var (
	LogConn = logs.NewLogger(1000)
	LogMail = logs.NewLogger(1000)
	Enc     = mahonia.NewEncoder("gb18030")
)

func initLogConn() {
	LogConn.SetLogger(logs.AdapterConn, fmt.Sprintf(`{"net":"tcp","addr":"127.0.0.1:%d","reconnect":true}`, DEBUG_PORT))
	LogConn.SetLevel(logs.LevelNotice)
}

func InitEmail() {
	LogMail.Async()
	LogMail.EnableFuncCallDepth(true)
	err := LogMail.SetLogger(logs.AdapterMail, `{"level":7,"username":"sunqida@126.com","password":"","fromAddress":"sunqida@126.com","subject":"", "host":"smtp.126.com:994","sendTos":["sunqida@foxmail.com"]}`) //654/994
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ZXJY_API初始化邮件报警")
	if beego.BConfig.RunMode == "dev" {
		LogMail.Notice("Api Test系统开始运行：%v", time.Now())
	} else {
		LogMail.Notice("Api Prod系统开始运行：%v", time.Now())
	}
}

var (
	DebugList map[string]*tcp_server.Client
)
var DEBUG_PORT = 8888

func ServerTcpReplay() {
	go serverTcpReplay()
	initLogConn()
}
func serverTcpReplay() {
	fmt.Printf("调试 在 %d 监听...\r\n", DEBUG_PORT)
	DebugList = make(map[string]*tcp_server.Client)
	server := tcp_server.New(fmt.Sprintf("0.0.0.0:%d", DEBUG_PORT))
	// utf-8=>gb18030
	//dec := mahonia.NewDecoder("GB18030")
	// gb18030=>utf-8
	//enc := mahonia.Newutil.Encoder("GB18030")
	server.OnNewClient(func(c *tcp_server.Client) {
		c.Send(fmt.Sprintf("Welcome %s \n", c.GetConn().RemoteAddr().String()))
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {

		// 中文处理 //
		for _, v := range DebugList {
			v.Send(message)
		}
		if message == "debug\r\n" {
			fmt.Printf("A Debugger:%+v\r\n", c)
			DebugList[c.GetConn().RemoteAddr().String()] = c
			c.Send("Welcome Debugger\r\n")
		}
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		fmt.Printf("调试端断开\r\n")
		delete(DebugList, c.GetConn().RemoteAddr().String())
	})
	server.Listen()
}
