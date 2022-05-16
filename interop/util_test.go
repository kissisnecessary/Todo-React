package interop

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Fatal(err error, t *testing.T) {
	if err != nil {
