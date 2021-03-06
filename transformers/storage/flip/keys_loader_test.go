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

package flip_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	mcdStorage "github.com/makerdao/vdb-mcd-transformers/transformers/storage"
	"github.com/makerdao/vdb-mcd-transformers/transformers/storage/flip"
	"github.com/makerdao/vdb-mcd-transformers/transformers/storage/test_helpers"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/storage"
	vdbStorage "github.com/makerdao/vulcanizedb/libraries/shared/storage"
	"github.com/makerdao/vulcanizedb/pkg/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flip storage keys loader", func() {
	var (
		storageRepository *test_helpers.MockMakerStorageRepository
		storageKeysLoader storage.KeysLoader
	)

	BeforeEach(func() {
		storageRepository = &test_helpers.MockMakerStorageRepository{}
		storageKeysLoader = flip.NewKeysLoader(storageRepository, constants.GetContractAddress("MCD_FLIP_ETH_A"))
	})

	It("returns value metadata for static keys", func() {
		mappings, err := storageKeysLoader.LoadMappings()

		Expect(err).NotTo(HaveOccurred())
		Expect(mappings[flip.VatKey]).To(Equal(flip.VatMetadata))
		Expect(mappings[flip.IlkKey]).To(Equal(flip.IlkMetadata))
		Expect(mappings[flip.BegKey]).To(Equal(flip.BegMetadata))
		Expect(mappings[flip.TtlAndTauStorageKey]).To(Equal(flip.TtlAndTauMetadata))
		Expect(mappings[flip.KicksKey]).To(Equal(flip.KicksMetadata))
	})

	Describe("bid", func() {
		Describe("when loading bid IDs fails", func() {
			It("returns error", func() {
				storageRepository.GetFlipBidIdsError = fakes.FakeError

				_, err := storageKeysLoader.LoadMappings()

				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(fakes.FakeError))
			})
		})

		Describe("when loading bid IDs succeeds", func() {
			var (
				fakeBidId       = "1"
				fakeHexBidId, _ = shared.ConvertIntStringToHex(fakeBidId)
				bidBidKey       = common.BytesToHash(crypto.Keccak256(common.FromHex(fakeHexBidId + flip.BidsMappingIndex)))
				mappings        map[common.Hash]vdbStorage.ValueMetadata
			)

			BeforeEach(func() {
				storageRepository.FlipBidIds = []string{fakeBidId}
				var err error
				mappings, err = storageKeysLoader.LoadMappings()
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns value metadata for bid bid", func() {
				expectedMetadata := vdbStorage.ValueMetadata{
					Name: mcdStorage.BidBid,
					Keys: map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type: vdbStorage.Uint256,
				}

				Expect(mappings[bidBidKey]).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid lot", func() {
				bidLotKey := vdbStorage.GetIncrementedKey(bidBidKey, 1)
				expectedMetadata := vdbStorage.ValueMetadata{
					Name: mcdStorage.BidLot,
					Keys: map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type: vdbStorage.Uint256,
				}

				Expect(mappings[bidLotKey]).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid guy + tic + end packed slot", func() {
				bidGuyKey := vdbStorage.GetIncrementedKey(bidBidKey, 2)
				expectedMetadata := vdbStorage.ValueMetadata{
					Name:        mcdStorage.Packed,
					Keys:        map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type:        vdbStorage.PackedSlot,
					PackedTypes: map[int]vdbStorage.ValueType{0: vdbStorage.Address, 1: vdbStorage.Uint48, 2: vdbStorage.Uint48},
					PackedNames: map[int]string{0: mcdStorage.BidGuy, 1: mcdStorage.BidTic, 2: mcdStorage.BidEnd},
				}

				Expect(mappings[bidGuyKey]).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid usr", func() {
				bidUsrKey := vdbStorage.GetIncrementedKey(bidBidKey, 3)
				expectedMetadata := vdbStorage.ValueMetadata{
					Name: mcdStorage.BidUsr,
					Keys: map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type: vdbStorage.Address,
				}

				Expect(mappings[bidUsrKey]).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid gal", func() {
				bidGalKey := vdbStorage.GetIncrementedKey(bidBidKey, 4)
				expectedMetadata := vdbStorage.ValueMetadata{
					Name: mcdStorage.BidGal,
					Keys: map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type: vdbStorage.Address,
				}

				Expect(mappings[bidGalKey]).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid tab", func() {
				bidTabKey := vdbStorage.GetIncrementedKey(bidBidKey, 5)
				expectedMetadata := vdbStorage.ValueMetadata{
					Name: mcdStorage.BidTab,
					Keys: map[vdbStorage.Key]string{constants.BidId: fakeBidId},
					Type: vdbStorage.Uint256,
				}

				Expect(mappings[bidTabKey]).To(Equal(expectedMetadata))
			})
		})
	})
})
