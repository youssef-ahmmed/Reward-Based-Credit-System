package bridge

type WalletPort interface {
	AddToWallet(userID string, credits, points int) error
}
