package main

type State int
type CharType int

const (
	START_SPACE State = iota
	STATE_SIGN
	STATE_POINT_WITHOUT_NUMBER
	STATE_INTEGER
	STATE_POINT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_INTEGER
	STATE_END_SPACE
)

const (
	SPACE CharType = iota
	POINT
	NUMBER
	EXP
	SIGN
	ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case ' ':
		return SPACE
	case '.':
		return POINT
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER
	case 'e', 'E':
		return EXP
	case '+', '-':
		return SIGN
	default:
		return ILLEGAL
	}
}

// 映射有 map[state, char]->state 和 map[state]->map[char]->state  后者更简洁好用
func isNumber(s string) bool {
	m := map[State]map[CharType]State{
		START_SPACE: map[CharType]State{
			SPACE:  START_SPACE,
			NUMBER: STATE_INTEGER,
			POINT:  STATE_POINT_WITHOUT_NUMBER,
			SIGN:   STATE_SIGN,
		},
		STATE_SIGN: {
			POINT:  STATE_POINT_WITHOUT_NUMBER,
			NUMBER: STATE_INTEGER,
		},
		STATE_INTEGER: {
			NUMBER: STATE_INTEGER,
			SPACE:  STATE_END_SPACE,
			POINT:  STATE_POINT,
			EXP:    STATE_EXP,
		},
		STATE_POINT_WITHOUT_NUMBER: {
			NUMBER: STATE_FRACTION,
		},
		STATE_POINT: {
			SPACE:  STATE_END_SPACE,
			EXP:    STATE_EXP,
			NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: {
			NUMBER: STATE_FRACTION,
			SPACE:  STATE_END_SPACE,
			EXP:    STATE_EXP,
		},
		STATE_EXP: {
			SIGN:   STATE_EXP_SIGN,
			NUMBER: STATE_EXP_INTEGER,
		},
		STATE_EXP_SIGN: {
			NUMBER: STATE_EXP_INTEGER,
		},
		STATE_EXP_INTEGER: {
			NUMBER: STATE_EXP_INTEGER,
			SPACE:  STATE_END_SPACE,
		},
		STATE_END_SPACE: {
			SPACE: STATE_END_SPACE,
		},
	}

	state := START_SPACE
	for i := 0; i < len(s); i++ {
		v := toCharType(s[i])
		if _, ok := m[state][v]; !ok {
			return false
		} else {
			state = m[state][v]
		}
	}
	return state == STATE_END_SPACE || state == STATE_EXP_INTEGER || state == STATE_FRACTION || state == STATE_POINT || state == STATE_INTEGER
}
