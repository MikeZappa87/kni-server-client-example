package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/MikeZappa87/kni-server-client-example/pkg/apis/runtime/beta"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// How can this be used for the virtualized runtime use cases?
func main() {
	var cmd, protocol, sockAddr string

	flag.StringVar(&cmd, "cmd", "", "operation")
	flag.StringVar(&protocol, "protocol", "unix", "protocol")
	flag.StringVar(&sockAddr, "address", "/tmp/kni.sock", "socket address")

	flag.Parse()

	var (
		credentials = insecure.NewCredentials()
		dialer      = func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, protocol, addr)
		}
		options = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials),
			grpc.WithBlock(),
			grpc.WithContextDialer(dialer),
		}
	)

	conn, err := grpc.Dial(sockAddr, options...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := beta.NewKNIClient(conn)

	iso := &beta.Isolation{
		Path: "/var/run/netns/kni-12345",
		Type: "namespace",
		//Metadata: meta,
	}

	fmt.Printf("cmd is: '%s'\n", cmd)

	if cmd == "detach" {
		delreq := &beta.DetachNetworkRequest{
			Name:      "test",
			Id:        "123456",
			Isolation: iso,
			Namespace: "default",
			//Metadata: meta,
			//Annotations: anno,
			//Labels: labels,
		}

		fmt.Println("execute detach")

		_, err := client.DetachNetwork(context.TODO(), delreq)

		if err != nil {
			fmt.Print(err)
		}
	} else if cmd == "attach" {
		req := &beta.AttachNetworkRequest{
			Name:      "test",
			Id:        "123456",
			Isolation: iso,
			Namespace: "default",
			//Metadata: meta,
			//Annotations: anno,
			//Labels: labels,
		}

		fmt.Println("Execute attach")

		res, err := client.AttachNetwork(context.TODO(), req)

		for k, v := range res.Ipconfigs {
			fmt.Printf("name: %s ips: %s\n", k, strings.Join(v.Ip, ","))
		}

		if err != nil {
			fmt.Print(err)
		}
	} else if cmd == "querypod" {
		res, err := client.QueryPodNetwork(context.TODO(), &beta.QueryPodNetworkRequest{})

		if err != nil {
			fmt.Print(err)
		}

		for k, v := range res.Ipconfigs {
			fmt.Printf("int: %s ip: %s", k, v.Ip)
		}
	}
	/* else if cmd == "query" {
		res, err := client.QueryNetworks(context.TODO(), &beta.QueryNetworksRequest{})

		if err != nil {
			fmt.Print(err)
		}

		for _, v := range res.Names {
			fmt.Println(v)
		}
	} else if cmd == "status" {
		res, err := client.NetworkStatus(context.TODO(), &beta.NetworkStatusRequest{})

		if err != nil {
			fmt.Print(err)
		}

		for k, v := range res.Status {
			fmt.Printf("key: %s value: %s\n", k, v)
		}
	}
	*/
}
