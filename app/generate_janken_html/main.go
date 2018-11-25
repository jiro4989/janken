package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

// 環境ごとに異なる
var (
	// htmlのテンプレートのパス
	TEMPLATE_PATH = "../template/janken.html"
	// HTMLの出力先
	DEST_DIR = "./html"
)

func main() {
	if TEMPLATE_PATH == "" || DEST_DIR == "" {
		panic("環境未定義")
	}

	args := os.Args
	if isIllegalArgs(args) {
		panic("引数入れて")
		os.Exit(1)
	}

	userHand := args[1]
	enemyHand := args[2]

	status := winHand(userHand, enemyHand)
	html, err := generateHTML(userHand, enemyHand, status, TEMPLATE_PATH)
	if err != nil {
		panic(err)
	}

	fileName := userHand + ".html"
	dstFile := fmt.Sprintf("%s/%s", DEST_DIR, fileName)
	tmpFile := dstFile + ".tmp"

	// テンポラリを作成して移動
	if err := ioutil.WriteFile(tmpFile, []byte(html), 0644); err != nil {
		panic(err)
	}
	if err := os.Rename(tmpFile, dstFile); err != nil {
		panic(err)
	}
}

func isIllegalArgs(args []string) bool {
	if len(args) < 3 {
		return true
	}
	if !isHand(args[1]) {
		return true
	}
	if !isHand(args[2]) {
		return true
	}
	return false
}

func generateHTML(userHand, enemyHand string, status BattleStatus, templatePath string) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buff := new(bytes.Buffer)
	w := io.Writer(buff)
	m := map[string]string{
		"yourHand":  userHand,
		"enemyHand": enemyHand,
		"judge":     judgeText(status),
	}
	if err := tmpl.Execute(w, m); err != nil {
		return "", err
	}
	return buff.String(), nil
}
