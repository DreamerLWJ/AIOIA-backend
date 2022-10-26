package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	_testLockKey = "test_lock"
)

func TestRdsClient_TryLockWithTimeout(t *testing.T) {
	ctx := context.Background()
	client := getTestClient()
	isSuccess, err := client.TryLock(ctx, _testLockKey, time.Second*3)
	assert.Nil(t, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.TryLockWithTimeout(ctx, _testLockKey, time.Second*3, time.Second)
	assert.Nil(t, err)
	assert.Equal(t, false, isSuccess)
	isSuccess, err = client.TryLockWithTimeout(ctx, _testLockKey, time.Second*3, time.Second*3)
	assert.Nil(t, err)
	assert.Equal(t, true, isSuccess)
}

func TestRdsClient_TryLockWithTime(t *testing.T) {
	ctx := context.Background()
	client := getTestClient()
	isSuccess, err := client.TryLock(ctx, _testLockKey, time.Second)
	assert.Nil(t, err)
	assert.Equal(t, true, isSuccess)
	isSuccess, err = client.TryLockWithTime(ctx, _testLockKey, time.Second, 10)
	assert.Nil(t, err)
	assert.Equal(t, false, isSuccess)
	isSuccess, err = client.TryLockWithTime(ctx, _testLockKey, time.Second, 20)
	assert.Nil(t, err)
	assert.Equal(t, true, isSuccess)
}
