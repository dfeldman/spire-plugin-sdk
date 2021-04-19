// Code generated by protoc-gen-go-spire. DO NOT EDIT.

package notifierv1

import (
	pluginsdk "github.com/spiffe/spire-plugin-sdk/pluginsdk"
	grpc "google.golang.org/grpc"
)

func NotifierPluginServer(server NotifierServer) pluginsdk.PluginServer {
	return notifierPluginServer{NotifierServer: server}
}

type notifierPluginServer struct {
	NotifierServer
}

func (s notifierPluginServer) Type() string {
	return "Notifier"
}

func (s notifierPluginServer) GRPCServiceName() string {
	return "spire.plugin.server.notifier.v1.Notifier"
}

func (s notifierPluginServer) RegisterServer(server *grpc.Server) interface{} {
	RegisterNotifierServer(server, s.NotifierServer)
	return s.NotifierServer
}

type NotifierPluginClient struct {
	NotifierClient
}

func (s NotifierPluginClient) Type() string {
	return "Notifier"
}

func (c *NotifierPluginClient) IsInitialized() bool {
	return c.NotifierClient != nil
}

func (c *NotifierPluginClient) GRPCServiceName() string {
	return "spire.plugin.server.notifier.v1.Notifier"
}

func (c *NotifierPluginClient) InitClient(conn grpc.ClientConnInterface) interface{} {
	c.NotifierClient = NewNotifierClient(conn)
	return c.NotifierClient
}
