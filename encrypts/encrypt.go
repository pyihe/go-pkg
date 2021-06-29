package encrypts

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

//用于加密web的密码
func BcEncryptPass(plainPass string) (string, error) {
	var encryptPass string
	data, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return encryptPass, err
	}
	encryptPass = base64.StdEncoding.EncodeToString(data)
	return encryptPass, nil
}

//比较密码是否匹配
func BcComparePass(hashPass, plainPass string) error {
	hashBytes, err := base64.StdEncoding.DecodeString(hashPass)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hashBytes, []byte(plainPass))
}

//scrypt加密
func ScryptPass(plainPass, salt string) (string, error) {
	data, err := scrypt.Key([]byte(plainPass), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	encryptPass := base64.StdEncoding.EncodeToString(data)
	return encryptPass, nil
}
