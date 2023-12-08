package service

import (
	"context"
	"kni-server-client/pkg/apis/runtime/beta"

	"github.com/containerd/go-cni"
)

type KniService struct {
	c cni.CNI
}

func NewKniService() (beta.KNIServer, error) {
	opts := []cni.Opt{cni.WithLoNetwork, cni.WithDefaultConf}

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
	_, err := k.c.SetupSerially(ctx, req.Id, req.Isolation.Path)

	if err != nil {
		return nil, err
	}

	return &beta.AttachNetworkResponse{}, nil
}

func (k *KniService) DetachNetwork(ctx context.Context, req *beta.DetachNetworkRequest) (*beta.DetachNetworkResponse, error) {
	err := k.c.Remove(ctx, req.Id, req.Isolation.Path)

	if err != nil {
		return nil, err
	}

	return &beta.DetachNetworkResponse{}, nil
}

func (k *KniService) QueryNetworks(context.Context, *beta.QueryNetworksRequest) (*beta.QueryNetworksResponse, error) {
	var nets []string

	for _, v := range k.c.GetConfig().Networks {
		nets = append(nets, v.Config.Name)
	}

	return &beta.QueryNetworksResponse{
		Names: nets,
	}, nil
}

func (k *KniService) NetworkStatus(context.Context, *beta.NetworkStatusRequest) (*beta.NetworkStatusResponse, error) {
	data := make(map[string]string)
	data["node"] = "running"
	data["pod"] = "running"

	return &beta.NetworkStatusResponse{
		Status: data,
	}, nil
}

func (k *KniService) SetupNodeNetwork(context.Context, *beta.SetupNodeNetworkRequest) (*beta.SetupNodeNetworkResponse, error) {
	//Setup the initial node network

	return nil, nil
}

func (k *KniService) QueryPod(ctx context.Context,req *beta.QueryPodRequest) (*beta.QueryPodResponse, error) {
	
	
	
	
	return nil, nil
}