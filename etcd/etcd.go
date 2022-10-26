package etcd

import clientv3 "go.etcd.io/etcd/client/v3"

func NewEtcdClient() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            nil,
		AutoSyncInterval:     0,
		DialTimeout:          0,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		Context:              nil,
		Logger:               nil,
		LogConfig:            nil,
		PermitWithoutStream:  false,
	})
	if err != nil {

	}
	return client
}
