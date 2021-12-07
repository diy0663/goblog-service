package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 针对 针对路由里面的请求 进行限流

type MethodLimiter struct {
	// 嵌入最基础的限流器结构体
	*Limiter
}

// 返回值是个interface,实际返回是个struct 说明这个struct 实现了这个interface
func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		Limiter: &Limiter{
			limiterBuckets: make(map[string]*ratelimit.Bucket),
		},
	}
}
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	// strings.Index 函数, 用于获取指定子字符串的第一个实例 。如果未找到子字符串，则此方法将返回-1
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	// 其实就是按照 ? 截取 uri, ? 之前的
	// 这里有疑问, 对于 /api/user/1 和 /api/user/2 这种,估计是有问题的
	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// 新增多个令牌桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			//不存在就可以进行设置
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	// 结构体自身已实现了interface ,所以符合返回类型
	return l
}
