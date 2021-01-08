package currency_updater

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestgetRatesfromExternalApi(t *testing.T) {
	g := NewGomegaWithT(t)
	body, err := GetRates(urlRates)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(body).ShouldNot(BeNil())
}

func TestgetRatesfromExternalApiBTCandETH(t *testing.T) {
	g := NewGomegaWithT(t)
	body, err := GetRates(btcEthUrl)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(body).ShouldNot(BeNil())
}

func TestCurrencyLive(t *testing.T) {
	g := NewGomegaWithT(t)
	err := CurrencyLive()
	g.Expect(err).ShouldNot(HaveOccurred())
	for _, value := range Currencies {
		g.Expect(value).ShouldNot(BeZero())
	}
}
