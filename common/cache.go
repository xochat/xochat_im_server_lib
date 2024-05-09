package common

import (
	"sync"
	"time"

	"github.com/xochat/xochat_im_server_lib/pkg/redis"
)

// RedisCache redis缓存
type RedisCache struct {
	conn *redis.Conn
}

// NewRedisCache 创建
func NewRedisCache(addr string, password string, db int) *RedisCache {
	r := &RedisCache{}
	r.conn = redis.New(addr, password, db)
	return r
}

// Set Set
func (r *RedisCache) Set(key string, value string) error {
	return r.conn.Set(key, value)
}

// Delete 删除key
func (r *RedisCache) Delete(key string) error {
	return r.conn.Del(key)
}

// SetAndExpire 包含过期时间
func (r *RedisCache) SetAndExpire(key string, value string, expire time.Duration) error {
	return r.conn.SetAndExpire(key, value, expire)
}

// Get Get
func (r *RedisCache) Get(key string) (string, error) {
	return r.conn.GetString(key)
}

func (r *RedisCache) HGet(key string, field string) (string, error) {
	return r.conn.Hget(key, field)
}

func (r *RedisCache) HGetAll(key string) (map[string]string, error) {
	return r.conn.Hgetall(key)
}

// GetRedisConn 获取redis连接
func (r *RedisCache) GetRedisConn() *redis.Conn {
	return r.conn
}

// MemoryCache 内存缓存
type MemoryCache struct {
	cacheMap map[string]string
	sync.RWMutex
}

func (m *MemoryCache) HGet(key string, field string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryCache) HGetAll(key string) (map[string]string, error) {
	//TODO implement me
	panic("implement me")
}

// NewMemoryCache 创建
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cacheMap: map[string]string{},
	}
}

// Set Set
func (m *MemoryCache) Set(key string, value string) error {
	m.Lock()
	m.cacheMap[key] = value
	m.Unlock()
	return nil
}

// SetAndExpire SetAndExpire
func (m *MemoryCache) SetAndExpire(key string, value string, expire time.Duration) error {
	m.Lock()
	m.cacheMap[key] = value
	m.Unlock()
	return nil
}

// Get Get
func (m *MemoryCache) Get(key string) (string, error) {
	m.RLock()
	defer m.RUnlock()
	return m.cacheMap[key], nil
}

// Delete Delete
func (m *MemoryCache) Delete(key string) error {
	m.Lock()
	delete(m.cacheMap, key)
	m.Unlock()
	return nil
}
