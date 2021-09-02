package interop

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSM2(t *testing.T) {
	sourceDef := os.Getenv("SOURCE")
	tar