package flop_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/mcd_transformers/test_config"
	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
	"github.com/vulcanize/mcd_transformers/transformers/storage/flop"
	. "github.com/vulcanize/mcd_transformers/transformers/storage/test_helpers"
	"github.com/vulcanize/mcd_transformers/transformers/test_data/shared_behaviors"
	"github.com/vulcanize/vulcanizedb/libraries/shared/storage/utils"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"math/rand"
	"strconv"
)

var _ = Describe("Flop storage repository", func() {
	var (
		db              *postgres.DB
		repo            flop.FlopStorageRepository
		fakeBlockHash   string
		fakeBlockNumber int
	)

	BeforeEach(func() {
		fakeBlockNumber = rand.Int()
		fakeBlockHash = fakes.FakeHash.Hex()
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		repo = flop.FlopStorageRepository{ContractAddress: "0x668001c75a9c02d6b10c7a17dbd8aa4afff95037"}
		repo.SetDB(db)
	})

	It("panics if the metadata name is not recognized", func() {
		unrecognizedMetadata := utils.StorageValueMetadata{Name: "unrecognized"}
		flopCreate := func() {
			repo.Create(fakeBlockNumber, fakeBlockHash, unrecognizedMetadata, "")
		}

		Expect(flopCreate).Should(Panic())
	})
	Describe("Vat", func() {
		vatMetadata := utils.StorageValueMetadata{Name: flop.Vat}

		inputs := shared_behaviors.StorageVariableBehaviorInputs{
			FieldName:        flop.Vat,
			Value:            FakeAddress,
			StorageTableName: "maker.flop_vat",
			Repository:       &repo,
			Metadata:         vatMetadata,
		}

		shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
	})

	Describe("Gem", func() {
		gemMetadata := utils.StorageValueMetadata{Name: flop.Gem}
		inputs := shared_behaviors.StorageVariableBehaviorInputs{
			FieldName:        flop.Gem,
			Value:            FakeAddress,
			StorageTableName: "maker.flop_gem",
			Repository:       &repo,
			Metadata:         gemMetadata,
		}

		shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
	})

	Describe("Beg", func() {
		begMetadata := utils.StorageValueMetadata{Name: flop.Beg}
		fakeBeg := strconv.Itoa(rand.Int())

		inputs := shared_behaviors.StorageVariableBehaviorInputs{
			FieldName:        flop.Beg,
			Value:            fakeBeg,
			StorageTableName: "maker.flop_beg",
			Repository:       &repo,
			Metadata:         begMetadata,
		}

		shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)

		It("returns an error if inserting fails", func() {
			createErr := repo.Create(fakeBlockNumber, fakeBlockHash, begMetadata, "")
			Expect(createErr).To(HaveOccurred())
			Expect(createErr.Error()).To(MatchRegexp("pq: invalid input syntax for type numeric"))
		})
	})

	Describe("Ttl and Tau", func() {
		packedNames := make(map[int]string)
		packedNames[0] = flop.Ttl
		packedNames[1] = flop.Tau
		var ttlAndTauMetadata = utils.StorageValueMetadata{
			Name:        flop.Packed,
			PackedNames: packedNames,
		}

		var fakeTtl = strconv.Itoa(rand.Intn(100))
		var fakeTau = strconv.Itoa(rand.Intn(100))
		values := make(map[int]string)
		values[0] = fakeTtl
		values[1] = fakeTau

		It("persists a ttl record", func() {
			createErr := repo.Create(fakeBlockNumber, fakeBlockHash, ttlAndTauMetadata, values)
			Expect(createErr).NotTo(HaveOccurred())

			var ttlResult VariableRes
			getResErr := db.Get(&ttlResult, `SELECT block_number, block_hash, ttl AS value FROM maker.flop_ttl`)
			Expect(getResErr).NotTo(HaveOccurred())
			AssertVariable(ttlResult, fakeBlockNumber, fakeBlockHash, fakeTtl)
		})

		It("persists a tau record", func() {
			createErr := repo.Create(fakeBlockNumber, fakeBlockHash, ttlAndTauMetadata, values)
			Expect(createErr).NotTo(HaveOccurred())

			var tauResult VariableRes
			getResErr := db.Get(&tauResult, `SELECT block_number, block_hash, tau AS value FROM maker.flop_tau`)
			Expect(getResErr).NotTo(HaveOccurred())
			AssertVariable(tauResult, fakeBlockNumber, fakeBlockHash, fakeTau)
		})

		It("panics if the packed name is not recognized", func() {
			packedNames := make(map[int]string)
			packedNames[0] = "notRecognized"

			var badMetadata = utils.StorageValueMetadata{
				Name:        flop.Packed,
				PackedNames: packedNames,
			}

			createFunc := func() {
				_ = repo.Create(fakeBlockNumber, fakeBlockHash, badMetadata, values)
			}
			Expect(createFunc).To(Panic())
		})

		It("returns an error if inserting fails", func() {
			badValues := make(map[int]string)
			badValues[0] = ""
			createErr := repo.Create(fakeBlockNumber, fakeBlockHash, ttlAndTauMetadata, badValues)
			Expect(createErr).To(HaveOccurred())
			Expect(createErr.Error()).To(MatchRegexp("pq: invalid input syntax for type numeric"))
		})
	})

	Describe("Kicks", func() {
		var kicksMetadata = utils.StorageValueMetadata{Name: flop.Kicks}
		var fakeKicks = strconv.Itoa(rand.Int())
		inputs := shared_behaviors.StorageVariableBehaviorInputs{
			FieldName:        flop.Kicks,
			StorageTableName: "maker.flop_kicks",
			Repository:       &repo,
			Metadata:         kicksMetadata,
			Value:            fakeKicks,
		}

		shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
	})

	Describe("Live", func() {
		var liveMetadata = utils.StorageValueMetadata{Name: flop.Live}
		var fakeKicks = strconv.Itoa(rand.Intn(100))
		inputs := shared_behaviors.StorageVariableBehaviorInputs{
			FieldName:        flop.Live,
			StorageTableName: "maker.flop_live",
			Repository:       &repo,
			Metadata:         liveMetadata,
			Value:            fakeKicks,
		}

		shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
	})

	Describe("Bid", func() {
		var fakeBidId = strconv.Itoa(rand.Int())

		It("mappings returns an error if the metadata is missing the bid_id", func() {
			badMetadata := utils.StorageValueMetadata{
				Name: flop.BidBid,
				Keys: map[utils.Key]string{},
				Type: utils.Uint256,
			}
			createErr := repo.Create(fakeBlockNumber, fakeBlockHash, badMetadata, "")
			Expect(createErr).To(MatchError(utils.ErrMetadataMalformed{MissingData: constants.BidId}))
		})

		Describe("bid_bid", func() {
			var fakeBidValue = strconv.Itoa(rand.Int())
			var bidBidMetadata = utils.StorageValueMetadata{
				Name: flop.BidBid,
				Keys: map[utils.Key]string{constants.BidId: fakeBidId},
				Type: utils.Uint256,
			}
			inputs := shared_behaviors.StorageVariableBehaviorInputs{
				FieldName:        "bid",
				Value:            fakeBidValue,
				BidId:            fakeBidId,
				IsAMapping:       true,
				StorageTableName: "maker.flop_bid_bid",
				Repository:       &repo,
				Metadata:         bidBidMetadata,
			}

			shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
		})

		Describe("bid_lot", func() {
			var fakeLotValue = strconv.Itoa(rand.Int())
			var bidLotMetadata = utils.StorageValueMetadata{
				Name: flop.BidLot,
				Keys: map[utils.Key]string{constants.BidId: fakeBidId},
				Type: utils.Uint256,
			}
			inputs := shared_behaviors.StorageVariableBehaviorInputs{
				FieldName:        "lot",
				Value:            fakeLotValue,
				BidId:            fakeBidId,
				IsAMapping:       true,
				StorageTableName: "maker.flop_bid_lot",
				Repository:       &repo,
				Metadata:         bidLotMetadata,
			}

			shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
		})

		Describe("bid_guy", func() {
			var fakeGuyValue = FakeAddress
			var bidGuyMetadata = utils.StorageValueMetadata{
				Name: flop.BidGuy,
				Keys: map[utils.Key]string{constants.BidId: fakeBidId},
				Type: utils.Address,
			}
			inputs := shared_behaviors.StorageVariableBehaviorInputs{
				FieldName:        "guy",
				Value:            fakeGuyValue,
				BidId:            fakeBidId,
				IsAMapping:       true,
				StorageTableName: "maker.flop_bid_guy",
				Repository:       &repo,
				Metadata:         bidGuyMetadata,
			}

			shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
		})

		Describe("bid_tic", func() {
			var fakeTicValue = strconv.Itoa(rand.Intn(100))
			var bidTicMetadata = utils.StorageValueMetadata{
				Name: flop.BidTic,
				Keys: map[utils.Key]string{constants.BidId: fakeBidId},
				Type: utils.Uint48,
			}
			inputs := shared_behaviors.StorageVariableBehaviorInputs{
				FieldName:        "tic",
				Value:            fakeTicValue,
				BidId:            fakeBidId,
				IsAMapping:       true,
				StorageTableName: "maker.flop_bid_tic",
				Repository:       &repo,
				Metadata:         bidTicMetadata,
			}

			shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
		})

		Describe("bid_end", func() {
			var fakeEndValue = strconv.Itoa(rand.Intn(100))
			var bidEndMetadata = utils.StorageValueMetadata{
				Name: flop.BidEnd,
				Keys: map[utils.Key]string{constants.BidId: fakeBidId},
				Type: utils.Uint48,
			}
			inputs := shared_behaviors.StorageVariableBehaviorInputs{
				FieldName:        "\"end\"",
				Value:            fakeEndValue,
				BidId:            fakeBidId,
				IsAMapping:       true,
				StorageTableName: "maker.flop_bid_end",
				Repository:       &repo,
				Metadata:         bidEndMetadata,
			}

			shared_behaviors.SharedStorageRepositoryVariableBehaviors(&inputs)
		})
	})
})