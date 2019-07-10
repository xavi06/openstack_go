# openstack_go

openstack golang module

## Usage
```bash
go get github.com/xavi06/openstack_go
```

```go
import (
    "github.com/xavi06/openstack_go"
)

func main() {
    stackendpoint := "172.16.151.10:5000"
    stackuser := "demo"
    stackpass := "stack2017"
    stacktenant := "a75a30f4016a462f99d6a6ad6d9443ce"
    stackdomain := "default"
    stackregion := "RegionOne"
    client, _ := openstackgo.Conn(
        stackendpoint,
        stackuser,
        stackpass,
        stacktenant,
        stackdomain,
        stackregion,    
    )
    servers, err := openstackgo.GetServers(client)
    if err != nil {

    }
    ...

}
```
