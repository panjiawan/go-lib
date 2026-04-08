package main

import (
	"fmt"
	"github.com/panjiawan/go-lib/pkg/pstorage"
)

func main() {
	s := pstorage.New(pstorage.Qiniu, &pstorage.Config{
		SecretId:      "",
		SecretKey:     "",
		UseHTTPS:      false,
		UseCdnDomains: false,
	})

	//b := []byte("hello, this is qiniu cloud")
	//bio := bytes.NewReader(b)
	//res, err := s.PutFromStream("test_bucket", "avatar/test/test.log", bio, int64(len(b)))

	err := s.Delete("night-live-2", "picture/20260330/21.jpg")

	fmt.Println(err)
}
