package kvhelper

import "go.etcd.io/etcd/api/v3/mvccpb"

func ToMap(kvs ...*mvccpb.KeyValue) (result map[string]string) {
	result = map[string]string{}
	for _, i := range kvs {
		result[string(i.Key)] = string(i.Value)
	}
	return
}
