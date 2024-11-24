package main

import "math/rand"

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = base62Chars[rand.Intn(len(base62Chars))]
	}
	return string(code)
}