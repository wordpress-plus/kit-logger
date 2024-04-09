package limiter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wordpress-plus/kit-common/api/response"
	"github.com/wordpress-plus/kit-common/kg"
	"net/http"
)

const (
	localCache = iota
	localCacheV2
	redisCache
)

var ErrLimited = errors.New("rate limit exceeded")

type CheckFn = func(key string, expire int, limit int) (err error)

type LimitInterface interface {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey(c *gin.Context) string

	// CheckOrMark 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark(key string, expire int, limit int) error
}

type LimitBase struct {
	LimitInterface
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l *LimitBase) GenerationKey(c *gin.Context) string {
	return "wordpress-plus_Limit" + c.ClientIP()
}

func (l *LimitBase) Process(check CheckFn, expire int, limit int) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := check(l.GenerationKey(c), expire, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func Limiter() gin.HandlerFunc {
	cacheType := kg.C.System.LimitType

	if localCache == cacheType {
		return LocalLimiter()
	}

	if localCacheV2 == cacheType {
		return LocalLimiterV2()
	}

	if redisCache == cacheType {
		return RedisLimiter()
	}

	panic("未配置缓存类型")
}