# bigkey
## 危害
key 或者 value过大，会造成内存、cpu、网络消耗增大

## 定义bigkey
* Key本身的数据量过大。例如，一个String类型的Key，它的值为5 MB。

* Key中的成员数过多。例如，一个ZSET类型的Key，它的成员数量为10,000个。

* Key中成员的数据量过大。例如，一个Hash类型的Key，它的成员数量虽然只有2,000个但这些成员的Value（值）总大小为100 MB。

[https://www.alibabacloud.com/help/zh/tair/user-guide/identify-and-handle-large-keys-and-hotkeys](https://www.alibabacloud.com/help/zh/tair/user-guide/identify-and-handle-large-keys-and-hotkeys)

## 解决
### 1. 压缩 key
序列化压缩，例如 pb
```
brew install protobuf  //保持最新稳定版
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest //保持最新稳定版
protoc --go_out=. --go_opt=paths=source_relative bigkey/pb/product/product.proto
```
执行
```
go run main.go
jsonLens: 121
pbLens: 41
```


### 2. 拆分 key
基于版本号机制拆分
```go
// 计算字节切片的MD5值，并返回后6位
func md5LastSixBytes(data []byte) string {
	hash := md5.Sum(data)
	// 将MD5的16字节数组转换为32位的十六进制字符串
	md5String := hex.EncodeToString(hash[:])
	// 取MD5字符串的后6位
	lastSix := md5String[len(md5String)-6:]
	return lastSix
}

// 数据元信息
type MetaInfo struct {
	Data     []byte   `json:"data"` // 如果不是大Key直接取这个字段
	IsBigKey bool     `json:"is_big_key"`
	Keys     []string `json:"keys"`
}

// 将大Value按字节大小拆分后存入Redis
func storeValueInRedis(ctx context.Context, key string, value []byte, chunkSize int) error {
	// 计算需要多少个chunk
	totalChunks := len(value) / chunkSize
	if len(value)%chunkSize != 0 {
		totalChunks++
	}
	// 默认小key
	meta := MetaInfo{IsBigKey: false, Data: value}
	// 大key处理
	if totalChunks > 1 {
		// md5后6位作为数据版本号
		version := md5LastSixBytes(value)
		keys := make([]string, 0, totalChunks)
		// 创建Pipeline
		pipe := redisClient.Pipeline()
		// 存储每个chunk
		for i := 0; i < totalChunks; i++ {
			start := i * chunkSize
			end := (i + 1) * chunkSize
			if end > len(value) {
				end = len(value)
			}
			chunk := value[start:end]

			// 构造每个chunk的key
			chunkKey := fmt.Sprintf("%s:%s:%d", key, version, i)
			keys = append(keys, chunkKey)
			// 将chunk存入Pipeline
			pipe.Set(ctx, chunkKey, chunk, 0)

		}
		// 执行Pipeline中的所有命令
		_, err := pipe.Exec(ctx)
		if err != nil {
			return err
		}
		meta = MetaInfo{IsBigKey: true, Keys: keys, Data: nil}
	}
	metaByte, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	// 获取原来的数据元信息
	oldMetaByte, err := redisClient.Get(ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	_, err = redisClient.Set(ctx, key, metaByte, 0).Result()
	if err != nil {
		return err
	}
	if len(oldMetaByte) > 0 {
		var oldMetaInfo MetaInfo
		err = json.Unmarshal(oldMetaByte, &oldMetaInfo)
		if err != nil {
			return err
		}
		if oldMetaInfo.IsBigKey {
			// 获取旧key,设置旧key过期时间，比如说10分钟，防止服务端还有旧数据在读
		}
	}
	return nil
}

// 从Redis获取数据，使用Pipeline机制
func getDataFromRedis(ctx context.Context, key string) ([]byte, error) {
	// 获取数据元信息
	metaByte, err := redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var metaInfo MetaInfo
	err = json.Unmarshal(metaByte, &metaInfo)
	if err != nil {
		return nil, err
	}

	if !metaInfo.IsBigKey {
		// 如果不是大Key，直接返回Data字段
		return metaInfo.Data, nil
	}

	// 如果是大Key，使用Pipeline从多个键中获取数据
	pipe := redisClient.Pipeline()

	// 将所有Get操作添加到Pipeline
	for _, chunkKey := range metaInfo.Keys {
		pipe.Get(ctx, chunkKey)
	}

	// 执行Pipeline中的所有命令
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	// 处理Pipeline的结果
	var data []byte
	for _, cmd := range cmds {
		if cmd.Err() != nil {
			return nil, cmd.Err()
		}

		chunkData := []byte(cmd.String())
		if err != nil {
			return nil, err
		}
		data = append(data, chunkData...)
	}

	return data, nil
}

```