package rediss

func (conn *myRedisConn) SADD(key string, members ...interface{}) error {
	err := conn.sAdd(key, members...)
	return err
}

func (conn *myRedisConn) SIsMember(key string, member interface{}) (bool, error) {
	result, err := conn.sIsMember(key, member)
	return result, err
}

func (conn *myRedisConn) SCard(key string) (int, error) {
	count, err := conn.sCARD(key)
	return count, err
}

func (conn *myRedisConn) Smembers(key string) ([]interface{}, error) {
	result, err := conn.sMembers(key)
	return result, err
}
