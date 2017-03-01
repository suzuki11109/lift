package lift

import (
	"fmt"
	"strings"
)

func New() *Lift {
	return &Lift{floor: 1}
}

type Lift struct {
	floor int
}

func (l *Lift) Status() string {
	return fmt.Sprintf("at %d", l.floor)
}

func (l *Lift) CallFrom(from int) string {
	return moveLiftTo(l, from)
}

func (l *Lift) GoTo(to int) string {
	return moveLiftTo(l, to)
}

func moveLiftTo(l *Lift, dest int) string {
	status := []string{}
	if l.floor == dest {
		return fmt.Sprintf("Lift is already at %d", l.floor)
	}

	for l.floor != dest {
		if l.floor < dest {
			l.floor += 1
		} else {
			l.floor -= 1
		}

		status = append(status, fmt.Sprintf("at %d", l.floor))
	}

	return strings.Join(status, "\n")
}
