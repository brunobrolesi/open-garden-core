package service

import (
	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHashService struct {
	hashCost int
}

func NewBcryptHashService(hashCoast int) gateway.HashService {
	return bcryptHashService{
		hashCost: hashCoast,
	}
}

func (b bcryptHashService) GenerateHash(s string) (gateway.Hash, error) {
	bytes := []byte(s)
	hash, err := bcrypt.GenerateFromPassword(bytes, b.hashCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (bcryptHashService) CompareStringAndHash(s, hash string) error {
	stringBytes := []byte(s)
	hashBytes := []byte(hash)
	return bcrypt.CompareHashAndPassword(hashBytes, stringBytes)
}
