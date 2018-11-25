package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHand(t *testing.T) {
	assert.Equal(t, true, isHand("rock"))
	assert.Equal(t, true, isHand("paper"))
	assert.Equal(t, true, isHand("scissors"))
	assert.Equal(t, false, isHand("ROCK"))
	assert.Equal(t, false, isHand("gu"))
	assert.Equal(t, false, isHand(""))
}

func TestWinHand(t *testing.T) {
	assert.Equal(t, STATUS_SAME, winHand("rock", "rock"))
	assert.Equal(t, STATUS_LOSE, winHand("rock", "paper"))
	assert.Equal(t, STATUS_WIN, winHand("rock", "scissors"))
	assert.Equal(t, STATUS_WIN, winHand("paper", "rock"))
	assert.Equal(t, STATUS_SAME, winHand("paper", "paper"))
	assert.Equal(t, STATUS_LOSE, winHand("paper", "scissors"))
	assert.Equal(t, STATUS_LOSE, winHand("scissors", "rock"))
	assert.Equal(t, STATUS_WIN, winHand("scissors", "paper"))
	assert.Equal(t, STATUS_SAME, winHand("scissors", "scissors"))
	assert.Equal(t, STATUS_ILLEGAL, winHand("scissors", ""))
}
