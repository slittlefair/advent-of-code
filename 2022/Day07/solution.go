package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

type node struct {
	id     string
	parent *node
	size   int
}

type dirMap map[string]*node

func (dm dirMap) handleFileOrDir(line string, currentDir *node) error {
	if line[:3] == "dir" {
		// If the line starts with "dir" then we have a directory - add this to the map along with
		// its parent
		var name string
		_, err := fmt.Sscanf(line, "dir %s", &name)
		if err != nil {
			return fmt.Errorf("addFileOrDir on %s: %w", line, err)
		}
		newSys := &node{
			id:     fmt.Sprintf("%s/%s", currentDir.id, name),
			parent: currentDir,
		}
		dm[newSys.id] = newSys
	} else {
		// Otherwise we have a file, add its size to its parent, that dir's parent and so on
		var n string
		var size int
		_, err := fmt.Sscanf(line, "%d %s", &size, &n)
		if err != nil {
			return fmt.Errorf("addFileOrDir on %s: %w", line, err)
		}
		currentDir.size += size
		for currentDir.parent != nil {
			currentDir.parent.size += size
			currentDir = currentDir.parent
		}
	}
	return nil
}

func (dm dirMap) getSub100000Dirs() int {
	sum := 0
	for _, dir := range dm {
		if dir.size <= 100000 {
			sum += dir.size
		}
	}
	return sum
}

func parseInput(input []string) (dirMap, error) {
	// Set up the root node and add it to a map of all nodes
	currentDir := &node{
		id: "/",
	}
	dm := dirMap{
		"/": currentDir,
	}
	for i := 1; i < len(input); i++ {
		line := input[i]
		if string(line[0]) != "$" {
			// If the line is not a command, it must be listing a directory or file
			err := dm.handleFileOrDir(line, currentDir)
			if err != nil {
				return nil, err
			}
		} else if line == "$ cd .." {
			// Move up to the parent directory
			currentDir = currentDir.parent
		} else if line == "$ ls" {
			// Don't do anything, we handle following list in the above
			continue
		} else {
			// If none of the above we must be going into one of the child directories
			var dir string
			_, err := fmt.Sscanf(line, "$ cd %s", &dir)
			if err != nil {
				return nil, fmt.Errorf("parseInput on %s: %w", line, err)
			}
			currentDir = dm[fmt.Sprintf("%s/%s", currentDir.id, dir)]
		}
	}
	return dm, nil
}

func (dm dirMap) findDirToDelete() int {
	minSpaceToDelete := 30000000 - (70000000 - dm["/"].size)
	sizeToDelete := maths.Infinity
	for _, dir := range dm {
		if dir.size >= minSpaceToDelete && dir.size < sizeToDelete {
			sizeToDelete = dir.size
		}
	}
	return sizeToDelete
}

func main() {
	input := file.Read()
	dm, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", dm.getSub100000Dirs())
	fmt.Println("Part 2:", dm.findDirToDelete())
}
