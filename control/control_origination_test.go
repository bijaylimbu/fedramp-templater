package control


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/fixtures"
)

var _ = Describe("controlOrignation", func() {
	Describe("newControlOrignation", func() {
		It("should return", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			tables, err := doc.SummaryTables()
			Expect(err).NotTo(HaveOccurred())
			st := NewSummaryTable(tables[0])
			co, err := newControlOrigination(st)
			Expect(err).ToNot(HaveOccurred())
			Expect(co.cell.Content()).To(Equal(""))
		})
	})
})