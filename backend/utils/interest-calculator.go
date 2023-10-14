package utils

func SimpleInterestCalculator(principal float64, rate float64, loanTime float64) float64 {

	SiNumerator := principal * rate * loanTime
	SiDenominator := 100

	SimpleInterest := SiNumerator / float64(SiDenominator)

	return SimpleInterest
}
