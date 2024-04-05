package expr

import "io"

type tokenKind int

const (
	tokenUndef tokenKind = iota
	// tokenLiteral 字面量，e.g. 50
	tokenLiteral
	// tokenOperator 操作符, e.g. + - * /
	tokenOperator
	tokenLParen
	tokenRParen // ）
)

type token struct {
	byteValue  []byte // Raw value of a token.
	delimValue byte
	// 类型，有 Literal、Operator 两种
	kind tokenKind
}

func (token token) String() string {
	switch token.kind {
	case tokenLParen:
		return "tokenLParen: " + string(token.delimValue)
	case tokenRParen:
		return "tokenRParen: " + string(token.delimValue)
	case tokenLiteral:
		return "tokenLiteral: " + string(token.byteValue)
	case tokenOperator:
		return "tokenOperator: " + string(token.delimValue)
	default:
		return "tokenUndef: " + string(token.byteValue)
	}
}

type Lexer struct {
	expression []byte
	start      int
	pos        int

	token      token
	fatalError error
}

func Parse(expression string) ([]token, error) {
	l := &Lexer{
		expression: []byte(expression),
	}
	tokens, err := l.parse()
	if err != nil && err != io.EOF {
		return nil, err
	}
	return tokens, nil
}

// parse 逐个字符扫描，得到一串 Token 序列
func (l *Lexer) parse() (tokens []token, err error) {
	for l.Ok() {
		l.scanToken()
		if l.Ok() {
			tokens = append(tokens, l.token)
			l.consume()
		}
	}
	return tokens, l.fatalError
}

// scanToken scans the next token if no token is currently available in the lexer.
func (l *Lexer) scanToken() {
	if l.token.kind != tokenUndef || l.fatalError != nil {
		return
	}
	l.fetchToken()
}

// consume resets the current token to allow scanning the next one.
func (l *Lexer) consume() {
	l.token.kind = tokenUndef
	l.token.delimValue = 0
	l.token.byteValue = l.token.byteValue[:0]
}

// Ok returns true if no error (including io.EOF) was encountered during scanning.
func (l *Lexer) Ok() bool {
	return l.fatalError == nil
}

func (l *Lexer) fetchToken() {
	l.token.kind = tokenUndef
	l.start = l.pos

	// Check if r.Data has r.pos element
	// If it doesn't, it means corrupted input data
	if len(l.expression) < l.pos {
		l.errParse("Unexpected end of data")
		return
	}
	for _, c := range l.expression[l.pos:] {
		switch c {
		case '(':
			l.token.kind = tokenLParen
			l.token.delimValue = c
			l.pos++
			return
		case ')':
			l.token.kind = tokenRParen
			l.token.delimValue = c
			l.pos++
			return
		case '+', '-', '*', '/':
			l.token.kind = tokenOperator
			l.token.delimValue = c
			l.pos++
			return
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.token.kind = tokenLiteral
			l.fetchNumber()
			return
		default:
			l.errSyntax()
			return
		}
	}
	l.fatalError = io.EOF
	return
}

const maxErrorContextLen = 13

func (l *Lexer) errParse(what string) {
	if l.fatalError == nil {
		var str string
		if len(l.expression)-l.pos <= maxErrorContextLen {
			str = string(l.expression)
		} else {
			str = string(l.expression[l.pos:l.pos+maxErrorContextLen-3]) + "..."
		}
		l.fatalError = &LexerError{
			Reason: what,
			Offset: l.pos,
			Data:   str,
		}
	}
}

func (l *Lexer) fetchNumber() {
	hasE := false
	afterE := false
	hasDot := false

	l.pos++
	for i, c := range l.expression[l.pos:] {
		switch {
		case c >= '0' && c <= '9':
			afterE = false
		case c == '.' && !hasDot:
			hasDot = true
		case (c == 'e' || c == 'E') && !hasE:
			hasE = true
			hasDot = true
			afterE = true
		case (c == '+' || c == '-') && afterE:
			afterE = false
		default:
			l.pos += i
			l.token.byteValue = l.expression[l.start:l.pos]
			return
		}
	}

	l.pos = len(l.expression)
	l.token.byteValue = l.expression[l.start:]
}

func (l *Lexer) errSyntax() {
	l.errParse("syntax error")
}
