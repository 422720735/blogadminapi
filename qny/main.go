//package main
//
//import (
//	"fmt"
//
//	"github.com/qiniu/api.v7/auth/qbox"
//	"github.com/qiniu/api.v7/storage"
//)
//
//func main() {
//	accessKey := "TgVGKnpCMLDI6hSS4rSWE3g-FZjMPf6ZbcX0Kd7c"
//	secretKey := "zqZvH3fNVaggw00oc9wCrcWKgeeiV7WITFTFds7H"
//	bucket := "wangshubotest"
//	key := "2.log"
//
//	mac := qbox.NewMac(accessKey, secretKey)
//
//	cfg := storage.Config{}
//	cfg.Zone = &storage.ZoneHuadong
//	cfg.UseHTTPS = false
//	cfg.UseCdnDomains = false
//
//	bucketManager := storage.NewBucketManager(mac, &cfg)
//
//	// Get file info
//	fmt.Println("------Get file info------")
//	fileInfo, sErr := bucketManager.Stat(bucket, key)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//
//	// Change the mime of the file
//	fmt.Println("------Change the mime of the file------")
//	newMime := "image/x-png"
//	err := bucketManager.ChangeMime(bucket, key, newMime)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fileInfo, sErr = bucketManager.Stat(bucket, key)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//
//	// Change filetype
//	fmt.Println("------Change filetype------")
//	fileType := 1
//	err = bucketManager.ChangeType(bucket, key, fileType)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fileInfo, sErr = bucketManager.Stat(bucket, key)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//
//	//Copy file
//	fmt.Println("------Copy file------")
//	destBucket := bucket
//	destKey := "2_copy.log"
//	force := false
//	err = bucketManager.Copy(bucket, key, destBucket, destKey, force)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fileInfo, sErr = bucketManager.Stat(bucket, destKey)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//
//	//Rename file
//	fmt.Println("------Rename file------")
//	destKey = "2_new.log"
//	force = false
//	err = bucketManager.Move(bucket, key, destBucket, destKey, force)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fileInfo, sErr = bucketManager.Stat(bucket, destKey)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//
//	// Delete file
//	fmt.Println("------Delete file------")
//	err = bucketManager.Delete(bucket, key)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fileInfo, sErr = bucketManager.Stat(bucket, key)
//	if sErr != nil {
//		fmt.Println(sErr)
//		return
//	}
//	fmt.Println(fileInfo.String())
//	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
//}