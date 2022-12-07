package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var root = &node{
	id:   "/",
	size: 48381165,
}
var a = &node{
	id:     "//a",
	parent: root,
	size:   94853,
}
var d = &node{
	id:     "//d",
	parent: root,
	size:   24933642,
}
var ae = &node{
	id:     "//a/e",
	parent: a,
	size:   584,
}
var aocExampleDM = dirMap{
	"/":     root,
	"//a":   a,
	"//d":   d,
	"//a/e": ae,
}

func TestHandleFileOrDir(t *testing.T) {
	t.Run("returns an error if a directory the line isn't in the right format", func(t *testing.T) {
		currentDir := &node{
			id: "/",
		}
		dm := dirMap{
			"/": currentDir,
		}
		err := dm.handleFileOrDir("directory name", currentDir)
		assert.Error(t, err)
	})

	t.Run("adds a directory to a root only map", func(t *testing.T) {
		currentDir := &node{
			id: "/",
		}
		dm := dirMap{
			"/": currentDir,
		}
		err := dm.handleFileOrDir("dir a", currentDir)
		assert.NoError(t, err)
		want := dirMap{
			"/": &node{
				id: "/",
			},
			"//a": &node{
				id: "//a",
				parent: &node{
					id: "/",
				},
			},
		}
		assert.Equal(t, want, dm)
	})

	t.Run("adds a directory to a fuller map", func(t *testing.T) {
		currentDir := &node{
			id: "//a",
			parent: &node{
				id: "/",
			},
		}
		dm := dirMap{
			"/": &node{
				id: "/",
			},
			"//a": &node{
				id: "//a",
				parent: &node{
					id: "/",
				},
			},
			"//d": &node{
				id: "//d",
				parent: &node{
					id: "/",
				},
			},
		}
		err := dm.handleFileOrDir("dir e", currentDir)
		assert.NoError(t, err)
		want := dirMap{
			"/": &node{
				id: "/",
			},
			"//a": &node{
				id: "//a",
				parent: &node{
					id: "/",
				},
			},
			"//d": &node{
				id: "//d",
				parent: &node{
					id: "/",
				},
			},
			"//a/e": &node{
				id: "//a/e",
				parent: &node{
					id: "//a",
					parent: &node{
						id: "/",
					},
				},
			},
		}
		assert.Equal(t, want, dm)
	})

	t.Run("returns an error if a non directory line is not formatted correctly", func(t *testing.T) {
		currentDir := &node{
			id: "/",
		}
		dm := dirMap{
			"/": currentDir,
		}
		err := dm.handleFileOrDir("file of size 345", currentDir)
		assert.Error(t, err)
	})

	t.Run("adds a file to a map and updates parent sizes", func(t *testing.T) {
		root := &node{
			id: "/",
		}
		a := &node{
			id:     "//a",
			parent: root,
		}
		d := &node{
			id:     "//d",
			parent: root,
		}
		ae := &node{
			id:     "//a/e",
			parent: a,
		}
		dm := dirMap{
			"/":     root,
			"//a":   a,
			"//d":   d,
			"//a/e": ae,
		}
		currentDir := ae
		err := dm.handleFileOrDir("234 e.txt", currentDir)
		assert.NoError(t, err)
		want := dirMap{
			"/": &node{
				id:   "/",
				size: 234,
			},
			"//a": &node{
				id: "//a",
				parent: &node{
					id:   "/",
					size: 234,
				},
				size: 234,
			},
			"//d": &node{
				id: "//d",
				parent: &node{
					id:   "/",
					size: 234,
				},
			},
			"//a/e": &node{
				id: "//a/e",
				parent: &node{
					id: "//a",
					parent: &node{
						id:   "/",
						size: 234,
					},
					size: 234,
				},
				size: 234,
			},
		}
		assert.Equal(t, want, dm)
	})

	t.Run("adds a file to a map and updates parent sizes that already have some", func(t *testing.T) {
		root := &node{
			id:   "/",
			size: 1643,
		}
		a := &node{
			id:     "//a",
			parent: root,
			size:   456,
		}
		d := &node{
			id:     "//d",
			parent: root,
			size:   200,
		}
		ae := &node{
			id:     "//a/e",
			parent: a,
			size:   111,
		}
		dm := dirMap{
			"/":     root,
			"//a":   a,
			"//d":   d,
			"//a/e": ae,
		}
		currentDir := ae
		err := dm.handleFileOrDir("234 e.txt", currentDir)
		assert.NoError(t, err)
		want := dirMap{
			"/": &node{
				id:   "/",
				size: 1877,
			},
			"//a": &node{
				id: "//a",
				parent: &node{
					id:   "/",
					size: 1877,
				},
				size: 690,
			},
			"//d": &node{
				id: "//d",
				parent: &node{
					id:   "/",
					size: 1877,
				},
				size: 200,
			},
			"//a/e": &node{
				id: "//a/e",
				parent: &node{
					id: "//a",
					parent: &node{
						id:   "/",
						size: 1877,
					},
					size: 690,
				},
				size: 345,
			},
		}
		assert.Equal(t, want, dm)
	})
}

func TestParseInput(t *testing.T) {
	t.Run("returns an error if an input line contains an error for a dir line", func(t *testing.T) {
		_, err := parseInput([]string{
			"$ ls",
			"dir a",
			"$ cd a",
			"$ ls",
			"234 e.txt",
			"directory w",
			"$ cd ..",
		})
		assert.Error(t, err)
	})

	t.Run("returns an error if an input line contains an error for a cd dir line", func(t *testing.T) {
		_, err := parseInput([]string{
			"$ ls",
			"dir a",
			"$cd a",
			"$ ls",
			"234 e.txt",
			"directory w",
			"$ cd ..",
		})
		assert.Error(t, err)
	})

	t.Run("returns a populated directory map, advent of code example", func(t *testing.T) {
		input := []string{
			"$ cd /",
			"$ ls",
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
			"$ cd a",
			"$ ls",
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
			"$ cd e",
			"$ ls",
			"584 i",
			"$ cd ..",
			"$ cd ..",
			"$ cd d",
			"$ ls",
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k",
		}
		dm, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, aocExampleDM, dm)
	})
}

func TestGetSub100000Dirs(t *testing.T) {
	t.Run("returns sum of directory sizes less than 100000, advent of code example", func(t *testing.T) {
		got := aocExampleDM.getSub100000Dirs()
		assert.Equal(t, 95437, got)
	})
}

func TestFindDirToDelete(t *testing.T) {
	t.Run("finds the smallest file that can be deleted, advent of code example", func(t *testing.T) {
		got := aocExampleDM.findDirToDelete()
		assert.Equal(t, 24933642, got)
	})
}
