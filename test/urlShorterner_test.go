package test

import (
	"testing"

	"urlapp/internal/utilities"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	testLink_1 := "https://github.com/iamrachanagupta/Go"
	key := utilities.GenerateKey(testLink_1)

	assert.Equal(t, "MTM2MTYx", key)

}

func TestGenerateKeyWithEmptyString(t *testing.T) {
	testLink_1 := ""
	key := utilities.GenerateKey(testLink_1)

	assert.Equal(t, "MTE4NTk1", key)

}
