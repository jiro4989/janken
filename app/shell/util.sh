#!/bin/bash

LOCK_FILE_ROCK=/var/run/janken/rock.lock
LOCK_FILE_PAPER=/var/run/janken/paper.lock
LOCK_FILE_SCISSORS=/var/run/janken/scissors.lock
LOCK_FILE_STOP=/var/run/janken/stop.lock

PRIORITY_ROCK=local0
PRIORITY_PAPER=local1
PRIORITY_SCISSORS=local2

# アプリ実行時の引数で上書きされる
PRIORITY=
LOCK_FILE=

info() {
  logger -p "$PRIORITY.info" "$@"
}

warning() {
  logger -p "$PRIORITY.warning" "$@"
}

err() {
  logger -p "$PRIORITY.err" "$@"
}

init() {
  hand=$1
  case "$hand" in
    "rock")
      PRIORITY=$PRIORITY_ROCK
      LOCK_FILE=$LOCK_FILE_ROCK
      return 0
      ;;
    "paper")
      PRIORITY=$PRIORITY_PAPER
      LOCK_FILE=$LOCK_FILE_PAPER
      return 0
      ;;
    "scissors")
      PRIORITY=$PRIORITY_SCISSORS
      LOCK_FILE=$LOCK_FILE_SCISSORS
      return 0
      ;;
  esac
  return 1
}
