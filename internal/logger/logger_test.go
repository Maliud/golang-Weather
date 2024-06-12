package logger_test

import (
	"errors"
	"testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/Maliud/golang-Weather/internal/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Logger Suite", func() {
	var loggerInstance *logger.Logger

	BeforeEach(func() {
		loggerInstance = logger.NewLogger(true)
	})

	var _ = Describe("Logging", func()  {
		It("DEBUG günlük mesajlarını saklamalı", func() {
			loggerInstance.Debug("Bu bir testtir.")
			loggerInstance.Debug("Bu bir başka testtir.")

			Expect(len(loggerInstance.Flush())).To(Equal(2))
		})

		It("Error Mesajları Saklanmalıdır", func()  {
			loggerInstance.Debug("Bu bir Error testtir.")
			loggerInstance.Error(errors.New("Bu bir Error Testidir."))

			msgs := loggerInstance.Flush()

			Expect(len(msgs)).To(Equal(2))
			Expect(msgs[1]).To(Equal("Bu bir Test hatasıdır."))
		})
	})
})