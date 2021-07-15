# my-cmdb

## Introduction
A cmdb written in Go.

## Arch
my-cmdb is decomposed into 4 layers:
- API gateway: offer Restful API to use scene services.
- Scene services: service that operating resources (maintained in data-server).
    - host-server: manager physical machines.
- Core services: monitor-server, admin-server
    - monitor-server: monitoring all services.
    - admin-server: service registry and discovery.
- Storage: data persistor.

others:
- Services use grpc to communicate with each other. 
- Use grpc-gateway to offer grpc/Restful-API conversion. 
- Use etcd for service registry and discovery. 

## Code Dir
- `api/`: Save proto files of services and scripts. 
- `gateway/`: grpc-gateway.
- `scene/`: Scene servers.
- `utils/`: something useful. 

