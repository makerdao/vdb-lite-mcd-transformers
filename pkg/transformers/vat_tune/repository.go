package vat_tune

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
)

type VatTuneRepository struct {
	db *postgres.DB
}

func (repository VatTuneRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Begin()
	if dBaseErr != nil {
		return dBaseErr
	}
	for _, model := range models {
		vatTune, ok := model.(VatTuneModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, VatTuneModel{})
		}

		ilkID, ilkErr := shared.GetOrCreateIlkInTransaction(vatTune.Ilk, tx)
		if ilkErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return ilkErr
		}

		_, execErr := tx.Exec(
			`INSERT into maker.vat_tune (header_id, ilk, urn, v, w, dink, dart, tx_idx, log_idx, raw_log)
	   VALUES($1, $2, $3, $4, $5, $6::NUMERIC, $7::NUMERIC, $8, $9, $10)
		ON CONFLICT (header_id, tx_idx, log_idx) DO UPDATE SET ilk = $2, urn = $3, v = $4, w = $5, dink = $6, dart = $7, raw_log = $10;`,
			headerID, ilkID, vatTune.Urn, vatTune.V, vatTune.W, vatTune.Dink, vatTune.Dart, vatTune.TransactionIndex, vatTune.LogIndex, vatTune.Raw,
		)
		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := shared.MarkHeaderCheckedInTransaction(headerID, tx, constants.VatTuneChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}
	return tx.Commit()
}

func (repository VatTuneRepository) MarkHeaderChecked(headerID int64) error {
	return shared.MarkHeaderChecked(headerID, repository.db, constants.VatTuneChecked)
}

func (repository VatTuneRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	return shared.MissingHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.VatTuneChecked)
}

func (repository VatTuneRepository) RecheckHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	return shared.RecheckHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.VatTuneChecked)
}

func (repository *VatTuneRepository) SetDB(db *postgres.DB) {
	repository.db = db
}
