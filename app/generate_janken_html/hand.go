package main

const (
	HAND_ROCK     = "rock"
	HAND_PAPER    = "paper"
	HAND_SCISSORS = "scissors"
)

type BattleStatus int

const (
	STATUS_WIN BattleStatus = iota
	STATUS_LOSE
	STATUS_SAME    // あいこ
	STATUS_ILLEGAL // 不正
	WIN_TEXT       = "あなたの勝ち組です"
	LOSE_TEXT      = "あなたは負け犬です"
	SAME_TEXT      = "あなたは勝っても負けてもいません"
	ILLEGAL_TEXT   = "不正な入力は迷惑です"
)

func isHand(hand string) bool {
	switch hand {
	case HAND_ROCK:
		return true
	case HAND_PAPER:
		return true
	case HAND_SCISSORS:
		return true
	}
	return false
}

func winHand(userHand, enemyHand string) BattleStatus {
	switch userHand {
	case HAND_ROCK:
		switch enemyHand {
		case HAND_ROCK:
			return STATUS_SAME
		case HAND_PAPER:
			return STATUS_LOSE
		case HAND_SCISSORS:
			return STATUS_WIN
		}
	case HAND_PAPER:
		switch enemyHand {
		case HAND_ROCK:
			return STATUS_WIN
		case HAND_PAPER:
			return STATUS_SAME
		case HAND_SCISSORS:
			return STATUS_LOSE
		}
	case HAND_SCISSORS:
		switch enemyHand {
		case HAND_ROCK:
			return STATUS_LOSE
		case HAND_PAPER:
			return STATUS_WIN
		case HAND_SCISSORS:
			return STATUS_SAME
		}
	}
	return STATUS_ILLEGAL
}

func judgeText(status BattleStatus) string {
	switch status {
	case STATUS_WIN:
		return WIN_TEXT
	case STATUS_LOSE:
		return LOSE_TEXT
	case STATUS_SAME:
		return SAME_TEXT
	}
	return ILLEGAL_TEXT
}
