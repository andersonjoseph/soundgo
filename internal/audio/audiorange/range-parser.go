package audiorange

import (
	"fmt"
	"strconv"
)

type Range struct {
	Unit  string
	Start int64
	End   int64
}

type Parser struct {
	lexer        Lexer
	ranges       []Range
	currentIndex uint
}

func (p *Parser) Parse(i string) ([]Range, error) {
	lexer := NewLexer(i)
	p.ranges = make([]Range, 0)

	tok := lexer.NextToken()
	if tok.Type != alpha {
		return nil, fmt.Errorf("range should start with a unit")
	}
	unit := tok.Literal

	tok = lexer.NextToken()
	if tok.Type != equal {
		return nil, fmt.Errorf("expected =, received: %s", tok.Literal)
	}

	for tok = lexer.NextToken(); tok.Type != end; tok = lexer.NextToken() {
		startRange, err := parseStart(tok)
		if err != nil {
			return nil, err
		}

		if tok.Type != separator {
			if tok = lexer.NextToken(); tok.Type != separator {
				return nil, fmt.Errorf("expected separator -, received: %s", tok.Literal)
			}
		}

		tok = lexer.NextToken()
		endRange, err := parseEnd(tok)

		if endRange != -1 && startRange >= endRange {
			return nil, fmt.Errorf("start should be smaller than end range")
		}

		p.ranges = append(p.ranges, Range{
			Unit:  unit,
			Start: startRange,
			End:   endRange,
		})

		tok = lexer.NextToken()

		if tok.Type != comma && tok.Type != end {
			return nil, fmt.Errorf("expected end of input or comma (,) received: %s", tok.Literal)
		}
	}

	if len(p.ranges) == 0 {
		return nil, fmt.Errorf("range should have at least one value")
	}

	return p.ranges, nil
}

func parseNumber(t Token) (int64, error) {
	if t.Type != number {
		return 0, fmt.Errorf("invalid start range, received: %s", t.Literal)
	}

	v, err := strconv.ParseInt(t.Literal, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("error parsing range number: %s: %w", t.Literal, err)
	}
	return v, nil
}

func parseStart(t Token) (int64, error) {
	if t.Type == separator {
		return -1, nil
	}

	return parseNumber(t)
}

func parseEnd(t Token) (int64, error) {
	if t.Type == end {
		return -1, nil
	}

	return parseNumber(t)
}
