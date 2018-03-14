package tools

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

var DB *sqlx.DB
var INCLIENT *redis.Database
var Sess *sessions.Sessions

func InitDB() {
	DB, _ = sqlx.Connect("mysql", "root:1q2w3e@tcp(127.0.0.1:3306)/txwzglxt?charset=utf8")

}

func Initclient() {
	INCLIENT = redis.New(service.Config{
		Network:     service.DefaultRedisNetwork,
		Addr:        "localhost:6379",
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      ""})

	Sess = sessions.New(sessions.Config{
		Cookie:                      "mysessionid",
		Expires:                     time.Hour * 2,
		DisableSubdomainPersistence: false,
	})

}
