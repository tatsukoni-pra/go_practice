package event

import (
	"github.com/aws/aws-lambda-go/events"
)

type S3Info struct {
	Bucket string
	Key  string
}

func GetS3TrigerInfo(S3Event events.S3Event) *S3Info {
	for _, record := range S3Event.Records {
		return &S3Info {
			Bucket: record.S3.Bucket.Name,
			Key: record.S3.Object.Key}
	}
	return &S3Info {
		Bucket: "no-info",
		Key: "no-info"}
}
