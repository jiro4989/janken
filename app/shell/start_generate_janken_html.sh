#!/bin/bash

start_time=$(date +%N)

source /opt/janken/shell/util.sh

# ログプライオリティ、など引数によって変数を決定する初期化処理
user_hand=$1
init "$user_hand"
ret=$?
if [ "$ret" -ne 0 ]; then
  err "引数が不正 user_hand=$user_hand"
  exit 1
fi

info "HTML生成 開始 user_hand=$user_hand"

if [ -e "$LOCK_FILE_STOP" ]; then
  info "${LOCK_FILE_STOP}が存在するためスキップ"
  exit 0
fi

trap 'rmdir "$LOCK_FILE"' ERR

mkdir "$LOCK_FILE"
ret=$?
if [ "$ret" -ne 0 ]; then
  info "${LOCK_FILE}が存在するためスキップ"
  exit 0
fi

# アプリバイナリ起動 (メイン処理)
enemy_hand=$(shuf -re rock paper scissors | head -n 1)
/opt/janken/bin/generate_janken_html.sh "$user_hand" "$enemy_hand"
ret=$?
if [ "$ret" -ne 0 ]; then
  err "HTML生成に失敗 exit=$ret"
  exit 1
fi

end_time=$(date +%N)
diff_time=$(awk 'BEGIN{print '$SECONDS' + (('$end_time' - '$start_time') / 1000 / 1000 / 1000)}' /dev/zero)
echo "HTML生成 正常終了 全体処理時間=${diff_time}ms"
