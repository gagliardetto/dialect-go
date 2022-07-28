// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dialect

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Dialect"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_CreateMetadata = ag_binary.TypeID([8]byte{30, 35, 117, 134, 196, 139, 44, 25})

	Instruction_CloseMetadata = ag_binary.TypeID([8]byte{10, 220, 196, 138, 19, 60, 204, 130})

	Instruction_CreateDialect = ag_binary.TypeID([8]byte{220, 72, 83, 96, 55, 126, 232, 148})

	Instruction_CloseDialect = ag_binary.TypeID([8]byte{66, 28, 68, 124, 204, 27, 39, 116})

	Instruction_SubscribeUser = ag_binary.TypeID([8]byte{102, 17, 164, 232, 26, 65, 43, 134})

	Instruction_SendMessage = ag_binary.TypeID([8]byte{57, 40, 34, 178, 189, 10, 65, 26})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_CreateMetadata:
		return "CreateMetadata"
	case Instruction_CloseMetadata:
		return "CloseMetadata"
	case Instruction_CreateDialect:
		return "CreateDialect"
	case Instruction_CloseDialect:
		return "CloseDialect"
	case Instruction_SubscribeUser:
		return "SubscribeUser"
	case Instruction_SendMessage:
		return "SendMessage"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"create_metadata", (*CreateMetadata)(nil),
		},
		{
			"close_metadata", (*CloseMetadata)(nil),
		},
		{
			"create_dialect", (*CreateDialect)(nil),
		},
		{
			"close_dialect", (*CloseDialect)(nil),
		},
		{
			"subscribe_user", (*SubscribeUser)(nil),
		},
		{
			"send_message", (*SendMessage)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
