package func_demo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFuncDemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FuncDemo Suite")
}
