// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package consensus implements different Ethereum consensus engines.
package consensus

import (
	"math/big"

	"github.com/holiman/uint256"

	"github.com/ledgerwatch/erigon-lib/log/v3"

	"github.com/ledgerwatch/erigon-lib/chain"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/core/state"
	"github.com/ledgerwatch/erigon/core/tracing"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/rlp"
	"github.com/ledgerwatch/erigon/rpc"
)

// ChainHeaderReader defines a small collection of methods needed to access the local
// blockchain during header verification.
//
//go:generate mockgen -typed=true -destination=./chain_header_reader_mock.go -package=consensus . ChainHeaderReader
type ChainHeaderReader interface {
	// Config retrieves the blockchain's chain configuration.
	Config() *chain.Config

	// CurrentHeader retrieves the current header from the local chain.
	CurrentHeader() *types.Header

	// CurrentFinalizedHeader retrieves the current finalized header from the local chain.
	CurrentFinalizedHeader() *types.Header

	// CurrentSafeHeader retrieves the current safe header from the local chain.
	CurrentSafeHeader() *types.Header

	// GetHeader retrieves a block header from the database by hash and number.
	GetHeader(hash libcommon.Hash, number uint64) *types.Header

	// GetHeaderByNumber retrieves a block header from the database by number.
	GetHeaderByNumber(number uint64) *types.Header

	// GetHeaderByHash retrieves a block header from the database by its hash.
	GetHeaderByHash(hash libcommon.Hash) *types.Header

	// GetTd retrieves the total difficulty from the database by hash and number.
	GetTd(hash libcommon.Hash, number uint64) *big.Int

	// Number of blocks frozen in the block snapshots
	FrozenBlocks() uint64
	FrozenBorBlocks() uint64

	// Byte string representation of a bor span with given ID
	BorSpan(spanId uint64) []byte
}

// ChainReader defines a small collection of methods needed to access the local
// blockchain during header and/or uncle verification.
type ChainReader interface {
	ChainHeaderReader

	// GetBlock retrieves a block from the database by hash and number.
	GetBlock(hash libcommon.Hash, number uint64) *types.Block

	HasBlock(hash libcommon.Hash, number uint64) bool

	BorEventsByBlock(hash libcommon.Hash, number uint64) []rlp.RawValue
	BorStartEventID(hash libcommon.Hash, number uint64) uint64
}

type SystemCall func(contract libcommon.Address, data []byte) ([]byte, error)

// Use more options to call contract
type SysCallCustom func(contract libcommon.Address, data []byte, ibs *state.IntraBlockState, header *types.Header, constCall bool) ([]byte, error)
type Call func(contract libcommon.Address, data []byte) ([]byte, error)

// RewardKind - The kind of block reward.
// Depending on the consensus engine the allocated block reward might have
// different semantics which could lead e.g. to different reward values.
type RewardKind uint16

const (
	// RewardAuthor - attributed to the block author.
	RewardAuthor RewardKind = 0
	// RewardEmptyStep - attributed to the author(s) of empty step(s) included in the block (AuthorityRound engine).
	RewardEmptyStep RewardKind = 1
	// RewardExternal - attributed by an external protocol (e.g. block reward contract).
	RewardExternal RewardKind = 2
	// RewardUncle - attributed to the block uncle(s) with given difference.
	RewardUncle RewardKind = 3
)

type Reward struct {
	Beneficiary libcommon.Address
	Kind        RewardKind
	Amount      uint256.Int
}

// Engine is an algorithm agnostic consensus engine.
type Engine interface {
	EngineReader
	EngineWriter
}

// EngineReader are read-only methods of the consensus engine
// All of these methods should have thread-safe implementations
type EngineReader interface {
	// Author retrieves the Ethereum address of the account that minted the given
	// block, which may be different from the header's coinbase if a consensus
	// engine is based on signatures.
	Author(header *types.Header) (libcommon.Address, error)

	// Service transactions are free and don't pay baseFee after EIP-1559
	IsServiceTransaction(sender libcommon.Address, syscall SystemCall) bool

	Type() chain.ConsensusName

	CalculateRewards(config *chain.Config, header *types.Header, uncles []*types.Header, syscall SystemCall,
	) ([]Reward, error)

	// Close terminates any background threads, DB's etc maintained by the consensus engine.
	Close() error
}

// EngineReader are write methods of the consensus engine
type EngineWriter interface {
	// VerifyHeader checks whether a header conforms to the consensus rules of a
	// given engine. Verifying the seal may be done optionally here, or explicitly
	// via the VerifySeal method.
	VerifyHeader(chain ChainHeaderReader, header *types.Header, seal bool) error

	// VerifyUncles verifies that the given block's uncles conform to the consensus
	// rules of a given engine.
	VerifyUncles(chain ChainReader, header *types.Header, uncles []*types.Header) error

	// Prepare initializes the consensus fields of a block header according to the
	// rules of a particular engine. The changes are executed inline.
	Prepare(chain ChainHeaderReader, header *types.Header, state *state.IntraBlockState) error

	// Initialize runs any pre-transaction state modifications (e.g. epoch start)
	Initialize(config *chain.Config, chain ChainHeaderReader, header *types.Header,
		state *state.IntraBlockState, syscall SysCallCustom, logger log.Logger, eLogger *tracing.Hooks)

	// Finalize runs any post-transaction state modifications (e.g. block rewards)
	// but does not assemble the block.
	//
	// Note: The block header and state database might be updated to reflect any
	// consensus rules that happen at finalization (e.g. block rewards).
	Finalize(config *chain.Config, header *types.Header, state *state.IntraBlockState,
		txs types.Transactions, uncles []*types.Header, receipts types.Receipts, withdrawals []*types.Withdrawal, requests types.Requests, chain ChainReader, syscall SystemCall, logger log.Logger,
	) (types.Transactions, types.Receipts, types.Requests, error)

	// FinalizeAndAssemble runs any post-transaction state modifications (e.g. block
	// rewards) and assembles the final block.
	//
	// Note: The block header and state database might be updated to reflect any
	// consensus rules that happen at finalization (e.g. block rewards).
	FinalizeAndAssemble(config *chain.Config, header *types.Header, state *state.IntraBlockState,
		txs types.Transactions, uncles []*types.Header, receipts types.Receipts, withdrawals []*types.Withdrawal, requests types.Requests, chain ChainReader, syscall SystemCall, call Call, logger log.Logger,
	) (*types.Block, types.Transactions, types.Receipts, error)

	// Seal generates a new sealing request for the given input block and pushes
	// the result into the given channel.
	//
	// Note, the method returns immediately and will send the result async. More
	// than one result may also be returned depending on the consensus algorithm.
	Seal(chain ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error

	// SealHash returns the hash of a block prior to it being sealed.
	SealHash(header *types.Header) libcommon.Hash

	// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
	// that a new block should have.
	CalcDifficulty(chain ChainHeaderReader, time, parentTime uint64, parentDifficulty *big.Int, parentNumber uint64,
		parentHash, parentUncleHash libcommon.Hash, parentAuRaStep uint64) *big.Int

	GenerateSeal(chain ChainHeaderReader, currnt, parent *types.Header, call Call) []byte

	// APIs returns the RPC APIs this consensus engine provides.
	APIs(chain ChainHeaderReader) []rpc.API
}

// PoW is a consensus engine based on proof-of-work.
type PoW interface {
	Engine

	// Hashrate returns the current mining hashrate of a PoW consensus engine.
	Hashrate() float64
}
