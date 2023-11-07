package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword 密码比对
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}
