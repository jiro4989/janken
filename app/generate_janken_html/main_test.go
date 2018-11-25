package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIllegalArgs(t *testing.T) {
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "rock", "rock"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "rock", "paper"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "rock", "scissors"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "paper", "rock"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "paper", "paper"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "paper", "scissors"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "scissors", "rock"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "scissors", "paper"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "scissors", "scissors"}))
	assert.Equal(t, false, isIllegalArgs([]string{"bin", "scissors", "scissors", "xx"}))
	assert.Equal(t, true, isIllegalArgs([]string{"bin", "scissors"}))
	assert.Equal(t, true, isIllegalArgs([]string{"bin"}))
	assert.Equal(t, true, isIllegalArgs([]string{"bin", "scissors", "x"}))
	assert.Equal(t, true, isIllegalArgs([]string{"bin", "x", "rock"}))
	assert.Equal(t, true, isIllegalArgs([]string{"rock", "rock"}))
}

func TestGenerateHTML(t *testing.T) {
	expect := `<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8">
    <title>じゃんけん - rock</title>
  </head>
  <body>
    <h1>じゃんけん</h1>
    <hr>
    あなたの手：rock
    <br>
    相手の手：paper
    <br>
    勝敗：あなたは負け犬です
  </body>
</html>
`
	got, err := generateHTML("rock", "paper", STATUS_LOSE, "../template/janken.html")
	assert.Equal(t, expect, got, "正常系")
	assert.NoError(t, err)

	expect = `<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8">
    <title>じゃんけん - &lt;script&gt;</title>
  </head>
  <body>
    <h1>じゃんけん</h1>
    <hr>
    あなたの手：&lt;script&gt;
    <br>
    相手の手：paper
    <br>
    勝敗：あなたは負け犬です
  </body>
</html>
`
	got, err = generateHTML("<script>", "paper", STATUS_LOSE, "../template/janken.html")
	assert.Equal(t, expect, got, "タグ文字がエスケープされること")
	assert.NoError(t, err)

	got, err = generateHTML("rock", "paper", STATUS_LOSE, "")
	assert.Equal(t, "", got, "エラーが返ること")
	assert.Error(t, err)
}
