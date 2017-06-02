package kvs

import (
	logxi "gopkg.in/mgutz/logxi/v1"
)

var logger logxi.Logger

func init() {
	logger = logxi.New("dat.cache")
}
