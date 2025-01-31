package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func GenUUID() int64 {
	timestamp := time.Now().Unix() % 1e9
	randomBits, _ := rand.Int(rand.Reader, big.NewInt(1e2)) // Random 2 chữ số (0-99)

	return timestamp*100 + randomBits.Int64()
}

func GenTime() time.Time {
	return time.Now().UTC()
}

func GenPassWord() int64 {
	max := big.NewInt(99999999 - 10000000 + 1)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}

	return n.Int64() + 10000000
}

func GenPasswordString(length int) string {
	if length < 8 {
		length = 8
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		password[i] = charset[index.Int64()]
	}

	return string(password)
}

func GenOTP(length int) int64 {
	if length < 4 {
		length = 4
	} else if length > 10 {
		length = 10
	}

	max := big.NewInt(1)
	for i := 0; i < length; i++ {
		max.Mul(max, big.NewInt(10))
	}

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}

	return n.Int64()
}

func ValidatePassword(password string) bool {
	var (
		hasUpper  bool
		hasLower  bool
		hasNumber bool
		hasSymbol bool
	)

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char >= '!' && char <= '/':
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func GenerateConfigFile() {
	content := `
	{
    "data_source": "host=localhost user=postgres password=1234 dbname=rices port=5432 sslmode=disable TimeZone=Asia/Shanghai",
    "port": "8080",
    "access_secret": "secretAAAAAAaal;kjmnopiaassssdsv",
    "expire_access": "24h",
    "address_redis":"127.0.0.1:6379",
    "password_redis":"",
    "database_redis_index":0,
    "key_aes":"y-in-y-srkss-u-dgr-y1ie32ncelv-ohee-aare-tv",
    "email":"tranhuythang9999@gmail.com",
    "app_key":"jkqr axuy tjie ziyl",
    "smtp_host":"smtp.gmail.com",
    "smtp_port":"587"
	}
	`

	if err := os.MkdirAll("configs", os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err := os.OpenFile("configs/configs.json", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("File already exists, skipping creation.")
		} else {
			fmt.Println("Error opening file:", err)
		}
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func FormatTime(req time.Time) string {
	if req.IsZero() {
		return "Invalid time"
	}

	return req.Format("2006-01-02 15:04:05")
}
