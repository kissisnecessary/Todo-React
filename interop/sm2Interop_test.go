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

	base_format := "2