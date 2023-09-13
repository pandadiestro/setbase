package parser
import (
	"errors"
	"fmt"
	"strings"
)

// if this gets somehow implemented with user input
// in mind, it should let this list clear

type kind string
const (
    proposition kind = "PROP"
    operator    kind = "OPER"
    delimiter   kind = "DELM"
)

var specials = [...]struct{
    special string
    replace string
}{
    {
        special: "=>",
        replace: ">",
    },
    {
        special: "<=>",
        replace: "<",
    },
    {
        special: "/\\",
        replace: "&",
    },
    {
        special: "\\/",
        replace: "|",
    },
    {
        special: "!",
        replace: "~",
    },
}

var whitespaces = [...]string{
    "\n",
    "\t",
    " ",
}

type state rune
const (
    closed state = 'C'
    opened state = 'O'
)

var delimitersOpen = [...]rune{
    '(',
    '[',
    '{',
}

var delimitersClosed = [...]rune{
    ')',
    ']',
    '}',
}

var operators = [...]rune{
    '>',
    '+',
    '*',
    '~',
}

//---

type token struct {
    which   kind
    name    rune
    state   *state //used only at delimiters
}

//---

func removeWhitespaces(in string) string {
    str := in
    for _, whitespace := range whitespaces {
        str = strings.ReplaceAll(str, whitespace, "")
    }
    return str
}

func changeSpecialSymbols(in string) string {
    str := in
    for _, specialPair := range specials {
        str = strings.ReplaceAll(str, specialPair.special, specialPair.replace)
    }
    return str
}

func isDelimiterOpen(symbol rune) bool {
    for _, e := range delimitersOpen {
        if e == symbol {
            return true
        }
    }
    return false
}

func isDelimiterClosed(symbol rune) bool {
    for _, e := range delimitersClosed {
        if e == symbol {
            return true
        }
    }
    return false
}

func isOperator(symbol rune) bool {
    for _, e := range operators {
        if e == symbol {
            return true
        }
    }
    return false
}


func CleanString(in string) error {
    //var depth int = 0
    var finalStr string

    finalStr = changeSpecialSymbols(in)
    finalStr = removeWhitespaces(finalStr)

    fmt.Println(finalStr)

    // it should not break with your shitty unicode, even though
    // theres literally no special character to use here
    runesArr := []rune(finalStr)
    syntax := &Stack{}
    for _, symbol := range runesArr {
        if isDelimiterOpen(symbol) {
            syntax.Push()
        } else if isDelimiterClosed(symbol) {
            err := syntax.Pop()
            if err != nil {
                return errors.New("fuck you!!!!")
            }
        } else {
            continue
        }
    }

    if syntax.IsEmpty() == false {
        return errors.New("fuck you!!!!")
    }
    return nil
}

