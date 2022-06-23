package main

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainNoFlags(t *testing.T) {
	cmd := exec.Command("./go_league_cli")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	assert.NotNil(t, err, "Error should not be nil if we don't pass flags")
}

func TestMain(t *testing.T) {
	cmd := exec.Command("./go_league_cli", "-file", "test_score.txt")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if assert.Nil(t, err) {

		resultStr := out.String()
		assert.Contains(t, resultStr, "1. Lions, 6 pts")
		assert.Contains(t, resultStr, "2. Tarantulas, 6 pts")
		assert.Contains(t, resultStr, "3. FC Awesome, 1 pts")
		assert.Contains(t, resultStr, "4. Grouches, 1 pts")
		assert.Contains(t, resultStr, "5. Snakes, 1 pts")

	}
}
