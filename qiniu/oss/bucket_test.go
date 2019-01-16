package oss

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")

	mac           *qbox.Mac
	bucketManager *storage.BucketManager
	manager       *Manager
)

func init() {
	if accessKey == "" || secretKey == "" {
		panic("please set QINIU_ACCESS_KEY & QINIU_SECRET_KEY first")
	}

	mac = qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{}
	cfg.Zone = &storage.Zone_z1
	cfg.UseCdnDomains = true
	bucketManager = storage.NewBucketManager(mac, &cfg)
	manager = NewManager(bucketManager)
	rand.Seed(time.Now().Unix())
}

//Test get bucket list
func TestSharedBuckets(t *testing.T) {
	buckets, err := manager.Buckets(true)
	if err != nil {
		t.Fatalf("Buckets() error, %s", err)
	}

	for _, bucket := range buckets {
		t.Log(bucket)
	}
}

func TestCreateBucket(t *testing.T) {
	if err := manager.CreateBucket("test", "z1"); err != nil {
		t.Fatalf("Create() error, [%d]:%s", responseCode(err), err)
	}
}

func TestBucketExists(t *testing.T) {
	exist, err := manager.BucketExists("test")
	if err != nil {
		t.Fatalf("Create() error, [%d]:%s", responseCode(err), err)
	}

	if !exist {
		t.Fatalf("Bucket test should exists, but not")
	}
}
