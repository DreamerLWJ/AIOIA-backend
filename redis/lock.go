package redis

import (
	"context"
	"errors"
	"time"
)

var (
	nilResError = errors.New("client reply result nil error")
)

const (
	_lockTag           = 1
	_lockSpinFrequency = time.Millisecond * 50
)

func (c *RdsClient) Lock(ctx context.Context, lock string, expire time.Duration) (bool, error) {

	return false, nil
}

// TryLockWithTime try lock and spin times
func (c *RdsClient) TryLockWithTime(ctx context.Context, lock string, expire time.Duration, tries int) (ok bool, err error) {
	for !ok && tries > 0 {
		ok, err = c.TryLock(ctx, lock, expire)
		if err != nil {
			return false, err
		}
		tries--
		time.Sleep(_lockSpinFrequency)
	}
	return ok, nil
}

// TryLockWithTimeout try lock until timeout
func (c *RdsClient) TryLockWithTimeout(ctx context.Context, lock string, expire time.Duration, timeout time.Duration) (ok bool, err error) {
	start := time.Now()

	for !ok {
		ok, err = c.TryLock(ctx, lock, expire)
		if err != nil {
			return false, err
		}
		current := time.Now()
		if current.Sub(start) >= timeout {
			break
		}
	}
	return ok, nil
}

// TryLock try lock one time
func (c *RdsClient) TryLock(ctx context.Context, lock string, expire time.Duration) (bool, error) {
	res := c.SetNX(ctx, lock, _lockTag, expire)
	if res == nil {
		return false, nilResError
	}
	if res.Err() != nil {
		return false, res.Err()
	}
	return res.Result()
}

func (c *RdsClient) UnLock(ctx context.Context, lock string) (bool, error) {
	res := c.Del(ctx, lock)
	if res == nil {
		return false, nilResError
	}
	if res.Err() != nil {
		return false, res.Err()
	}

	return res.Val() == 1, nil
}
