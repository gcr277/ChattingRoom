package dao
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var GlobalPool *redis.Pool
func GlobalPoolInit(protocol string, 
		addrPort string, 
		maxIdle int, 
		maxActive int, 
		idleTimeout time.Duration) {

			GlobalPool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func()(redis.Conn, error){
			return redis.Dial(protocol, addrPort)
		},
	}
}

func GlobalPoolClose(){
	GlobalPool.Close()
}