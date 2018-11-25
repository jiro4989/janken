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
		return "あなたの勝ち組です"
	case STATUS_LOSE:
		return "あなたは負け犬です"
	case STATUS_SAME:
		return "あなたは勝っても負けてもいません"
	}
	return "不正な入力は迷惑です"
}
