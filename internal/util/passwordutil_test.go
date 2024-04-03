package util

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {

	mytest := "password"
	encrypt, err := Encrypt(mytest, "abc&1*~#^2^#s0^=)^^7%b34")
	if err != nil {
		t.Fatal(err)
	}

	res, _ := Decrypt(encrypt, "abc&1*~#^2^#s0^=)^^7%b34")

	if mytest != res {
		fmt.Println(fmt.Sprintf("Problem with decryption: %s != %s", mytest, res))
		t.Fail()
	}

}
