package model

import (
	"encoding/json"
	"fmt"
	"go_dev/day9/exercises/chat/model"
	"time"

	"github.com/garyburd/redigo/redis"
)

type UserMgr struct {
	pool *redis.Pool
}

var (
	mgr *UserMgr
)

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)

var (
	UserTable = "users"
)

func InitUserMgr() {
	mgr = NewUserMgr(pool)
}

func NewUserMgr(pool *redis.Pool) (mgr *UserMgr) {
	mgr = &UserMgr{
		pool: pool,
	}
	return
}

func (p *UserMgr) getUser(conn redis.Conn, id int) (user *model.User, err error) {

	result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}
		return
	}

	user = &model.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return

}

func (p *UserMgr) Login(id int, passwd string) (user *model.User, err error) {
	conn := p.pool.Get()
	defer conn.Close()

	user, err = p.getUser(conn, id)
	if err != nil {
		fmt.Println("Get user failed, err:", err)
		return
	}
	if user.UserID != id || user.Passwd != passwd {
		err = ErrInvalidPasswd
	}
	user.Status = UserStatusOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now())

	return

}

func (p *UserMgr) Register(user *model.User) (err error) {
	conn := p.pool.Get()
	defer conn.Close()

	if user == nil {
		fmt.Println("invaild user")
		err = ErrInvalidPrarms
	}
	_, err = p.getUser(conn, user.UserID)

	if err == nil {
		err = ErrUserExist
		return
	}

	if err != ErrUserNotExist {
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("Hset", UserTable, fmt.Sprintf("%d", user.UserID), string(data))

	if err != nil {
		fmt.Println("redis Hset err:", err)
		return
	}

	return
}
