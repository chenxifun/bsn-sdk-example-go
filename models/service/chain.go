package service

import "time"

type TransInfo struct {
	BlockHash   string
	BlockNumber uint64

	CreateName string
	CreateTime time.Time
}
