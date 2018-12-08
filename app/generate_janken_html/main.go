package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log/syslog"
	"os"
	"time"
)

var logger *syslog.Writer

// 環境ごとに異なる
var (
	// htmlのテンプレートのパス
	TEMPLATE_PATH = "../template/janken.html"
	// HTMLの出力先
	DEST_DIR = "./html"
)

func init() {
	var err error
	// 引数不正の場合にもログを出力したいので、初期値rockとしてロガーをセット
	logger, err = syslog.New(syslog.LOG_LOCAL0, "")
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	logger.Info(fmt.Sprintf("HTML生成 開始 args=%v", args))
	start := time.Now()

	if isIllegalArgs(args) {
		msg := fmt.Sprintf("引数が不正 args=%v", args)
		logger.Err(msg)
		os.Exit(1)
	}

	userHand := args[1]
	enemyHand := args[2]

	if err := setLogger(userHand); err != nil {
		logger.Err(err.Error())
		os.Exit(1)
	}

	status := winHand(userHand, enemyHand)
	html, err := generateHTML(userHand, enemyHand, status, TEMPLATE_PATH)
	if err != nil {
		msg := fmt.Sprintf("HTMLの生成に失敗 userHand=%s enemyHand=%s status=%d TEMPLATE_PATH=%s err=%v", userHand, enemyHand, status, TEMPLATE_PATH, err)
		logger.Err(msg)
		os.Exit(2)
	}

	fileName := userHand + ".html"
	dstFile := fmt.Sprintf("%s/%s", DEST_DIR, fileName)
	tmpFile := dstFile + ".tmp"

	// テンポラリを作成して移動
	if err := ioutil.WriteFile(tmpFile, []byte(html), 0644); err != nil {
		msg := fmt.Sprintf("%sの生成に失敗 err=%v", tmpFile, err)
		logger.Err(msg)
		os.Exit(3)
	}
	if err := os.Rename(tmpFile, dstFile); err != nil {
		msg := fmt.Sprintf("%sを%sに移動するのに失敗 err=%v", tmpFile, dstFile, err)
		logger.Err(msg)
		os.Exit(4)
	}

	end := time.Now()
	diff := end.Sub(start)
	logger.Info(fmt.Sprintf("HTML生成 正常終了 処理時間=%fs", diff.Seconds()))
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

func setLogger(userHand string) error {
	var err error
	switch userHand {
	case HAND_ROCK:
		logger, err = syslog.New(syslog.LOG_LOCAL0, "")
	case HAND_PAPER:
		logger, err = syslog.New(syslog.LOG_LOCAL1, "")
	case HAND_SCISSORS:
		logger, err = syslog.New(syslog.LOG_LOCAL2, "")
	default:
		err = errors.New(fmt.Sprintf("userHandが不正 err=%s", userHand))
	}
	return err
}
