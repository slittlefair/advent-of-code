package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var masterDisk = Disk{
	0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9,
}

func Test_parseInput(t *testing.T) {
	t.Run("returns an error if the input line contains a non-number", func(t *testing.T) {
		input := []string{"2333133121414131a02"}

		disk, err := parseInput(input)
		assert.Nil(t, disk)
		assert.Error(t, err)
	})

	t.Run("returns a Disk for a given input", func(t *testing.T) {
		input := []string{"2333133121414131402"}

		disk, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, masterDisk, disk)
	})
}

func TestDisk_copy(t *testing.T) {
	t.Run("returns a copy of the disk being called on", func(t *testing.T) {
		disk := masterDisk.copy()
		assert.Equal(t, masterDisk, disk)
	})

	t.Run("returns a copy of the disk being called on that doesn't change when the original does", func(t *testing.T) {
		disk := masterDisk.copy()
		assert.Equal(t, masterDisk, disk)
		masterDisk[2] = 100
		assert.NotEqual(t, masterDisk, disk)
		// revert the change to masterDisk
		masterDisk[2] = -1
	})
}

func TestDisk_compact(t *testing.T) {
	t.Run("compacts a disk according to part1 rules", func(t *testing.T) {
		disk := masterDisk.copy()
		disk.compact()

		want := Disk{
			0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}

		assert.Equal(t, want, disk)
	})
}

func TestDisk_compact2(t *testing.T) {
	t.Run("compacts a disk according to part2 rules", func(t *testing.T) {
		disk := masterDisk.copy()
		disk.compact2()

		want := Disk{
			0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1,
		}

		assert.Equal(t, want, disk)
	})
}

func TestDisk_calculateChecksum(t *testing.T) {
	t.Run("returns a checksum for a given disk, 1", func(t *testing.T) {
		disk := Disk{
			0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}
		got := disk.calculateChecksum()
		assert.Equal(t, 1928, got)
	})

	t.Run("returns a checksum for a given disk, 1", func(t *testing.T) {
		disk := Disk{
			0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1,
		}
		got := disk.calculateChecksum()
		assert.Equal(t, 2858, got)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns an error if the input doesn't contains a non-number", func(t *testing.T) {
		input := []string{"23331,3121414131402"}
		_, _, err := findSolutions(input)
		assert.Error(t, err)
	})

	t.Run("returns correct solutions to parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{"2333133121414131402"}
		part1, part2, err := findSolutions(input)
		assert.NoError(t, err)
		assert.Equal(t, 1928, part1)
		assert.Equal(t, 2858, part2)
	})
}
