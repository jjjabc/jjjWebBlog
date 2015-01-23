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

func (ju *JJJuser) SigupUser(pwd string) error {
	//判断用户名是否存在
	if _, err := GetUid(ju.Name); err.Error() != "redigo: nil returned" {
		return err
	}
	uid, err := generateUid()
	if err != nil {
		return err
	}
	orm.Red.Send("SET", "account:uname:"+ju.Name, strconv.Itoa(uid))
	orm.Red.Send("SET", "account:nickname:"+strconv.Itoa(uid), ju.NickName)
	orm.Red.Send("SET", "account:des:"+strconv.Itoa(uid), ju.Description)
	orm.Red.Send("SADD", "account:UidSets", strconv.Itoa(uid))
	if err := savePassword(uid, pwd); err != nil {
		return err
	}
	return orm.Red.Flush()
}
func (this *JJJuser) Reflush() bool {
	ju := GetUser(this.Id)
	if ju != nil {
		this = ju
		return true
	}
	return false
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

func generateUid() (int, error) {
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
	dbPwdHa, err := getPasswordHash(uid)
	if err != nil {
		return false
	}
	if dbPwdHa != passwordHash(password) {
		return false
	}
	return true
}
func getPasswordHash(uid int) (string, error) {
	return redis.String(orm.Red.Do("GET", "account:password:"+strconv.Itoa(uid)))
}
func passwordHash(pwd string) string {
	return pwd
}
func savePassword(uid int, pwd string) error {
	_, err := orm.Red.Do("SET", "account:password:"+strconv.Itoa(uid), passwordHash(pwd))
	return err
}
