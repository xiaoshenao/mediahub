package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看https://console.cloud.tencent.com/ ，关于地域的详情见https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://mediahubdev-1328920593.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{

			SecretID:  "AKID56CWp8Vp6JYn5mbKXp7JG3iOBA5ATyjT",
			SecretKey: "ycMHCA7ygcYXqIrWAmUAhWHzR3wzgLyc",
		},
	})
	// Case1 使用 Put 上传对象
	key := "dev.config.yaml" //key是cos文件名称
	f, err := os.Open("./dev.config.yaml")
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "private",
		},
	}
	_, err = client.Object.Put(context.Background(), key, f, opt)
	if err != nil {
		panic(err)
	}

	// // Case 2 使用 PUtFromFile 上传本地文件到 COS
	// filepath := "./test"
	// _, err = client.Object.PutFromFile(context.Background(), key, filepath, opt)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	// // Case 3 上传 0 字节文件, 设置输入流长度为 0
	// _, err = client.Object.Put(context.Background(), key, strings.NewReader(""),
	//
	//	nil)
	//
	//	if err != nil {
	//		// ERROR
	//	}
}
