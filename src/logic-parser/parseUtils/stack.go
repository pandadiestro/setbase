package parser

import "errors"

type delimNode struct {
    Next *delimNode
}

type Stack struct {
    Head *delimNode
}

func (p *Stack) IsEmpty() bool {
    return (p.Head == nil)
}

func (p *Stack) Push() {
    temp := &delimNode{
        Next: nil,
    }

    if p.Head == nil {
        p.Head = temp
        return
    }

    temp.Next = p.Head
    p.Head = temp
}

func (p *Stack) Pop() error {
    if p.Head == nil {
        return errors.New("cant pop from empty(nil) delimiter stack")
    }

    temp := p.Head.Next
    p.Head = temp
    return nil
}



