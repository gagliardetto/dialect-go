package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	sendAndConfirmTransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func waitForConfirmation(
	wsClient *ws.Client,
	sig solana.Signature,
) (*ws.SignatureResult, solana.Signature, error) {
	sub, err := wsClient.SignatureSubscribe(
		sig,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return nil, sig, err
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			return nil, sig, err
		}
		if got.Value.Err != nil {
			return nil, sig, fmt.Errorf("transaction confirmation failed: %v", got.Value.Err)
		} else {
			return got, sig, nil
		}
	}
}

func sendTransaction(
	cluster rpc.Cluster,
	client *rpc.Client,
	wallet solana.PrivateKey,
	instructions []solana.Instruction,
	signers []solana.PrivateKey,
) (solana.Signature, error) {

	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(wallet.PublicKey()),
	)

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, err
	}
	spew.Dump(tx)

	wsClient, err := ws.Connect(context.Background(), cluster.WS)
	if err != nil {
		return solana.Signature{}, err
	}

	return sendAndConfirmTransaction.SendAndConfirmTransactionWithOpts(
		context.TODO(),
		client,
		wsClient,
		tx,
		rpc.TransactionOpts{
			SkipPreflight: true,
		},
	)
}
