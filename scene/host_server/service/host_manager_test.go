package service

import (
	"context"
	rpc "my-cmdb/api/host_server"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestHostManagerCURD(t *testing.T) {
	//	todo: use goconvey
	//
	//	Test:
	//		Create 2 host object and then fetch all of them.
	//
	//
	// init
	hm := defaultHostManager()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	go hm.Serve(ctx, "localhost:8880")

	conn, err := grpc.Dial("localhost:8880", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.FailNow()
	}
	defer conn.Close()

	client := rpc.NewHostManagementClient(conn)
	tctx, c := context.WithTimeout(ctx, 5*time.Second)
	defer c()

	hosts := []*rpc.HostObject{
		{
			HostId:    "host-1",
			OsVersion: "Ubuntu 20.04 LTS",
			IpAddr:    "192.168.0.2",
		},
		{
			HostId:    "host-2",
			OsVersion: "Ubuntu 20.04 LTS",
			IpAddr:    "192.168.0.3",
		},
	}

	// create 2 HostObject
	for i, h := range hosts {
		if _, err := client.UpdateHost(tctx, &rpc.RequestUpdateHost{
			Meta: &rpc.RequestMeta{TraceId: int64(i)},
			Host: h,
		}); err != nil {
			t.Fatal()
		}
	}

	// Fetch all hostObject
	tid := int64(len(hosts))
	if resp, err := client.FetchAllHost(
		tctx,
		&rpc.RequestFetchAllHost{
			Meta: &rpc.RequestMeta{TraceId: tid},
		}); err != nil {

		t.Fatal(err.Error())

	} else {
		if resp.Meta.TraceId != tid {
			t.FailNow()
		}

		if len(resp.Result.Hosts) != 2 {
			t.FailNow()
		}

	}

}
