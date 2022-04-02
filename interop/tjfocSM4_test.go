package interop

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/Hyperledger-TWGC/tjfoc-gm/sm4"
)

/**
1. 执行Hyperledger-TWGC/java-gm中的SM4UtilTest.sm4Interaction4Encrypt()方法，获得Java sm4加密密文，更新代码中待解密参数；
2. 执行本测试方法，核对解密结果，并获得golang sm4加密密文，更新Hyperledger-TWGC/java-gm中的SM4