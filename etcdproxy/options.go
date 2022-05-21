package etcdproxy

import (
	"time"

	"github.com/Golang-Tools/optparams"
	clientv3 "go.etcd.io/etcd/client/v3"
)

//Option 设置key行为的选项
//@attribute MaxTTL time.Duration 为0则不设置过期
//@attribute AutoRefresh string 需要为crontab格式的字符串,否则不会自动定时刷新
type Options struct {
	*clientv3.Config
	QueryTimeout     time.Duration
	ParallelCallback bool
}

var DefaultOptions = Options{
	QueryTimeout: time.Duration(50) * time.Millisecond,
	Config:       &clientv3.Config{},
}

//WithQueryTimeout 设置最大过期时间,单位ms
//@params timeout int 请求etcd的最大超时,单位ms
func WithQueryTimeout(timeout int) optparams.Option[Options] {
	return optparams.NewFuncOption(func(o *Options) {
		o.QueryTimeout = time.Duration(timeout) * time.Millisecond
	})
}

//WithParallelCallback 设置callback并行执行
func WithParallelCallback() optparams.Option[Options] {
	return optparams.NewFuncOption(func(o *Options) {
		o.ParallelCallback = true
	})
}

//WithEtcdConnConfig 设置etcd的连接项
//@params conf *clientv3.Config etcd的连接配置
func WithEtcdConfig(conf *clientv3.Config) optparams.Option[Options] {
	return optparams.NewFuncOption(func(o *Options) {
		o.Config = conf
	})
}

//WithAutoSyncInterval 设置etcd连接的最大自动同步周期
//@params iterval int 最大自动同步周期,单位ms
func WithAutoSyncInterval(iterval int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.AutoSyncInterval = time.Duration(iterval) * time.Millisecond
	})
}

//WithDialTimeout 设置etcd连接的拨号超时时间
//@params timeout int 拨号超时时间,单位ms
func WithDialTimeout(timeout int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.DialTimeout = time.Duration(timeout) * time.Millisecond
	})
}

//WithDialKeepAliveTime 设置etcd连接的拨号存活时长
//@params alivetime int 拨号存活时长,单位ms
func WithDialKeepAliveTime(alivetime int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.DialKeepAliveTime = time.Duration(alivetime) * time.Millisecond
	})
}

//WithDialKeepAliveTimeout 设置etcd连接的拨号存活超时时长
//@params timeout int 拨号存活超时时长,单位ms
func WithDialKeepAliveTimeout(timeout int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.DialKeepAliveTimeout = time.Duration(timeout) * time.Millisecond
	})
}

//WithMaxCallSendMsgSize 设置etcd连接的最大发送消息大小
//@params size int 最大发送消息大小,单位bytes
func WithMaxCallSendMsgSize(size int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.MaxCallSendMsgSize = size
	})
}

//WithMaxCallRecvMsgSize 设置etcd连接的最大接收消息大小
//@params size int 最大接收消息大小,单位bytes
func WithMaxCallRecvMsgSize(size int) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.MaxCallRecvMsgSize = size
	})
}

//WithRejectOldCluster 设置etcd连接拒绝旧集群
func WithRejectOldCluster() optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.RejectOldCluster = true
	})
}

//WithPermitWithoutStream 设置etcd连接允许没有流
func WithPermitWithoutStream() optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.PermitWithoutStream = true
	})
}

//WithUsername 设置etcd连接的用户名
//@params username string 用户名
func WithUsername(username string) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.Username = username
	})
}

//WithPassword 设置etcd连接的密码
//@params pwd string 密码
func WithPassword(pwd string) optparams.Option[clientv3.Config] {
	return optparams.NewFuncOption(func(o *clientv3.Config) {
		o.Password = pwd
	})
}

//UseEtcdOpts 设置etcd的连接项
func UseEtcdOpts(opts ...optparams.Option[clientv3.Config]) optparams.Option[Options] {
	return optparams.NewFuncOption(func(o *Options) {
		if o.Config == nil {
			o.Config = new(clientv3.Config)
		}
		optparams.GetOption(o.Config, opts...)
	})
}
