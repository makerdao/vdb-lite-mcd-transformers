package vat_fold_test

import (
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVatFold(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VatFold Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(ioutil.Discard)
})
