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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/mcd_transformers/test_config"
	"github.com/vulcanize/mcd_transformers/transformers/events/new_cdp"
	mcdConstants "github.com/vulcanize/mcd_transformers/transformers/shared/constants"
	"github.com/vulcanize/mcd_transformers/transformers/test_data"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	"github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

var _ = Describe("NewCdp Transformer", func() {
	var (
		db         *postgres.DB
		blockChain core.BlockChain
	)

	BeforeEach(func() {
		rpcClient, ethClient, err := getClients(ipc)
		Expect(err).NotTo(HaveOccurred())
		blockChain, err = getBlockChain(rpcClient, ethClient)
		Expect(err).NotTo(HaveOccurred())
		db = test_config.NewTestDB(blockChain.Node())
		test_config.CleanTestDB(db)
	})

	newCdpConfig := transformer.EventTransformerConfig{
		TransformerName:   mcdConstants.NewCdpLabel,
		ContractAddresses: []string{test_data.CdpManagerAddress()},
		ContractAbi:       mcdConstants.CdpManagerABI(),
		Topic:             mcdConstants.NewCdpSignature(),
	}

	It("fetches a transforms a NewCdp event from Kovan chain", func() {
		blockNumber := int64(12975121)
		newCdpConfig.StartingBlockNumber = blockNumber
		newCdpConfig.EndingBlockNumber = blockNumber

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		tr := event.Transformer{
			Config:     newCdpConfig,
			Converter:  &new_cdp.NewCdpConverter{},
			Repository: &new_cdp.NewCdpRepository{},
		}.NewTransformer(db)

		logFetcher := fetcher.NewLogFetcher(blockChain)
		logs, err := logFetcher.FetchLogs(
			transformer.HexStringsToAddresses(newCdpConfig.ContractAddresses),
			[]common.Hash{common.HexToHash(newCdpConfig.Topic)},
			header)
		Expect(err).NotTo(HaveOccurred())

		err = tr.Execute(logs, header)
		Expect(err).NotTo(HaveOccurred())

		var dbResult []new_cdp.NewCdpModel
		queryErr := db.Select(&dbResult, `SELECT usr, own, cdp FROM maker.new_cdp`)
		Expect(queryErr).NotTo(HaveOccurred())

		Expect(len(dbResult)).To(Equal(1))
		Expect(dbResult[0].Usr).To(Equal("0xA9fCcB07DD3f774d5b9d02e99DE1a27f47F91189"))
		Expect(dbResult[0].Own).To(Equal("0xA9fCcB07DD3f774d5b9d02e99DE1a27f47F91189"))
		Expect(dbResult[0].Cdp).To(Equal("21"))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.CdpManagerAddress())
		abi, parseErr := geth.ParseAbi(mcdConstants.CdpManagerABI())
		Expect(parseErr).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &new_cdp.NewCdpEntity{}

		var eventLog = test_data.EthNewCdpLog

		unpackErr := contract.UnpackLog(entity, "NewCdp", eventLog)
		Expect(unpackErr).NotTo(HaveOccurred())

		expectedEntity := test_data.NewCdpEntity
		Expect(entity.Usr).To(Equal(expectedEntity.Usr))
		Expect(entity.Own).To(Equal(expectedEntity.Own))
		Expect(entity.Cdp).To(Equal(entity.Cdp))
	})
})