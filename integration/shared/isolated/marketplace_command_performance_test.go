package isolated

import (
	"fmt"

	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var benchOrgWasCreated = false

var _ = Describe("cf marketplace performance", func() {
	firstSetup := true

	BeforeEach(func() {
		if !firstSetup {
			return
		}
		firstSetup = false
		//benchOrgWasCreated = true

		helpers.LoginCF()

		// 0. create space-1, and space-2
		//helpers.CreateOrgAndSpace("bench-org", "bench-space-1")
		//helpers.CreateSpace("bench-space-2")

		// 0.1. target space-1
		//helpers.TargetOrgAndSpace("bench-org", "bench-space-1")
		helpers.TargetOrgAndSpace("acceptance", "dev")

		// TODO: automate performance test setup here in the future story
		//       right now, it fails for Rack/Ruby issue:
		//       https://github.com/rack/rack/issues/318

		// 1. create normal broker with 10 services:
		//domain := helpers.DefaultSharedDomain()
		////helpers.CreateColossalBroker(domain)
		//helpers.
		//	NewColossalBroker("INTEGRATION-SERVICE-BROKER-ce9746fa-d830-4b1f-51a7-656496b1f579", domain).
		//	Configure().
		//	Create()

		//// 1.1. make one service fully public
		//Eventually(helpers.CF("enable-service-access", "service-0")).Should(Exit(0))

		//// 1.2. make second service partially public
		//for i := 0; i < 15; i++ {
		//	Eventually(helpers.CF(
		//		"enable-service-access", "service-1",
		//		"-p", "plan-"+strconv.Itoa(i))).Should(Exit(0))
		//}

		//// 1.3. make third service fully private
		//Eventually(helpers.CF("disable-service-access", "service-2")).Should(Exit(0))

		//// 2.1. make bunch of services fully avalable to org
		//for i := 3; i < 18; i++ {
		//	Eventually(helpers.CF(
		//		"enable-service-access", "service-"+strconv.Itoa(i),
		//		"-o", "acceptance")).Should(Exit(0))
		//}

		//// 2.2. make bunch of services partially available to org
		//for i := 18; i < 30; i++ {
		//	for j := 0; j < 20; j++ {
		//		Eventually(helpers.CF(
		//			"enable-service-access", "service-"+strconv.Itoa(i),
		//			"-o", "acceptance",
		//			"-p", "plan-"+strconv.Itoa(j))).Should(Exit(0))
		//	}
		//}

		//// 3. create space-scoped broker with 3 services
	})

	PMeasure("cf marketplace performance", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			fmt.Printf("cf marketplace...")
			session := helpers.CF("marketplace")
			session.Wait()
			fmt.Printf(" DONE.\n")
			Expect(session).Should(Exit(0))
		})

		Expect(runtime.Seconds()).To(BeNumerically("<", 100), "cf marketplace should not be too slow")
	}, 10)
})
