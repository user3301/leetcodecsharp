package tests

import (
	"gotest.tools/assert"
	"github.com/leetcode/questions"
	"testing"
)

func TestValidPalindromeII1(t *testing.T) {
	input := "abca"
	actual := questions.ValidPalindromeII(input)
	assert.Equal(t, true,actual)
}

func TestValidPalindromeII2(t *testing.T) {
	input := "abc"
	actual := questions.ValidPalindromeII(input)
	assert.Equal(t, false,actual)
}

func TestValidPalindromeII3(t *testing.T) {
	input := "atbbga"
	actual := questions.ValidPalindromeII(input)
	assert.Equal(t, false,actual)
}