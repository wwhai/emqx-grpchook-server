package main

import (
	"context"
	"emqx-grpchook-server/grpchook"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

type rpc_hook_server struct {
	grpchook.UnimplementedRpcHookServer
}

func (s *rpc_hook_server) ClientConnect(ctx context.Context, request *grpchook.ConnectRequest) (response *grpchook.ConnectResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.ConnectResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientConnack(ctx context.Context, request *grpchook.ConnackRequest) (response *grpchook.ConnackResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.ConnackResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientConnected(ctx context.Context, request *grpchook.ConnectedRequest) (response *grpchook.ConnectedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.ConnectedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientDisconnected(ctx context.Context, request *grpchook.DisconnectedRequest) (response *grpchook.DisconnectedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.DisconnectedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientAuthenticate(ctx context.Context, request *grpchook.AuthenticateRequest) (response *grpchook.AuthenticateResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.AuthenticateResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientCheckAcl(ctx context.Context, request *grpchook.CheckAclRequest) (response *grpchook.CheckAclResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.CheckAclResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientSubscribe(ctx context.Context, request *grpchook.SubscribeRequest) (response *grpchook.SubscribeResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.SubscribeResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) ClientUnsubscribe(ctx context.Context, request *grpchook.UnsubscribeRequest) (response *grpchook.UnsubscribeResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.UnsubscribeResponse)
	response.Result = true
	return response, nil
}

// Session
func (s *rpc_hook_server) SessionCreated(ctx context.Context, request *grpchook.SessionCreatedRequest) (response *grpchook.SessionCreatedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.SessionCreatedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) SessionSubscribed(ctx context.Context, request *grpchook.SessionSubscribedRequest) (response *grpchook.SessionSubscribedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.SessionSubscribedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) SessionUnsubscribed(ctx context.Context, request *grpchook.SessionUnsubscribedRequest) (response *grpchook.SessionUnsubscribedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.SessionUnsubscribedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) SessionTerminated(ctx context.Context, request *grpchook.SessionTerminatedRequest) (response *grpchook.SessionTerminatedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.SessionTerminatedResponse)
	response.Result = true
	return response, nil
}

// Message
func (s *rpc_hook_server) MessagePublish(ctx context.Context, request *grpchook.MessagePublishRequest) (response *grpchook.MessagePublishResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.MessagePublishResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) MessageDelivered(ctx context.Context, request *grpchook.MessageDeliveredRequest) (response *grpchook.MessageDeliveredResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.MessageDeliveredResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) MessageAcked(ctx context.Context, request *grpchook.MessageAckedRequest) (response *grpchook.MessageAckedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.MessageAckedResponse)
	response.Result = true
	return response, nil
}
func (s *rpc_hook_server) MessageDropped(ctx context.Context, request *grpchook.MessageDroppedRequest) (response *grpchook.MessageDroppedResponse, err error) {

	log.Debug("========>", request)
	response = new(grpchook.MessageDroppedResponse)
	response.Result = true
	return response, nil
}
func startServer() {
	listener, err := net.Listen("tcp", ":1994")
	if err != nil {
		log.Fatal(err)
		return
	}
	rpcServer := grpc.NewServer()
	grpchook.RegisterRpcHookServer(rpcServer, new(rpc_hook_server))
	go func(c context.Context) {
		log.Info("rpcCodecServer started on", listener.Addr())
		rpcServer.Serve(listener)
	}(context.TODO())

}

//go:generate ./gen_proto.sh

func main() {

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	startServer()
	<-channel
	log.Debug("Exit.")
}
