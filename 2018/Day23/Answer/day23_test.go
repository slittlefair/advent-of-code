package day23_test

// import (
// 	"testing"

// 	"bitbucket.org/thezeez/advent-of-code-2018/pkg/day23"
// )

// func TestStrongestReachable(t *testing.T) {
// 	tests := []struct {
// 		in  []string
// 		out int
// 	}{
// 		{[]string{
// 			"pos=<0,0,0>, r=4",
// 			"pos=<1,0,0>, r=1",
// 			"pos=<4,0,0>, r=3",
// 			"pos=<0,2,0>, r=1",
// 			"pos=<0,5,0>, r=3",
// 			"pos=<0,0,3>, r=1",
// 			"pos=<1,1,1>, r=1",
// 			"pos=<1,1,2>, r=1",
// 			"pos=<1,3,1>, r=1",
// 		}, 7},
// 	}

// 	for _, test := range tests {
// 		actual := day23.StrongestReachable(day23.NewBots(test.in))
// 		if actual != test.out {
// 			t.Errorf("StrongestReachable(%q) => %d, want %d", test.in, actual, test.out)
// 		}
// 	}
// }

// func TestClosestSuccess(t *testing.T) {
// 	tests := []struct {
// 		in  []string
// 		out int
// 	}{
// 		{[]string{
// 			"pos=<10,12,12>, r=2",
// 			"pos=<12,14,12>, r=2",
// 			"pos=<16,12,12>, r=4",
// 			"pos=<14,14,14>, r=6",
// 			"pos=<50,50,50>, r=200",
// 			"pos=<10,10,10>, r=5",
// 		}, 36},
// 	}

// 	for _, test := range tests {
// 		actual := day23.ClosestSuccess(day23.NewBots(test.in))
// 		if actual != test.out {
// 			t.Errorf("ClosestSuccess(%q) => %d, want %d", test.in, actual, test.out)
// 		}
// 	}
// }
