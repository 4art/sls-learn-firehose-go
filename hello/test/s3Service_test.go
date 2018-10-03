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
}
