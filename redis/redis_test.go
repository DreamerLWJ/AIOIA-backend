package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestClient() *RdsClient {
	client, err := NewRedisClientToml("./redis_test.toml")
	if err != nil {
		panic(err)
	}
	return client
}

func TestNewRedisClient(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()
	res1 := client.Set(ctx, "test_string", "123", 0)
	assert.NotEqual(t, nil, res1)
	assert.Equal(t, nil, res1.Err())

	res2 := client.Get(ctx, "test_string")
	assert.NotEqual(t, nil, res2)
	assert.Equal(t, nil, res2.Err())
	val := res2.Val()
	assert.Equal(t, "123", val)
}

func TestRdsClient_TryLock(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()
	isSuccess, err := client.TryLock(ctx, "seq", time.Second*30)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.TryLock(ctx, "seq", time.Second*30)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, isSuccess)
	isSuccess, err = client.UnLock(ctx, "seq")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.TryLock(ctx, "seq", time.Second*30)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.UnLock(ctx, "seq")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.UnLock(ctx, "seq")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, isSuccess)
}
