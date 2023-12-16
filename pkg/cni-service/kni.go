package cniservice

import (
	"context"
	"fmt"

	"github.com/MikeZappa87/kni-server-client-example/pkg/apis/runtime/beta"

	"github.com/containerd/go-cni"
)

type KniService struct {
	c cni.CNI
}

func NewKniService() (beta.KNIServer, error) {
	opts := []cni.Opt{
		cni.WithInterfacePrefix("kni"),
		 cni.WithDefaultConf,
		 cni.WithLoNetwork} 
	
	cni, err := cni.New()

	if err != nil {
		return nil, err
	}

	kni := &KniService{
		c: cni,
	}

	err = kni.c.Load(opts...)

	if err != nil {
		return nil, err
	}

	return kni, nil
}

func (k *KniService) AttachNetwork(ctx context.Context, req *beta.AttachNetworkRequest) (*beta.AttachNetworkResponse, error) {
	//This is nice. In the container runtime if you want to add one you need to contribute this to the container runtime
	//This way you are in complete control. Better capability support with KNI -> CNI
	
	opts := []cni.NamespaceOpts{
		cni.WithArgs("IgnoreUnknown", "1"),
		cni.WithLabels(req.Labels),
		cni.WithLabels(req.Annotations),
		cni.WithLabels(req.Metadata),
	}

	if val, ok := req.Annotations["netns"]; ok {
		fmt.Printf("annotations netns: %s\n", val)
	}
	
	fmt.Printf("id: %s netns: %s\n", req.Id, req.Isolation.Path)

	res, err := k.c.SetupSerially(ctx, req.Id, req.Isolation.Path, opts...)

	fmt.Println("ATTACH RECEIVED")

	if err != nil {
		fmt.Println(fmt.Errorf("issue attaching with cni: %s",err.Error()))

		return nil, err
	}

	ip := make(map[string]*beta.IPConfig)

	for outk, outv := range res.Interfaces {
		data := &beta.IPConfig{}
		ip[outk] = data
		data.Mac = outv.Mac

		for _, v := range outv.IPConfigs {
			data.Ip = append(data.Ip, v.IP.String())
			fmt.Printf("interface: %s ip: %s\n", outk, v.IP.String())
		}
	}
	
	if err != nil {
		fmt.Println(fmt.Errorf("issue attaching with json marshalling: %s",err.Error()))
		return nil, err
	}

	return &beta.AttachNetworkResponse{
		Ipconfigs: ip,
	}, nil
}

func (k *KniService) DetachNetwork(ctx context.Context, req *beta.DetachNetworkRequest) (*beta.DetachNetworkResponse, error) {
	
	opts := []cni.NamespaceOpts{
		cni.WithArgs("IgnoreUnknown", "1"),
		cni.WithLabels(req.Labels),
		cni.WithLabels(req.Annotations),
		cni.WithLabels(req.Metadata),
	}

	if req.Isolation == nil {
		return &beta.DetachNetworkResponse{}, nil
	}

	fmt.Printf("id: %s netns: %s\n", req.Id, req.Isolation.Path)
	
	err := k.c.Remove(ctx, req.Id, req.Isolation.Path, opts...)

	if err != nil {
		fmt.Println(fmt.Errorf("issue detaching with cni: %s",err.Error()))

		return nil, err
	}

	fmt.Println("DETACH RECEIVED")

	return &beta.DetachNetworkResponse{}, nil
}

func (k *KniService) SetupNodeNetwork(context.Context, *beta.SetupNodeNetworkRequest) (*beta.SetupNodeNetworkResponse, error) {
	//Setup the initial node network

	return nil, nil
}

func (k *KniService) QueryPodNetwork(ctx context.Context,req *beta.QueryPodNetworkRequest) (*beta.QueryPodNetworkResponse, error) {
	
	var data map[string]*beta.IPConfig

	return &beta.QueryPodNetworkResponse{
		Ipconfigs: data,
	}, nil
}

func (k *KniService) QueryNodeNetworks(ctx context.Context, req *beta.QueryNodeNetworksRequest) (*beta.QueryNodeNetworksResponse, error) {
	return nil, nil
}