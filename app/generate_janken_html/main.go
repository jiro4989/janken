package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
)

// 環境ごとに異なる
var TEMPLATE_PATH string

func main() {
	if TEMPLATE_PATH == "" {
		panic("環境未定義")
	}

	args := os.Args
	if isIllegalArgs(args) {
		panic("引数入れて")
		os.Exit(1)
	}

	userHand := args[1]
	enemyHand := args[2]
	fmt.Println(userHand, enemyHand)

	status := winHand(userHand, enemyHand)
	html, err := generateHTML(userHand, enemyHand, status, TEMPLATE_PATH)
	if err != nil {
		panic(err)
	}
	fmt.Println(html)
	// html := generateHTML(userHand, enemyHand, judge)
	// craeteTmpFile
	// mv tmp to location
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
