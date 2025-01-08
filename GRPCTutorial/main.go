package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Edu58/Learn-Go/GRPCTutorial/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(_ context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	log.Printf("RECEIVED REQUEST FOR AMOUNT %d FROM %s to %s", req.Amount.Amount, req.From, req.To)
	log.Printf("PROCESSING CREATE REQUEST at %s", time.Now())

	return &invoicer.CreateResponse{
		Pdf:  []byte("test RESPONSE"),
		Docx: []byte("document RESPONSE"),
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
