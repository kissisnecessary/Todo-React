package interop

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSM2(t *testing.T) {
	sourceDef := os.Getenv("SOURCE")
	targetDef := os.Getenv("TARGET")
	action := os.Getenv("ACTION")

	base_format := "2006-01-02 15:04:05"
	time := time.Now()
	str_time :