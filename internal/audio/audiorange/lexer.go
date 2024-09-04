package audiorange

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	illegal   TokenType = "illegal"
	end       TokenType = "end"
	number    TokenType = "number"
	alpha     TokenType = "alpha"
	equal     TokenType = "equal"
	separator TokenType = "separator"
	comma     TokenType = "comma"
)

type Lexer struct {
	input        string
	currentIndex uint64
}

func NewLexer(i string) Lexer {
	return Lexer{
		input:        i + " ",
		currentIndex: 0,
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token
	if int(l.currentIndex+1) == len(l.input) {
		tok.Type = end
		return tok
	}

	switch l.input[l.currentIndex] {
	case '=':
		tok.Type = equal
		tok.Literal = "="

		l.currentIndex++
	case '-':
		tok.Type = separator
		tok.Literal = "-"

		l.currentIndex++
	case ',':
		tok.Type = comma
		tok.Literal = ","

		l.currentIndex++
	default:
		if isDigit(l.input[l.currentIndex]) {
			tok.Type = number
			tok.Literal = l.readNumber()
		} else if isAlpha(l.input[l.currentIndex]) {
			tok.Type = alpha
			tok.Literal = l.readAlpha()
		} else {
			tok.Type = illegal
		}
	}

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.input[l.currentIndex] == ' ' {
		if !l.hasNextChar() {
			break
		}
		l.currentIndex++
	}
}

func (l *Lexer) readNumber() string {
	pos := l.currentIndex
	for isDigit(l.input[l.currentIndex]) {
		if !l.hasNextChar() {
			break
		}
		l.currentIndex++
	}

	return l.input[pos:l.currentIndex]
}

func (l *Lexer) readAlpha() string {
	pos := l.currentIndex
	for isAlpha(l.input[l.currentIndex]) {
		if !l.hasNextChar() {
			break
		}
		l.currentIndex++
	}

	return l.input[pos:l.currentIndex]
}

func (l *Lexer) hasNextChar() bool {
	return int(l.currentIndex+1) < len(l.input)
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
