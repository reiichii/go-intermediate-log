package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func readFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(bytes), err
}

func main_() {
	idToken, err := readFile("./idtoken.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	dataArray := strings.Split(idToken, ".")
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("errors:", err)
		return
	}

	// fmt.Println("header: ", string(headerData))
	// fmt.Println("payload: ", string(payloadData))

	// 公開鍵構造体を作る
	// curl https://www.googleapis.com/oauth2/v3/certs で取得し、 kidが一致する方の値を取得
	N := "or83anRxFNTbjOy47m4SRDZQ7WpX_yjJdqN_LgNUBfbb_VnBwIUv_k4E1tXOE1yQC704YAT6JQ4AJtvLw598NxSuyXSvo-JCQ4pNugjVZ0w2MErJtARcxCu4LI6gsA_xSfSfuNVVSdrHqg8G-wsog0BS6N4M5IJtUlRR6UtjLaJxgqFGzV5sHWAfmpBekqCC5l19OXtE9J00r_Wjo4kfleonpVlEHszx5KUzShfGTGwgoeryNcp4yBULh8El8vt50a4SP_D74gCL5YINUl4E8hfQoqbPoxLj33oXYEvMKL34xYErEF5Tw39oAEfky3OgTXsCQvAp5il7HQjRY1JGow"
	E := "AQAB"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)
	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	message := sha256.Sum256([]byte(header + "." + payload))
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
	}
}
