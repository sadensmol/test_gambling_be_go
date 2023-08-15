package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/sadensmol/test_gambling_be_go/api/proto/v1"
	"github.com/sadensmol/test_gambling_be_go/internal/adapters/controllers"
	"github.com/sadensmol/test_gambling_be_go/internal/services"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	go setupGrpc()
	go setupGrpcGW()
	setupWebsocket()

}

func setupGrpc() {
	println("fasdfasd")
	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	walletService := services.NewWalletService()
	walletHandler := controllers.NewWalletHandler(walletService)

	v1.RegisterWalletServiceServer(s, walletHandler)

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9090")
	log.Fatalln(s.Serve(lis))
}

func setupGrpcGW() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := v1.RegisterWalletServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	httpServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(httpServer.ListenAndServe())
}

func setupWebsocket() {
	wsHandler := controllers.NewWSHanlder()
	http.Handle("/ws", websocket.Handler(wsHandler.Handle))
	log.Println("Serving WS on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
