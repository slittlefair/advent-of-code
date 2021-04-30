package waypoint

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2020/Day12/ship"
	"fmt"
	"reflect"
	"testing"
)

type turnWaypointTestCase struct {
	name string
	w    *Waypoint
	val  int
	want *Waypoint
}

func generateTurnLeftTest(w, want *Waypoint, val int) turnWaypointTestCase {
	return generateTest(w, want, val, "left")
}

func generateTurnRightTest(w, want *Waypoint, val int) turnWaypointTestCase {
	return generateTest(w, want, val, "right")
}

func generateTest(w, want *Waypoint, val int, dir string) turnWaypointTestCase {
	return turnWaypointTestCase{
		name: fmt.Sprintf("correctly rotates waypoint %s %d degrees", dir, val),
		w:    w,
		val:  val,
		want: want,
	}
}

func TestWaypoint_turnWaypointLeft(t *testing.T) {
	tests := []turnWaypointTestCase{
		generateTurnLeftTest(&Waypoint{X: 19, Y: 33}, &Waypoint{X: -33, Y: 19}, 90),
		generateTurnLeftTest(&Waypoint{X: 33, Y: 17}, &Waypoint{X: -33, Y: -17}, 180),
		generateTurnLeftTest(&Waypoint{X: -99, Y: -303}, &Waypoint{X: -303, Y: 99}, 270),
		generateTurnLeftTest(&Waypoint{X: -112, Y: 63}, &Waypoint{X: -112, Y: 63}, 360),
		generateTurnLeftTest(&Waypoint{X: 0, Y: 2}, &Waypoint{X: -2, Y: 0}, 450),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.turnWaypointLeft(tt.val)
			if !reflect.DeepEqual(tt.w, tt.want) {
				t.Errorf("parseBag() w = %v, want %v", tt.w, tt.want)
			}
		})
	}
}

func TestWaypoint_turnWaypointRight(t *testing.T) {
	tests := []struct {
		name string
		w    *Waypoint
		val  int
		want *Waypoint
	}{
		generateTurnRightTest(&Waypoint{X: 19, Y: 33}, &Waypoint{X: 33, Y: -19}, 90),
		generateTurnRightTest(&Waypoint{X: 33, Y: -17}, &Waypoint{X: -33, Y: 17}, 180),
		generateTurnRightTest(&Waypoint{X: 0, Y: 303}, &Waypoint{X: -303, Y: 0}, 270),
		generateTurnRightTest(&Waypoint{X: -112, Y: 63}, &Waypoint{X: -112, Y: 63}, 360),
		generateTurnRightTest(&Waypoint{X: 0, Y: 0}, &Waypoint{X: 0, Y: 0}, 450),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.turnWaypointRight(tt.val)
			if !reflect.DeepEqual(tt.w, tt.want) {
				t.Errorf("parseBag() w = %v, want %v", tt.w, tt.want)
			}
		})
	}
}

func TestWaypoint_MoveWaypoint(t *testing.T) {
	var baseShip = &ship.Ship{
		Co:        helpers.Co{X: 99, Y: 13},
		FacingDir: "E",
	}
	type args struct {
		s   *ship.Ship
		d   string
		val int
	}
	tests := []struct {
		name  string
		w     *Waypoint
		args  args
		want  *Waypoint
		want1 *ship.Ship
	}{
		{
			name: "correctly moves the waypoint north the given amount when told to",
			w:    &Waypoint{X: 87, Y: -23},
			args: args{
				s:   baseShip,
				d:   "N",
				val: 75,
			},
			want:  &Waypoint{X: 87, Y: 52},
			want1: baseShip,
		},
		{
			name: "correctly moves the waypoint east the given amount when told to",
			w:    &Waypoint{X: 87, Y: -23},
			args: args{
				s:   baseShip,
				d:   "E",
				val: 12,
			},
			want:  &Waypoint{X: 99, Y: -23},
			want1: baseShip,
		},
		{
			name: "correctly moves the waypoint south the given amount when told to",
			w:    &Waypoint{X: 879, Y: 2},
			args: args{
				s:   baseShip,
				d:   "S",
				val: 7,
			},
			want:  &Waypoint{X: 879, Y: -5},
			want1: baseShip,
		},
		{
			name: "correctly moves the waypoint west the given amount when told to",
			w:    &Waypoint{X: 287, Y: -23},
			args: args{
				s:   baseShip,
				d:   "W",
				val: 588,
			},
			want:  &Waypoint{X: -301, Y: -23},
			want1: baseShip,
		},
		{
			name: "correctly rotates the waypoint left the given amount when told to",
			w:    &Waypoint{X: 87, Y: -23},
			args: args{
				s:   baseShip,
				d:   "L",
				val: 90,
			},
			want:  &Waypoint{X: 23, Y: 87},
			want1: baseShip,
		},
		{
			name: "correctly rotates the waypoint right the given amount when told to",
			w:    &Waypoint{X: 807, Y: 1113},
			args: args{
				s:   baseShip,
				d:   "R",
				val: 450,
			},
			want:  &Waypoint{X: 1113, Y: -807},
			want1: baseShip,
		},
		{
			name: "correctly moves the waypoint and ship north the given amount when told to",
			w:    &Waypoint{X: 0, Y: 23},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 99, Y: 13},
					FacingDir: "E",
				},
				d:   "F",
				val: 4,
			},
			want: &Waypoint{X: 0, Y: 23},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 99, Y: 105},
				FacingDir: "E",
			},
		},
		{
			name: "correctly moves the waypoint and ship east the given amount when told to",
			w:    &Waypoint{X: 70, Y: 0},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: -6, Y: 22},
					FacingDir: "S",
				},
				d:   "F",
				val: 1,
			},
			want: &Waypoint{X: 70, Y: 0},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 64, Y: 22},
				FacingDir: "S",
			},
		},
		{
			name: "correctly moves the waypoint and ship south the given amount when told to",
			w:    &Waypoint{X: 0, Y: -6},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 2, Y: -7},
					FacingDir: "E",
				},
				d:   "F",
				val: 7,
			},
			want: &Waypoint{X: 0, Y: -6},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 2, Y: -49},
				FacingDir: "E",
			},
		},
		{
			name: "correctly moves the waypoint and ship west the given amount when told to",
			w:    &Waypoint{X: -12, Y: 0},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 17, Y: 1333},
					FacingDir: "N",
				},
				d:   "F",
				val: 12,
			},
			want: &Waypoint{X: -12, Y: 0},
			want1: &ship.Ship{
				Co:        helpers.Co{X: -127, Y: 1333},
				FacingDir: "N",
			},
		},
		{
			name: "correctly moves the waypoint and ship north-east the given amount when told to",
			w:    &Waypoint{X: 10, Y: 3},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 99, Y: 103},
					FacingDir: "S",
				},
				d:   "F",
				val: 5,
			},
			want: &Waypoint{X: 10, Y: 3},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 149, Y: 118},
				FacingDir: "S",
			},
		},
		{
			name: "correctly moves the waypoint and ship south-east the given amount when told to",
			w:    &Waypoint{X: 7, Y: -9},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: -9, Y: -9},
					FacingDir: "S",
				},
				d:   "F",
				val: 2,
			},
			want: &Waypoint{X: 7, Y: -9},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 5, Y: -27},
				FacingDir: "S",
			},
		},
		{
			name: "correctly moves the waypoint and ship south-west the given amount when told to",
			w:    &Waypoint{X: -6, Y: -33},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 1, Y: 12},
					FacingDir: "S",
				},
				d:   "F",
				val: 10,
			},
			want: &Waypoint{X: -6, Y: -33},
			want1: &ship.Ship{
				Co:        helpers.Co{X: -59, Y: -318},
				FacingDir: "S",
			},
		},
		{
			name: "correctly moves the waypoint and ship north-west the given amount when told to",
			w:    &Waypoint{X: -4, Y: 3},
			args: args{
				s: &ship.Ship{
					Co:        helpers.Co{X: 99, Y: 103},
					FacingDir: "E",
				},
				d:   "F",
				val: 1,
			},
			want: &Waypoint{X: -4, Y: 3},
			want1: &ship.Ship{
				Co:        helpers.Co{X: 95, Y: 106},
				FacingDir: "E",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.MoveWaypoint(tt.args.s, tt.args.d, tt.args.val)
			if !reflect.DeepEqual(tt.w, tt.want) {
				t.Errorf("parseBag() w = %v, want %v", tt.w, tt.want)
			}
			if !reflect.DeepEqual(tt.args.s, tt.want1) {
				t.Errorf("parseBag() w = %v, want %v", tt.args.s, tt.want1)
			}
		})
	}
}
