package redispkg

import (
	"fmt"
	"time"

	"github.com/pyihe/go-pkg/serialize"

	"github.com/garyburd/redigo/redis"
)

type RedisConn interface {
	Close() error
	//string
	GetKeys(pattern string) (keys []string, err error)
	GetString(key string) (value string, err error)
	SetString(key, value string) error
	GetBytes(key string) (value []byte, err error)
	SetBytes(key string, value []byte) error
	GetInt(key string) (value int, err error)
	GetInt64(key string) (value int64, err error)
	SetInt(key string, value int64) error
	GetStruct(key string, data interface{}) (err error)
	SetStruct(key string, data interface{}) error

	//hash
	HGet(key string, field string) ([]byte, error)
	HSet(key string, field string, value interface{}) error
	HGetAll(key string) ([]interface{}, error)
	HKeys(key string) ([]string, error)
	HMset(key string, fieldValues ...interface{}) error
	HDel(key, field string) (int, error)

	//list
	RPush(key string, values ...interface{}) error
	RPushX(key string, values ...interface{}) error
	LPush(key string, values ...interface{}) error
	LPushX(key string, values ...interface{}) error
	LPop(key string) (result []byte, err error)
	RPop(key string) (result []byte, err error)

	//set
	SADD(key string, members ...interface{}) error
	SIsMember(key string, member interface{}) (bool, error)
	SCard(key string) (int, error)
	Smembers(key string) ([]interface{}, error)
}

type RedisPool interface {
	Get() (RedisConn, error)
	Close() error
}

type myPool struct {
	prefix  string
	net     string
	addr    string
	pass    string
	db      int
	p       *redis.Pool
	encoder serialize.Serializer
}

type InitOptions func(m *myPool)

func WithEncoding(encoder serialize.Serializer) InitOptions {
	return func(m *myPool) {
		m.encoder = encoder
	}
}

func WithPrefix(prefix string) InitOptions {
	return func(m *myPool) {
		m.prefix = prefix
	}
}

func WithNetWork(net string) InitOptions {
	return func(m *myPool) {
		m.net = net
	}
}

func WithAddr(addr string) InitOptions {
	return func(m *myPool) {
		m.addr = addr
	}
}

func WithPass(pass string) InitOptions {
	return func(m *myPool) {
		m.pass = pass
	}
}

func WithDBIndex(db int) InitOptions {
	return func(m *myPool) {
		m.db = db
	}
}

func NewPool(opts ...InitOptions) (RedisPool, error) {
	defaultPool := &myPool{}
	for _, op := range opts {
		op(defaultPool)
	}
	if defaultPool.addr == "" {
		return nil, fmt.Errorf("no redispkg address")
	}
	if defaultPool.db == 0 {
		defaultPool.db = 1
	}
	if defaultPool.net == "" {
		defaultPool.net = "tcp"
	}
	defaultPool.p = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial(defaultPool.net, defaultPool.addr, redis.DialDatabase(defaultPool.db), redis.DialPassword(defaultPool.pass))
		},
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 120 * time.Second,
		Wait:        true,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return defaultPool, nil
}

func (m *myPool) Get() (RedisConn, error) {
	conn := m.p.Get()
	if conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}
	c := &myRedisConn{
		prefix:  m.prefix,
		conn:    conn,
		encoder: m.encoder,
	}
	return c, nil
}

func (m *myPool) Close() error {
	return m.p.Close()
}
