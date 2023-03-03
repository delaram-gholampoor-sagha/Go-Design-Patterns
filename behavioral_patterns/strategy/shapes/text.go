package shapes

import "github.com/delaram-gholampoor-sagha/Go-Design-Patterns/behavioral_patterns/strategy"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
