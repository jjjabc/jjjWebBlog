package user

import (
	"github.com/garyburd/redigo/redis"
	"jjjBlog/orm"
	"strconv"
)

type JJJuser struct {
	Name        string
	Id          int
	Description string
	NickName    string
}

func SigupUser(ju JJJuser) error {
	orm.Red.Send("SET", "account:uname:"+ju.Name, ju.Id)
	orm.Red.Send("SET", "account:nickname:"+strconv.Itoa(ju.Id), ju.NickName)
	orm.Red.Send("SET", "account:des:"+strconv.Itoa(ju.Id), ju.Description)
	return orm.Red.Flush()
}
func GetUser(uid int) *JJJuser {
	ju := JJJuser{
		Id: uid,
	}
	var err error
	ju.NickName, err = redis.String(orm.Red.Do("GET", "account:nickname:"+strconv.Itoa(uid)))
	if err != nil {
		return nil
	}
	ju.Description, _ = redis.String(orm.Red.Do("GET", "account:des:"+strconv.Itoa(uid)))
	return &ju
}

func GenerateUid() (int, error) {
	return redis.Int(orm.Red.Do("INCR", "account:count"))
}

func GetUid(username string) (int, error) {
	return redis.Int(orm.Red.Do("GET", "account:uname:"+username))
}

func CheckUser(username string, password string) bool {
	uid, err := GetUid(username)
	if err != nil {
		return false
	}
	dbpwd, err := getPassword(uid)
	if err != nil {
		return false
	}
	if dbpwd != password {
		return false
	}
	return true
}
func getPassword(uid int) (string, error) {
	return redis.String(orm.Red.Do("GET", "account:password:"+strconv.Itoa(uid)))
}
func passwordHash(pwd string) string {
	return pwd
}
