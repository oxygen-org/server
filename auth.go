/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:06:38
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)


const (
	SALT_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	DEFAULT_PBKDF2_ITERATIONS = 50000
)

// GeneratePasswordHash Gen pw-hash
func GeneratePasswordHash(password string, saltLen int)string{
	salt := 
	byteResult := pbkdf2.Key([]byte("some password"), []byte(salt), 4096, 32, sha256.New)
	return string(byteResult)
}

// Check password
func CheckPasswordHash(passwordHash,password string)bool{

}

func BasicAuth() {

}

func CustomBasicAuth(){

}


