package isolated

import (
	. "code.cloudfoundry.org/cli/cf/util/testhelpers/matchers"
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/fakeservicebroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("delete-service-broker command", func() {
	Context("Help", func() {
		It("appears in cf help -a", func() {
			session := helpers.CF("help", "-a")
			Eventually(session).Should(Exit(0))
			Expect(session).To(HaveCommandInCategoryWithDescription("delete-service-broker", "SERVICE ADMIN", "Delete a service broker"))
		})

		It("displays the help information", func() {
			session := helpers.CF("delete-service-broker", "--help")
			Eventually(session).Should(Say(`NAME:`))
			Eventually(session).Should(Say(`delete-service-broker - Delete a service broker\n`))
			Eventually(session).Should(Say(`\n`))

			Eventually(session).Should(Say(`USAGE:`))
			Eventually(session).Should(Say(`cf delete-service-broker SERVICE_BROKER \[-f\]\n`))
			Eventually(session).Should(Say(`\n`))

			Eventually(session).Should(Say(`OPTIONS:`))
			Eventually(session).Should(Say(`-f\s+Force deletion without confirmation`))
			Eventually(session).Should(Say(`\n`))

			Eventually(session).Should(Say(`SEE ALSO:`))
			Eventually(session).Should(Say(`delete-service, purge-service-offering, service-brokers`))

			Eventually(session).Should(Exit(0))
		})
	})

	When("an api is targeted and the user is logged in", func() {
		BeforeEach(func() {
			helpers.LoginCF()
		})

		FWhen("there is a service broker without any instances", func() {
			var (
				orgName   string
				spaceName string
				broker    *fakeservicebroker.FakeServiceBroker
			)

			BeforeEach(func() {
				// TODO: why do we create these when the broker helper also does it?
				orgName = helpers.NewOrgName()
				spaceName = helpers.NewSpaceName()
				helpers.SetupCF(orgName, spaceName)
				broker = fakeservicebroker.New().Register()

				helpers.ClearTarget()
			})

			AfterEach(func() {
				//broker.Destroy()
				helpers.QuickDeleteOrg(orgName)
			})

			It("should delete the service broker", func() {
				// TODO: should this not be part of the broker helper?
				session := helpers.CF("enable-service-access", broker.ServiceName())
				Eventually(session).Should(Exit(0))

				// Check our setup before testing
				helpers.TargetOrgAndSpace(orgName, spaceName)

				session = helpers.CF("service-brokers")
				Eventually(session).Should(Say(broker.Name()))
				Eventually(session).Should(Exit(0))

				session = helpers.CF("marketplace")
				Eventually(session).Should(Say(broker.ServicePlanName()))
				Eventually(session).Should(Exit(0))

				// Do the action
				helpers.ClearTarget()
				session = helpers.CF("delete-service-broker", broker.Name(), "-f")
				Eventually(session).Should(Exit(0))

				// Check the world has changed
				helpers.TargetOrgAndSpace(orgName, spaceName)

				session = helpers.CF("service-brokers")
				Consistently(session).ShouldNot(Say(broker.Name()))
				Eventually(session).Should(Exit(0))

				session = helpers.CF("marketplace")
				Consistently(session).ShouldNot(Say(broker.ServicePlanName()))
				Eventually(session).Should(Exit(0))
			})
		})

		When("the service broker doesn't exist", func() {
			It("should exit 0 (idempotent case)", func() {
				session := helpers.CF("delete-service-broker", "not-a-broker", "-f")
				Eventually(session).Should(Say(`Service broker 'not-a-broker' does not exist.`))
				Eventually(session).Should(Exit(0))
			})
		})

		When("the service broker is not specified", func() {
			It("displays error and exits 1", func() {
				session := helpers.CF("delete-service-broker")
				Eventually(session.Err).Should(Say("Incorrect Usage: the required argument `SERVICE_BROKER` was not provided\n"))
				Eventually(session.Err).Should(Say("\n"))
				Eventually(session).Should(Say("NAME:\n"))
				Eventually(session).Should(Exit(1))
			})
		})
	})
})
