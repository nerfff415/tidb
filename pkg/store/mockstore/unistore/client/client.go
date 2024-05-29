package client

import (
	"context"
	"hash/crc32"
	"time"

	"github.com/tikv/client-go/v2/tikvrpc"
)

type Client interface {
	Close() error
	SendRequest(ctx context.Context, addr string, req *tikvrpc.Request, timeout time.Duration) (*tikvrpc.Response, error)
}

func NewShardingStrategy(key []byte) uint64 {
	hash := crc32.ChecksumIEEE(key)
	return uint64(hash)
}

type Region struct {
	ID   uint64
	Name string
}

type ClientImpl struct {
	regions map[uint64]*Region
}

func (c *ClientImpl) GetRegionByKey(key []byte) *Region {
	shardID := NewShardingStrategy(key)
	return c.searchRegion(shardID)
}

func (c *ClientImpl) searchRegion(shardID uint64) *Region {
	return c.regions[shardID]
}
