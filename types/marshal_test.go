package types

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/avinashpandit/trustwallet-golibs/coin"
	"github.com/avinashpandit/trustwallet-golibs/mock"
	"github.com/stretchr/testify/assert"
)

var (
	txJSON, _ = mock.JsonStringFromFilePath("mocks/nim_tx.json")
	txModel   = Tx{
		ID:     "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
		Coin:   coin.NIMIQ,
		From:   "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
		To:     "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
		Fee:    "138",
		Date:   1548954343,
		Block:  419040,
		Status: StatusCompleted,
		Meta: &Transfer{
			Value: "5004160",
		},
	}
)

func TestTx_UnmarshalJSON(t *testing.T) {
	// Expect to get txModel, but with type set
	expected := txModel
	expected.Type = TxTransfer

	// Unmarshal source
	var got Tx
	err := json.Unmarshal([]byte(txJSON), &got)
	if err != nil {
		t.Fatal(err)
	}

	// Compare got and expected
	if !reflect.DeepEqual(expected, got) {
		t.Error("txs not equal")
	}
}

func TestTx_MarshalJSON(t *testing.T) {
	// Input is txModel without type
	input := txModel

	// Marshal transaction
	got, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	// After marshal, the type should be set
	if input.Type == "" {
		t.Fatal("type was not set")
	} else if input.Type != TxTransfer {
		t.Fatal("wrong type set")
	}

	// Compare expected and output JSON
	bytes.Equal(got, []byte(txJSON))
}

func TestSortTxPage(t *testing.T) {
	tests := []struct {
		name string
		page Txs
		want Txs
	}{
		{"test sort 1", Txs{{Date: 5}, {Date: 2}, {Date: 1}, {Date: 4}, {Date: 3}}, Txs{{Date: 5}, {Date: 4}, {Date: 3}, {Date: 2}, {Date: 1}}},
		{"test sort 2", Txs{{Date: 100}, {Date: 2}, {Date: 33}, {Date: 409}}, Txs{{Date: 409}, {Date: 100}, {Date: 33}, {Date: 2}}},
		{"test sort 3", Txs{{Date: 100}, {Date: 2}, {Date: 100}}, Txs{{Date: 100}, {Date: 100}, {Date: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.page)
			assert.Equal(t, tt.want, tt.page)
		})
	}
}
