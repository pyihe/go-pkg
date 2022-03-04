package redispkg

func (conn *myRedisConn) RPush(key string, values ...interface{}) error {
	err := conn.rpush(key, values...)
	return err
}

func (conn *myRedisConn) RPushX(key string, values ...interface{}) error {
	err := conn.rpushx(key, values...)
	return err
}

func (conn *myRedisConn) LPush(key string, values ...interface{}) error {
	err := conn.lpush(key, values...)
	return err
}

func (conn *myRedisConn) LPushX(key string, values ...interface{}) error {
	err := conn.lpushx(key, values...)
	return err
}

func (conn *myRedisConn) LPop(key string) (result []byte, err error) {
	result, err = conn.lpop(key)
	return result, err
}

func (conn *myRedisConn) RPop(key string) (result []byte, err error) {
	result, err = conn.rpop(key)
	return result, err
}
