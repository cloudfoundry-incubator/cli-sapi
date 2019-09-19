package isolated

import (
	"fmt"

	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/fakeservicebroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("space command", func() {
	var (
		orgName   string
		spaceName string
	)

	BeforeEach(func() {
		orgName = helpers.NewOrgName()
		spaceName = helpers.NewSpaceName()
	})

	Describe("help", func() {
		When("--help flag is set", func() {
			It("Displays command usage to output", func() {
				session := helpers.CF("space", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("space - Show space info"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say(`cf space SPACE \[--guid\]`))
				Eventually(session).Should(Say("OPTIONS:"))
				Eventually(session).Should(Say(`--guid\s+Retrieve and display the given space's guid\.  All other output for the space is suppressed\.`))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("set-space-isolation-segment, space-quota, space-users"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	When("the environment is not setup correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(true, false, ReadOnlyOrg, "space", "space-name")
		})
	})

	When("the environment is set up correctly", func() {
		BeforeEach(func() {
			helpers.LoginCF()
			helpers.SetupCF(orgName, spaceName)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
			helpers.ClearTarget()
		})

		When("the space does not exist", func() {
			It("displays not found and exits 1", func() {
				badSpaceName := fmt.Sprintf("%s-1", spaceName)
				session := helpers.CF("space", badSpaceName)
				userName, _ := helpers.GetCredentials()
				Eventually(session).Should(Say(`Getting info for space %s in org %s as %s\.\.\.`, badSpaceName, orgName, userName))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Space '%s' not found.", badSpaceName))
				Eventually(session).Should(Exit(1))
			})
		})

		When("the space exists", func() {
			When("the --guid flag is used", func() {
				It("displays the space guid", func() {
					session := helpers.CF("space", "--guid", spaceName)
					Eventually(session).Should(Say(`[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}`))
					Eventually(session).Should(Exit(0))
				})
			})

			When("the --guid flag is not used", func() {
				When("no flags are used", func() {
					var (
						appName              string
						isolationSegmentName string
					)

					BeforeEach(func() {
						appName = helpers.PrefixedRandomName("app")
						helpers.WithHelloWorldApp(func(appDir string) {
							Eventually(helpers.CF("push", appName, "-p", appDir, "-b", "staticfile_buildpack")).Should(Exit(0))
						})
					})

					It("displays a table with space name, org and apps", func() {
						session := helpers.CF("space", spaceName)
						userName, _ := helpers.GetCredentials()
						Eventually(session).Should(Say(`Getting info for space %s in org %s as %s\.\.\.`, spaceName, orgName, userName))

						Eventually(session).Should(Say(`name:\s+%s`, spaceName))
						Eventually(session).Should(Say(`org:\s+%s`, orgName))
						Eventually(session).Should(Say(`apps:\s+%s`, appName))
						Eventually(session).Should(Say(`services:`))
						Eventually(session).Should(Exit(0))
					})

					When("isolations segments are enabled", func() {
						BeforeEach(func() {
							isolationSegmentName = helpers.NewIsolationSegmentName()
							Eventually(helpers.CF("create-isolation-segment", isolationSegmentName)).Should(Exit(0))
							Eventually(helpers.CF("enable-org-isolation", orgName, isolationSegmentName)).Should(Exit(0))
							Eventually(helpers.CF("set-space-isolation-segment", spaceName, isolationSegmentName)).Should(Exit(0))
						})

						It("displays the isolation segment in the table", func() {
							session := helpers.CF("space", spaceName)
							userName, _ := helpers.GetCredentials()

							Eventually(session).Should(Say(`Getting info for space %s in org %s as %s\.\.\.`, spaceName, orgName, userName))
							Eventually(session).Should(Say(`isolation segment:\s+%s`, isolationSegmentName))
						})
					})
				})

				When("the space does not have an isolation segment and its org has a default isolation segment", func() {
					var orgIsolationSegmentName string

					BeforeEach(func() {
						orgIsolationSegmentName = helpers.NewIsolationSegmentName()
						Eventually(helpers.CF("create-isolation-segment", orgIsolationSegmentName)).Should(Exit(0))
						Eventually(helpers.CF("enable-org-isolation", orgName, orgIsolationSegmentName)).Should(Exit(0))
						Eventually(helpers.CF("set-org-default-isolation-segment", orgName, orgIsolationSegmentName)).Should(Exit(0))
					})

					It("shows the org default isolation segment", func() {
						session := helpers.CF("space", spaceName)
						Eventually(session).Should(Say(`isolation segment:\s+%s \(org default\)`, orgIsolationSegmentName))
						Eventually(session).Should(Exit(0))
					})
				})

				When("the space has service instances", func() {
					var (
						service             string
						servicePlan         string
						serviceInstanceName string
						broker              *fakeservicebroker.FakeServiceBroker
					)

					BeforeEach(func() {
						broker = fakeservicebroker.New().EnsureBrokerIsAvailable()
						service = broker.ServiceName()
						servicePlan = broker.ServicePlanName()
						serviceInstanceName = helpers.NewServiceInstanceName()
						Eventually(helpers.CF("enable-service-access", service)).Should(Exit(0))
						Eventually(helpers.CF("create-service", service, servicePlan, serviceInstanceName)).Should(Exit(0))
					})

					AfterEach(func() {
						broker.Destroy()
					})

					It("shows the service instance", func() {
						session := helpers.CF("space", spaceName)
						Eventually(session).Should(Say(`services:\s+%s`, serviceInstanceName))
						Eventually(session).Should(Exit(0))
					})
				})
			})
		})
	})
})
