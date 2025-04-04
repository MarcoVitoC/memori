package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const TEST_ENV_KEY = "TEST_ENV_KEY"

func TestGetString(t *testing.T) {
	require.NoError(t, os.Setenv(TEST_ENV_KEY, "TEST_ENV_VALUE"))
	t.Cleanup(func() {
		os.Unsetenv(TEST_ENV_KEY)
	})

	t.Run("should get set env value", func(t *testing.T) {
		val := GetString(TEST_ENV_KEY, "FALLBACK")
		assert.Equal(t, "TEST_ENV_VALUE", val)
	})

	t.Run("should get env fallback", func(t *testing.T) {
		fallback := GetString("invalid_key", "FALLBACK")
		assert.Equal(t, "FALLBACK", fallback)
	})
}

func TestGetInt(t *testing.T) {
	require.NoError(t, os.Setenv(TEST_ENV_KEY, "10"))
	t.Cleanup(func() {
		os.Unsetenv(TEST_ENV_KEY)
	})

	t.Run("should get set env value", func(t *testing.T) {
		valAsInt := GetInt(TEST_ENV_KEY, -1)
		assert.Equal(t, 10, valAsInt)
	})

	t.Run("should get env fallback", func(t *testing.T) {
		fallback := GetInt("invalid_key", -1)
		assert.Equal(t, -1, fallback)
	})
}
