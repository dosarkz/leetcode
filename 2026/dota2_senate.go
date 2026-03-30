package _026

func predictPartyVictory(senate string) string {
	var radiant, dire []int

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		i, j := radiant[0], dire[0]
		radiant = radiant[1:]
		dire = dire[1:]

		if i < j {
			radiant = append(radiant, i+len(senate))
		} else {
			dire = append(dire, i+len(senate))
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	}

	return "Dire"
}
