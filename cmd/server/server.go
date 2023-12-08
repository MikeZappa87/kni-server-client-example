package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	cniservice "github.com/MikeZappa87/kni-server-client-example/pkg/cni-service"

	"github.com/MikeZappa87/kni-server-client-example/pkg/apis/runtime/beta"

	"google.golang.org/grpc"
)

func main() {
	var cmd, protocol, sockAddr string

	flag.StringVar(&cmd, "cmd", "cni", "backend")
	flag.StringVar(&protocol, "protocol", "unix", "protocol")
	flag.StringVar(&sockAddr, "address", "/tmp/kni.sock", "socket address")

	flag.Parse()

	if _, err := os.Stat(sockAddr); !os.IsNotExist(err) {
		if err := os.RemoveAll(sockAddr); err != nil {
			log.Fatal(err)
		}
	}

	listener, err := net.Listen(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	server := grpc.NewServer()

	kni, err := GetBackend(cmd)

	if err != nil {
		log.Fatal(err)
		return
	}

	beta.RegisterKNIServer(server, kni)

	fmt.Println("Running")

	server.Serve(listener)
}

func GetBackend(cmd string) (beta.KNIServer, error) {
	if cmd == "cni" {
		return cniservice.NewKniService()
	}
	return nil, errors.New("implementation not found")
}
