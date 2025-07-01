package uuidgenerator

import "github.com/google/uuid"

type UuidGenerator interface {
	Generate() string
}

type uuidGenerator struct {
}

func NewUuidGenerator() UuidGenerator {
	return uuidGenerator{}
}

func (g uuidGenerator) Generate() string {
	return uuid.New().String()
}
