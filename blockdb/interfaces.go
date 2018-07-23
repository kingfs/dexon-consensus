// Copyright 2018 The dexon-consensus-core Authors
// This file is part of the dexon-consensus-core library.
//
// The dexon-consensus-core library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus-core library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus-core library. If not, see
// <http://www.gnu.org/licenses/>.

package blockdb

import (
	"errors"

	"github.com/dexon-foundation/dexon-consensus-core/common"
	"github.com/dexon-foundation/dexon-consensus-core/core/types"
)

var (
	// ErrBlockExists is the error when block eixsts.
	ErrBlockExists = errors.New("block exists")
	// ErrBlockDoesNotExist is the error when block does not eixst.
	ErrBlockDoesNotExist = errors.New("block does not exist")
	// ErrValidatorDoesNotExist is the error when validator does not eixst.
	ErrValidatorDoesNotExist = errors.New("validator does not exist")
)

// BlockDatabase is the interface for a BlockDatabase.
type BlockDatabase interface {
	Reader
	Writer
}

// Reader defines the interface for reading blocks into DB.
type Reader interface {
	Has(hash common.Hash) bool
	Get(hash common.Hash) (types.Block, error)
	GetByValidatorAndHeight(vID types.ValidatorID, height uint64) (types.Block, error)
}

// Writer defines the interface for writing blocks into DB.
type Writer interface {
	Update(block types.Block) error
	Put(block types.Block) error
}