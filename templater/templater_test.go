package templater_test

import (
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Templater", func() {
	Describe("TemplatizeSSP", func() {
		It("fills in the Responsible Role fields", func() {
			sspPath := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			s, err := ssp.Load(sspPath)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			openControlDir := filepath.Join("..", "fixtures", "opencontrols")
			openControlDir, err = filepath.Abs(openControlDir)
			Expect(err).NotTo(HaveOccurred())
			openControlData, errors := opencontrols.LoadFrom(openControlDir)
			for _, err := range errors {
				Expect(err).NotTo(HaveOccurred())
			}

			err = TemplatizeSSP(s, openControlData)

			Expect(err).NotTo(HaveOccurred())
			content := s.Content()
			Expect(content).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})
})