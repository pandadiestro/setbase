// ref: https://codereview.stackexchange.com/a/145608
// this is my first ever attempt at actually parsing,
// tokenizing and using an ast

package parser

import (
	"errors"
)

type Expr interface {
    Eval([]rune) bool
}

type Variable struct {
    Expr
    Name rune
    Value bool
}

type Constant struct {
    Expr
    Value bool
}

type UnaryOp struct {
    Expr
    Operator rune
    Operand Expr
}

type BinaryOp struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

type Token struct {
    Value rune
    Next *Token
}

const TOKEN_END rune = 0x00

func iterate(str string) (*Token, error) {
    var tokens *Token
    var head *Token

    for _, ch := range str {
        if tokens == nil {
            tokens = &Token{
                Value: ch,
                Next: nil,
            }

            head = tokens
            continue
        }

        tokens.Next = &Token{
            Value: ch,
            Next: nil,
        }

        tokens = tokens.Next
    }

    if tokens == nil {
        return tokens, errors.New("failed to tokenize input string(\"" + str + "\")")
    }

    tokens.Next = &Token{
        Value: TOKEN_END,
        Next: nil,
    }

    tokens = head
    return tokens, nil
}

var Variables = []rune {
    'a',
    'b',
    'c',
    'd',
    'e',
    'f',
    'g',
    'h',
    'i',
    'j',
    'k',
    'l',
    'm',
    'n',
    'o',
    'p',
    'q',
    'r',
    's',
    't',
    'u',
    'v',
    'w',
    'x',
    'y',
    'z',
    // uppercase
    'A',
    'B',
    'C',
    'D',
    'E',
    'F',
    'G',
    'H',
    'I',
    'J',
    'K',
    'L',
    'M',
    'N',
    'O',
    'P',
    'Q',
    'R',
    'S',
    'T',
    'U',
    'V',
    'W',
    'X',
    'Y',
    'Z',
}

var operators = []rune{
    '&',
    '|',
    '>',
}

var constants = []rune{
    '0',
    '1',
}

func Parse(str string) (Expr, error) {
    // this is the head of the recursive functions
    var disjunction func() Expr
    var conjunction func() Expr
    var implication func() Expr
    var unaryExpr func() Expr

    tokens, tokenErr := iterate(str)

    if tokenErr != nil {
        return &Proposition{}, tokenErr
    }

    match := func(compare []rune) bool {
        for _, el := range compare {
            if (*tokens).Value == el {
                tokens = (tokens).Next
                return true
            }
        }
        return false
    }

    term := func() Expr {
        current := (*tokens).Value

        if match(Variables) {
            return &Proposition{Name: current}
        } else if match(constants) {
            return &ConstantProposition{Name: current}
        } else if match([]rune{'('}) {
            tree := disjunction()
            if match([]rune{')'}) {
                return tree
            } else {
                return &Proposition{}
            }
        } else {
            return &Proposition{}
        }
    }

    binaryExpr := func(left Expr, operators []rune, right func() Expr) Expr {
        current := (*tokens).Value
        leftExpr := left

        if match(operators) {
            rightExpr := right()

            switch (current) {
            case '|':
                return &DisjunctionExpr{
                    Left: leftExpr,
                    Operator: current,
                    Right: rightExpr,
                }
            case '&':
                return &ConjunctionExpr{
                    Left: leftExpr,
                    Operator: current,
                    Right: rightExpr,
                }
            case '>':
                return &ImplicationExpr{
                    Left: leftExpr,
                    Operator: current,
                    Right: rightExpr,
                }
            default:
                return &BinaryOp{
                    Left: leftExpr,
                    Operator: current,
                    Right: rightExpr,
                }
            }

        } else {
            return leftExpr
        }
    }

    unaryExpr = func() Expr {
        if match([]rune{'~'}) {
            return &NegatedProposition{
                Operator: '~',
                Operand: unaryExpr(),
            }
        } else {
            return term()
        }
    }

    implication = func() Expr {
        return binaryExpr(unaryExpr(), []rune{'>'}, implication)
    }

    conjunction = func() Expr {
        return binaryExpr(implication(), []rune{'&'}, conjunction)
    }

    disjunction = func() Expr {
        return binaryExpr(conjunction(), []rune{'|'}, disjunction)
    }

    var tree = disjunction()
    return tree, nil
}

