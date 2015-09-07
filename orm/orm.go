package orm

import (
	"github.com/garyburd/redigo/redis"
)

var Red interface {
	redis.Conn
}

func InitPass(ip string, port string,pass string) error {
	var err error
	Red, err = redis.Dial("tcp", ip+":"+port,redis.DialPassword(pass))
	return err
}
