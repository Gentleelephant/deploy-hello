package global

import (
	"testing"
)

func TestLogger(t *testing.T) {

	logger := Logger

	logger.Info("info")
	logger.Debug("debug")

}
