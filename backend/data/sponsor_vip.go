package data

// EffectiveSponsorVipLevel 始终返回 VIP 有效状态，解锁全部功能。
func EffectiveSponsorVipLevel() (level int, active bool) {
	return 100, true
}
