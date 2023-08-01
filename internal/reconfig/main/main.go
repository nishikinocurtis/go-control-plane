package main

import (
	"bytes"
	"fmt"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	any1 "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"io"
	"net/http"
	"time"
)

const (
	NewUpstreamHost = "127.0.0.1"
	NewUpstreamPort = 10992
	ClusterName     = "udp_cluster"
)

func makeCluster(clusterName string) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       durationpb.New(5 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_STATIC},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
	}
}

func makeEndpoint(clusterName string) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_UDP,
									Address:  NewUpstreamHost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: NewUpstreamPort,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func makeMessage(clusterName string) *discovery.DiscoveryResponse {
	cls, err := anypb.New(makeCluster(clusterName))
	if err != nil {
		panic(err)
	}
	return &discovery.DiscoveryResponse{
		VersionInfo: "2",
		Resources: []*any1.Any{
			cls,
		},
		TypeUrl: "type.googleapis.com/envoy.config.cluster.v3.Cluster",
	}
}

func main() {
	cls, err := proto.Marshal(makeMessage(ClusterName))
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://127.0.0.1:9903/rr_cluster",
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(cls))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Print(body)
}
