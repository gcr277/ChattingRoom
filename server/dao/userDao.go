package dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"encoding/json"
)

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) SetPool(p *redis.Pool){
	(*this).pool = p
	return
}

//////////////////////////////////////////////////////
// 根据用户ID，返回一个用户实例和error
func (this *UserDao) GetUserByID(userID int)(obj.User, error){
	conn := (*this).pool.Get()
	defer conn.Close()
	var user obj.User
	res, redisErr := redis.String(conn.Do("HGET", "users", userID))
	if redisErr != nil {
		if redisErr == redis.ErrNil{
			return user, obj.ERROR_USER_NOTEXIST
		}else{
			return user, redisErr
		}
	}
	unmarshallErr := json.Unmarshal([]byte(res), &user)
	if unmarshallErr != nil {
		return user, unmarshallErr
	}
	return user, nil
}

// 
func (this *UserDao)SetNewUser(registerMes obj.RegisterMes)error{
	conn := (*this).pool.Get()
	defer conn.Close()
	_, redisErr1 := redis.String(conn.Do("HGET", "users", registerMes.UserID))
	if redisErr1 != nil {
		if redisErr1 == redis.ErrNil{// 说明数据库查不到，是新键，继续set
			registerMesSli, marshalErr := json.Marshal(registerMes)
			if marshalErr != nil{
				return marshalErr
			}
			res, redisErr2 := redis.Int64(conn.Do("HSET", "users", registerMes.UserID, string(registerMesSli)))
			fmt.Printf("[debug-%v]:(%v)|err:(%v)\n", info.CurrFuncName(), res, redisErr2)
			if redisErr2 != nil{
				return redisErr2
			}
			return nil
		}else{ // 其他的错误
			return redisErr1
		}
	}
	return obj.ERROR_USER_EXISTS
}

func (this *UserDao)SetOldUser(registerMes obj.RegisterMes)error{

	return nil
}