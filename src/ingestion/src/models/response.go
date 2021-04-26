package models

type Response struct {
	Error error `json:"err,omitempty"`
}

func (r Response) error() error { return r.Error }
