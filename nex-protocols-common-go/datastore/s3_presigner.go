package datastore

import (
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type S3PresignerInterface interface {
	GetObject(bucket, key string, lifetime time.Duration) (*url.URL, error)
	PostObject(bucket, key string, lifetime time.Duration) (*url.URL, map[string]string, error)
}

type S3Presigner struct {
	minio *minio.Client
}

func (p *S3Presigner) GetObject(bucket, key string, lifetime time.Duration) (*url.URL, error) {
	reqParams := make(url.Values)

	return p.minio.PresignedGetObject(context.Background(), bucket, key, lifetime, reqParams)
}

func (p *S3Presigner) PostObject(bucket, key string, lifetime time.Duration) (*url.URL, map[string]string, error) {
	policy := minio.NewPostPolicy()

	err := policy.SetBucket(bucket)
	if err != nil {
		return nil, nil, err
	}

	err = policy.SetKey(key)
	if err != nil {
		return nil, nil, err
	}

	err = policy.SetExpires(time.Now().UTC().Add(lifetime).UTC())
	if err != nil {
		return nil, nil, err
	}

	return p.minio.PresignedPostPolicy(context.Background(), policy)
}

func NewS3Presigner(minioClient *minio.Client) *S3Presigner {
	return &S3Presigner{
		minio: minioClient,
	}
}
