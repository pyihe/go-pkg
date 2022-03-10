package rediss

func (conn *myRedisConn) GetKeys(pattern string) (keys []string, err error) {
	keys, err = conn.getKeys(pattern)
	return
}

func (conn *myRedisConn) GetString(key string) (value string, err error) {
	return conn.getString(key)
}

func (conn *myRedisConn) SetString(key, value string) error {
	return conn.setString(key, value)
}

func (conn *myRedisConn) GetBytes(key string) (value []byte, err error) {
	value, err = conn.getBytes(key)
	return value, err
}

func (conn *myRedisConn) SetBytes(key string, value []byte) error {
	return conn.setBytes(key, value)
}

func (conn *myRedisConn) GetInt(key string) (value int, err error) {
	value, err = conn.getInt(key)
	return value, err
}

func (conn *myRedisConn) GetInt64(key string) (value int64, err error) {
	value, err = conn.getInt64(key)
	return value, err
}

func (conn *myRedisConn) SetInt(key string, value int64) error {
	err := conn.setInt(key, value)
	return err
}

func (conn *myRedisConn) GetStruct(key string, data interface{}) (err error) {
	err = conn.getStruct(key, data)
	return err
}

func (conn *myRedisConn) SetStruct(key string, data interface{}) error {
	err := conn.setStruct(key, data)
	return err
}
