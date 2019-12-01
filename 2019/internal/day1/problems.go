package day1

import (
	"strconv"
	"strings"
)

// Problem1 reads the input and calculates the required total fuel
// to launch all of the modules. Each line of input is the module Mass(m),
// required fuel = round(m/3) - 2
func Problem1() (int64, error) {
	var result int64
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			break
		}
		modMass, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return modMass, err
		}
		result += calcFuelRequirement(modMass)
	}

	return result, nil
}

// Problem2 requires calculating the total fuel for each module plus the
// additional fuel required to account for the added mass of fuel.
func Problem2() (int64, error) {
	var result int64
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			break
		}
		modMass, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return modMass, err
		}
		var modFuel int64
		mass := modMass
		for {
			addFuel := calcFuelRequirement(mass)
			if addFuel < 0 {
				break
			}
			modFuel += addFuel
			mass = addFuel
		}
		result += modFuel
	}

	return result, nil
}

func calcFuelRequirement(mass int64) int64 {
	return int64((mass / 3) - 2)
}
