package orm

import (
	"github.com/garyburd/redigo/redis"
)

var Red interface {
	redis.Conn
}

func Init(ip string, port string) error {
	var err error
	Red, err = redis.Dial("tcp", ip+":"+port)
	return err
}
func InitPass(ip string, port string,password string) error {
	var err error
	Red, err = redis.Dial("tcp", ip+":"+port,redis.DialPassword(password))
	return err
}
