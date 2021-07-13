package hostserver

import (
	rpc "my-cmdb/api/host_server/gen/go"
	"sync"
)

type HostObject = rpc.HostObject
type HostID = string
type SearchResult = rpc.SearchResult

// Concurrent-safe persistent storage interface
type PersistentStorage interface {
	SetHostObject(*HostObject) error
	GetHostObject(HostID) *HostObject
	DelHostObject(HostID) *HostObject
	GetAllHostObject() ([]*HostObject, error)
}

type inMemStroage struct {
	mapper sync.Map
}

func (im *inMemStroage) SetHostObject(host *HostObject) error {
	key := host.HostId
	im.mapper.Store(key, host)
	return nil
}

func (im *inMemStroage) GetHostObject(hostID HostID) (host *HostObject) {
	v, _ := im.mapper.Load(hostID)
	host, _ = v.(*HostObject)
	return host
}

func (im *inMemStroage) DelHostObject(hostID HostID) (host *HostObject) {
	v, _ := im.mapper.LoadAndDelete(hostID)
	host, _ = v.(*HostObject)
	return host
}

func (im *inMemStroage) GetAllHostObject() (res []*HostObject, err error) {
	res = make([]*HostObject, 0, 4)
	im.mapper.Range(func(_, value interface{}) bool {
		host, ok := value.(*HostObject)
		if ok {
			res = append(res, host)
		}
		return true
	})
	return res, err
}

func NewInMemStorage() *inMemStroage {
	return &inMemStroage{}
}

var _ PersistentStorage = (*inMemStroage)(nil)
