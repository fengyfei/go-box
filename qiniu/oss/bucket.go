package oss

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/qiniu/api.v7/conf"
	"github.com/qiniu/api.v7/storage"
)

// Manager wraps the official BucketManager.
type Manager struct {
	*storage.BucketManager
}

// NewManager creates a bucket manager via BucketManager.
func NewManager(bm *storage.BucketManager) *Manager {
	return &Manager{
		bm,
	}
}

func (m *Manager) context() context.Context {
	return context.WithValue(context.TODO(), "mac", m.Mac)
}

func (m *Manager) requsetHost() string {
	return m.Cfg.RsReqHost()
}

// BucketExists checks if a bucket already exists.
func (m *Manager) BucketExists(name string) (bool, error) {
	buckets, err := m.Buckets(true)
	if err != nil {
		return false, err
	}

	for _, bucket := range buckets {
		if bucket == name {
			return true, nil
		}
	}

	return false, nil
}

// CreateBucket a bucket by the given name.
// zone values
// 	z0 华东
// 	z1 华北
// 	z2 华南
// 	na0 北美
// 	as0 东南亚
func (m *Manager) CreateBucket(name, zone string) error {
	if zone == "" {
		zone = "z0"
	}

	ctx := m.context()
	reqURL := fmt.Sprintf("%s/mkbucketv2/%s/region/%s", m.requsetHost(), base64.StdEncoding.EncodeToString([]byte(name)), zone)

	headers := http.Header{}
	headers.Add("Content-Type", conf.CONTENT_TYPE_FORM)
	err := m.Client.Call(ctx, nil, "POST", reqURL, headers)

	if responseCode(err) == errCodeBucketAlreadyExists {
		return nil
	}

	return err
}
