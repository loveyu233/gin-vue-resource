package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"server/model/system"
	"time"
)

var GormClient *gorm.DB

var RedisClient *redis.Client

var MongoClient *system.MongoClient

const (
	VerCodeTime = 3 * time.Minute
)
