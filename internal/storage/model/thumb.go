package model

type Thumb struct {
	ID         string
	Resolution string
	Data       []byte
}

type SuccessThumbTX struct{}

type FailedThumbTX struct{}
