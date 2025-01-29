package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"ladipage_server/apis/entities"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const googlePublicKeyURL = "https://www.googleapis.com/oauth2/v3/certs"

func getGooglePublicKeys() (map[string]*rsa.PublicKey, error) {
	resp, err := http.Get(googlePublicKeyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var keyData struct {
		Keys []struct {
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.Unmarshal(body, &keyData); err != nil {
		return nil, err
	}

	keys := make(map[string]*rsa.PublicKey)
	for _, key := range keyData.Keys {
		n, err := base64.RawURLEncoding.DecodeString(key.N)
		if err != nil {
			return nil, err
		}

		e, err := base64.RawURLEncoding.DecodeString(key.E)
		if err != nil {
			return nil, err
		}

		keys[key.Kid] = &rsa.PublicKey{
			N: new(big.Int).SetBytes(n),
			E: int(new(big.Int).SetBytes(e).Int64()),
		}
	}

	return keys, nil
}

func VerifyGoogleToken(idToken string) (*entities.GoogleClaims, error) {
	publicKeys, err := getGooglePublicKeys()
	if err != nil {
		return nil, fmt.Errorf("error getting Google public keys: %v", err)
	}

	token, err := jwt.ParseWithClaims(idToken, &entities.GoogleClaims{}, func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}

		key, ok := publicKeys[kid]
		if !ok {
			return nil, errors.New("key not found for the given kid")
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*entities.GoogleClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !strings.HasPrefix(claims.Iss, "https://accounts.google.com") {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}
