package test_helpers

import (
	"math/big"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/mcd_transformers/transformers/storage"
)

type MockMakerStorageRepository struct {
	DaiKeys          []string
	GemKeys          []storage.Urn
	GetDaiKeysCalled bool
	GetDaiKeysError  error
	GetGemKeysCalled bool
	GetGemKeysError  error
	GetIlksCalled    bool
	GetIlksError     error
	GetMaxFlipCalled bool
	GetMaxFlipError  error
	GetSinKeysCalled bool
	GetSinKeysError  error
	GetUrnsCalled    bool
	GetUrnsError     error
	Ilks             []string
	MaxFlip          *big.Int
	SinKeys          []string
	Urns             []storage.Urn
}

func (repository *MockMakerStorageRepository) GetDaiKeys() ([]string, error) {
	repository.GetDaiKeysCalled = true
	return repository.DaiKeys, repository.GetDaiKeysError
}

func (repository *MockMakerStorageRepository) GetGemKeys() ([]storage.Urn, error) {
	repository.GetGemKeysCalled = true
	return repository.GemKeys, repository.GetGemKeysError
}

func (repository *MockMakerStorageRepository) GetIlks() ([]string, error) {
	repository.GetIlksCalled = true
	return repository.Ilks, repository.GetIlksError
}

func (repository *MockMakerStorageRepository) GetMaxFlip() (*big.Int, error) {
	repository.GetMaxFlipCalled = true
	return repository.MaxFlip, repository.GetMaxFlipError
}

func (repository *MockMakerStorageRepository) GetSinKeys() ([]string, error) {
	repository.GetSinKeysCalled = true
	return repository.SinKeys, repository.GetSinKeysError
}

func (repository *MockMakerStorageRepository) GetUrns() ([]storage.Urn, error) {
	repository.GetUrnsCalled = true
	return repository.Urns, repository.GetUrnsError
}

func (repository *MockMakerStorageRepository) SetDB(db *postgres.DB) {}