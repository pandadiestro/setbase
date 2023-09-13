package ast

// Binary operators

type DisjunctionExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *DisjunctionExpr) R() bool {
    return (ex.Left.R() || ex.Right.R())
}

type ConjunctionExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *ConjunctionExpr) R() bool {
    return (ex.Left.R() && ex.Right.R())
}

type ImplicationExpr struct {
    Expr
    Left Expr
    Operator rune
    Right Expr
}

func (ex *ImplicationExpr) R() bool {
    return (!ex.Left.R() || ex.Right.R())
}

// Unary operators

type Proposition struct {
    Expr
    Name rune
}

func (pr *Proposition) R() bool {
    return pr.Name == '1'
}

type ConstantProposition struct {
    Expr
    Name rune
}

func (pr *ConstantProposition) R() bool {
    return pr.Name == '1'
}

type NegatedProposition struct {
    Expr
    Operator rune
    Operand Expr
}

func (pr *NegatedProposition) R() bool {
    return !(pr.Operand.R())
}


