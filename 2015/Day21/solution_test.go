package main

import "testing"

// func Test_runFights(t *testing.T) {
// 	tests := []struct {
// 		name                string
// 		arg                 []string
// 		wantBoss            *cmb.Combatant
// 		wantCheapestVictory int
// 		wantDearestLoss     int
// 		wantErr             bool
// 	}{
// 		{
// 			name:    "returns an error if there is an error parsing the Boss",
// 			wantErr: true,
// 		},
// 		// don't have a good example, so just use the real advent of code questions since we know
// 		// we got the correct solution. It's too much to compare everything, so just make sure the
// 		// Boss is parsed and we will end up with the right Cost arrays
// 		{
// 			name: "runs fights against given boss",
// 			arg: []string{
// 				"Hit Points: 109",
// 				"Damage: 8",
// 				"Armor: 2",
// 			},
// 			wantBoss: &cmb.Combatant{
// 				HitPoints: 109,
// 				Damage:    8,
// 				Armour:    2,
// 			},
// 			wantCheapestVictory: 111,
// 			wantDearestLoss:     188,
// 			wantErr:             false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := runFights(tt.arg)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("runFights() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if tt.wantErr == true {
// 				return
// 			}
// 			if !reflect.DeepEqual(got.Boss, tt.wantBoss) {
// 				t.Errorf("runFights().Boss = %v, want %v", got.Boss, tt.wantBoss)
// 			}
// 			sort.Ints(got.SuccessfulCosts)
// 			if cheapestVictory := got.SuccessfulCosts[0]; cheapestVictory != tt.wantCheapestVictory {
// 				t.Errorf("runFights().CheapestVictory = %v, want %v", cheapestVictory, tt.wantCheapestVictory)
// 			}
// 			sort.Sort(sort.Reverse(sort.IntSlice(got.UnsuccessfulCosts)))
// 			if dearestLoss := got.UnsuccessfulCosts[0]; dearestLoss != tt.wantDearestLoss {
// 				t.Errorf("runFights().CheapestVictory = %v, want %v", dearestLoss, tt.wantDearestLoss)
// 			}
// 		})
// 	}
// }

func Test_runFights(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if there is an error parsing the Boss",
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		// don't have a good example, so just use the real advent of code questions since we know
		// we got the correct solution. It's too much to compare everything, so just make sure the
		// Boss is parsed and we will end up with the right Cost arrays
		{
			name: "runs fights against given boss",
			arg: []string{
				"Hit Points: 109",
				"Damage: 8",
				"Armor: 2",
			},
			want:    111,
			want1:   188,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := runFights(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("runFights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runFights() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("runFights() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
