package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/fakeservicebroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("create-service-broker command", func() {
	var brokerName string

	BeforeEach(func() {
		// TODO: remove that when capi-release is cut with v3 create-service-broker functionality
		helpers.SkipIfV7AndVersionLessThan("3.72.0")

		brokerName = helpers.NewServiceBrokerName()

		helpers.LoginCF()
	})

	Describe("help", func() {
		When("--help flag is set", func() {
			It("displays command usage to output", func() {
				session := helpers.CF("create-service-broker", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("\\s+create-service-broker - Create a service broker"))

				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("\\s+cf create-service-broker SERVICE_BROKER USERNAME PASSWORD URL \\[--space-scoped\\]"))

				Eventually(session).Should(Say("ALIAS:"))
				Eventually(session).Should(Say("\\s+csb"))

				Eventually(session).Should(Say("OPTIONS:"))
				Eventually(session).Should(Say("\\s+--space-scoped      Make the broker's service plans only visible within the targeted space"))

				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("\\s+enable-service-access, service-brokers, target"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	When("not logged in", func() {
		BeforeEach(func() {
			helpers.LogoutCF()
		})

		It("displays an informative error that the user must be logged in", func() {
			session := helpers.CF("create-service-broker", brokerName, "user", "pass", "http://example.com")
			Eventually(session).Should(Say("FAILED"))
			Eventually(session.Err).Should(Say("Not logged in. Use 'cf login' to log in."))
			Eventually(session).Should(Exit(1))
		})
	})

	When("logged in", func() {
		When("all arguments are provided", func() {
			When("no org or space is targeted", func() {
				var (
					orgName   string
					spaceName string
					broker    *fakeservicebroker.FakeServiceBroker
				)

				BeforeEach(func() {
					orgName = helpers.NewOrgName()
					spaceName = helpers.NewSpaceName()
					helpers.SetupCF(orgName, spaceName)
					broker = fakeservicebroker.New().WithName(brokerName).Deploy()
					helpers.ClearTarget()
				})

				AfterEach(func() {
					Eventually(helpers.CF("delete-service-broker", brokerName, "-f")).Should(Exit(0))
					helpers.QuickDeleteOrg(orgName)
				})

				It("registers the broker", func() {
					session := helpers.CF("create-service-broker", brokerName, "username", "password", broker.URL())
					Eventually(session).Should(Say("Creating service broker %s as admin...", brokerName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("service-brokers")
					Eventually(session).Should(Say(brokerName))
				})
			})

			FWhen("the --space-scoped flag is passed", func() {
				BeforeEach(func() {
					// TODO: replace skip with versioned skip when
					// https://www.pivotaltracker.com/story/show/166063310 is resolved.
					// helpers.SkipIfV7()
				})

				When("no org or space is targeted", func() {
					BeforeEach(func() {
						helpers.ClearTarget()
					})

					FIt("displays an informative error that a space must be targeted", func() {
						session := helpers.CF("create-service-broker", "space-scoped-broker", "username", "password", "http://example.com", "--space-scoped")
						Eventually(session).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("No org targeted, use 'cf target -o ORG' to target an org."))
						Eventually(session).Should(Exit(1))
					})
				})

				When("both org and space are targeted", func() {
					var (
						orgName   string
						spaceName string
						broker    *fakeservicebroker.FakeServiceBroker
					)

					BeforeEach(func() {
						orgName = helpers.NewOrgName()
						spaceName = helpers.NewSpaceName()
						helpers.SetupCF(orgName, spaceName)

						broker = fakeservicebroker.New().WithName(brokerName).Deploy()
					})

					AfterEach(func() {
						Eventually(helpers.CF("delete-service-broker", brokerName, "-f")).Should(Exit(0))
						helpers.QuickDeleteOrg(orgName)
					})

					It("registers the broker and exposes its services only to the targeted space", func() {
						session := helpers.CF("create-service-broker", brokerName, "username", "password", broker.URL(), "--space-scoped")
						Eventually(session).Should(Say("Creating service broker " + brokerName + " in org " + orgName + " / space " + spaceName + " as admin..."))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Exit(0))

						session = helpers.CF("service-brokers")
						Eventually(session).Should(Say(brokerName))

						session = helpers.CF("marketplace")
						Eventually(session).Should(Say(broker.ServicePlanName()))

						helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
						session = helpers.CF("marketplace")
						Eventually(session).ShouldNot(Say(broker.ServicePlanName()))
					})
				})
			})
		})
	})

	FWhen("the broker URL is invalid", func() {
		BeforeEach(func() {
			// TODO: replace skip with versioned skip when
			// https://www.pivotaltracker.com/story/show/166215494 is resolved.
			// helpers.SkipIfV7()
		})

		It("displays a relevant error", func() {
			session := helpers.CF("create-service-broker", brokerName, "user", "pass", "not-a-valid-url")
			Eventually(session).Should(Say("FAILED"))
			Eventually(session.Err).Should(Say("not-a-valid-url is not a valid URL"))
			Eventually(session).Should(Exit(1))
		})
	})

	When("no arguments are provided", func() {
		It("displays an error, naming each of the missing args and the help text", func() {
			session := helpers.CF("create-service-broker")
			Eventually(session.Err).Should(Say("Incorrect Usage: the required arguments `SERVICE_BROKER`, `USERNAME`, `PASSWORD` and `URL` were not provided"))

			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Say("\\s+create-service-broker - Create a service broker"))

			Eventually(session).Should(Say("USAGE:"))
			Eventually(session).Should(Say("\\s+cf create-service-broker SERVICE_BROKER USERNAME PASSWORD URL \\[--space-scoped\\]"))

			Eventually(session).Should(Say("ALIAS:"))
			Eventually(session).Should(Say("\\s+csb"))

			Eventually(session).Should(Say("OPTIONS:"))
			Eventually(session).Should(Say("\\s+--space-scoped      Make the broker's service plans only visible within the targeted space"))

			Eventually(session).Should(Say("SEE ALSO:"))
			Eventually(session).Should(Say("\\s+enable-service-access, service-brokers, target"))

			Eventually(session).Should(Exit(1))
		})
	})
})
