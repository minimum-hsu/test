package main

import (
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"

	"github.com/Cepave/common/model"
	"github.com/toolkits/net"
)

type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

var (
	req       model.NqmPingTaskRequest
	rpcClient SingleConnRpcClient
)

func (this *SingleConnRpcClient) close() {
	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

func (this *SingleConnRpcClient) insureConn() {
	if this.rpcClient != nil {
		return
	}

	var err error
	var retry int = 1

	for {
		if this.rpcClient != nil {
			return
		}

		this.rpcClient, err = net.JsonRpcClient("tcp", this.RpcServer, this.Timeout)
		if err == nil {
			return
		}

		log.Println("[ hbs ] Error on query:", err)

		if retry > 6 {
			retry = 1
		}

		time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)

		retry++
	}
}

func (this *SingleConnRpcClient) Call(method string, args interface{}, reply interface{}) error {

	this.Lock()
	defer this.Unlock()

	this.insureConn()

	timeout := time.Duration(50 * time.Second)
	done := make(chan error)

	go func() {
		err := this.rpcClient.Call(method, args, reply)
		done <- err
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
		this.close()
	case err := <-done:
		if err != nil {
			this.close()
			return err
		}
	}

	return nil
}

func InitRPC() {
	rpcClient.RpcServer = GetGeneralConfig().Hbs.RPCServer
	req = model.NqmPingTaskRequest{
		Hostname:     GetGeneralConfig().Hostname,
		IpAddress:    GetGeneralConfig().IPAddress,
		ConnectionId: GetGeneralConfig().ConnectionID,
	}
}
