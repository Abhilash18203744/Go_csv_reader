package customerimporter

import (
	"testing"
)

func TestBasicDomainCounterEmpty(t *testing.T) {
	domainCount, err := basicDomainCounter("customers_empty.csv")
	_ = err
	if len(domainCount) != 0 {
		t.Errorf("Expected empty map domainCounts.")
	}
}

func TestConcDomainCounterEmpty(t *testing.T) {
	domainCount, err := concWPDomainCounter("customers_empty.csv")
	_ = err
	if len(domainCount) != 0 {
		t.Errorf("Expected empty map domainCounts.")
	}
}

func TestBasicDomainCounterSmall(t *testing.T) {
	domainCount, err := basicDomainCounter("customers.csv")
	_ = domainCount
	if err != nil {
		t.Errorf("Expected Domain counts, received %v.", err)
	}
}

func TestConcDomainCounterSmall(t *testing.T) {
	domainCount, err := concWPDomainCounter("customers.csv")
	_ = domainCount
	if err != nil {
		t.Errorf("Expected Domain counts, received %v.", err)
	}
}

func TestBasicDomainCounterMissing(t *testing.T) {
	domainCount, err := basicDomainCounter("custom.csv")
	_ = domainCount
	if err == nil {
		t.Errorf("Expected an error for missing file.")
	}
}

func TestConcDomainCounterMissing(t *testing.T) {
	domainCount, err := concWPDomainCounter("custom.csv")
	_ = domainCount
	if err == nil {
		t.Errorf("Expected an error for missing file.")
	}
}

func TestBasicDomainCounterWrong(t *testing.T) {
	domainCount, err := basicDomainCounter("customers_wrong.csv")
	_ = domainCount
	if err == nil {
		t.Errorf("Expected column missing error.")
	}
}

func TestConcDomainCounterWrong(t *testing.T) {
	domainCount, err := concWPDomainCounter("customers_wrong.csv")
	_ = domainCount
	if err == nil {
		t.Errorf("Expected column missing error.")
	}
}

// func TestBasicDomainCounterLarge(t *testing.T) {
// 	domainCount, err := basicDomainCounter("customers_large.csv")
// 	_ = domainCount
// 	if err != nil {
// 		t.Errorf("Expected Domain counts, received %v", err)
// 	}
// }

// func TestConcDomainCounterLarge(t *testing.T) {
// 	domainCount, err := concWPDomainCounter("customers_large.csv")
// 	_ = domainCount
// 	if err != nil {
// 		t.Errorf("Expected Domain counts, received %v", err)
// 	}
// }
