package queries

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/mcd_transformers/test_config"
	"github.com/vulcanize/mcd_transformers/transformers/component_tests/queries/test_helpers"
	"github.com/vulcanize/mcd_transformers/transformers/events/vat_frob"
	"github.com/vulcanize/mcd_transformers/transformers/test_data"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"math/rand"
	"strconv"
)

var _ = Describe("Frobs query", func() {
	var (
		db         *postgres.DB
		frobRepo   vat_frob.VatFrobRepository
		headerRepo repositories.HeaderRepository
		fakeUrn    = test_data.RandomString(5)
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		headerRepo = repositories.NewHeaderRepository(db)
		frobRepo = vat_frob.VatFrobRepository{}
		frobRepo.SetDB(db)
	})

	AfterEach(func() {
		closeErr := db.Close()
		Expect(closeErr).NotTo(HaveOccurred())
	})

	Describe("urn_frobs", func() {
		It("returns frobs for relevant ilk/urn", func() {
			headerOne := fakes.GetFakeHeaderWithTimestamp(int64(111111111), 1)

			headerOneId, err := headerRepo.CreateOrUpdateHeader(headerOne)
			Expect(err).NotTo(HaveOccurred())

			frobBlockOne := test_data.VatFrobModelWithPositiveDart
			frobBlockOne.Ilk = test_helpers.FakeIlk.Hex
			frobBlockOne.Urn = fakeUrn
			frobBlockOne.Dink = strconv.Itoa(rand.Int())
			frobBlockOne.Dart = strconv.Itoa(rand.Int())

			irrelevantFrob := test_data.VatFrobModelWithPositiveDart
			irrelevantFrob.Ilk = test_helpers.AnotherFakeIlk.Hex
			irrelevantFrob.Urn = fakeUrn
			irrelevantFrob.Dink = strconv.Itoa(rand.Int())
			irrelevantFrob.Dart = strconv.Itoa(rand.Int())
			irrelevantFrob.TransactionIndex = frobBlockOne.TransactionIndex + 1

			err = frobRepo.Create(headerOneId, []interface{}{frobBlockOne, irrelevantFrob})
			Expect(err).NotTo(HaveOccurred())

			// New block
			headerTwo := fakes.GetFakeHeaderWithTimestamp(int64(222222222), 2)
			headerTwo.Hash = "anotherHash"
			headerTwoId, err := headerRepo.CreateOrUpdateHeader(headerTwo)
			Expect(err).NotTo(HaveOccurred())

			frobBlockTwo := test_data.VatFrobModelWithPositiveDart
			frobBlockTwo.Ilk = test_helpers.FakeIlk.Hex
			frobBlockTwo.Urn = fakeUrn
			frobBlockTwo.Dink = strconv.Itoa(rand.Int())
			frobBlockTwo.Dart = strconv.Itoa(rand.Int())

			err = frobRepo.Create(headerTwoId, []interface{}{frobBlockTwo})
			Expect(err).NotTo(HaveOccurred())

			var actualFrobs []test_helpers.FrobEvent
			err = db.Select(&actualFrobs, `SELECT ilk_identifier, urn_guy, dink, dart FROM api.urn_frobs($1, $2)`, test_helpers.FakeIlk.Identifier, fakeUrn)
			Expect(err).NotTo(HaveOccurred())

			Expect(actualFrobs).To(ConsistOf(
				test_helpers.FrobEvent{IlkIdentifier: test_helpers.FakeIlk.Identifier, UrnGuy: fakeUrn, Dink: frobBlockOne.Dink, Dart: frobBlockOne.Dart},
				test_helpers.FrobEvent{IlkIdentifier: test_helpers.FakeIlk.Identifier, UrnGuy: fakeUrn, Dink: frobBlockTwo.Dink, Dart: frobBlockTwo.Dart},
			))
		})

		It("fails if no argument is supplied (STRICT)", func() {
			_, err := db.Exec(`SELECT * FROM api.urn_frobs()`)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("function api.urn_frobs() does not exist"))
		})

		It("fails if only one argument is supplied (STRICT)", func() {
			_, err := db.Exec(`SELECT * FROM api.urn_frobs($1::text)`, test_helpers.FakeIlk.Identifier)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("function api.urn_frobs(text) does not exist"))
		})
	})

	Describe("all_frobs", func() {
		It("returns all frobs for a whole ilk", func() {
			headerOne := fakes.GetFakeHeader(1)

			headerOneId, err := headerRepo.CreateOrUpdateHeader(headerOne)
			Expect(err).NotTo(HaveOccurred())

			frobOne := test_data.VatFrobModelWithPositiveDart
			frobOne.Ilk = test_helpers.FakeIlk.Hex
			frobOne.Urn = fakeUrn
			frobOne.Dink = strconv.Itoa(rand.Int())
			frobOne.Dart = strconv.Itoa(rand.Int())

			anotherUrn := "anotherUrn"
			frobTwo := test_data.VatFrobModelWithPositiveDart
			frobTwo.Ilk = test_helpers.FakeIlk.Hex
			frobTwo.Urn = anotherUrn
			frobTwo.Dink = strconv.Itoa(rand.Int())
			frobTwo.Dart = strconv.Itoa(rand.Int())
			frobTwo.TransactionIndex = frobOne.TransactionIndex + 1

			err = frobRepo.Create(headerOneId, []interface{}{frobOne, frobTwo})
			Expect(err).NotTo(HaveOccurred())

			var actualFrobs []test_helpers.FrobEvent
			err = db.Select(&actualFrobs, `SELECT ilk_identifier, urn_guy, dink, dart FROM api.all_frobs($1)`, test_helpers.FakeIlk.Identifier)
			Expect(err).NotTo(HaveOccurred())

			Expect(actualFrobs).To(ConsistOf(
				test_helpers.FrobEvent{IlkIdentifier: test_helpers.FakeIlk.Identifier, UrnGuy: fakeUrn, Dink: frobOne.Dink, Dart: frobOne.Dart},
				test_helpers.FrobEvent{IlkIdentifier: test_helpers.FakeIlk.Identifier, UrnGuy: anotherUrn, Dink: frobTwo.Dink, Dart: frobTwo.Dart},
			))
		})

		It("fails if no argument is supplied (STRICT)", func() {
			_, err := db.Exec(`SELECT * FROM api.all_frobs()`)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("function api.all_frobs() does not exist"))
		})
	})
})