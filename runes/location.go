package runes

type Line struct {
	start int
	end   int
}

type Location struct {
	line   int
	column int
}

func NewLocation(line, column int) Location {
	return Location{
		line:   line,
		column: column,
	}
}

func (l Location) Empty() bool {
	return l.column == 0 && l.line == 0
}

func (l *Location) Line() int {
	return l.line
}

func (l *Location) SetLine(val int) {
	l.line = val
}

func (l *Location) Column() int {
	return l.column
}

func (l *Location) SetColumn(val int) {
	l.column = val
}

func (l Location) Equal(k Location) bool {
	return l.line == k.line && l.column == k.column
}

func (l Location) Get() (int, int) {
	return l.line, l.column
}
