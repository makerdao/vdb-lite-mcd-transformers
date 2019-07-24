package flip_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/libraries/shared/storage"
	"github.com/vulcanize/vulcanizedb/libraries/shared/storage/utils"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/mcd_transformers/transformers/shared"
	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
	"github.com/vulcanize/mcd_transformers/transformers/storage/flip"
	"github.com/vulcanize/mcd_transformers/transformers/storage/test_helpers"
)

var _ = Describe("Flip storage mappings", func() {

	var (
		storageRepository *test_helpers.MockMakerStorageRepository
		mappings          flip.StorageKeysLookup
	)

	BeforeEach(func() {
		storageRepository = &test_helpers.MockMakerStorageRepository{}
		mappings = flip.StorageKeysLookup{StorageRepository: storageRepository, ContractAddress: constants.EthFlipContractAddressA()}
	})

	Describe("looking up static keys", func() {
		It("returns value metadata if key exists", func() {
			Expect(mappings.Lookup(flip.VatKey)).To(Equal(flip.VatMetadata))
			Expect(mappings.Lookup(flip.IlkKey)).To(Equal(flip.IlkMetadata))
			Expect(mappings.Lookup(flip.BegKey)).To(Equal(flip.BegMetadata))
			Expect(mappings.Lookup(flip.TtlAndTauStorageKey)).To(Equal(flip.TtlAndTauMetadata))
			Expect(mappings.Lookup(flip.KicksKey)).To(Equal(flip.KicksMetadata))
		})

		It("returns error if key does not exist", func() {
			_, err := mappings.Lookup(fakes.FakeHash)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(utils.ErrStorageKeyNotFound{Key: fakes.FakeHash.Hex()}))
		})
	})

	Describe("looking up dynamic keys", func() {
		It("refreshes mappings from repository if key not found", func() {
			_, _ = mappings.Lookup(fakes.FakeHash)

			Expect(storageRepository.GetFlipBidIdsCalledWith).To(Equal(mappings.ContractAddress))
		})

		It("returns error if bid ID lookup fails", func() {
			storageRepository.GetFlipBidIdsError = fakes.FakeError

			_, err := mappings.Lookup(fakes.FakeHash)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(fakes.FakeError))
		})

		Describe("bid", func() {
			fakeBidId := "42"
			fakeHexBidId, _ := shared.ConvertIntStringToHex(fakeBidId)
			var bidBidKey = common.BytesToHash(crypto.Keccak256(common.FromHex(fakeHexBidId + flip.BidsMappingIndex)))

			BeforeEach(func() {
				storageRepository.FlipBidIds = []string{fakeBidId}
			})

			It("returns value metadata for bid bid", func() {
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidBid,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Uint256,
				}
				Expect(mappings.Lookup(bidBidKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid lot", func() {
				bidLotKey := storage.GetIncrementedKey(bidBidKey, 1)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidLot,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Uint256,
				}
				Expect(mappings.Lookup(bidLotKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid guy", func() {
				bidGuyKey := storage.GetIncrementedKey(bidBidKey, 2)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidGuy,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Address,
				}
				Expect(mappings.Lookup(bidGuyKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid tic", func() {
				bidTicKey := storage.GetIncrementedKey(bidBidKey, 3)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidTic,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Uint48,
				}
				Expect(mappings.Lookup(bidTicKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid end", func() {
				bidEndKey := storage.GetIncrementedKey(bidBidKey, 4)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidEnd,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Uint48,
				}
				Expect(mappings.Lookup(bidEndKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid usr", func() {
				bidUsrKey := storage.GetIncrementedKey(bidBidKey, 5)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidUsr,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Address,
				}
				Expect(mappings.Lookup(bidUsrKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid gal", func() {
				bidGalKey := storage.GetIncrementedKey(bidBidKey, 6)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidGal,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Address,
				}
				Expect(mappings.Lookup(bidGalKey)).To(Equal(expectedMetadata))
			})

			It("returns value metadata for bid tab", func() {
				bidTabKey := storage.GetIncrementedKey(bidBidKey, 7)
				expectedMetadata := utils.StorageValueMetadata{
					Name: flip.BidTab,
					Keys: map[utils.Key]string{constants.BidId: fakeBidId},
					Type: utils.Uint256,
				}
				Expect(mappings.Lookup(bidTabKey)).To(Equal(expectedMetadata))
			})
		})
	})
})