package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var svc = createService()

func ListBuckets() ([]string, error) {
	listBucketsOutput, e := svc.ListBuckets(nil)
	buckets := listBucketsOutput.Buckets
	var sBuckets []string
	for _, e := range buckets {
		sBuckets = append(sBuckets, *e.Name)
	}
	return sBuckets, e
}

func createService() *s3.S3 {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	return svc
}