package hostserver

import (
	"context"
	rpc "my-cmdb/api/host_server/gen/go"
	"net"

	"google.golang.org/grpc"
)

type hostManager struct {
	*rpc.UnimplementedHostManagementServer

	db PersistentStorage
}

// Serve will block until ctx.Done() or GrpcServer error occurs.
func (hm *hostManager) Serve(ctx context.Context, addr string) error {

	// todo: add graceful shutdown

	sAddr := ":8880"
	conn, err := net.Listen("tcp", sAddr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	rpc.RegisterHostManagementServer(s, hm)

	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Serve(conn)
	}()

	select {
	case e := <-errCh:
		err = e
	case <-ctx.Done():
		s.GracefulStop()
		err = nil
	}
	return err
}

func (hm *hostManager) FetchAllHost(ctx context.Context, req *rpc.RequestFetchAllHost) (resp *rpc.ResponseFetchAllHost, err error) {
	resp = &rpc.ResponseFetchAllHost{
		Meta: &rpc.ResponseMeta{
			TraceId: req.Meta.TraceId,
		},
	}

	hosts, err := hm.db.GetAllHostObject()
	if err != nil {
		resp.Meta.ErrCode = rpc.ResponseMeta_InternalError
		resp.Meta.ErrHint = err.Error()
	}

	resp.Result = &rpc.SearchResult{
		Hosts: hosts,
	}

	return resp, err
}

func (hm *hostManager) UpdateHost(ctx context.Context, req *rpc.RequestUpdateHost) (resp *rpc.ResponseUpdateHost, err error) {
	resp = &rpc.ResponseUpdateHost{
		Meta: &rpc.ResponseMeta{
			TraceId: req.Meta.TraceId,
		},
	}

	host := req.Host

	err = hm.db.SetHostObject(host)
	if err != nil {
		resp.Meta.ErrCode = rpc.ResponseMeta_InternalError
		resp.Meta.ErrHint = err.Error()
		return resp, err
	}

	return resp, err
}

var _ rpc.HostManagementServer = (*hostManager)(nil)

func defaultHostManager() *hostManager {
	return &hostManager{
		db: NewInMemStorage(),
	}
}
