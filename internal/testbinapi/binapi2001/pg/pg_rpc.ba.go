// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package pg

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  pg.
type RPCService interface {
	PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error)
	PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error)
	PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error) {
	out := new(PgCaptureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error) {
	out := new(PgCreateInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error) {
	out := new(PgEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
