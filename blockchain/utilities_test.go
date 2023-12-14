package blockchain

import (
	"testing"

	"github.com/AwespireTech/RXCA-Backend/config"
)

func TestMain(m *testing.M) {
	err := Init("https://eth-sepolia.g.alchemy.com/v2/s65QiLyN-74IJZtUJgtWJiZ9gGzUfxOm")
	if err != nil {
		panic(err)
	}
	config.CONTRACT_ADDRESS = "0xcdF367bb783bC7C3681df313364fdf9b1E82A7aD"
	m.Run()
}
func TestDecodeMintTransaction(t *testing.T) {
	t.Log("DecodeMintTransaction")
	txHash := "0x81d3d31e62a69929b777f4dce33eb5a6067f0e4dd5f3f65aff60e38324889ab5"
	txTarget, tid, err := DecodeMintTransaction(txHash)
	if err != nil {
		t.Error(err)
	}
	t.Logf("txTarget: %s, tid: %d", txTarget, tid)
	//Test if given txHash is not a mint transaction
	txHash = "0xd79d028836333e3df2ad4dbb4af75db8a9d7e8d786b2771c4b02e6caf2d700c7"
	_, _, err = DecodeMintTransaction(txHash)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestDecodeBurnTransaction(t *testing.T) {
	t.Log("DecodeBurnTransaction")
	txHash := "0xd79d028836333e3df2ad4dbb4af75db8a9d7e8d786b2771c4b02e6caf2d700c7"
	txTarget, tid, err := DecodeBurnTransaction(txHash)
	if err != nil {
		t.Error(err)
	}
	t.Logf("txTarget: %s, tid: %d", txTarget, tid)
	//Test if given txHash is not a burn transaction
	txHash = "0x81d3d31e62a69929b777f4dce33eb5a6067f0e4dd5f3f65aff60e38324889ab5"
	_, _, err = DecodeBurnTransaction(txHash)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
