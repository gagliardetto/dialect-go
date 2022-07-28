// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dialect

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type CyclicByteBuffer struct {
	ReadOffset  uint16
	WriteOffset uint16
	ItemsCount  uint16
	Buffer      [8192]uint8
}

func (obj CyclicByteBuffer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ReadOffset` param:
	err = encoder.Encode(obj.ReadOffset)
	if err != nil {
		return err
	}
	// Serialize `WriteOffset` param:
	err = encoder.Encode(obj.WriteOffset)
	if err != nil {
		return err
	}
	// Serialize `ItemsCount` param:
	err = encoder.Encode(obj.ItemsCount)
	if err != nil {
		return err
	}
	// Serialize `Buffer` param:
	err = encoder.Encode(obj.Buffer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CyclicByteBuffer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ReadOffset`:
	err = decoder.Decode(&obj.ReadOffset)
	if err != nil {
		return err
	}
	// Deserialize `WriteOffset`:
	err = decoder.Decode(&obj.WriteOffset)
	if err != nil {
		return err
	}
	// Deserialize `ItemsCount`:
	err = decoder.Decode(&obj.ItemsCount)
	if err != nil {
		return err
	}
	// Deserialize `Buffer`:
	err = decoder.Decode(&obj.Buffer)
	if err != nil {
		return err
	}
	return nil
}

type Subscription struct {
	Pubkey  ag_solanago.PublicKey
	Enabled bool
}

func (obj Subscription) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Pubkey` param:
	err = encoder.Encode(obj.Pubkey)
	if err != nil {
		return err
	}
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Subscription) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Pubkey`:
	err = decoder.Decode(&obj.Pubkey)
	if err != nil {
		return err
	}
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

type Member struct {
	PublicKey ag_solanago.PublicKey
	Scopes    [2]bool
}

func (obj Member) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PublicKey` param:
	err = encoder.Encode(obj.PublicKey)
	if err != nil {
		return err
	}
	// Serialize `Scopes` param:
	err = encoder.Encode(obj.Scopes)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Member) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PublicKey`:
	err = decoder.Decode(&obj.PublicKey)
	if err != nil {
		return err
	}
	// Deserialize `Scopes`:
	err = decoder.Decode(&obj.Scopes)
	if err != nil {
		return err
	}
	return nil
}

type ErrorCode ag_binary.BorshEnum

const (
	ErrorCodeDialectOwnerIsNotAdmin ErrorCode = iota
)

func (value ErrorCode) String() string {
	switch value {
	case ErrorCodeDialectOwnerIsNotAdmin:
		return "DialectOwnerIsNotAdmin"
	default:
		return ""
	}
}