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
2. 执行本测试方法，核对解密结果，并获得golang sm4加密密文，更新Hyperledger-TWGC/java-gm中的SM4UtilTest.sm4Interaction4Decrypt()
	中待解密参数；
*/
func TestSM4InteractionWithJava(t *testing.T) {
	// 128比特密钥
	key := []byte("1234567890abcdef")
	fmt.Println("key = " + string(key))
	// 128比特iv
	//iv := make([]byte, sm4.BlockSize)
	iv := []byte("ilovegolangjava.")
	fmt.Println("iv = " + string(iv))

	// 加密明文以供Java sm4解密验证
	data := []byte("I am encrypted by golang SM4.")
	ciphertxt, err := sm4Encrypt(key, iv, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("加密结果: %x