package oss

import (
	"github.com/qiniu/api.v7/storage"
)

const (
	errCodeNotExists           = -1
	errCodeInvalidBucketName   = 400
	errCodeBucketWithoutAuth   = 401
	errCodeTooManyBucket       = 630
	errCodeBucketAlreadyExists = 614 // returns if at least one region contains the bucket
)

func responseCode(err error) int {
	errInfo, ok := err.(*storage.ErrorInfo)
	if !ok {
		return errCodeNotExists
	}

	return errInfo.HttpCode()
}
