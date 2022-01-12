package main

import (
	"Advent-of-Code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	version    int
	typeID     int
	value      int
	subPackets []Packet
}

func parseInput(input string) ([]string, error) {
	b := strings.Builder{}
	for _, r := range input {
		s, err := strconv.ParseUint(string(r), 16, 64)
		if err != nil {
			return nil, err
		}
		b.WriteString(fmt.Sprintf("%04b", s))
	}
	return strings.Split(b.String(), ""), nil
}

func getVersionOrTypeID(bits []string, i *int) (int, error) {
	val, err := strconv.ParseUint(strings.Join(bits[*i:*i+3], ""), 2, 64)
	if err != nil {
		return -1, err
	}
	*i += 3
	return int(val), nil
}

func (p *Packet) getVersion(bits []string, i *int) error {
	val, err := getVersionOrTypeID(bits, i)
	if err != nil {
		return err
	}
	p.version = val
	return nil
}

func (p *Packet) getTypeID(bits []string, i *int) error {
	val, err := getVersionOrTypeID(bits, i)
	if err != nil {
		return err
	}
	p.typeID = val
	return nil
}

func (p *Packet) getLiteralValue(bits []string, i *int) error {
	val := strings.Builder{}
	for {
		val.WriteString(strings.Join(bits[*i+1:*i+5], ""))
		*i += 5
		if bits[*i-5] == "0" {
			break
		}
	}
	v, err := strconv.ParseUint(val.String(), 2, 64)
	if err != nil {
		return err
	}
	p.value = int(v)
	return nil
}

func (p *Packet) evaluateOperatorPacket(bits []string, i *int) error {
	var lengthBits []string
	lengthTypeID := bits[*i]
	*i++
	switch lengthTypeID {
	case "0":
		lengthBits = bits[*i : *i+15]
		*i += 15
	case "1":
		lengthBits = bits[*i : *i+11]
		*i += 11
	default:
		return fmt.Errorf("invalid character at bit %d, expected 0 or 1, got %s", i, lengthTypeID)
	}
	length, err := strconv.ParseUint(strings.Join(lengthBits, ""), 2, 64)
	if err != nil {
		return err
	}
	subPackets := []Packet{}
	if lengthTypeID == "0" {
		end := *i + int(length)
		for *i < end {
			sp := Packet{value: -1}
			err := sp.evaluatePacket(bits, i)
			if err != nil {
				return err
			}
			subPackets = append(subPackets, sp)
		}
	} else {
		for len(subPackets) < int(length) {
			sp := Packet{value: -1}
			err := sp.evaluatePacket(bits, i)
			if err != nil {
				return err
			}
			subPackets = append(subPackets, sp)
		}
	}
	p.subPackets = subPackets
	return nil
}

func (p *Packet) evaluatePacket(bits []string, i *int) error {
	err := p.getVersion(bits, i)
	if err != nil {
		return err
	}
	err = p.getTypeID(bits, i)
	if err != nil {
		return err
	}
	if p.typeID == 4 {
		err = p.getLiteralValue(bits, i)
		if err != nil {
			return err
		}
	} else {
		err = p.evaluateOperatorPacket(bits, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Packet) sumVersions() int {
	v := p.version
	for _, sp := range p.subPackets {
		v += sp.sumVersions()
	}
	return v
}

func (p *Packet) getValue() error {
	if p.value != -1 {
		return nil
	}
	switch p.typeID {
	case 0:
		p.value = 0
		for _, sp := range p.subPackets {
			err := sp.getValue()
			if err != nil {
				return err
			}
			p.value += sp.value
		}
	case 1:
		p.value = 1
		for _, sp := range p.subPackets {
			err := sp.getValue()
			if err != nil {
				return err
			}
			p.value *= sp.value
		}
	case 2:
		p.value = utils.Infinty
		for _, sp := range p.subPackets {
			if err := sp.getValue(); err != nil {
				return err
			} else if sp.value < p.value {
				p.value = sp.value
			}
		}
	case 3:
		p.value = 0
		for _, sp := range p.subPackets {
			if err := sp.getValue(); err != nil {
				return err
			} else if sp.value > p.value {
				p.value = sp.value
			}
		}
	case 4:
		break
	case 5:
		if len(p.subPackets) != 2 {
			return fmt.Errorf("expected 2 subpackets, got %v", p.subPackets)
		}
		err := p.subPackets[0].getValue()
		if err != nil {
			return err
		}
		err = p.subPackets[1].getValue()
		if err != nil {
			return err
		}
		p.value = 0
		if p.subPackets[0].value > p.subPackets[1].value {
			p.value = 1
		}
	case 6:
		if len(p.subPackets) != 2 {
			return fmt.Errorf("expected 2 subpackets, got %v", p.subPackets)
		}
		err := p.subPackets[0].getValue()
		if err != nil {
			return err
		}
		err = p.subPackets[1].getValue()
		if err != nil {
			return err
		}
		p.value = 0
		if p.subPackets[0].value < p.subPackets[1].value {
			p.value = 1
		}
	case 7:
		if len(p.subPackets) != 2 {
			return fmt.Errorf("expected 2 subpackets, got %v", p.subPackets)
		}
		err := p.subPackets[0].getValue()
		if err != nil {
			return err
		}
		err = p.subPackets[1].getValue()
		if err != nil {
			return err
		}
		p.value = 0
		if p.subPackets[0].value == p.subPackets[1].value {
			p.value = 1
		}
	default:
		return fmt.Errorf("got invalid typeID: %d", p.typeID)
	}
	return nil
}

func findSolutions(input string) (int, int, error) {
	bits, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	i := 0
	p := &Packet{value: -1}
	err = p.evaluatePacket(bits, &i)
	if err != nil {
		return -1, -1, err
	}
	err = p.getValue()
	if err != nil {
		return -1, -1, err
	}
	return p.sumVersions(), p.value, nil
}

func main() {
	input := utils.ReadFile()[0]
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
