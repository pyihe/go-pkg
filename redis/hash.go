package redis

func (conn *myRedisConn) HGet(key string, field string) ([]byte, error) {
	value, err := conn.hGet(key, field)
	return value, err
}

func (conn *myRedisConn) HSet(key string, field string, value interface{}) error {
	err := conn.hSet(key, field, value)
	return err
}

func (conn *myRedisConn) HGetAll(key string) ([]interface{}, error) {
	result, err := conn.hGetAll(key)
	return result, err
}

func (conn *myRedisConn) HKeys(key string) ([]string, error) {
	keys, err := conn.hKeys(key)
	return keys, err
}

func (conn *myRedisConn) HMset(key string, fieldValues ...interface{}) error {
	err := conn.hMset(key, fieldValues...)
	return err
}

func (conn *myRedisConn) HDel(key, field string) (int, error) {
	num, err := conn.hDel(key, field)
	return num, err
}
