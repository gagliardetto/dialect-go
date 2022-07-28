// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dialect

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type MetadataAccount struct {
	User          ag_solanago.PublicKey
	Subscriptions [32]Subscription
}

var MetadataAccountDiscriminator = [8]byte{32, 224, 226, 224, 77, 64, 109, 234}

func (obj MetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MetadataAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Subscriptions` param:
	err = encoder.Encode(obj.Subscriptions)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MetadataAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[32 224 226 224 77 64 109 234]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Subscriptions`:
	err = decoder.Decode(&obj.Subscriptions)
	if err != nil {
		return err
	}
	return nil
}

type DialectAccount struct {
	Members              [2]Member
	Messages             CyclicByteBuffer
	LastMessageTimestamp uint32
	Encrypted            bool
}

var DialectAccountDiscriminator = [8]byte{157, 38, 120, 189, 93, 204, 119, 18}

func (obj DialectAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DialectAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Members` param:
	err = encoder.Encode(obj.Members)
	if err != nil {
		return err
	}
	// Serialize `Messages` param:
	err = encoder.Encode(obj.Messages)
	if err != nil {
		return err
	}
	// Serialize `LastMessageTimestamp` param:
	err = encoder.Encode(obj.LastMessageTimestamp)
	if err != nil {
		return err
	}
	// Serialize `Encrypted` param:
	err = encoder.Encode(obj.Encrypted)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DialectAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DialectAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[157 38 120 189 93 204 119 18]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Members`:
	err = decoder.Decode(&obj.Members)
	if err != nil {
		return err
	}
	// Deserialize `Messages`:
	err = decoder.Decode(&obj.Messages)
	if err != nil {
		return err
	}
	// Deserialize `LastMessageTimestamp`:
	err = decoder.Decode(&obj.LastMessageTimestamp)
	if err != nil {
		return err
	}
	// Deserialize `Encrypted`:
	err = decoder.Decode(&obj.Encrypted)
	if err != nil {
		return err
	}
	return nil
}
