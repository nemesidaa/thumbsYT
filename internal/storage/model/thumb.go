package model

type Thumb struct {
	ID         string
	Resolution string
	Data       string
}

type SuccessThumbTX struct{}

type FailedThumbTX struct{}
