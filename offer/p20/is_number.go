package main

// 有限状态机

type state int

/*
有限状态：
起始
空格
符号
整数部分的数字
有整数部分的小数点
无整数部分的小数点
小数部分的数字
e|E
指数符号
指数部分的数字
空格
结束
*/
const (
	stateInit state = iota
	stateSpaceLeading
	stateSign
	stateInt
	statePoint
	statePointWithoutInt
	stateFraction
	stateExp
	stateExpSign
	stateExpNumber
	stateSpaceFollowing
)

type charType int

/*
所能遇到的字符：
空格
+-
数字0-9
小数点
e|E
其他非法字符
*/
const (
	charSpace charType = iota
	charSign
	charNumber
	charPoint
	charExp
	charIllegal
)

var transfer []map[charType]state

func init() {
	transfer = make([]map[charType]state, stateSpaceFollowing+1)

	transfer[stateInit] = map[charType]state{
		charSpace:  stateSpaceLeading,
		charSign:   stateSign,
		charNumber: stateInt,
		charPoint:  statePointWithoutInt,
	}
	transfer[stateSpaceLeading] = map[charType]state{
		charSpace:  stateSpaceLeading,
		charSign:   stateSign,
		charNumber: stateInt,
		charPoint:  statePointWithoutInt,
	}
	transfer[stateSign] = map[charType]state{
		charNumber: stateInt,
		charPoint:  statePointWithoutInt,
	}
	transfer[stateInt] = map[charType]state{
		charNumber: stateInt,
		charPoint:  statePoint,
		charExp:    stateExp,
		charSpace:  stateSpaceFollowing,
	}
	transfer[statePoint] = map[charType]state{
		charNumber: stateFraction,
		charExp:    stateExp,
		charSpace:  stateSpaceFollowing,
	}
	transfer[statePointWithoutInt] = map[charType]state{
		charNumber: stateFraction, // 没有整数的小数点后面只能跟小数
	}
	transfer[stateFraction] = map[charType]state{
		charNumber: stateFraction,
		charExp:    stateExp,
		charSpace:  stateSpaceFollowing,
	}
	transfer[stateExp] = map[charType]state{
		charSign:   stateExpSign,
		charNumber: stateExpNumber,
	}
	transfer[stateExpSign] = map[charType]state{
		charNumber: stateExpNumber,
	}
	transfer[stateExpNumber] = map[charType]state{
		charNumber: stateExpNumber,
		charSpace:  stateSpaceFollowing,
	}
	transfer[stateSpaceFollowing] = map[charType]state{
		charSpace: stateSpaceFollowing,
	}
}

func completeState(s state) bool {
	switch s {
	case stateSpaceFollowing, stateInt, statePoint, stateFraction, stateExpNumber:
		return true
	default:
		return false
	}
}

func isNumber(s string) bool {
	currentState := stateInit

	var t charType
	for i := range s {
		t = typeOfChar(s[i])
		nextState, ok := transfer[currentState][t]
		if !ok {
			return false
		}
		currentState = nextState
	}

	return completeState(currentState)
}

func typeOfChar(c byte) charType {
	switch c {
	case ' ':
		return charSpace
	case '+', '-':
		return charSign
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return charNumber
	case '.':
		return charPoint
	case 'e', 'E':
		return charExp
	default:
		return charIllegal
	}
}
