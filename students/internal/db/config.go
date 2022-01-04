package db

import "time"

type Config struct {
	Name              string
	URL               string
	ConnectionTimeout time.Duration
}
