package main

import (
	"flag"
	"fmt"
	"net/http"
	"novel/websocket/server/internal/config"

	"github.com/tal-tech/go-zero/core/conf"

	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/server-api.yaml", "the config file")

func main() {

	// flag.Parse()

	//	var c rest.RestConf
	//	//logx.Disable()
	//	json.Unmarshal([]byte(`
	//{
	//	"Host":"0.0.0.0",
	//    "Port":8889,
	//	"Log":{
	//        "ServiceName":"",
	//        "Mode":"console",
	//        "TimeFormat":"",
	//        "Path":"logs",
	//        "Level":"info",
	//        "Compress":false,
	//        "KeepDays":0,
	//        "StackCooldownMillis":100
	//    },
	//
	//    "Name":"server-api",
	//    "Mode":"pro",
	//    "MetricsUrl":"",
	//    "Prometheus":{
	//        "Host":"",
	//        "Port":0,
	//        "Path":""
	//    },
	//    "Telemetry":{
	//        "Name":"",
	//        "Endpoint":"",
	//        "Sampler":0,
	//        "Batcher":""
	//    },
	//
	//    "CertFile":"",
	//    "KeyFile":"",
	//    "Verbose":false,
	//    "MaxConns":0,
	//    "MaxBytes":0,
	//    "Timeout":3000,
	//    "CpuThreshold":900,
	//    "Signature":{
	//        "Strict":false,
	//        "Expiry":0,
	//        "PrivateKeys":null
	//    }
	//}
	//`), &c)
	//c := rest.RestConf{
	//	ServiceConf: service.ServiceConf{
	//		Log: logx.LogConf{
	//			Mode: "console",
	//		},
	//		Mode: "pro",
	//		Name: "server-api",
	//	},
	//	Port: 8889,
	//	Host: "0.0.0.0",
	//	//MaxConns: 	  1000,
	//	//MaxBytes:     1000,
	//
	//	// Timeout:      *timeout,
	//	// CpuThreshold: *cpu,
	//}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//data, err := json.Marshal(c.RestConf)
	//fmt.Println(string(data), err)
	//return

	// ctx := svc.NewServiceContext(c)
	// fmt.Println(c)
	server := rest.MustNewServer(c.RestConf)
	//fmt.Printf("server type is %T\n", server)
	defer server.Stop()

	// handler.RegisterHandlers(server, ctx)
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws/:token",
		Handler: func(writer http.ResponseWriter, request *http.Request) {
			fmt.Printf("w's type is %T\n", writer)
		},
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
