package proxy

import (
	"context"
	"strings"

	log "github.com/Golang-Tools/loggerhelper/v2"
	"github.com/Golang-Tools/optparams"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var Logger *log.Log

func init() {
	log.Set(log.WithExtFields(log.Dict{"module": "etcd-proxy"}))
	Logger = log.Export()
	log.Set(log.WithExtFields(log.Dict{}))
}

//Callback redis操作的回调函数
type Callback func(cli *clientv3.Client) error

//EtcdProxyredis客户端的代理
type EtcdProxy struct {
	*clientv3.Client
	Opt       Options
	callBacks []Callback
}

// New 创建一个新的数据库客户端代理
func New() *EtcdProxy {
	proxy := new(EtcdProxy)
	defaultopt := DefaultOptions
	proxy.Opt = defaultopt
	return proxy
}

// IsOk 检查代理是否已经可用
func (proxy *EtcdProxy) IsOk() bool {
	return proxy.Client != nil
}

//SetConnect 设置连接的客户端
//@params cli UniversalClient 满足redis.UniversalClient接口的对象的指针
func (proxy *EtcdProxy) SetConnect(cli *clientv3.Client) error {
	if proxy.IsOk() {
		return ErrProxyAllreadySettedClient
	}
	proxy.Client = cli
	if proxy.Opt.ParallelCallback {
		for _, cb := range proxy.callBacks {
			go func(cb Callback) {
				err := cb(proxy.Client)
				if err != nil {
					Logger.Error("regist callback get error", log.Dict{"err": err})
				} else {
					Logger.Debug("regist callback done")
				}
			}(cb)
		}
	} else {
		for _, cb := range proxy.callBacks {
			err := cb(proxy.Client)
			if err != nil {
				Logger.Error("regist callback get error", log.Dict{"err": err})
			} else {
				Logger.Debug("regist callback done")
			}
		}
	}
	return nil
}

//Init 从配置条件初始化代理对象
//@params endpoints string 设置etcd连接的地址端点,以`,`分隔
//@params opts ...optparams.Option[Options]
func (proxy *EtcdProxy) Init(endpoints string, opts ...optparams.Option[Options]) error {
	optparams.GetOption(&proxy.Opt, opts...)
	eps := strings.Split(endpoints, ",")
	proxy.Opt.Config.Endpoints = eps
	cli, err := clientv3.New(*proxy.Opt.Config)
	if err != nil {
		return err
	}
	return proxy.SetConnect(cli)
}

// Regist 注册回调函数,在init执行后执行回调函数
//如果对象已经设置了被代理客户端则无法再注册回调函数
//@params cb ...Callback 回调函数
func (proxy *EtcdProxy) Regist(cb ...Callback) error {
	if proxy.IsOk() {
		return ErrProxyAllreadySettedClient
	}
	proxy.callBacks = append(proxy.callBacks, cb...)
	return nil
}

// NewCtx 根据注册的超时时间构造一个上下文
func (proxy *EtcdProxy) NewCtx() (ctx context.Context, cancel context.CancelFunc) {
	if proxy.Opt.QueryTimeout > 0 {
		timeout := proxy.Opt.QueryTimeout
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	return
}

//Default 默认的etcd代理对象
var Default = New()
