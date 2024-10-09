package helper

import "kenalbatik-be/internal/domain"

func GetNextExp(currentTier domain.UserTier) int {
	switch currentTier {
	case domain.TIER1:
		return 25
	case domain.TIER2:
		return 50
	case domain.TIER3:
		return 75
	case domain.TIER4:
		return 100
	case domain.TIER5:
		return 100
	default:
		return 0
	}
}