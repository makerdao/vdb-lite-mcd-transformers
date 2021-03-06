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
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vdb-mcd-transformers/transformers/events/jug_file/ilk"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/fetcher"
	"github.com/makerdao/vulcanizedb/libraries/shared/transformer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jug File Ilk EventTransformer", func() {
	BeforeEach(func() {
		test_config.CleanTestDB(db)
	})

	jugFileIlkConfig := transformer.EventTransformerConfig{
		TransformerName:   constants.JugFileIlkTable,
		ContractAddresses: []string{test_data.JugAddress()},
		ContractAbi:       constants.JugABI(),
		Topic:             constants.JugFileIlkSignature(),
	}

	It("transforms jug file ilk log events", func() {
		blockNumber := int64(8928358)
		jugFileIlkConfig.StartingBlockNumber = blockNumber
		jugFileIlkConfig.EndingBlockNumber = blockNumber

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		initializer := event.Transformer{
			Config:    jugFileIlkConfig,
			Converter: ilk.Converter{},
		}
		tr := initializer.NewTransformer(db)

		f := fetcher.NewLogFetcher(blockChain)
		logs, err := f.FetchLogs(
			transformer.HexStringsToAddresses(jugFileIlkConfig.ContractAddresses),
			[]common.Hash{common.HexToHash(jugFileIlkConfig.Topic)},
			header)
		Expect(err).NotTo(HaveOccurred())

		headerSyncLogs := test_data.CreateLogs(header.Id, logs, db)

		err = tr.Execute(headerSyncLogs)
		Expect(err).NotTo(HaveOccurred())

		var dbResult []jugFileIlkModel
		err = db.Select(&dbResult, `SELECT ilk_id, what, data FROM maker.jug_file_ilk`)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(dbResult)).To(Equal(1))
		ilkID, err := shared.GetOrCreateIlk("0x4554482d41000000000000000000000000000000000000000000000000000000", db)
		Expect(err).NotTo(HaveOccurred())
		Expect(dbResult[0].Ilk).To(Equal(strconv.FormatInt(ilkID, 10)))
		Expect(dbResult[0].What).To(Equal("duty"))
		Expect(dbResult[0].Data).To(Equal("1000000001243680656318820312"))
	})
})

type jugFileIlkModel struct {
	Ilk  string `db:"ilk_id"`
	What string
	Data string
}
