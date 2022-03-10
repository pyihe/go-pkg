package rediss

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/serialize"
)

var (
	ErrInvalidKey      = errors.New("invalid key")
	ErrInvalidEncoder  = errors.New("not figure encoder")
	ErrInvalidConn     = errors.New("invalid redis conn")
	ErrInvalidParamNum = errors.New("invalid param num")
)

type myRedisConn struct {
	conn    redis.Conn //redis连接池
	encoder serialize.Serializer
	prefix  string
}

func (conn *myRedisConn) checkKey(key string) (string, error) {
	if len(key) == 0 {
		return "", ErrInvalidKey
	}
	key = conn.prefix + key
	return key, nil
}

func (conn *myRedisConn) Close() error {
	if conn.conn != nil {
		return conn.conn.Close()
	}
	return ErrInvalidConn
}

/*****************************string操作**********************************/
//获取指定模式的key
func (conn *myRedisConn) getKeys(pattern string) (keys []string, err error) {
	if pattern == "" {
		pattern = conn.prefix + "*"
	}

	keys, err = redis.Strings(conn.conn.Do("KEYS", pattern))
	return keys, err
}

//string
func (conn *myRedisConn) getString(key string) (value string, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return "", err
	}
	value, err = redis.String(conn.conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return value, err
	}
	return value, err
}

func (conn *myRedisConn) setString(key, value string) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	_, err = redis.String(conn.conn.Do("SET", key, value))
	return err
}

//[]byte
func (conn *myRedisConn) getBytes(key string) (value []byte, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	value, err = redis.Bytes(conn.conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return value, err
}

//
func (conn *myRedisConn) setBytes(key string, value []byte) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}

	_, err = redis.String(conn.conn.Do("SET", key, value))
	if err != nil {
		return err
	}
	return nil
}

//int
func (conn *myRedisConn) getInt(key string) (value int, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return 0, err
	}
	value, err = redis.Int(conn.conn.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (conn *myRedisConn) getInt64(key string) (value int64, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return 0, err
	}
	value, err = redis.Int64(conn.conn.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (conn *myRedisConn) setInt(key string, value int64) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	_, err = redis.String(conn.conn.Do("SET", key, value))
	return err
}

//struct
func (conn *myRedisConn) getStruct(key string, data interface{}) (err error) {
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	bytes, err := redis.Bytes(conn.conn.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
		return err
	}
	if err != nil {
		return err
	}
	err = conn.encoder.Decode(bytes, data)
	if err != nil {
		return err
	}
	return nil
}

func (conn *myRedisConn) setStruct(key string, data interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if conn.encoder == nil {
		return ErrInvalidEncoder
	}
	bytes, err := conn.encoder.Encode(data)
	if err != nil {
		return err
	}
	_, err = redis.String(conn.conn.Do("SET", key, bytes))
	return err
}

func (conn *myRedisConn) del(key string) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	_, err = redis.String(conn.conn.Do("DEL", key))
	return err
}

/******************************set操作*********************************/
//往key对应的set添加元素
func (conn *myRedisConn) sAdd(key string, members ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	_, err = conn.conn.Do("MULTI")
	if err != nil {
		return err
	}
	for _, m := range members {
		_, err = conn.conn.Do("SADD", key, m)
		if err != nil {
			return err
		}
	}
	_, err = conn.conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

//判断元素是否为set的元素
func (conn *myRedisConn) sIsMember(key string, member interface{}) (bool, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return false, err
	}
	result, err := redis.Int(conn.conn.Do("SISMEMBER", key, member))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return false, err
	}
	return result == 1, err
}

//随机从集合中获取元素
func (conn *myRedisConn) sRandMember(key string, count uint) (value []interface{}, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	value, err = redis.Values(conn.conn.Do("SRANDMEMBER", key, count))
	return
}

//返回集合中的元素数量
func (conn *myRedisConn) sCARD(key string) (int, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return 0, err
	}
	result, err := redis.Int(conn.conn.Do("SCARD", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

//返回集合中的所有元素
func (conn *myRedisConn) sMembers(key string) ([]interface{}, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	result, err := redis.Values(conn.conn.Do("SMEMBERS", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

/*****************************hash操作**********************************/
func (conn *myRedisConn) hSet(key, field string, value interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	_, err = conn.conn.Do("HSET", key, field, value)
	return err
}

//获取指定域的值
func (conn *myRedisConn) hGet(key, field string) ([]byte, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	result, err := redis.Bytes(conn.conn.Do("HGET", key, field))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}

//返回hash中所有的域
func (conn *myRedisConn) hKeys(key string) (keys []string, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	keys, err = redis.Strings(conn.conn.Do("HKEYS", key))
	return keys, err
}

//返回key对应的所有域和值
func (conn *myRedisConn) hGetAll(key string) (value []interface{}, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	value, err = redis.Values(conn.conn.Do("HGETALL", key))
	return value, err
}

//设置多对field-value
func (conn *myRedisConn) hMset(key string, fields ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if len(fields)%2 != 0 {
		return ErrInvalidParamNum
	}
	_, err = conn.conn.Do("MULTI")
	if err != nil {
		return err
	}
	var field, value interface{}
	for i, v := range fields {
		if i%2 != 0 {
			value = v
		} else {
			field = v
		}
		_, err = redis.String(conn.conn.Do("HMSET", key, field, value))
		if err != nil {
			break
		}
	}
	_, execErr := conn.conn.Do("EXEC")
	return execErr
}

//删除
func (conn *myRedisConn) hDel(key string, field string) (num int, err error) {
	if key, err = conn.checkKey(key); err != nil {
		return 0, err
	}
	num, err = redis.Int(conn.conn.Do("HDEL", key, field))
	return num, err
}

/*****************************list操作**********************************/
//将value插入到list头部
func (conn *myRedisConn) lpush(key string, values ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(conn.conn.Do("LPUSH", key, v))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (conn *myRedisConn) lpushx(key string, values ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(conn.conn.Do("LPUSHX", key, v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (conn *myRedisConn) rpush(key string, values ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(conn.conn.Do("RPUSH", key, v))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (conn *myRedisConn) rpushx(key string, values ...interface{}) error {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return err
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(conn.conn.Do("RPUSHX", key, v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//移除列表的头元素，及左边的那个元素
func (conn *myRedisConn) lpop(key string) ([]byte, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	result, err := redis.Bytes(conn.conn.Do("LPOP", key))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}

func (conn *myRedisConn) rpop(key string) ([]byte, error) {
	var err error
	if key, err = conn.checkKey(key); err != nil {
		return nil, err
	}
	result, err := redis.Bytes(conn.conn.Do("RPOP", key))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}
