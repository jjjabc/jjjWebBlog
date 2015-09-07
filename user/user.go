package user

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/jjjabc/jjjWebBlog/orm"
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
	if _, err := GetUid(ju.Name); err == nil {
		return errors.New("username exist！")
	}
	uid, err := generateUid()
	if err != nil {
		return err
	}
	orm.Red.Send("SET", "account:uname:"+ju.Name, strconv.Itoa(uid))
	orm.Red.Send("SET", "account:uid:"+strconv.Itoa(uid), ju.Name)
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
func (this *JJJuser) DelUser() error {
	if _, err := GetUid(this.Name); err != nil {
		return err
	}
	orm.Red.Send("DEL", "account:uname:"+this.Name)
	orm.Red.Send("DEL", "account:uid:"+strconv.Itoa(this.Id))
	orm.Red.Send("DEL", "account:nickname:"+strconv.Itoa(this.Id))
	orm.Red.Send("DEL", "account:des:"+strconv.Itoa(this.Id))
	orm.Red.Send("SREM", "account:UidSets", strconv.Itoa(this.Id))
	return orm.Red.Flush()
}
func (this *JJJuser) Updata() error {
	if _, err := GetUid(this.Name); err.Error() != "redigo: nil returned" {
		return err
	}
	orm.Red.Send("SET", "account:nickname:"+strconv.Itoa(this.Id), this.NickName)
	orm.Red.Send("SET", "account:des:"+strconv.Itoa(this.Id), this.Description)
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
	ju.Name, err = redis.String(orm.Red.Do("GET", "account:uid:"+strconv.Itoa(uid)))
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
func GetUsers(pageNum int, number int) ([]JJJuser, error) {
	Sets := "account:UidSets"
	all, err := redis.Strings(orm.Red.Do("SMEMBERS", Sets))
	if err != nil {
		return nil, errors.New("DB error")
	}
	if len(all) == 0 {
		return make([]JJJuser, 0), nil
	}
	juSets := make([]JJJuser, 0)
	start := (pageNum - 1) * number
	last := len(all) - 1

	for i := start; (i < (start + number)) && (i <= last); i++ {
		uId, _ := strconv.Atoi(all[i])
		ju := GetUser(uId)
		juSets = append(juSets, *ju)
	}
	return juSets, nil

}
func GetAllUsers() ([]JJJuser, error) {
	//2147483647是32位系统中Int的最大值
	return GetUsers(1, 2000000000)
}
