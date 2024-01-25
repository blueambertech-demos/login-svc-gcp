package data

import "time"

const (
	ProjectID   = "<GCP Project ID here>"
	ServiceName = "login-svc-with-gcp"
)

type LoginDetails struct {
	UserName    string
	PassHash    string
	Salt        string
	DateCreated time.Time
}
