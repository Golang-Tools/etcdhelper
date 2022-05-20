package proxy

import (
	"testing"

	"github.com/Golang-Tools/etcdhelper/kvmap"
	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdProxy_Get_Set(t *testing.T) {
	Default.Init(UseEtcdOpts(WithEndpoints("localhost:12379")))
	defer Default.Close()

	ctx, cancel := Default.NewCtx()
	defer cancel()

	_, err := Default.Put(ctx, "foo", "bar")
	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Put wrong")
	}
	resp, err := Default.Get(ctx, "foo")
	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Get wrong")
	}
	assert.Equal(t, "bar", string(resp.Kvs[0].Value))
}

func TestEtcdProxy_Get_Set_With_Prefix(t *testing.T) {
	Default.Init(UseEtcdOpts(WithEndpoints("localhost:12379")))
	defer Default.Close()

	ctx, cancel := Default.NewCtx()
	defer cancel()

	_, err := Default.Put(ctx, "/foo/x", "bar1")
	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Put wrong")
	}
	_, err = Default.Put(ctx, "/foo/y", "bar2")
	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Put wrong")
	}
	resp, err := Default.Get(ctx, "/foo", clientv3.WithPrefix())

	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Get wrong")
	}
	result := kvmap.KVToMap(resp.Kvs...)

	assert.Equal(t, "bar1", result["/foo/x"])
	assert.Equal(t, "bar2", result["/foo/y"])
}

func TestEtcdProxy_Get_Empty(t *testing.T) {
	Default.Init(UseEtcdOpts(WithEndpoints("localhost:12379")))
	defer Default.Close()

	ctx, cancel := Default.NewCtx()
	defer cancel()

	// _, err := Default.Put(ctx, "foo", "bar")
	// if err != nil {
	// 	assert.FailNow(t, err.Error(), "cli.Put wrong")
	// }
	resp, err := Default.Get(ctx, "empty")
	if err != nil {
		assert.FailNow(t, err.Error(), "cli.Get wrong")
	}
	result := kvmap.KVToMap(resp.Kvs...)
	assert.Equal(t, 0, len(result))

}
