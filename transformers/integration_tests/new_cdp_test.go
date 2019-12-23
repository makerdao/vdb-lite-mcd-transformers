// VulcanizeDB
// Copyright © 2019 Vulcanize

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

package integration_tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vdb-mcd-transformers/transformers/events/new_cdp"
	mcdConstants "github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/fetcher"
	"github.com/makerdao/vulcanizedb/libraries/shared/transformer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewCdp Transformer", func() {
	BeforeEach(func() {
		test_config.CleanTestDB(db)
	})

	newCdpConfig := transformer.EventTransformerConfig{
		TransformerName:   mcdConstants.NewCdpTable,
		ContractAddresses: []string{test_data.CdpManagerAddress()},
		ContractAbi:       mcdConstants.CdpManagerABI(),
		Topic:             mcdConstants.NewCdpSignature(),
	}

	It("fetches and transforms a NewCdp event from Kovan chain", func() {
		blockNumber := int64(8930579)
		newCdpConfig.StartingBlockNumber = blockNumber
		newCdpConfig.EndingBlockNumber = blockNumber

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		tr := event.Transformer{
			Config:    newCdpConfig,
			Converter: new_cdp.Converter{},
		}.NewTransformer(db)

		logFetcher := fetcher.NewLogFetcher(blockChain)
		logs, err := logFetcher.FetchLogs(
			transformer.HexStringsToAddresses(newCdpConfig.ContractAddresses),
			[]common.Hash{common.HexToHash(newCdpConfig.Topic)},
			header)
		Expect(err).NotTo(HaveOccurred())
		headerSyncLogs := test_data.CreateLogs(header.Id, logs, db)

		err = tr.Execute(headerSyncLogs)
		Expect(err).NotTo(HaveOccurred())

		var dbResult []NewCdpModel
		queryErr := db.Select(&dbResult, `SELECT usr, own, cdp FROM maker.new_cdp ORDER BY cdp DESC LIMIT 1`)
		Expect(queryErr).NotTo(HaveOccurred())

		Expect(len(dbResult)).To(Equal(1))
		Expect(dbResult[0].Usr).To(Equal("0x094766D0C35300C1fc2D8A7DA8d641886Df0e5FD"))
		Expect(dbResult[0].Own).To(Equal("0xB2AA786Be6264c3F29f96513f13bA927f8b350c2"))
		Expect(dbResult[0].Cdp).To(Equal("99"))
	})
})

type NewCdpModel struct {
	Usr      string
	Own      string
	Cdp      string
	LogID    int64 `db:"log_id"`
	HeaderID int64
}
