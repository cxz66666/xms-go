package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	plaintext:=[]byte("cxz666")
	AesKey:= []byte("12RTmCGWwr3LVYbxqUs1!")
	AesIv:= []byte("8a7oawe32CSIGEIuivbvQcvR0")

	//hex is 87adb96ee9c5ee609e7ed10eb47f0597
	encryptText,err:=AesCBCEncrypt(plaintext,AesKey,AesIv)
	if err!=nil{
		t.Fatalf("encrypt error %v",err)
	}
	fmt.Println(hex.EncodeToString(encryptText))

	p,err:=AesCBCDecrypt(encryptText,AesKey,AesIv)
	if err!=nil{
		t.Fatalf("decrypt error %v",err)

	}

	if !bytes.Equal(plaintext,p) {
		t.Fatal("not equal")
	}
}
