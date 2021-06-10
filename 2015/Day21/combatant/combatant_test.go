package combatant

import "testing"

func TestCombatant_IsDead(t *testing.T) {
	tests := []struct {
		name      string
		HitPoints int
		want      bool
	}{
		{
			name:      "returns true if combatant has negative hit point",
			HitPoints: -1,
			want:      true,
		},
		{
			name:      "returns true if combatant has zero hit point",
			HitPoints: 0,
			want:      true,
		},
		{
			name:      "returns false if combatant has positive hit point",
			HitPoints: 1,
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Combatant{
				HitPoints: tt.HitPoints,
			}
			if got := c.IsDead(); got != tt.want {
				t.Errorf("Combatant.IsDead() = %v, want %v", got, tt.want)
			}
		})
	}
}
