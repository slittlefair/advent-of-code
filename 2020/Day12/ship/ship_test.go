package ship

import (
	helpers "Advent-of-Code"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type turnShipTestArgs struct {
	d   string
	val int
}

type turnShipTestCase struct {
	name      string
	facingDir string
	args      turnShipTestArgs
	want      string
}

func generateTest(facingDir, d string, val int, want string) turnShipTestCase {
	return turnShipTestCase{
		name:      fmt.Sprintf("faces %s when turning %d degrees %s from %s", want, val, d, facingDir),
		facingDir: facingDir,
		args: turnShipTestArgs{
			d:   strings.ToUpper(string(d[0])),
			val: val,
		},
		want: want,
	}
}

func TestShip_turnShip(t *testing.T) {
	tests := []turnShipTestCase{
		generateTest("N", "left", 90, "W"),
		generateTest("N", "left", 180, "S"),
		generateTest("N", "left", 270, "E"),
		generateTest("N", "left", 360, "N"),
		generateTest("N", "left", 450, "W"),
		generateTest("E", "left", 90, "N"),
		generateTest("E", "left", 180, "W"),
		generateTest("E", "left", 270, "S"),
		generateTest("E", "left", 360, "E"),
		generateTest("E", "left", 450, "N"),
		generateTest("S", "left", 90, "E"),
		generateTest("S", "left", 180, "N"),
		generateTest("S", "left", 270, "W"),
		generateTest("S", "left", 360, "S"),
		generateTest("S", "left", 450, "E"),
		generateTest("W", "left", 90, "S"),
		generateTest("W", "left", 180, "E"),
		generateTest("W", "left", 270, "N"),
		generateTest("W", "left", 360, "W"),
		generateTest("W", "left", 450, "S"),
		generateTest("N", "right", 90, "E"),
		generateTest("N", "right", 180, "S"),
		generateTest("N", "right", 270, "W"),
		generateTest("N", "right", 360, "N"),
		generateTest("N", "right", 450, "E"),
		generateTest("E", "right", 90, "S"),
		generateTest("E", "right", 180, "W"),
		generateTest("E", "right", 270, "N"),
		generateTest("E", "right", 360, "E"),
		generateTest("E", "right", 450, "S"),
		generateTest("S", "right", 90, "W"),
		generateTest("S", "right", 180, "N"),
		generateTest("S", "right", 270, "E"),
		generateTest("S", "right", 360, "S"),
		generateTest("S", "right", 450, "W"),
		generateTest("W", "right", 90, "N"),
		generateTest("W", "right", 180, "E"),
		generateTest("W", "right", 270, "S"),
		generateTest("W", "right", 360, "W"),
		generateTest("W", "right", 450, "N"),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				FacingDir: tt.facingDir,
			}
			s.turnShip(tt.args.d, tt.args.val)
			if !reflect.DeepEqual(s.FacingDir, tt.want) {
				t.Errorf("parseBag() s.FacingDir = %s, want %s", s.FacingDir, tt.want)
			}
		})
	}
}

func TestShip_MoveShip(t *testing.T) {
	type fields struct {
		Co        helpers.Coordinate
		FacingDir string
	}
	type args struct {
		d   string
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ship
	}{
		{
			name: "correctly moves the ship north the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 3, Y: 11},
				FacingDir: "E",
			},
			args: args{
				d:   "N",
				val: 34,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 3, Y: 45},
				FacingDir: "E",
			},
		},
		{
			name: "correctly moves the ship east the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 21, Y: 37},
				FacingDir: "N",
			},
			args: args{
				d:   "E",
				val: 88,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 109, Y: 37},
				FacingDir: "N",
			},
		},
		{
			name: "correctly moves the ship south the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 12, Y: 67},
				FacingDir: "N",
			},
			args: args{
				d:   "S",
				val: 49,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 12, Y: 18},
				FacingDir: "N",
			},
		},
		{
			name: "correctly moves the ship west the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 41, Y: 7},
				FacingDir: "E",
			},
			args: args{
				d:   "W",
				val: 53,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: -12, Y: 7},
				FacingDir: "E",
			},
		},
		{
			name: "correctly turns the ship left the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 21, Y: 79},
				FacingDir: "W",
			},
			args: args{
				d:   "L",
				val: 270,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 21, Y: 79},
				FacingDir: "N",
			},
		},
		{
			name: "correctly turns the ship right the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: 21, Y: 79},
				FacingDir: "N",
			},
			args: args{
				d:   "R",
				val: 180,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 21, Y: 79},
				FacingDir: "S",
			},
		},
		{
			name: "correctly moves the ship forward the given amount when told to",
			fields: fields{
				Co:        helpers.Coordinate{X: -19, Y: 66},
				FacingDir: "E",
			},
			args: args{
				d:   "F",
				val: 187,
			},
			want: &Ship{
				Co:        helpers.Coordinate{X: 168, Y: 66},
				FacingDir: "E",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				Co:        tt.fields.Co,
				FacingDir: tt.fields.FacingDir,
			}
			s.MoveShip(tt.args.d, tt.args.val)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("parseBag() s = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestShip_CalculateDistance(t *testing.T) {
	type fields struct {
		Co        helpers.Coordinate
		FacingDir string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 when ship is on the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 0, Y: 0},
			},
			want: 0,
		},
		{
			name: "returns distance when ship is north of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 0, Y: 17},
			},
			want: 17,
		},
		{
			name: "returns distance when ship is east of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 99, Y: 0},
			},
			want: 99,
		},
		{
			name: "returns distance when ship is south of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 0, Y: -981},
			},
			want: 981,
		},
		{
			name: "returns distance when ship is west of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: -1, Y: 0},
			},
			want: 1,
		},
		{
			name: "returns distance when ship is north-east of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 31, Y: 17},
			},
			want: 48,
		},
		{
			name: "returns distance when ship is south-east of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: 42, Y: -65},
			},
			want: 107,
		},
		{
			name: "returns distance when ship is south-west of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: -301, Y: -67},
			},
			want: 368,
		},
		{
			name: "returns distance when ship is north-west of the origin",
			fields: fields{
				Co: helpers.Coordinate{X: -10, Y: 54},
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				Co:        tt.fields.Co,
				FacingDir: tt.fields.FacingDir,
			}
			if got := s.CalculateDistance(); got != tt.want {
				t.Errorf("Ship.CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
