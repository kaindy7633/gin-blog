package app

import (
	"gin-blog/pkg/logging"

	"github.com/beego/beego/v2/core/validation"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
