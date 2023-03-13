package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	ans := 0
	for i := range energy {
		if initialEnergy <= energy[i] {
			ans += energy[i] - initialEnergy + 1
			initialEnergy = energy[i] + 1
		}
		if initialExperience <= experience[i] {
			ans += experience[i] - initialExperience + 1
			initialExperience = experience[i] + 1
		}
		initialExperience += experience[i]
		initialEnergy -= energy[i]
	}
	return ans
}
