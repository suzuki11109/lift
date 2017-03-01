package lift_test

import (
	"testing"

	"github.com/suzuki11109/assert"
	"github.com/suzuki11109/lift"
)

func TestLiftStartWithFirstFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, l.Status(), "at 1")
}

func TestCallFromSecondFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, "at 2", l.CallFrom(2))
}

func TestGoToFirstFloorFromSecondFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, "at 2", l.CallFrom(2))
	assert.Equal(t, "at 1", l.GoTo(1))
}

func TestCallAtThirdFloorFromFirstFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, "at 2\nat 3", l.CallFrom(3))
}

func TestCallAtFirstFloorFromThirdFloor(t *testing.T) {
	l := lift.New()
	l.CallFrom(3)
	assert.Equal(t, "at 2\nat 1", l.CallFrom(1))
}

func TestGoToThirdFloorFromFirstFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, "at 2\nat 3", l.GoTo(3))
}

func TestGoToFirstFloorFromFirstFloor(t *testing.T) {
	l := lift.New()
	assert.Equal(t, "Lift is already at 1", l.GoTo(1))
}
