package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEnvFromFile(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	if noError := assert.NoError(t, err, "No error whlie the environment"); noError {
		assert.NotNil(t, env, "Environment has been loaded")
	}
}

func TestCreateEnvFromFileFailure(t *testing.T) {
	_, err := CreateEnvFromFile("bad_file_name")
	assert.Error(t, err, "Error whlie the environment")
}

func TestGet(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	assert.NoError(t, err, "No error whlie the environment")
	assert.NotNil(t, env, "Environment has been loaded")

	assert.NotEqual(t, env.Get("APP_DEBUG_LEVEL"), "", "Environment variables are loaded")
}

func TestSet(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	assert.NoError(t, err, "No error whlie the environment")
	assert.NotNil(t, env, "Environment has been loaded")

	env.Set("TEST_VAR", "TEST_VAR_VALUES")
	defer os.Unsetenv("TEST_VAR") // cleanup

	assert.NotEqual(t, env.Get("TEST_VAR"), "", "Environment variables are loaded")
}

func TestGetMissing(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	assert.NoError(t, err, "No error whlie the environment")
	assert.NotNil(t, env, "Environment has been loaded")

	env.Set("TEST_VAR_1", "TEST_VAR_VALUES_1")
	defer os.Unsetenv("TEST_VAR_1") // cleanup

	missing := env.GetMissing([]string{"TEST_VAR_1", "TEST_VAR_2"})
	assert.Len(t, missing, 1, "Found missing environment variables")
	assert.Equal(t, missing[0], "TEST_VAR_2", "Correct missing var is found")
}

func TestCheckEnv(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	assert.NoError(t, err, "No error whlie the environment")
	assert.NotNil(t, env, "Environment has been loaded")

	env.Set("TEST_VAR_1", "TEST_VAR_VALUES_1")
	defer os.Unsetenv("TEST_VAR_1") //cleanup

	env.Set("TEST_VAR_2", "TEST_VAR_VALUES_2")
	defer os.Unsetenv("TEST_VAR_2") // cleanup
	err = env.CheckEnv([]string{"TEST_VAR_1", "TEST_VAR_2"})
	assert.NoError(t, err, "Found no missing environment variables")
}

func TestCheckEnvError(t *testing.T) {
	env, err := CreateEnvFromFile("../../.env")
	assert.NoError(t, err, "No error whlie the environment")
	assert.NotNil(t, env, "Environment has been loaded")

	env.Set("TEST_VAR_1", "TEST_VAR_VALUES_1")
	defer os.Unsetenv("TEST_VAR_1") // cleanup

	err = env.CheckEnv([]string{"TEST_VAR_1", "TEST_VAR_2"})
	assert.Error(t, err, "Found missing environment variables")
}

// The following test is dummy. Because a working directory is different from root in 'go test' the function will
// not work here, but will work when a non-test go program is run. This test is only written to raise test coverage
func TestCreateEnv(t *testing.T) {
	_, err := CreateEnv()
	assert.Error(t, err, "Error whlie the environment")
}
