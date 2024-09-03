package model

type Thumb struct {
	ID         string
	Path       string
	Resolution string
}

type SuccessThumbTX struct{}

type FailedThumbTX struct{}
