# lua

## 通过 lua 可以实现原子操作，可用于秒杀场景库存计算等

lua.go

```go
var incrBy = redis.NewScript(`
local key = KEYS[1]
local change = ARGV[1]

local value = redis.call("GET", key)
if not value then
  value = 0
end

value = value + change
redis.call("SET", key, value)

return value
`)
```

You can then run the script like this:

```go
keys := []string{"my_counter"}
values := []interface{}{+1}
num, err := incrBy.Run(ctx, rdb, keys, values...).Int()
```


## 可用于限速器（令牌桶）

ratelimit.go

可用于 http  中间件场景 ： ratelimit_middleware.go

**如果在秒杀场景限制用户尝试下单，可以按照 product:userid 组装 key 进行限制**

```go
func rateLimit(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
    return func(w http.ResponseWriter, req bunrouter.Request) error {
        res, err := limiter.Allow(req.Context(), "project:123", redis_rate.PerMinute(10))
        if err != nil {
            return err
        }

        h := w.Header()
        h.Set("RateLimit-Remaining", strconv.Itoa(res.Remaining))

        if res.Allowed == 0 {
            // We are rate limited.

            seconds := int(res.RetryAfter / time.Second)
            h.Set("RateLimit-RetryAfter", strconv.Itoa(seconds))

            // Stop processing and return the error.
            return ErrRateLimited
        }

        // Continue processing as normal.
        return next(w, req)
    }
}
```


## 参考

1. https://redis.uptrace.dev/guide/lua-scripting.html#redis-script
2. https://redis.uptrace.dev/guide/go-redis-rate-limiting.html
3. https://github.com/go-redis/redis_rate/blob/v10/lua.go 【ratelimit lua 脚本】