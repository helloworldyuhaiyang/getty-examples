/******************************************************
# DESC    : echo server
# AUTHOR  : Alex Stocks
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2016-09-04 15:49
# FILE    : server.go
******************************************************/

package main

import (
	// "flag"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	// "strings"
	"syscall"
	"time"
)

import (
	"github.com/AlexStocks/getty"
	gxlog "github.com/AlexStocks/goext/log"
	gxnet "github.com/AlexStocks/goext/net"
	log "github.com/AlexStocks/log4go"
)

const (
	pprofPath = "/debug/pprof/"
)

var (
// host  = flag.String("host", "127.0.0.1", "local host address that server app will use")
// ports = flag.String("ports", "12345,12346,12347", "local host port list that the server app will bind")
)

var (
	serverList []getty.Server
	taskPool   *getty.TaskPool
)

func main() {
	// flag.Parse()
	// if *host == "" || *ports == "" {
	// 	panic(fmt.Sprintf("Please intput local host ip or port lists"))
	// }

	initConf()

	initProfiling()

	initServer()
	gxlog.CInfo("%s starts successfull! its version=%s, its listen ends=%s:%s\n",
		conf.AppName, Version, conf.Host, conf.Ports)
	log.Info("%s starts successfull! its version=%s, its listen ends=%s:%s\n",
		conf.AppName, Version, conf.Host, conf.Ports)

	initSignal()
}

func initProfiling() {
	var (
		addr string
	)

	// addr = *host + ":" + "10000"
	addr = gxnet.HostAddress(conf.Host, conf.ProfilePort)
	log.Info("App Profiling startup on address{%v}", addr+pprofPath)
	go func() {
		log.Info(http.ListenAndServe(addr, nil))
	}()
}

func newSession(session getty.Session) error {
	var (
		ok      bool
		tcpConn *net.TCPConn
	)

	if conf.GettySessionParam.CompressEncoding {
		session.SetCompressType(getty.CompressZip)
	}

	if tcpConn, ok = session.Conn().(*net.TCPConn); !ok {
		panic(fmt.Sprintf("%s, session.conn{%#v} is not tcp connection\n", session.Stat(), session.Conn()))
	}

	tcpConn.SetNoDelay(conf.GettySessionParam.TcpNoDelay)
	tcpConn.SetKeepAlive(conf.GettySessionParam.TcpKeepAlive)
	if conf.GettySessionParam.TcpKeepAlive {
		tcpConn.SetKeepAlivePeriod(conf.GettySessionParam.keepAlivePeriod)
	}
	tcpConn.SetReadBuffer(conf.GettySessionParam.TcpRBufSize)
	tcpConn.SetWriteBuffer(conf.GettySessionParam.TcpWBufSize)

	session.SetName(conf.GettySessionParam.SessionName)
	session.SetMaxMsgLen(conf.GettySessionParam.MaxMsgLen)
	session.SetPkgHandler(echoPkgHandler)
	session.SetEventListener(echoMsgHandler)
	session.SetRQLen(conf.GettySessionParam.PkgRQSize)
	session.SetWQLen(conf.GettySessionParam.PkgWQSize)
	session.SetReadTimeout(conf.GettySessionParam.tcpReadTimeout)
	session.SetWriteTimeout(conf.GettySessionParam.tcpWriteTimeout)
	session.SetCronPeriod((int)(conf.sessionTimeout.Nanoseconds() / 1e6))
	session.SetWaitTime(conf.GettySessionParam.waitTimeout)
	session.SetTaskPool(taskPool)
	log.Debug("app accepts new session:%s\n", session.Stat())

	return nil
}

func initServer() {
	var (
		addr     string
		portList []string
		server   getty.Server
	)

	if conf.TaskPoolSize != 0 {
		taskPool = getty.NewTaskPool(
			getty.WithTaskPoolTaskPoolSize(conf.TaskPoolSize),
			getty.WithTaskPoolTaskQueueLength(conf.TaskQueueLength),
			getty.WithTaskPoolTaskQueueNumber(conf.TaskQueueNumber),
		)
	}

	// if *host == "" {
	// 	panic("host can not be nil")
	// }
	// if *ports == "" {
	// 	panic("ports can not be nil")
	// }

	// portList = strings.Split(*ports, ",")

	portList = conf.Ports
	if len(portList) == 0 {
		panic("portList is nil")
	}
	for _, port := range portList {
		addr = gxnet.HostAddress2(conf.Host, port)
		server = getty.NewTCPServer(
			getty.WithLocalAddress(addr),
		)
		// run server
		server.RunEventLoop(newSession)
		log.Debug("server bind addr{%s} ok!", addr)
		serverList = append(serverList, server)
	}
}

func uninitServer() {
	for _, server := range serverList {
		server.Close()
	}
	taskPool.Close()
}

func initSignal() {
	// signal.Notify的ch信道是阻塞的(signal.Notify不会阻塞发送信号), 需要设置缓冲
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		log.Info("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
		// reload()
		default:
			go time.AfterFunc(conf.failFastTimeout, func() {
				// log.Warn("app exit now by force...")
				// os.Exit(1)
				log.Exit("app exit now by force...")
				log.Close()
			})

			// 要么fastFailTimeout时间内执行完毕下面的逻辑然后程序退出，要么执行上面的超时函数程序强行退出
			uninitServer()
			// fmt.Println("app exit now...")
			log.Exit("app exit now...")
			log.Close()
			return
		}
	}
}
