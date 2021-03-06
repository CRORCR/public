package mainserver

import (
	"card_public/server/modes"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type RPCServer struct {
	strPort string
}

func (this *RPCServer) Start() {

	merchant := new(modes.Merchant)
	rpc.Register(merchant)

	staff := new(modes.Staff)
	rpc.Register(staff)

	transaction := new(modes.TransactionFoot)
	rpc.Register(transaction)

	rate := new(modes.YoawoRate)
	rpc.Register(rate)

	revenue := new(modes.YoawoRevenue)
	rpc.Register(revenue)

	with := new(modes.MWithdrawalFoot)
	rpc.Register(with)

	withdraw := new(modes.WithdrawalAccount)
	rpc.Register(withdraw)

	temp := new(modes.TemplateBanner)
	rpc.Register(temp)

	banner := new(modes.Banner)
	rpc.Register(banner)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7003")
	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	modes.ReviewBanner()
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		//需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
