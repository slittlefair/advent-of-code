package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

type sysType int

const (
	sysFile sysType = iota
	sysDir
)

type sys struct {
	id       string
	parent   *sys
	children []*sys
	size     int
	sysType  sysType
}

type sysMap map[string]*sys

func (sm sysMap) addFileOrDir(line string, currentDir *sys) error {
	if line[:3] == "dir" {
		var name string
		_, err := fmt.Sscanf(line, "dir %s", &name)
		if err != nil {
			return fmt.Errorf("addFileOrDir on %s: %w", line, err)
		}
		newSys := &sys{
			id:      fmt.Sprintf("%s/%s", currentDir.id, name),
			parent:  currentDir,
			sysType: sysDir,
		}
		currentDir.children = append(currentDir.children, newSys)
		sm[newSys.id] = newSys
		currentDir = newSys
	} else {
		var name string
		var size int
		_, err := fmt.Sscanf(line, "%d %s", &size, &name)
		if err != nil {
			return fmt.Errorf("addFileOrDir on %s: %w", line, err)
		}
		newSys := &sys{
			id:      fmt.Sprintf("%s/%s", currentDir.id, name),
			parent:  currentDir,
			size:    size,
			sysType: sysFile,
		}
		currentDir.children = append(currentDir.children, newSys)
		currentDir.size += size
		for currentDir.parent != nil {
			currentDir.parent.size += size
			currentDir = currentDir.parent
		}
		sm[newSys.id] = newSys
	}
	return nil
}

func (sm sysMap) getSub100000Dirs() int {
	sum := 0
	for _, sys := range sm {
		if sys.sysType == sysDir && sys.size <= 100000 {
			sum += sys.size
		}
	}
	return sum
}

func parseInput(input []string) (sysMap, error) {
	currentDir := &sys{
		id:      "/",
		sysType: sysDir,
	}
	sm := sysMap{
		"/": currentDir,
	}
	for i := 1; i < len(input); i++ {
		line := input[i]
		if string(line[0]) != "$" {
			err := sm.addFileOrDir(line, currentDir)
			if err != nil {
				return nil, err
			}
		} else if line == "$ cd .." {
			currentDir = currentDir.parent
		} else if line == "$ ls" {
			continue
		} else {
			var dir string
			_, err := fmt.Sscanf(line, "$ cd %s", &dir)
			if err != nil {
				return nil, fmt.Errorf("parseInput on %s: %w", line, err)
			}
			currentDir = sm[fmt.Sprintf("%s/%s", currentDir.id, dir)]
		}
	}
	return sm, nil
}

func (sm sysMap) findDirToDelete() int {
	minSpaceToDelete := 30000000 - (70000000 - sm["/"].size)
	sizeToDelete := maths.Infinity
	for _, sys := range sm {
		if sys.sysType == sysDir && sys.size >= minSpaceToDelete && sys.size < sizeToDelete {
			sizeToDelete = sys.size
		}
	}
	return sizeToDelete
}

func main() {
	input := file.Read()
	sm, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sm.getSub100000Dirs())
	fmt.Println("Part 2:", sm.findDirToDelete())
}
