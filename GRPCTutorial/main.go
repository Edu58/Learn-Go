package main

import (
	"context"
	"log"
	"net"

	"github.com/Edu58/Learn-Go/GRPCTutorial/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("document"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalln(err)
	}

	serviceRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serviceRegistrar, service)

	err = serviceRegistrar.Serve(listener)

	if err != nil {
		log.Fatalln(err)
	}
}
