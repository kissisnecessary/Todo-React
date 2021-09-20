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
	str_time := time.Format(base_format)
	msg := []byte(str_time)
	// generate key from source
	var source, target SM2
	var err error
	fmt.Println("source lib " + sourceDef)
	if sourceDef == "TJ" {
		source, err = Ne