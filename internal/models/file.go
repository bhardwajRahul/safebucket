package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name      string
	Extension string
	Uploaded  bool
	BucketId  uint
	Bucket    Bucket
}
