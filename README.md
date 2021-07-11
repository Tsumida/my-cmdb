# my-cmdb

## Introduction
A cmdb written in Go.

## Arch

my-cmdb is decomposed into 4 layers:
- API gateway: offer Restful API to use scene services.
- Scene services: service that operating resources (maintained in data-server).
- Core services: monitor-server, admin-server and data-server:
    - monitor-server: monitoring all services.
    - admin-server: service registry and discovery.
    - data-server: maintaining IT resources such as hosts, IPs.
- Storage: data persistor.

All services communicates with each other using grpc internally. My-cmdb uses grpc-gateway to offer grpc\Restful conversion. 

## Code Dir
(todo)