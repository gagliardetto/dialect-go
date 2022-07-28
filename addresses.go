package dialect

import (
	"bytes"
	"sort"

	"github.com/gagliardetto/solana-go"
)

var DEVNET = solana.MustPublicKeyFromBase58("2YFyZAg8rBtuvzFFiGvXwPHFAQJ2FXZoS7bYCKticpjk")
var MAINNET = solana.MustPublicKeyFromBase58("CeNUxGUsSeb5RuAGvaMLNx3tEZrpBwQqA7Gs99vMPCAb")

func sortPublicKeys(keys []solana.PublicKey) {
	// TODO: sort alphabetically?
	sort.Slice(keys, func(i, j int) bool {
		return bytes.Compare(keys[i][:], keys[j][:]) < 0
	})
}

func GetDialectThreadPDA(program solana.PublicKey, members solana.PublicKeySlice) (solana.PublicKey, uint8, error) {
	sortPublicKeys(members)
	seeds := [][]byte{
		[]byte("dialect"),
	}
	for _, member := range members {
		seeds = append(seeds, member.Bytes())
	}
	return solana.FindProgramAddress(seeds, program)
}
