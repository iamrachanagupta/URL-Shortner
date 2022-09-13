package test

import (
	"os"
	"testing"
	"urlapp/internal/utilities"

	"github.com/stretchr/testify/assert"
)

func TestGenerateURL(t *testing.T) {
	utilities.ReadTestConfig()
	testLink := "https://github.com/iamrachanagupta/Go"
	generatedUrl := utilities.GenerateURL(testLink)
	expectedUrl := utilities.Data["appconfig"].AppURL + "shorturl/MTM2MTYx"

	assert.Equal(t, expectedUrl, generatedUrl)
}

func TestGetUrlList(t *testing.T) {
	utilities.GetUrlList()
}

func TestReadConfig_PanicTest(t *testing.T) {
	pwd, _ := os.Getwd()
	os.Chdir("../../..")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function should panic")
		}
	}()

	utilities.ReadConfig()
	os.Chdir(pwd)
}

func TestReadTestConfig(t *testing.T) {
	utilities.ReadTestConfig()
	generatedUrl := utilities.Data["appconfig"].AppURL
	expectedUrl := "http://localhost:8081/"
	assert.Equal(t, expectedUrl, generatedUrl)
}
