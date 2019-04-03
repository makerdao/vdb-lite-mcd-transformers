// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package test_data

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/mcd_transformers/transformers/events/jug_file/base"
	ilk2 "github.com/vulcanize/mcd_transformers/transformers/events/jug_file/ilk"
	"github.com/vulcanize/mcd_transformers/transformers/events/jug_file/vow"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"math/big"
)

var EthJugFileIlkLog = types.Log{
	Address: common.HexToAddress(KovanJugContractAddress),
	Topics: []common.Hash{
		common.HexToHash(KovanJugFileIlkSignature),
		common.HexToHash("0x000000000000000000000000793c60586f0be45d9e5c13dfa982dd6d3b5bc9a4"),
		common.HexToHash("0x4554482d41000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x6475747900000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000641a0b287e4554482d4100000000000000000000000000000000000000000000000000000064757479000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000033b2e3caaae446a355df2a700000000000000000000000000000000000000000000000000000000"),
	BlockNumber: 10590493,
	TxHash:      common.HexToHash("0x3eb60a65e5da5b25f988ecc4e6b4b07e68890af8b20926256993c16f12fc5a54"),
	TxIndex:     2,
	BlockHash:   fakes.FakeHash,
	Index:       1,
	Removed:     false,
}

var rawJugFileIlkLog, _ = json.Marshal(EthJugFileIlkLog)
var JugFileIlkModel = ilk2.JugFileIlkModel{
	Ilk:              "4554482d41000000000000000000000000000000000000000000000000000000",
	What:             "duty",
	Data:             "1000000000782997609082909351",
	LogIndex:         EthJugFileIlkLog.Index,
	TransactionIndex: EthJugFileIlkLog.TxIndex,
	Raw:              rawJugFileIlkLog,
}

var EthJugFileBaseLog = types.Log{
	Address: common.HexToAddress(KovanJugContractAddress),
	Topics: []common.Hash{
		common.HexToHash(KovanJugFileBaseSignature),
		common.HexToHash("0x00000000000000000000000064d922894153be9eef7b7218dc565d1d0ce2a092"),
		common.HexToHash("0x66616b6520776861740000000000000000000000000000000000000000000000"),
		common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000007b"),
	},
	Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000004429ae811466616b6520776861740000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b"),
	BlockNumber: 36,
	TxHash:      common.HexToHash("0xeeaa16de1d91c239b66773e8c2116a26cfeaaf5d962b31466c9bf047a5caa20f"),
	TxIndex:     13,
	BlockHash:   fakes.FakeHash,
	Index:       16,
	Removed:     false,
}

var rawJugFileBaseLog, _ = json.Marshal(EthJugFileBaseLog)
var JugFileBaseModel = base.JugFileBaseModel{
	What:             "fake what",
	Data:             big.NewInt(123).String(),
	LogIndex:         EthJugFileBaseLog.Index,
	TransactionIndex: EthJugFileBaseLog.TxIndex,
	Raw:              rawJugFileBaseLog,
}

var EthJugFileVowLog = types.Log{
	Address: common.HexToAddress(KovanJugContractAddress),
	Topics: []common.Hash{
		common.HexToHash(KovanJugFileVowSignature),
		common.HexToHash("0x0000000000000000000000003652c2af10cbbdb753c3b46489db5226b73e6497"),
		common.HexToHash("0x766f770000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x17560834075da3db54f737db74377e799c865821000000000000000000000000"),
	},
	Data:        hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000044e9b674b9766f77000000000000000000000000000000000000000000000000000000000017560834075da3db54f737db74377e799c86582100000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	BlockNumber: 51,
	TxHash:      common.HexToHash("0x586e26b71b41fcd6905044dbe8f0cca300517542278f74a9b925c4f800fed85c"),
	TxIndex:     14,
	BlockHash:   fakes.FakeHash,
	Index:       17,
	Removed:     false,
}

var rawJugFileVowLog, _ = json.Marshal(EthJugFileVowLog)
var JugFileVowModel = vow.JugFileVowModel{
	What:             "vow",
	Data:             "0x17560834075da3db54f737db74377e799c865821000000000000000000000000",
	LogIndex:         EthJugFileVowLog.Index,
	TransactionIndex: EthJugFileVowLog.TxIndex,
	Raw:              rawJugFileVowLog,
}
