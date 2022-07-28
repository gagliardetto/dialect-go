package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/dialect-go"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	. "github.com/gagliardetto/utilz"
)

func main() {
	dialectProgramAddress := dialect.DEVNET
	cluster := rpc.DevNet
	// NOTE: VERY IMPORTANT:
	dialect.SetProgramID(dialectProgramAddress)

	rpcClient := rpc.NewWithHeaders(
		cluster.RPC,
		map[string]string{},
	)

	wsClient, err := ws.Connect(context.Background(), cluster.WS)
	if err != nil {
		panic(err)
	}

	// NOTE: the sender must have some SOL on the account (at least 0.06 SOL)
	sender, err := solana.PrivateKeyFromSolanaKeygenFile(os.ExpandEnv("$HOME/.config/solana/id.json"))
	if err != nil {
		panic(err)
	}
	fmt.Println("sender public key:", sender.PublicKey())

	if false {
		// airdrop some SOL to the sender so it can send transactions:
		out, err := rpcClient.RequestAirdrop(
			context.TODO(),
			sender.PublicKey(),
			solana.LAMPORTS_PER_SOL*5,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("airdrop transaction signature:", out)

		res, _, err := waitForConfirmation(wsClient, out)
		if err != nil {
			panic(err)
		}
		spew.Dump(res)
	}

	// The person/entity to whom you're sending the message:
	receiver := solana.MustPublicKeyFromBase58("71hbcejgN7FdCB7w2TxQvogVqzLgvq2zv7vBdEgVpHH5")

	// dialectPDA is the chat thread between the two accounts.
	dialectPDA, dialectNonce, err := dialect.GetDialectThreadPDA(
		dialectProgramAddress,
		[]solana.PublicKey{
			sender.PublicKey(),
			receiver,
		})
	if err != nil {
		panic(err)
	}

	// The account containing the thread (if already existing)
	dialectData, err := rpcClient.GetAccountInfo(context.Background(), dialectPDA)
	if err != nil {
		if err == rpc.ErrNotFound {
			fmt.Println("chat thread account not found - creating it")
		} else {
			panic(err)
		}
	}

	if dialectData == nil {
		// Create dialect (chat thread) account:
		// https://github.com/dialectlabs/protocol/blob/525b8e34cb87fd168714f1c2dd4f60969c78d7e6/src/api/index.ts#L399
		var instructions []solana.Instruction
		var signers []solana.PrivateKey
		if true {
			encrypted := false
			scopes := [2][2]bool{
				{true, true},
				{false, true},
			}
			owner := sender.PublicKey()
			createDialectInstruction := dialect.NewCreateDialectInstructionBuilder().
				SetDialectNonce(dialectNonce).
				SetEncrypted(encrypted).
				SetScopes(scopes).
				SetOwnerAccount(owner).
				SetMember0Account(sender.PublicKey()).
				SetMember1Account(receiver).
				SetDialectAccount(dialectPDA).
				SetRentAccount(solana.SysVarRentPubkey).
				SetSystemProgramAccount(solana.SystemProgramID)

			instructions = append(instructions, createDialectInstruction.Build())

			signers = []solana.PrivateKey{sender}
		}

		sig, err := sendTransaction(
			cluster,
			rpcClient,
			sender,
			instructions,
			signers,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("createDialect transaction signature:", sig)
	}

	{
		// Send message:
		// https://github.com/dialectlabs/protocol/blob/525b8e34cb87fd168714f1c2dd4f60969c78d7e6/src/api/index.ts#L470
		var instructions []solana.Instruction
		var signers []solana.PrivateKey

		// TODO:
		// - get the dialect account data and parse it
		// - get dialect options from the data (encrypted, etc.)
		// - reproduce text encryption: https://github.com/dialectlabs/protocol/blob/525b8e34cb87fd168714f1c2dd4f60969c78d7e6/src/api/text-serde.ts
		{
			text := []byte("Hello world from Dialect: 🥰🔥")
			sendMessageInstruction := dialect.NewSendMessageInstructionBuilder().
				SetDialectNonce(dialectNonce).
				SetText(text).
				SetSenderAccount(sender.PublicKey()).
				SetDialectAccount(dialectPDA).
				SetRentAccount(solana.SysVarRentPubkey).
				SetSystemProgramAccount(solana.SystemProgramID)

			instructions = append(instructions, sendMessageInstruction.Build())
			signers = []solana.PrivateKey{sender}
		}

		sig, err := sendTransaction(
			cluster,
			rpcClient,
			sender,
			instructions,
			signers,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("sendMessage transaction signature:", sig)

	}
	Successf("Success")
}
