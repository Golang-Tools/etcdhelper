# etcdhelper

etcd的帮助程序,主要是代理对象,使用泛型所以只支持go 1.18+

## 构成

子模块| 功能
---|---
`proxy`| 代理etcd客户端
`kvmap`| 提取`[]*mvccpb.KeyValue`中的键值对构造成map
