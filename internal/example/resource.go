// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package example

import (
	v3 "github.com/cncf/xds/go/xds/core/v3"
	v32 "github.com/cncf/xds/go/xds/type/matcher/v3"
	udp_proxyv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/udp_proxy/v3"
	any1 "github.com/golang/protobuf/ptypes/any"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

const (
	ClusterName  = "svc-a"
	RouteName    = "local_route"
	ListenerName = "listener_1"
	ListenerPort = 10980
	UpstreamHost = "10.214.96.108"
	UpstreamPort = 10730
)

func makeCluster(clusterName string) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       durationpb.New(5 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_EDS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		//LoadAssignment:       makeEndpoint(clusterName),
		DnsLookupFamily: cluster.Cluster_V4_ONLY,
		EdsClusterConfig: &cluster.Cluster_EdsClusterConfig{
			EdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_Ads{
					//	ApiConfigSource: &core.ApiConfigSource{
					//		ApiType: core.ApiConfigSource_GRPC,
					//		GrpcServices: []*core.GrpcService{{
					//			TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					//				EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: "xds_cluster"},
					//			},
					//		}},
					//		TransportApiVersion: core.ApiVersion_V3,
					//	},
				},
				InitialFetchTimeout: &durationpb.Duration{Seconds: 10},
			},
			ServiceName: clusterName,
		},
	}
}

func makeEndpoint(clusterName string, upHost string, upPort uint32, protocol core.SocketAddress_Protocol) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: protocol,
									Address:  upHost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: upPort,
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

func makeRoute(routeName string, clusterName string) *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: routeName,
		VirtualHosts: []*route.VirtualHost{{
			Name:    "local_service",
			Domains: []string{"*"},
			Routes: []*route.Route{{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{
							Cluster: clusterName,
						},
						HostRewriteSpecifier: &route.RouteAction_HostRewriteLiteral{
							HostRewriteLiteral: UpstreamHost,
						},
					},
				},
			}},
		}},
	}
}

func makeHTTPListener(listenerName string, routeClusterA string, listenPort uint32) *listener.Listener {
	routerConfig, _ := anypb.New(&router.Router{})
	// HTTP filter configuration
	stateConfig := &any1.Any{TypeUrl: "type.googleapis.com/envoy.extensions.filters.http.states_replication.v3.StatesReplication"}
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: &route.RouteConfiguration{
				Name: "local_route",
				VirtualHosts: []*route.VirtualHost{
					{
						Name:    "local_service",
						Domains: []string{"*"},
						Routes: []*route.Route{
							{
								Match: &route.RouteMatch{
									PathSpecifier: &route.RouteMatch_Prefix{
										Prefix: "/colibri/v2/conference",
									},
								},
								Action: &route.Route_Route{
									Route: &route.RouteAction{
										ClusterSpecifier: &route.RouteAction_Cluster{
											Cluster: routeClusterA,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		HttpFilters: []*hcm.HttpFilter{
			{
				Name:       wellknown.StatesReplication,
				ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: stateConfig},
			},
			{
				Name:       wellknown.Router,
				ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: routerConfig},
			},
		},
	}
	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	return &listener.Listener{
		Name: listenerName,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: listenPort,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}

func makeConfigSource() *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
		ApiConfigSource: &core.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   core.ApiConfigSource_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*core.GrpcService{{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: "xds_cluster"},
				},
			}},
		},
	}
	return source
}

func makeUDPListener(listenerName string, udpClusterName string, udpListenerPort uint32) *listener.Listener {
	routeAction, err := anypb.New(&udp_proxyv3.Route{Cluster: udpClusterName})
	if err != nil {
		panic(err)
	}
	udpProxyAny, err := anypb.New(&udp_proxyv3.UdpProxyConfig{
		StatPrefix:                "service",
		UsePerPacketLoadBalancing: true,
		RouteSpecifier: &udp_proxyv3.UdpProxyConfig_Matcher{
			Matcher: &v32.Matcher{
				OnNoMatch: &v32.Matcher_OnMatch{
					OnMatch: &v32.Matcher_OnMatch_Action{
						Action: &v3.TypedExtensionConfig{
							Name:        "envoy.extensions.filters.udp.udp_proxy.v3.Route",
							TypedConfig: routeAction,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return &listener.Listener{
		Name: listenerName,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_UDP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: udpListenerPort,
					},
				},
			},
		},
		ListenerFilters: []*listener.ListenerFilter{
			&listener.ListenerFilter{
				Name: "envoy.filters.udp_listener.udp_proxy",
				ConfigType: &listener.ListenerFilter_TypedConfig{
					TypedConfig: udpProxyAny,
				},
			},
		},
	}
}

func GenerateSnapshot() *cache.Snapshot {
	snap, _ := cache.NewSnapshot("1",
		map[resource.Type][]types.Resource{
			resource.ClusterType: {
				makeCluster("udp-jvb"),
				makeCluster("udp-local"),
				makeCluster("local-jvb"),
				makeCluster("jicofo-jvb"),
			},
			// resource.RouteType:    {makeRoute(RouteName, ClusterName)},
			resource.ListenerType: {
				makeUDPListener("gw-udp", "udp-jvb", 10230),
				makeUDPListener("jvb-udp", "udp-local", 10232),
				makeHTTPListener("jicofo-http", "jicofo-jvb", 20231),
				makeHTTPListener("jvb-http", "local-jvb", 20232),
			},
			resource.EndpointType: {
				makeEndpoint("udp-jvb", "10.214.96.112", 10232, core.SocketAddress_UDP),
				makeEndpoint("udp-local", "127.0.0.1", 10000, core.SocketAddress_UDP),
				makeEndpoint("jicofo-jvb", "10.214.96.112", 20232, core.SocketAddress_TCP),
				makeEndpoint("local-jvb", "127.0.0.1", 9900, core.SocketAddress_TCP),
			},
		},
	)
	return snap
}
