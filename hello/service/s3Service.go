package service

import (
	"../model"
	"encoding/json"
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

func IsBucketExist(bucket string) bool {
	bucketNames, e := ListBuckets()
	if e != nil {
		panic(e)
	}
	for _, e := range bucketNames {
		if e == bucket {
			return true
		}
	}
	return false
}

func IsBucketEmpty(bucket string) bool {
	return len(getListObjectsOutput(bucket).Contents) == 0
}

func ClearBucket(bucket string) {
	for _, v := range getListObjectsOutput(bucket).Contents {
		svc.DeleteObject(&s3.DeleteObjectInput{Key: v.Key})
	}
}

func PullObjectsFromS3(bucket string) []string {
	var objects []string
	for _, v := range getListObjectsOutput(bucket).Contents {
		objects = append(objects, *v.Key)
	}
	return objects
}

func PullCities(bucket string, key string) []model.City {
	var cities []model.City
	jsonStr := QueryJson(bucket, key, "SELECT * FROM s3Object")
	err := json.Unmarshal([]byte(jsonStr), &cities)
	if err != nil {
		panic(err)
	}
	return cities
}

func QueryJson(bucket string, key string, query string) string {
	compressionType := "GZIP"
	jsonType := "LINES"
	recordDelimiter := "\n"
	selectObjectContentOutput, e := svc.SelectObjectContent(&s3.SelectObjectContentInput{
		Bucket:     &bucket,
		Key:        &key,
		Expression: &query,
		InputSerialization: &s3.InputSerialization{
			CompressionType: &compressionType,
			JSON:            &s3.JSONInput{Type: &jsonType},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{RecordDelimiter: &recordDelimiter},
		},
	})
	if e != nil {
		panic(e)
	}
	return selectObjectContentOutput.GoString()
}

func getListObjectsOutput(bucket string) *s3.ListObjectsOutput {
	listObjectsInput := s3.ListObjectsInput{Bucket: &bucket}
	listObjectsOutput, e := svc.ListObjects(&listObjectsInput)
	if e != nil {
		panic(e)
	}
	return listObjectsOutput
}

func createService() *s3.S3 {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	return svc
}
