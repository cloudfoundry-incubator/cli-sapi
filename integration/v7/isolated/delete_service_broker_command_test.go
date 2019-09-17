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

var _ = FDescribe("delete-service-broker command", func() {
	Context("Help", func() {
		It("appears in cf help -a", func() {
			session := helpers.CF("help", "-a")
			Eventually(session).Should(Exit(0))
			Expect(session).To(HaveCommandInCategoryWithDescription("delete-service-broker", "SERVICE ADMIN", "Delete a service broker"))
		})

		It("displays the help information", func() {
			session := helpers.CF("delete-route", "--help")
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

		// TODO: See ticket 166502005 for details, we are unsure if this behaviour is expected or not
		//When("the environment is not setup correctly", func() {
		//	It("fails with the appropriate errors", func() {
		//		By("checking the org is targeted correctly")
		//		session := helpers.CF("delete-service-broker", "service-broker-name", "-f")
		//		Eventually(session).Should(Say("FAILED"))
		//		Eventually(session.Out).Should(Say("No org and space targeted, use 'cf target -o ORG -s SPACE' to target an org and space"))
		//		Eventually(session).Should(Exit(1))
		//
		//		By("checking the space is targeted correctly")
		//		helpers.TargetOrg(ReadOnlyOrg)
		//		session = helpers.CF("delete-service-broker", "service-broker-name", "-f")
		//		Eventually(session).Should(Say("FAILED"))
		//		Eventually(session.Out).Should(Say(`No space targeted, use 'cf target -s' to target a space\.`))
		//		Eventually(session).Should(Exit(1))
		//	})
		//})

		When("an org and space are targeted", func() {
			var (
				orgName   string
				spaceName string
			)

			BeforeEach(func() {
				orgName = helpers.NewOrgName()
				spaceName = helpers.NewSpaceName()
				helpers.CreateOrgAndSpace(orgName, spaceName)
				helpers.TargetOrgAndSpace(orgName, spaceName)
			})

			AfterEach(func() {
				helpers.QuickDeleteOrg(orgName)
			})

			When("there is a service broker without any instances", func() {
				var (
					service     string
					servicePlan string
					broker      *fakeservicebroker.FakeServiceBroker
				)

				BeforeEach(func() {
					broker = fakeservicebroker.New().Register()
					service = broker.ServiceName()
					servicePlan = broker.ServicePlanName()

					Eventually(helpers.CF("enable-service-access", service)).Should(Exit(0))
				})

				AfterEach(func() {
					broker.Destroy()
				})

				It("should delete the service broker", func() {
					session := helpers.CF("delete-service-broker", broker.Name(), "-f")
					Eventually(session).Should(Exit(0))

					session = helpers.CF("service-brokers")
					Consistently(session).ShouldNot(Say(broker.Name()))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("marketplace")
					Consistently(session).ShouldNot(Say(servicePlan))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("services")
					Consistently(session).ShouldNot(Say(service))
					Eventually(session).Should(Exit(0))
				})
			})

			When("the service broker doesn't exist", func() {
				It("should exit 0 (idempotent case)", func() {
					session := helpers.CF("delete-service-broker", "not-a-broker", "-f")
					Eventually(session).Should(Say(`Service broker 'non-existent-broker' does not exist.`))
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
})
