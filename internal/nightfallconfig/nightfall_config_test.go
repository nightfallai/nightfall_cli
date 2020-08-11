package nightfallconfig_test

import (
	"os"
	"path"
	"testing"

	"github.com/nightfallai/jenkins_test/internal/nightfallconfig"
	nightfallAPI "github.com/nightfallai/nightfall_go_client/generated"
	"github.com/stretchr/testify/assert"
)

const testFileName = "nightfall_test_config.json"
const excludedCreditCardRegex = "4242-4242-4242-[0-9]{4}"
const excludedApiToken = "xG0Ct4Wsu3OTcJnE1dFLAQfRgL6b8tIv"
const excludedIPRegex = "^127\\."

func TestGetNightfallConfig(t *testing.T) {
	workspaceConfig, err := os.Getwd()
	assert.NoError(t, err, "Unexpected error when getting current directory")
	workspacePath := path.Join(workspaceConfig, "../../test/data")
	expectedConfig := &nightfallconfig.NightfallConfigFileStructure{
		Detectors: nightfallconfig.DetectorConfig{
			nightfallAPI.CREDIT_CARD_NUMBER: nightfallAPI.POSSIBLE,
			nightfallAPI.PHONE_NUMBER:       nightfallAPI.LIKELY,
			nightfallAPI.IP_ADDRESS:         nightfallAPI.POSSIBLE,
		},
		MaxNumberRoutines:  20,
		TokenExclusionList: []string{excludedCreditCardRegex, excludedApiToken, excludedIPRegex},
		FileInclusionList:  []string{".*"},
		FileExclusionList:  nil,
	}
	actualConfig, err := nightfallconfig.GetConfigFile(workspacePath, testFileName)
	assert.NoError(t, err, "Unexpected error when GetNightfallConfig")
	assert.Equal(t, expectedConfig, actualConfig, "Incorrect nightfall config")
}
