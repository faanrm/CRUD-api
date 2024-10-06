package utils

import "os"

func setEnv() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "crud")
}
