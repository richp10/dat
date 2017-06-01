package kvs

import (
	logxi "github.com/Janulka/logxi/v1"
)

var logger logxi.Logger

func init() {
	logger = logxi.New("dat.cache")
}
