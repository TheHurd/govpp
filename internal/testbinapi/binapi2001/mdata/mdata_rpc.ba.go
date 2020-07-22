// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package mdata

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  mdata.
type RPCService interface {
	MdataEnableDisable(ctx context.Context, in *MdataEnableDisable) (*MdataEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) MdataEnableDisable(ctx context.Context, in *MdataEnableDisable) (*MdataEnableDisableReply, error) {
	out := new(MdataEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
