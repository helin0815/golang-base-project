package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutomaticEnv(t *testing.T) {
	_ = os.Setenv("LI_ENV", "dev")
	automaticEnv()

	assert.Equal(t, "dev", Env)
}
