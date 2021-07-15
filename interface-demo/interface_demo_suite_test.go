package interface_demo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterfaceDemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InterfaceDemo Suite")
}
