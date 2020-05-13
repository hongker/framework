package redis

import (
	"github.com/hongker/framework/errors"
	"github.com/go-redis/redis"
)

// Connect 连接redis单机,返回连接与错误信息
func Connect(option *redis.Options) (*redis.Client, error) {
	connection := redis.NewClient(option)
	_, err := connection.Ping().Result()
	if err != nil {
		return nil, errors.RedisConnectFailed("%s", err.Error())
	}

	return connection, nil
}

// ConnectCluster 连接redis集群，返回连接与错误信息
func ConnectCluster(option *redis.ClusterOptions) (*redis.ClusterClient, error)  {
	rdb := redis.NewClusterClient(option)

	if err := rdb.Ping().Err(); err != nil {
		return nil, errors.RedisConnectFailed("%s", err.Error())
	}

	return rdb, nil
}