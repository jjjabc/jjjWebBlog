package orm

import (
	"github.com/garyburd/redigo/redis"
)

var Red interface {
	redis.Conn
}

func Init(ip string, port string, pass string) error {
	var err error
	if pass != "" {
		DailOption := redis.DialPassword(pass)
		Red, err = redis.Dial("tcp", ip+":"+port, DailOption)

	}else{
		Red, err = redis.Dial("tcp", ip+":"+port)
	}
	return err
}
