package parser

// Binary operators

type DisjunctionExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *DisjunctionExpr) Eval(ls []rune) bool {
    return (ex.Left.Eval(ls) || ex.Right.Eval(ls))
}

type ConjunctionExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *ConjunctionExpr) Eval(ls []rune) bool {
    return (ex.Left.Eval(ls) && ex.Right.Eval(ls))
}

type ImplicationExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *ImplicationExpr) Eval(ls []rune) bool {
    return (!ex.Left.Eval(ls) || ex.Right.Eval(ls))
}

// Unary operators

type Proposition struct {
    Expr
    Name rune
}

func (pr *Proposition) Eval(ls []rune) bool {
    for _, el := range ls {
        if el == pr.Name {
            return true
        }
    }

    return false
}

type ConstantProposition struct {
    Expr
    Name rune
}

func (pr *ConstantProposition) Eval(ls []rune) bool {
    return pr.Name == '1'
}

type NegatedProposition struct {
    Expr
    Operator rune
    Operand Expr
}

func (pr *NegatedProposition) Eval(ls []rune) bool {
    return !(pr.Operand.Eval(ls))
}


