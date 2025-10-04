package s3wrapper

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListObject_WithTestify(t *testing.T) {
	mockClient := new(MockS3Client)
	creationDate := time.Now()

	name := "test-bucket"

	expectedOutput := &s3.ListBucketsOutput{
		Buckets: []types.Bucket{
			{
				Name:         &name,
				CreationDate: &creationDate,
			},
		},
	}

	mockClient.On("ListBuckets", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedOutput, nil)

	wrapper := &S3Client{Client: mockClient}
	result, err := wrapper.ListBuckets()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Buckets))
	assert.Equal(t, "test-bucket", *result.Buckets[0].Name)

	mockClient.AssertExpectations(t)
}
