package currency

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestInitUpdater(t *testing.T) {
	g := NewGomegaWithT(t)
	err := Init()
	g.Expect(err).ShouldNot(HaveOccurred())
}

func TestNewCurrency(t *testing.T) {
	g := NewGomegaWithT(t)
	err := NewCurrency("AUD")
	g.Expect(err).ShouldNot(HaveOccurred())
	currencies := GetAllCurrencies()
	g.Expect(currencies["AUD"]).ShouldNot(BeNil())
}

func TestNewExistsCurrency(t *testing.T) {
	g := NewGomegaWithT(t)
	err := NewCurrency("USD")
	g.Expect(err).ShouldNot(BeNil())
}

func TestNewNonexistentCurrency(t *testing.T) {
	g := NewGomegaWithT(t)
	err := NewCurrency("AAA")
	g.Expect(err).ShouldNot(BeNil())
}

func TestDeleteCurrency(t *testing.T) {
	g := NewGomegaWithT(t)
	err := DeleteCurrency("BRL")
	g.Expect(err).ShouldNot(HaveOccurred())
	currencies := GetAllCurrencies()
	g.Expect(currencies["BRL"]).Should(BeZero())
}
