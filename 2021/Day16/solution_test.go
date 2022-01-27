package main

import (
	"Advent-of-Code/utils"
	"reflect"
	"testing"
)

var exampleBits1 = []string{"1", "1", "0", "1", "0", "0", "1", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "1", "0", "1", "0", "0", "0"}
var exampleBits2 = []string{"0", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "0", "1", "1", "1", "1", "0", "1", "0", "0", "0", "1", "0", "1", "0", "0", "1", "0", "1", "0", "0", "1", "0", "0", "0", "1", "0", "0", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
var exampleBits3 = []string{"1", "1", "1", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "0", "1", "0", "1", "0", "0", "0", "0", "0", "0", "1", "1", "0", "0", "1", "0", "0", "0", "0", "0", "1", "0", "0", "0", "1", "1", "0", "0", "0", "0", "0", "1", "1", "0", "0", "0", "0", "0"}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "returns an error if the hex line can't be parsed into binary",
			input:   "qqq",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns slice binary form of the given hexidecimal string, advent of code example 1",
			input:   "D2FE28",
			want:    exampleBits1,
			wantErr: false,
		},
		{
			name:    "returns slice binary form of the given hexidecimal string, advent of code example 2",
			input:   "38006F45291200",
			want:    exampleBits2,
			wantErr: false,
		},
		{
			name:    "returns slice binary form of the given hexidecimal string, advent of code example 3",
			input:   "EE00D40C823060",
			want:    exampleBits3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVersionOrTypeID(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if bits can't be converted from binary",
			args: args{
				bits: []string{"0", "1", "2", "0", "1"},
				i:    0,
			},
			want:    -1,
			want1:   0,
			wantErr: true,
		},
		{
			name: "returns correct value for given bits, advent of code example 1 version",
			args: args{
				bits: exampleBits1,
				i:    0,
			},
			want:    6,
			want1:   3,
			wantErr: false,
		},
		{
			name: "returns correct value for given bits, advent of code example 1 typeID",
			args: args{
				bits: exampleBits1,
				i:    3,
			},
			want:    4,
			want1:   6,
			wantErr: false,
		},
		{
			name: "returns correct value for given bits, advent of code example 2 version",
			args: args{
				bits: exampleBits2,
				i:    0,
			},
			want:    1,
			want1:   3,
			wantErr: false,
		},
		{
			name: "returns correct value for given bits, advent of code example 2 typeID",
			args: args{
				bits: exampleBits2,
				i:    3,
			},
			want:    6,
			want1:   6,
			wantErr: false,
		},
		{
			name: "returns correct value for given bits, advent of code example 3 version",
			args: args{
				bits: exampleBits3,
				i:    0,
			},
			want:    7,
			want1:   3,
			wantErr: false,
		},
		{
			name: "returns correct value for given bits, advent of code example 3 typeID",
			args: args{
				bits: exampleBits3,
				i:    3,
			},
			want:    3,
			want1:   6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVersionOrTypeID(tt.args.bits, &tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("getVersionOrTypeID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getVersionOrTypeID() = %v, want %v", got, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("getVersionOrTypeID() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_getVersion(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		packet  *Packet
		args    args
		want    *Packet
		want1   int
		wantErr bool
	}{
		{
			name:   "returns an error if the version bits cannot be converted from binary",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"1", "0", "q", "1", "1"},
				i:    0,
			},
			want:    &Packet{value: -1},
			want1:   0,
			wantErr: true,
		},
		{
			name:   "sets the version from the given bits and index, advent of code example 1",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits1,
				i:    0,
			},
			want:    &Packet{value: -1, version: 6},
			want1:   3,
			wantErr: false,
		},
		{
			name:   "sets the version from the given bits and index, advent of code example 2",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits2,
				i:    0,
			},
			want:    &Packet{value: -1, version: 1},
			want1:   3,
			wantErr: false,
		},
		{
			name:   "sets the version from the given bits and index, advent of code example 3",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits3,
				i:    0,
			},
			want:    &Packet{value: -1, version: 7},
			want1:   3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.getVersion(tt.args.bits, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Packet.getVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.getVersion() = %v, want %v", p, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("Packet.getVersion() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_getTypeID(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		packet  *Packet
		args    args
		want    *Packet
		want1   int
		wantErr bool
	}{
		{
			name:   "returns an error if the typeID bits cannot be converted from binary",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"1", "0", "1", "1", "3", "1", "0"},
				i:    3,
			},
			want:    &Packet{value: -1},
			want1:   3,
			wantErr: true,
		},
		{
			name:   "sets the typeID from the given bits and index, advent of code example 1",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits1,
				i:    3,
			},
			want:    &Packet{value: -1, typeID: 4},
			want1:   6,
			wantErr: false,
		},
		{
			name:   "sets the typeID from the given bits and index, advent of code example 2",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits2,
				i:    3,
			},
			want:    &Packet{value: -1, typeID: 6},
			want1:   6,
			wantErr: false,
		},
		{
			name:   "sets the typeID from the given bits and index, advent of code example 3",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits3,
				i:    3,
			},
			want:    &Packet{value: -1, typeID: 3},
			want1:   6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.getTypeID(tt.args.bits, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Packet.getTypeID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.getTypeID() = %v, want %v", p, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("Packet.getTypeID() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_getLiteralValue(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		packet  *Packet
		args    args
		want    *Packet
		want1   int
		wantErr bool
	}{
		{
			name:   "returns an error if the value bits cannot be converted from binary",
			packet: &Packet{value: -1, version: 8, typeID: 9},
			args: args{
				bits: []string{"1", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "s", "1", "0"},
				i:    6,
			},
			want:    &Packet{value: -1, version: 8, typeID: 9},
			want1:   16,
			wantErr: true,
		},
		{
			name:   "sets the value from the given bits and index, advent of code example 1",
			packet: &Packet{value: -1, version: 6, typeID: 4},
			args: args{
				bits: exampleBits1,
				i:    6,
			},
			want:    &Packet{value: 2021, version: 6, typeID: 4},
			want1:   21,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.getLiteralValue(tt.args.bits, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Packet.getLiteralValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.getLiteralValue() = %v, want %v", p, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("Packet.getLiteralValue() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_evaluateOperatorPacket(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		packet  *Packet
		args    args
		want    *Packet
		want1   int
		wantErr bool
	}{
		{
			name:   "returns an error if lengthTypeID is not valid",
			packet: &Packet{version: 2, typeID: 3, value: -1},
			args: args{
				bits: []string{"1", "0", "1", "1", "0", "0", "a", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0", "o", "0", "0"},
				i:    6,
			},
			want:    &Packet{version: 2, typeID: 3, value: -1},
			want1:   7,
			wantErr: true,
		},
		{
			name:   "returns an error if bits cannot be parsed",
			packet: &Packet{version: 2, typeID: 3, value: -1},
			args: args{
				bits: []string{"1", "0", "1", "1", "0", "0", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0", "o", "0", "0"},
				i:    6,
			},
			want:    &Packet{version: 2, typeID: 3, value: -1},
			want1:   18,
			wantErr: true,
		},
		{
			name:   "returns an error if subpacket evaluation for type 0 returns an error",
			packet: &Packet{version: 2, typeID: 3, value: -1},
			args: args{
				bits: []string{"1", "0", "1", "1", "0", "0", "0", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "0", "l", "0"},
				i:    6,
			},
			want:    &Packet{version: 2, typeID: 3, value: -1},
			want1:   22,
			wantErr: true,
		},
		{
			name:   "returns an error if subpacket evaluation for type 1 returns an error",
			packet: &Packet{version: 2, typeID: 3, value: -1},
			args: args{
				bits: []string{"1", "0", "1", "1", "0", "0", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "0", "l", "0"},
				i:    6,
			},
			want:    &Packet{version: 2, typeID: 3, value: -1},
			want1:   21,
			wantErr: true,
		},
		{
			name:   "correctly evaluates sub packets with length type 0",
			packet: &Packet{version: 1, typeID: 6, value: -1},
			args: args{
				bits: exampleBits2,
				i:    6,
			},
			want: &Packet{version: 1, typeID: 6, value: -1, subPackets: []Packet{
				{version: 6, typeID: 4, value: 10},
				{version: 2, typeID: 4, value: 20},
			}},
			want1:   49,
			wantErr: false,
		},
		{
			name:   "correctly evaluates sub packets with length type 1",
			packet: &Packet{version: 7, typeID: 3, value: -1},
			args: args{
				bits: exampleBits3,
				i:    6,
			},
			want: &Packet{version: 7, typeID: 3, value: -1, subPackets: []Packet{
				{version: 2, typeID: 4, value: 1},
				{version: 4, typeID: 4, value: 2},
				{version: 1, typeID: 4, value: 3},
			}},
			want1:   51,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.evaluateOperatorPacket(tt.args.bits, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Packet.evaluateOperatorPacket() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.evaluateOperatorPacket() = %v, want %v", p, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("Packet.evaluateOperatorPacket() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_evaluatePacket(t *testing.T) {
	type args struct {
		bits []string
		i    int
	}
	tests := []struct {
		name    string
		packet  *Packet
		args    args
		want    *Packet
		want1   int
		wantErr bool
	}{
		{
			name:   "returns an error if getVersion returns an error",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"0", "1", "z"},
				i:    0,
			},
			want:    &Packet{value: -1},
			want1:   0,
			wantErr: true,
		},
		{
			name:   "returns an error if getTypeID returns an error",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"0", "1", "1", "1", "0", "9"},
				i:    0,
			},
			want:    &Packet{version: 3, value: -1},
			want1:   3,
			wantErr: true,
		},
		{
			name:   "returns an error if getLiteralValue returns an error",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "a", "0"},
				i:    0,
			},
			want:    &Packet{version: 3, typeID: 4, value: -1},
			want1:   11,
			wantErr: true,
		},
		{
			name:   "returns an error if evaluateOperatorPacket returns an error",
			packet: &Packet{value: -1},
			args: args{
				bits: []string{"0", "1", "1", "1", "0", "1", "a", "1", "1", "0", "0"},
				i:    0,
			},
			want:    &Packet{version: 3, typeID: 5, value: -1},
			want1:   7,
			wantErr: true,
		},
		{
			name:   "correctly evaluates packet, advent of code example 1",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits1,
				i:    0,
			},
			want:    &Packet{version: 6, typeID: 4, value: 2021},
			want1:   21,
			wantErr: false,
		},
		{
			name:   "correctly evaluates packet, advent of code example 2",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits2,
				i:    0,
			},
			want: &Packet{version: 1, typeID: 6, value: -1, subPackets: []Packet{
				{version: 6, typeID: 4, value: 10},
				{version: 2, typeID: 4, value: 20},
			}},
			want1:   49,
			wantErr: false,
		},
		{
			name:   "correctly evaluates packet, advent of code example 3",
			packet: &Packet{value: -1},
			args: args{
				bits: exampleBits3,
				i:    0,
			},
			want: &Packet{version: 7, typeID: 3, value: -1, subPackets: []Packet{
				{version: 2, typeID: 4, value: 1},
				{version: 4, typeID: 4, value: 2},
				{version: 1, typeID: 4, value: 3},
			}},
			want1:   51,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.evaluatePacket(tt.args.bits, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Packet.evaluatePacket() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.evaluatePacket() = %v, want %v", p, tt.want)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("Packet.evaluatePacket() *i = %v, want %v", tt.args.i, tt.want1)
			}
		})
	}
}

func TestPacket_sumVersions(t *testing.T) {
	tests := []struct {
		name   string
		packet *Packet
		want   int
	}{
		{
			name: "it correctly sums all version numbers of the packet and subpackets, advent of code example 1",
			packet: &Packet{
				version: 4, subPackets: []Packet{
					{version: 1, subPackets: []Packet{
						{version: 5, subPackets: []Packet{
							{version: 6},
						}},
					}},
				},
			},
			want: 16,
		},
		{
			name: "it correctly sums all version numbers of the packet and subpackets, advent of code example 1",
			packet: &Packet{
				version: 7, subPackets: []Packet{
					{version: 1, subPackets: []Packet{
						{version: 5, subPackets: []Packet{
							{version: 5},
						}},
						{version: 5, subPackets: []Packet{
							{version: 3},
						}},
					}},
					{version: 1, subPackets: []Packet{
						{version: 5, subPackets: []Packet{
							{version: 6, subPackets: []Packet{
								{version: 2},
							}},
						}},
						{version: 2, subPackets: []Packet{
							{version: 6},
						}},
					}},
				},
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if got := p.sumVersions(); got != tt.want {
				t.Errorf("Packet.sumVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacket_getValue(t *testing.T) {
	tests := []struct {
		name    string
		packet  *Packet
		want    *Packet
		wantErr bool
	}{
		{
			name: "runs nothing on the packet if the value is already determined",
			packet: &Packet{typeID: 0, value: 9, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			want: &Packet{typeID: 0, value: 9, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			wantErr: false,
		},
		{
			name: "runs an error if typeID is invalid",
			packet: &Packet{typeID: 8, value: -1, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			want: &Packet{typeID: 8, value: -1, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 0 and getValue fails",
			packet: &Packet{typeID: 0, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 0, value: 0, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet if typeID is 0",
			packet: &Packet{typeID: 0, value: -1, subPackets: []Packet{
				{value: 9, typeID: 4},
				{value: 2, typeID: 4},
				{value: 6, typeID: 4},
			}},
			want: &Packet{typeID: 0, value: 17, subPackets: []Packet{
				{value: 9, typeID: 4},
				{value: 2, typeID: 4},
				{value: 6, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "returns an error if typeID is 1 and getValue fails",
			packet: &Packet{typeID: 1, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 1, value: 1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet if typeID is 1",
			packet: &Packet{typeID: 1, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 11, typeID: 4},
				{value: 13, typeID: 4},
			}},
			want: &Packet{typeID: 1, value: 1001, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 11, typeID: 4},
				{value: 13, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "returns an error if typeID is 2 and getValue fails",
			packet: &Packet{typeID: 2, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 2, value: utils.Infinity, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet if typeID is 2",
			packet: &Packet{typeID: 2, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 11, typeID: 4},
				{value: 13, typeID: 4},
			}},
			want: &Packet{typeID: 2, value: 7, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 11, typeID: 4},
				{value: 13, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "returns an error if typeID is 3 and getValue fails",
			packet: &Packet{typeID: 3, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 3, value: 0, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet if typeID is 3",
			packet: &Packet{typeID: 3, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			want: &Packet{typeID: 3, value: 13, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "runs nothing on the packet if the vtypeID is 4",
			packet: &Packet{typeID: 4, value: -1, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			want: &Packet{typeID: 4, value: -1, subPackets: []Packet{
				{value: -1, typeID: 7},
			}},
			wantErr: false,
		},
		{
			name: "returns an error if typeID is 5 and getValue fails",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 5 and there are fewer than 2 subpackets",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			want: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 5 and there are greater than 2 subpackets",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			want: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 5 and the first of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			want: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 5 and the second of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			want: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet to 1 if typeID is 5 and the first subpacket value is greater than the second subpacket value",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 27, typeID: 4},
				{value: 13, typeID: 4},
			}},
			want: &Packet{typeID: 5, value: 1, subPackets: []Packet{
				{value: 27, typeID: 4},
				{value: 13, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "correctly sets the value of the packet to 0 if typeID is 5 and the first subpacket value is less than the second subpacket value",
			packet: &Packet{typeID: 5, value: -1, subPackets: []Packet{
				{value: 27, typeID: 4},
				{value: 43, typeID: 4},
			}},
			want: &Packet{typeID: 5, value: 0, subPackets: []Packet{
				{value: 27, typeID: 4},
				{value: 43, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "returns an error if typeID is 6 and getValue fails",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 6 and there are fewer than 2 subpackets",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			want: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 6 and there are greater than 2 subpackets",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			want: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 6 and the first of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			want: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 6 and the second of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			want: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet to 1 if typeID is 6 and the first subpacket value is less than the second subpacket value",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
			}},
			want: &Packet{typeID: 6, value: 1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "correctly sets the value of the packet to 0 if typeID is 6 and the first subpacket value is greater than the second subpacket value",
			packet: &Packet{typeID: 6, value: -1, subPackets: []Packet{
				{value: 127, typeID: 4},
				{value: 43, typeID: 4},
			}},
			want: &Packet{typeID: 6, value: 0, subPackets: []Packet{
				{value: 127, typeID: 4},
				{value: 43, typeID: 4},
			}},
			wantErr: false,
		},

		{
			name: "returns an error if typeID is 7 and getValue fails",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			want: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: -1, typeID: 8},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 7 and there are fewer than 2 subpackets",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			want: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 7 and there are greater than 2 subpackets",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			want: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: 13, typeID: 4},
				{value: 11, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 7 and the first of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			want: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: -1, typeID: 9},
				{value: 2, typeID: 4},
			}},
			wantErr: true,
		},
		{
			name: "returns an error if typeID is 7 and the second of 2 subpackets returns an error for getValue",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			want: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 7, typeID: 4},
				{value: -1, typeID: 9},
			}},
			wantErr: true,
		},
		{
			name: "correctly sets the value of the packet to 1 if typeID is 7 and the first subpacket value is equal to the second subpacket value",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 137, typeID: 4},
				{value: 137, typeID: 4},
			}},
			want: &Packet{typeID: 7, value: 1, subPackets: []Packet{
				{value: 137, typeID: 4},
				{value: 137, typeID: 4},
			}},
			wantErr: false,
		},
		{
			name: "correctly sets the value of the packet to 0 if typeID is 6 and the first subpacket value is not equal to the second subpacket value",
			packet: &Packet{typeID: 7, value: -1, subPackets: []Packet{
				{value: 127, typeID: 4},
				{value: 43, typeID: 4},
			}},
			want: &Packet{typeID: 7, value: 0, subPackets: []Packet{
				{value: 127, typeID: 4},
				{value: 43, typeID: 4},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.packet
			if err := p.getValue(); (err != nil) != tt.wantErr {
				t.Errorf("Packet.getValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Packet.getValue() = %v, want %v", p, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if parseInput returns an error",
			input:   "QQQQQ",
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 1",
			input:   "D2FE28",
			want:    6,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 2",
			input:   "38006F45291200",
			want:    9,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 3",
			input:   "EE00D40C823060",
			want:    14,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 4",
			input:   "8A004A801A8002F478",
			want:    16,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 5",
			input:   "620080001611562C8802118E34",
			want:    12,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 6",
			input:   "C0015000016115A2E0802F182340",
			want:    23,
			wantErr: false,
		},
		{
			name:    "calculates sum of versions from given input, advent of code example 7",
			input:   "A0016C880162017C3686B18A3D4780",
			want:    31,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 1",
			input:   "C200B40A82",
			want1:   3,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 2",
			input:   "04005AC33890",
			want1:   54,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 3",
			input:   "880086C3E88112",
			want1:   7,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 4",
			input:   "CE00C43D881120",
			want1:   9,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 5",
			input:   "D8005AC2A8F0",
			want1:   1,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 6",
			input:   "F600BC2D8F",
			want1:   0,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 7",
			input:   "9C005AC2F8F0",
			want1:   0,
			wantErr: false,
		},
		{
			name:    "calculates value of outer packet from given input, advent of code example 8",
			input:   "9C0141080250320F1802104A08",
			want1:   1,
			wantErr: false,
		},
		// Throw in my puzzle input for good measure, since none of the advent of code examples provide solutions to both parts
		{
			name:    "calculates sum of versions and value of outer packet from given input, actual input",
			input:   "20546718027401204FE775D747A5AD3C3CCEEB24CC01CA4DFF2593378D645708A56D5BD704CC0110C469BEF2A4929689D1006AF600AC942B0BA0C942B0BA24F9DA8023377E5AC7535084BC6A4020D4C73DB78F005A52BBEEA441255B42995A300AA59C27086618A686E71240005A8C73D4CF0AC40169C739584BE2E40157D0025533770940695FE982486C802DD9DC56F9F07580291C64AAAC402435802E00087C1E8250440010A8C705A3ACA112001AF251B2C9009A92D8EBA6006A0200F4228F50E80010D8A7052280003AD31D658A9231AA34E50FC8010694089F41000C6A73F4EDFB6C9CC3E97AF5C61A10095FE00B80021B13E3D41600042E13C6E8912D4176002BE6B060001F74AE72C7314CEAD3AB14D184DE62EB03880208893C008042C91D8F9801726CEE00BCBDDEE3F18045348F34293E09329B24568014DCADB2DD33AEF66273DA45300567ED827A00B8657B2E42FD3795ECB90BF4C1C0289D0695A6B07F30B93ACB35FBFA6C2A007A01898005CD2801A60058013968048EB010D6803DE000E1C6006B00B9CC028D8008DC401DD9006146005980168009E1801B37E02200C9B0012A998BACB2EC8E3D0FC8262C1009D00008644F8510F0401B825182380803506A12421200CB677011E00AC8C6DA2E918DB454401976802F29AA324A6A8C12B3FD978004EB30076194278BE600C44289B05C8010B8FF1A6239802F3F0FFF7511D0056364B4B18B034BDFB7173004740111007230C5A8B6000874498E30A27BF92B3007A786A51027D7540209A04821279D41AA6B54C15CBB4CC3648E8325B490401CD4DAFE004D932792708F3D4F769E28500BE5AF4949766DC24BB5A2C4DC3FC3B9486A7A0D2008EA7B659A00B4B8ACA8D90056FA00ACBCAA272F2A8A4FB51802929D46A00D58401F8631863700021513219C11200996C01099FBBCE6285106",
			want:    955,
			want1:   158135423448,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != 0 && got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if tt.want1 != 0 && got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
