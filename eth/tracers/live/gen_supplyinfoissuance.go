// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package live

import (
	"encoding/json"
	"math/big"

	"github.com/qiruos/go-ethereum/common/hexutil"
)

var _ = (*supplyInfoIssuanceMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s supplyInfoIssuance) MarshalJSON() ([]byte, error) {
	type supplyInfoIssuance struct {
		GenesisAlloc *hexutil.Big `json:"genesisAlloc,omitempty"`
		Reward       *hexutil.Big `json:"reward,omitempty"`
		Withdrawals  *hexutil.Big `json:"withdrawals,omitempty"`
	}
	var enc supplyInfoIssuance
	enc.GenesisAlloc = (*hexutil.Big)(s.GenesisAlloc)
	enc.Reward = (*hexutil.Big)(s.Reward)
	enc.Withdrawals = (*hexutil.Big)(s.Withdrawals)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *supplyInfoIssuance) UnmarshalJSON(input []byte) error {
	type supplyInfoIssuance struct {
		GenesisAlloc *hexutil.Big `json:"genesisAlloc,omitempty"`
		Reward       *hexutil.Big `json:"reward,omitempty"`
		Withdrawals  *hexutil.Big `json:"withdrawals,omitempty"`
	}
	var dec supplyInfoIssuance
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.GenesisAlloc != nil {
		s.GenesisAlloc = (*big.Int)(dec.GenesisAlloc)
	}
	if dec.Reward != nil {
		s.Reward = (*big.Int)(dec.Reward)
	}
	if dec.Withdrawals != nil {
		s.Withdrawals = (*big.Int)(dec.Withdrawals)
	}
	return nil
}
