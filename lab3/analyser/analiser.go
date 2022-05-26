package analyser

import "fmt"

type Analyser struct {
	reader reader
}

func NewAnalyser(reader reader) *Analyser {
	return &Analyser{reader: reader}
}

func (a *Analyser) Parse() error {
	return a.program()
}

func (a *Analyser) program() error {
	return a.block()
}

func (a *Analyser) block() error {
	next, ok := a.reader.NextToken()
	if !ok || next != "{" {
		return fmt.Errorf("'{' was expected")
	}

	if err := a.operatorsList(); err != nil {
		return err
	}

	next, ok = a.reader.NextToken()
	if !ok || next != "}" {
		return fmt.Errorf("'}' was expected")
	}

	return nil
}

func (a *Analyser) operatorsList() error {
	if err := a.operator(); err != nil {
		return err
	}

	next, ok := a.reader.NextToken()
	if ok && next == ";" {
		a.reader.UnreadToken(next)
		return a.tail()
	}

	return nil
}

func (a *Analyser) tail() error {
	next, ok := a.reader.NextToken()
	if !ok || next != ";" {
		return fmt.Errorf("';' was expected")
	}

	if err := a.operator(); err != nil {
		return err
	}

	next, ok = a.reader.NextToken()
	if ok && next == ";" {
		a.reader.UnreadToken(next)
		return a.tail()
	}

	return nil
}

func (a *Analyser) operator() error {
	next, ok := a.reader.NextToken()
	if !ok || next != "id" {
		return fmt.Errorf("id was expected")
	}

	next, ok = a.reader.NextToken()
	if !ok || next != "=" {
		return fmt.Errorf("'=' was expected")
	}

	return a.expression()
}

func (a *Analyser) expression() error {
	if err := a.arithmeticExpression(); err != nil {
		return err
	}

	if err := a.compareSign(); err != nil {
		return err
	}

	return a.arithmeticExpression()
}

func (a *Analyser) arithmeticExpression() error {
	_ = a.sumSign()

	err := a.term()
	if err != nil {
		return err
	}

	_ = a.arithmeticExpression2()
	return nil
}

func (a *Analyser) arithmeticExpression2() error {
	err := a.sumSign()
	if err != nil {
		return err
	}

	err = a.term()
	if err != nil {
		return err
	}

	_ = a.arithmeticExpression2()
	return nil
}

func (a *Analyser) compareSign() error {
	next, ok := a.reader.NextToken()
	if !ok || (next != "<" && next != "<=" && next != "=" && next != ">=" && next != ">" && next != "<>") {
		a.reader.UnreadToken(next)
		return fmt.Errorf("compare sign was expected")
	}

	return nil
}

func (a *Analyser) term() error {
	if err := a.multiplier(); err != nil {
		return err
	}

	_ = a.term2()
	return nil
}

func (a *Analyser) term2() error {
	if err := a.multiplicationSign(); err != nil {
		return err
	}

	if err := a.multiplier(); err != nil {
		return err
	}

	_ = a.term2()
	return nil
}

func (a *Analyser) sumSign() error {
	next, ok := a.reader.NextToken()
	if !ok || (next != "+" && next != "-") {
		a.reader.UnreadToken(next)
		return fmt.Errorf("'+' or '-' were expected")
	}

	return nil
}

func (a *Analyser) multiplier() error {
	if err := a.primaryExpression(); err != nil {
		return err
	}

	_ = a.multiplier2()
	return nil
}

func (a *Analyser) multiplier2() error {
	next, ok := a.reader.NextToken()
	if !ok || next != "^" {
		a.reader.UnreadToken(next)
		return fmt.Errorf("'^' was expected")
	}

	if err := a.primaryExpression(); err != nil {
		return err
	}

	_ = a.multiplier2()
	return nil
}

func (a *Analyser) multiplicationSign() error {
	next, ok := a.reader.NextToken()
	if !ok || (next != "*" && next != "/" && next != "%") {
		a.reader.UnreadToken(next)
		return fmt.Errorf("multiplication sign was expected")
	}

	return nil
}

func (a *Analyser) primaryExpression() error {
	next, ok := a.reader.NextToken()
	if !ok {
		a.reader.UnreadToken(next)
		return fmt.Errorf("primary expression was expected")
	}

	switch next {
	case "number":
		return nil
	case "id":
		return nil
	case "(":
		if err := a.arithmeticExpression(); err != nil {
			return err
		}
		closeNext, ok := a.reader.NextToken()
		if !ok || closeNext != ")" {
			a.reader.UnreadToken(closeNext)
			a.reader.UnreadToken(next)
			return fmt.Errorf("')' was expected")
		}

		return nil
	}

	return fmt.Errorf("primary expression was expected")
}
