package test

import (
	"../service"
	"testing"
)

func TestListBuckets(t *testing.T) {
	strings, e := service.ListBuckets()
	if e != nil {
		t.Error("something went wrong", e)
	}

	if len(strings) == 0 {
		t.Error("list of buckets is empty")
	}
	for _, e := range strings {
		println(e)
	}
}

func TestIsBucketExist(t *testing.T) {
	exist := service.IsBucketExist("learning-serverless")
	if !exist {
		t.Error("something went wrong. Bucket should exist")
	}

	exist = service.IsBucketExist("serverless-learning")
	if exist {
		t.Error("something went wrong. Bucket should not exist")
	}
}
