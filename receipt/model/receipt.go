package model

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

var ReceiptDir = filepath.Join("uploads")

type Receipt struct {
	Name       string    `json:"name"`
	UploadDate time.Time `json:"uploadDate"`
}

func GetReceipts() ([]Receipt, error) {
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		receipts = append(receipts,
			Receipt{
				Name:       f.Name(),
				UploadDate: f.ModTime(),
			})
	}
	return receipts, nil
}
