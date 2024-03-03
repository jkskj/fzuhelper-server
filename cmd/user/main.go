package main

import (
	"flag"
	"fmt"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user/userservice"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/elastic/go-elasticsearch"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/west2-online/fzuhelper-server/cmd/template/dal"
	"github.com/west2-online/fzuhelper-server/cmd/template/rpc"
	"github.com/west2-online/fzuhelper-server/config"
	"github.com/west2-online/fzuhelper-server/pkg/constants"
	"github.com/west2-online/fzuhelper-server/pkg/eslogrus"
	"github.com/west2-online/fzuhelper-server/pkg/tracer"
	"github.com/west2-online/fzuhelper-server/pkg/utils"
)

var (
	serviceName = constants.UserServiceTableName
	path        *string
	listenAddr  string // listen port

	EsClient *elasticsearch.Client
)

func Init() {
	// config init
	path = flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path, serviceName)

	// dal
	dal.Init()

	// trace
	tracer.InitJaeger(serviceName)

	// rpc
	rpc.Init()

	// log
	EsInit()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))
}

func main() {
	Init() // 做一些中间件的初始化

	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	// get available port from config set
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available port from config")
		}
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)

	if err != nil {
		panic(err)
	}

	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
	)

	if err = svr.Run(); err != nil {
		panic(err)
	}
}

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.Elasticsearch.Host, logrus.DebugLevel, serviceName)
	if err != nil {
		panic(err)
	}

	return hook
}

// InitEs 初始化es
func EsInit() {
	esConn := fmt.Sprintf("http://%s", config.Elasticsearch.Addr)
	cfg := elasticsearch.Config{
		Addresses: []string{esConn},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	EsClient = client
}
