package level

func CalculateExperience(n int) int {
	switch n {
	case 1:
		return 0
	case 2:
		return 200
	case 3:
		return 1500
	case 4:
		return 4500
	case 5:
		return 10800
	case 6:
		return 28800
	default:
		return 0
	}
}

func CalculateLevel(experience int) int {
	if experience >= 0 && experience < 200 {
		return 1
	} else if experience >= 200 && experience < 1500 {
		return 2
	} else if experience >= 1500 && experience < 4500 {
		return 3
	} else if experience >= 4500 && experience < 10800 {
		return 4
	} else if experience >= 10800 && experience < 28800 {
		return 5
	} else {
		return 6
	}
}
