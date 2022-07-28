// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dialect

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateDialect is the `createDialect` instruction.
type CreateDialect struct {
	DialectNonce *uint8
	Encrypted    *bool
	Scopes       *[2][2]bool

	// [0] = [WRITE, SIGNER] owner
	//
	// [1] = [] member0
	//
	// [2] = [] member1
	//
	// [3] = [WRITE] dialect
	//
	// [4] = [] rent
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateDialectInstructionBuilder creates a new `CreateDialect` instruction builder.
func NewCreateDialectInstructionBuilder() *CreateDialect {
	nd := &CreateDialect{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetDialectNonce sets the "dialectNonce" parameter.
func (inst *CreateDialect) SetDialectNonce(dialectNonce uint8) *CreateDialect {
	inst.DialectNonce = &dialectNonce
	return inst
}

// SetEncrypted sets the "encrypted" parameter.
func (inst *CreateDialect) SetEncrypted(encrypted bool) *CreateDialect {
	inst.Encrypted = &encrypted
	return inst
}

// SetScopes sets the "scopes" parameter.
func (inst *CreateDialect) SetScopes(scopes [2][2]bool) *CreateDialect {
	inst.Scopes = &scopes
	return inst
}

// SetOwnerAccount sets the "owner" account.
func (inst *CreateDialect) SetOwnerAccount(owner ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CreateDialect) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMember0Account sets the "member0" account.
func (inst *CreateDialect) SetMember0Account(member0 ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(member0)
	return inst
}

// GetMember0Account gets the "member0" account.
func (inst *CreateDialect) GetMember0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMember1Account sets the "member1" account.
func (inst *CreateDialect) SetMember1Account(member1 ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(member1)
	return inst
}

// GetMember1Account gets the "member1" account.
func (inst *CreateDialect) GetMember1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetDialectAccount sets the "dialect" account.
func (inst *CreateDialect) SetDialectAccount(dialect ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(dialect).WRITE()
	return inst
}

// GetDialectAccount gets the "dialect" account.
func (inst *CreateDialect) GetDialectAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateDialect) SetRentAccount(rent ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateDialect) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateDialect) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateDialect {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateDialect) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst CreateDialect) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateDialect,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateDialect) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateDialect) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DialectNonce == nil {
			return errors.New("DialectNonce parameter is not set")
		}
		if inst.Encrypted == nil {
			return errors.New("Encrypted parameter is not set")
		}
		if inst.Scopes == nil {
			return errors.New("Scopes parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Member0 is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Member1 is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Dialect is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateDialect) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateDialect")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DialectNonce", *inst.DialectNonce))
						paramsBranch.Child(ag_format.Param("   Encrypted", *inst.Encrypted))
						paramsBranch.Child(ag_format.Param("      Scopes", *inst.Scopes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      member0", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      member1", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      dialect", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj CreateDialect) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DialectNonce` param:
	err = encoder.Encode(obj.DialectNonce)
	if err != nil {
		return err
	}
	// Serialize `Encrypted` param:
	err = encoder.Encode(obj.Encrypted)
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
func (obj *CreateDialect) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DialectNonce`:
	err = decoder.Decode(&obj.DialectNonce)
	if err != nil {
		return err
	}
	// Deserialize `Encrypted`:
	err = decoder.Decode(&obj.Encrypted)
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

// NewCreateDialectInstruction declares a new CreateDialect instruction with the provided parameters and accounts.
func NewCreateDialectInstruction(
	// Parameters:
	dialectNonce uint8,
	encrypted bool,
	scopes [2][2]bool,
	// Accounts:
	owner ag_solanago.PublicKey,
	member0 ag_solanago.PublicKey,
	member1 ag_solanago.PublicKey,
	dialect ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateDialect {
	return NewCreateDialectInstructionBuilder().
		SetDialectNonce(dialectNonce).
		SetEncrypted(encrypted).
		SetScopes(scopes).
		SetOwnerAccount(owner).
		SetMember0Account(member0).
		SetMember1Account(member1).
		SetDialectAccount(dialect).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram)
}
