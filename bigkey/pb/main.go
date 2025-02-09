package main

import (
	"fmt"
	"log"

	// 导入 protobuf 自动生成的包
	pb "bigkey/pb/product"
	"google.golang.org/protobuf/proto"
)

func jsonLens() {
	dataStr := []byte(`{
        "product_id":1,
        "name":"aaaa",
        "price":111,
        "url":"https://www.xxx.com/image.jpg"
    }`)
	fmt.Println("jsonLens:", len(dataStr))
}

func pbLens() {
	// 创建一个 Product 对象并填充数据
	product := &pb.Product{
		ProductId: 1,
		Name:      "aaaa",
		Price:     111,
		Url:       "https://www.xxx.com/image.jpg",
	}

	// 序列化为 protobuf 格式的字节流
	data, err := proto.Marshal(product)
	if err != nil {
		log.Fatalf("Error marshalling product: %v", err)
	}

	// 输出字节数
	fmt.Println("pbLens:", len(data))
}

func main() {
	jsonLens()
	pbLens()
}
