package day15

func Combat(input []string) int {
	cave := NewCave(input, defaultPower)
	for i := 1; true; i++ {
		hp, combat := cave.Status()

		if !combat {
			return (i - 1) * hp
		}

		if cleanRound, _ := cave.Tick(false); !cleanRound {
			i--
		}
	}
	return -1
}

func CheatingElves(input []string) int {
	elfDied := true
	for power := 4; elfDied; power++ {
		cave := NewCave(input, power)
		for i := 1; true; i++ {
			hp, combat := cave.Status()

			if !combat {
				return (i - 1) * hp
			}

			var cleanRound bool
			cleanRound, elfDied = cave.Tick(true)
			if elfDied {
				break
			}
			if !cleanRound {
				i--
			}
		}
	}
	return -1
}
