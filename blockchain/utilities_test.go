package blockchain

import "testing"

func TestMain(m *testing.M) {
	Init("https://eth-sepolia.g.alchemy.com/v2/s65QiLyN-74IJZtUJgtWJiZ9gGzUfxOm")
}
func TestDecodeMintTransaction(t *testing.T) {
	t.Log("DecodeMintTransaction")
	txHash := "0x81d3d31e62a69929b777f4dce33eb5a6067f0e4dd5f3f65aff60e38324889ab5"
	txTarget, tid, err := DecodeMintTransaction(txHash)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(txTarget, tid)

}
