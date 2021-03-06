package fakes

import (
	"io"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
)

type AzureClient struct {
	ListBlobsCall struct {
		CallCount int
		Receives  struct {
			ListBlobsParameters storage.ListBlobsParameters
		}
		Returns struct {
			BlobListResponse storage.BlobListResponse
			Error            error
		}
	}
	GetBlobSizeInBytesCall struct {
		CallCount int
		Receives  struct {
			BlobName string
			Snapshot time.Time
		}
		Returns struct {
			BlobSize int64
			Error    error
		}
	}
	GetCall struct {
		CallCount int
		Receives  struct {
			BlobName string
			Snapshot time.Time
		}
		Returns struct {
			BlobData []byte
			Error    error
		}
	}
	GetRangeCall struct {
		CallCount int
		Receives  []GetRangeCallReceives
		Returns   struct {
			BlobReader io.ReadCloser
			Error      error
		}
	}
	UploadFromStreamCall struct {
		CallCount int
		Stub      func(string, io.Reader) error
		Receives  struct {
			BlobName string
			Stream   io.Reader
		}
		Returns struct {
			Error error
		}
	}
	CreateSnapshotCall struct {
		CallCount int
		Receives  struct {
			BlobName string
		}
		Returns struct {
			Snapshot time.Time
			Error    error
		}
	}
	GetBlobURLCall struct {
		CallCount int
		Receives  struct {
			BlobName string
		}
		Returns struct {
			URL   string
			Error error
		}
	}
}

type GetRangeCallReceives struct {
	BlobName          string
	StartRangeInBytes uint64
	EndRangeInBytes   uint64
	Snapshot          time.Time
}

func (a *AzureClient) ListBlobs(params storage.ListBlobsParameters) (storage.BlobListResponse, error) {
	a.ListBlobsCall.CallCount++
	a.ListBlobsCall.Receives.ListBlobsParameters = params
	return a.ListBlobsCall.Returns.BlobListResponse, a.ListBlobsCall.Returns.Error
}

func (a *AzureClient) GetBlobSizeInBytes(blobName string, snapshot time.Time) (int64, error) {
	a.GetBlobSizeInBytesCall.CallCount++
	a.GetBlobSizeInBytesCall.Receives.BlobName = blobName
	a.GetBlobSizeInBytesCall.Receives.Snapshot = snapshot
	return a.GetBlobSizeInBytesCall.Returns.BlobSize, a.GetBlobSizeInBytesCall.Returns.Error
}

func (a *AzureClient) Get(blobName string, snapshot time.Time) ([]byte, error) {
	a.GetCall.CallCount++
	a.GetCall.Receives.BlobName = blobName
	a.GetCall.Receives.Snapshot = snapshot
	return a.GetCall.Returns.BlobData, a.GetCall.Returns.Error
}

func (a *AzureClient) GetRange(blobName string, startRangeInBytes, endRangeInBytes uint64, snapshot time.Time) (io.ReadCloser, error) {
	a.GetRangeCall.CallCount++
	a.GetRangeCall.Receives = append(a.GetRangeCall.Receives, GetRangeCallReceives{
		BlobName:          blobName,
		StartRangeInBytes: startRangeInBytes,
		EndRangeInBytes:   endRangeInBytes,
		Snapshot:          snapshot,
	})
	return a.GetRangeCall.Returns.BlobReader, a.GetRangeCall.Returns.Error
}

func (a *AzureClient) UploadFromStream(blobName string, stream io.Reader) error {
	a.UploadFromStreamCall.CallCount++
	a.UploadFromStreamCall.Receives.BlobName = blobName
	a.UploadFromStreamCall.Receives.Stream = stream

	if a.UploadFromStreamCall.Stub != nil {
		return a.UploadFromStreamCall.Stub(blobName, stream)
	}

	return a.UploadFromStreamCall.Returns.Error
}

func (a *AzureClient) CreateSnapshot(blobName string) (time.Time, error) {
	a.CreateSnapshotCall.CallCount++
	a.CreateSnapshotCall.Receives.BlobName = blobName
	return a.CreateSnapshotCall.Returns.Snapshot, a.CreateSnapshotCall.Returns.Error
}

func (a *AzureClient) GetBlobURL(blobName string) (string, error) {
	a.GetBlobURLCall.CallCount++
	a.GetBlobURLCall.Receives.BlobName = blobName
	return a.GetBlobURLCall.Returns.URL, a.GetBlobURLCall.Returns.Error
}
