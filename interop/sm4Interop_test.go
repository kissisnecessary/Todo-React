
package interop

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	ccs "github.com/Hyperledger-TWGC/ccs-gm/sm4"
	pku "github.com/Hyperledger-TWGC/pku-gm/gmssl"
	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm4"
)

func TestSM4(t *testing.T) {
	base_format := "2006-01-02 15:04:05"
	//random data
	time := time.Now()
	str_time := time.Format(base_format)
	msg := []byte(str_time)

	// hard code data for debug usage if needed
	// []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	// []byte("0123456789abcdef012345678")

	sourceDef := os.Getenv("SOURCE")
	targetDef := os.Getenv("TARGET")

	fmt.Println("source lib " + sourceDef)
	fmt.Println("target lib " + targetDef)
	fmt.Printf("msg = %x\n", msg)
	// generate key
	// hard code key for debug usage if needed
	key := []byte("1234567890abcdef")

	//keyLen, _ := pku.GetCipherKeyLength(pku.SMS4)
	//key, _ := pku.GenerateRandom(keyLen)
	fmt.Printf("key = %v\n", key)
	// ECB

	var ecbMsg []byte
	var ecbDec []byte
	var cbcMsg []byte
	var cbcDec []byte
	var err error
	iv := make([]byte, 16)
	// lib a encrypt
	switch sourceDef {
	case "TJ":
		{
			ecbMsg, err = tj.Sm4Ecb(key, msg, true)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
			cbcMsg, err = tj.Sm4Cbc(key, msg, true)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
		}
	case "CCS":
		{
			ecbMsg, err = ccs.Sm4Ecb(key, msg, ccs.ENC)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
			cbcMsg, err = ccs.Sm4Cbc(key, msg, ccs.ENC)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
		}
	case "PKU":
		{
			//ecbMsg
			inData := pkcs7Padding(msg)
			ecbMsg = make([]byte, len(inData))
			//loop for each 16 bytes of msg
			for i := 0; i < len(inData)/16; i++ {
				TempMsg, err := pku.CipherECBenc(inData[i*16:(i+1)*16], key)
				if err != nil {
					t.Errorf("sm4 enc error:%s", err)
					return
				}
				//append into ecbMsg
				copy(ecbMsg[i*16:i*16+16], TempMsg)
			}
			// cbcMsg
			encrypt, err := pku.NewCipherContext(pku.SMS4, key, iv, true)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
			ciphertext1, err := encrypt.Update(msg)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
			ciphertext2, err := encrypt.Final()
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
			cbcMsg = make([]byte, 0, len(ciphertext1)+len(ciphertext2))
			cbcMsg = append(cbcMsg, ciphertext1...)
			cbcMsg = append(cbcMsg, ciphertext2...)
		}
	default:
		t.Errorf("unsupported")
	}
	// lib b decrypt
	switch targetDef {
	case "TJ":
		{
			ecbDec, err = tj.Sm4Ecb(key, ecbMsg, false)
			if err != nil {
				t.Errorf("sm4 dec error:%s", err)
				return
			}
			cbcDec, err = tj.Sm4Cbc(key, cbcMsg, false)
			if err != nil {
				t.Errorf("sm4 enc error:%s", err)
				return
			}
		}
	case "CCS":
		{
			ecbDec, err = ccs.Sm4Ecb(key, ecbMsg, ccs.DEC)
			if err != nil {