package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/fakeservicebroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = FDescribe("purge-service-offering command", func() {
	Describe("help", func() {
		When("the --help flag is set", func() {
			It("displays command usage to output", func() {
				session := helpers.CF("purge-service-offering", "--help")

				Eventually(session).Should(Exit(0))
				Expect(session).To(Say("NAME:"))
				Expect(session).To(Say("purge-service-offering - Recursively remove a service and child objects from Cloud Foundry database without making requests to a service broker"))
				Expect(session).To(Say("USAGE:"))
				Expect(session).To(Say(`cf purge-service-offering SERVICE \[-b BROKER\] \[-f\]`))
				Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
				Expect(session).To(Say("OPTIONS:"))
				Expect(session).To(Say("-b\\s+Purge a service from a particular service broker. Required when service name is ambiguous"))
				Expect(session).To(Say("-f\\s+Force deletion without confirmation"))
				Expect(session).To(Say("SEE ALSO:"))
				Expect(session).To(Say("marketplace, purge-service-instance, service-brokers"))
			})
		})

		When("no args are passed", func() {
			It("displays an error message with help text", func() {
				session := helpers.CF("purge-service-offering")

				Eventually(session).Should(Exit(1))

				Expect(session.Err).To(Say("Incorrect Usage: the required argument `SERVICE` was not provided"))
				Expect(session).To(Say("NAME:"))
				Expect(session).To(Say("purge-service-offering - Recursively remove a service and child objects from Cloud Foundry database without making requests to a service broker"))
				Expect(session).To(Say("USAGE:"))
				Expect(session).To(Say(`cf purge-service-offering SERVICE \[-b BROKER\] \[-f\]`))
				Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
				Expect(session).To(Say("OPTIONS:"))
				Expect(session).To(Say("-b\\s+Purge a service from a particular service broker. Required when service name is ambiguous"))
				Expect(session).To(Say("-f\\s+Force deletion without confirmation"))
				Expect(session).To(Say("SEE ALSO:"))
				Expect(session).To(Say("marketplace, purge-service-instance, service-brokers"))
			})
		})

		When("more than required number of args are passed", func() {
			It("displays an error message with help text and exits 1", func() {
				session := helpers.CF("purge-service-offering", "service-name", "extra")

				Eventually(session).Should(Exit(1))
				Expect(session.Err).To(Say(`Incorrect Usage: unexpected argument "extra"`))
				Expect(session).To(Say("NAME:"))
				Expect(session).To(Say("purge-service-offering - Recursively remove a service and child objects from Cloud Foundry database without making requests to a service broker"))
				Expect(session).To(Say("USAGE:"))
				Expect(session).To(Say(`cf purge-service-offering SERVICE \[-b BROKER\] \[-f\]`))
				Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
				Expect(session).To(Say("OPTIONS:"))
				Expect(session).To(Say("-b\\s+Purge a service from a particular service broker. Required when service name is ambiguous"))
				Expect(session).To(Say("-f\\s+Force deletion without confirmation"))
				Expect(session).To(Say("SEE ALSO:"))
				Expect(session).To(Say("marketplace, purge-service-instance, service-brokers"))
			})
		})
	})

	When("logged in", func() {
		var orgName, spaceName string

		BeforeEach(func() {
			orgName = helpers.NewOrgName()
			spaceName = helpers.NewSpaceName()
			helpers.SetupCF(orgName, spaceName)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		When("the service exists", func() {
			var (
				broker *fakeservicebroker.FakeServiceBroker
				//buffer    *Buffer
			)

			BeforeEach(func() {
				//buffer = NewBuffer()

				broker = fakeservicebroker.New().EnsureBrokerIsAvailable()
			})

			When("the service name is ambiguous", func() {
				var secondBroker *fakeservicebroker.FakeServiceBroker

				BeforeEach(func() {
					secondBroker = fakeservicebroker.NewAlternate()
					secondBroker.Services[0].Name = broker.ServiceName()
					secondBroker.EnsureBrokerIsAvailable()
				})

				AfterEach(func() {
					secondBroker.Destroy()
				})

				It("fails, asking the user to disambiguate", func() {
					session := helpers.CF("purge-service-offering", broker.ServiceName(), "-f")

					Eventually(session).Should(Exit(1))
					Expect(session.Err).To(Say(`Service '%s' is provided by multiple service brokers.`, broker.ServiceName()))
					Expect(session.Err).To(Say(`Specify a broker from available brokers '%s', '%s' by using the '-b' flag.`, broker.Name(), secondBroker.Name()))
				})
			})
		})

		When("the service does not exist", func() {
			It("fails", func() {
				session := helpers.CF("purge-service-offering", "no-such-service", "-f")

				Eventually(session).Should(Exit(1))
				Expect(session.Err).To(Say(`Service offering 'no-such-service' not found.`))
			})
		})
	})

	When("the environment is not targeted correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(false, false, ReadOnlyOrg, "purge-service-offering", "-f", "foo")
		})
	})

	//		BeforeEach(func() {
	//			helpers.LoginCF()
	//		})
	//
	//		When("the service exists", func() {
	//			var (
	//				orgName   string
	//				spaceName string
	//				broker    *fakeservicebroker.FakeServiceBroker
	//				buffer    *Buffer
	//			)
	//
	//			BeforeEach(func() {
	//				buffer = NewBuffer()
	//
	//				orgName = helpers.NewOrgName()
	//				spaceName = helpers.NewSpaceName()
	//				helpers.SetupCF(orgName, spaceName)
	//
	//				broker = fakeservicebroker.New().EnsureBrokerIsAvailable()
	//			})
	//
	//			AfterEach(func() {
	//				broker.Destroy()
	//				helpers.QuickDeleteOrg(orgName)
	//			})
	//
	//			When("the user enters 'y'", func() {
	//				BeforeEach(func() {
	//					_, err := buffer.Write([]byte("y\n"))
	//					Expect(err).ToNot(HaveOccurred())
	//				})
	//
	//				It("purges the service offering, asking for confirmation", func() {
	//					session := helpers.CFWithStdin(buffer, "purge-service-offering", broker.ServiceName())
	//
	//					Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//					Expect(session).To(Say("Really purge service offering %s from Cloud Foundry?", broker.ServiceName()))
	//					Expect(session).To(Say("Purging service %s...", broker.ServiceName()))
	//					Expect(session).To(Say("OK"))
	//					Eventually(session).Should(Exit(0))
	//
	//					session = helpers.CF("marketplace")
	//					Expect(session).To(Say("OK"))
	//					Consistently(session).ShouldNot(Say(broker.ServiceName()))
	//					Eventually(session).Should(Exit(0))
	//				})
	//			})
	//
	//			When("the user enters something other than 'y' or 'yes'", func() {
	//				BeforeEach(func() {
	//					_, err := buffer.Write([]byte("wat\n\n"))
	//					Expect(err).ToNot(HaveOccurred())
	//				})
	//
	//				It("asks again", func() {
	//					session := helpers.CFWithStdin(buffer, "purge-service-offering", broker.ServiceName())
	//
	//					Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//					Expect(session).To(Say("Really purge service offering %s from Cloud Foundry?", broker.ServiceName()))
	//					Expect(session).To(Say(`invalid input \(not y, n, yes, or no\)`))
	//					Expect(session).To(Say("Really purge service offering %s from Cloud Foundry?", broker.ServiceName()))
	//					Eventually(session).Should(Exit(0))
	//				})
	//			})
	//
	//			When("the user enters 'n' or 'no'", func() {
	//				BeforeEach(func() {
	//					_, err := buffer.Write([]byte("n\n"))
	//					Expect(err).ToNot(HaveOccurred())
	//				})
	//
	//				It("does not purge the service offering", func() {
	//					session := helpers.CFWithStdin(buffer, "purge-service-offering", broker.ServiceName())
	//
	//					Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//					Expect(session).To(Say("Really purge service offering %s from Cloud Foundry?", broker.ServiceName()))
	//					Eventually(session).ShouldNot(Say("Purging service %s...", broker.ServiceName()))
	//					Eventually(session).ShouldNot(Say("OK"))
	//					Eventually(session).Should(Exit(0))
	//				})
	//			})
	//
	//			When("the -f flag is provided", func() {
	//				It("purges the service offering without asking for confirmation", func() {
	//					session := helpers.CF("purge-service-offering", broker.ServiceName(), "-f")
	//
	//					Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//					Expect(session).To(Say("Purging service %s...", broker.ServiceName()))
	//					Expect(session).To(Say("OK"))
	//					Eventually(session).Should(Exit(0))
	//				})
	//			})
	//		})
	//
	//		When("the service does not exist", func() {
	//			It("prints a message the service offering does not exist, exiting 0", func() {
	//				session := helpers.CF("purge-service-offering", "missing-service")
	//
	//				Expect(session).To(Say("Service offering 'missing-service' not found"))
	//				Eventually(session).Should(Exit(0))
	//			})
	//		})
	//
	//		When("the -p flag is provided", func() {
	//			It("prints a warning that this flag is no longer supported", func() {
	//				session := helpers.CF("purge-service-offering", "some-service", "-p", "some-provider")
	//
	//				Expect(session).To(Say("FAILED"))
	//				Eventually(session.Err).Should(Say("Flag '-p' is no longer supported"))
	//				Eventually(session).ShouldNot(Say("Purging service"))
	//				Eventually(session).ShouldNot(Say("OK"))
	//				Eventually(session).Should(Exit(1))
	//			})
	//		})
	//
	//		When("the -b flag is provided", func() {
	//			var (
	//				orgName   string
	//				spaceName string
	//				broker1   *fakeservicebroker.FakeServiceBroker
	//				broker2   *fakeservicebroker.FakeServiceBroker
	//				buffer    *Buffer
	//			)
	//
	//			It("fails when service broker is not registered yet", func() {
	//				session := helpers.CF("purge-service-offering", "some-service", "-b", "non-existent-broker")
	//
	//				Eventually(session.Err).Should(Say("Service broker 'non-existent-broker' not found"))
	//				Eventually(session.Err).Should(Say("TIP: Use 'cf service-brokers' to see a list of available brokers."))
	//				Expect(session).To(Say("FAILED"))
	//				Eventually(session).Should(Exit(1))
	//			})
	//
	//			When("the service is provided by multiple brokers", func() {
	//				BeforeEach(func() {
	//					buffer = NewBuffer()
	//					_, err := buffer.Write([]byte("y\n"))
	//					Expect(err).ToNot(HaveOccurred())
	//					orgName = helpers.NewOrgName()
	//					spaceName = helpers.NewSpaceName()
	//					helpers.SetupCF(orgName, spaceName)
	//
	//					broker1 = fakeservicebroker.New().EnsureBrokerIsAvailable()
	//					broker2 = fakeservicebroker.NewAlternate()
	//					broker2.Services[0].Name = broker1.ServiceName()
	//					broker2.EnsureBrokerIsAvailable()
	//
	//					session := helpers.CF("enable-service-access", broker1.ServiceName(), "-b", broker1.Name())
	//					Eventually(session).Should(Exit(0))
	//					session = helpers.CF("enable-service-access", broker1.ServiceName(), "-b", broker2.Name())
	//					Eventually(session).Should(Exit(0))
	//				})
	//
	//				AfterEach(func() {
	//					broker1.Destroy()
	//					broker2.Destroy()
	//					helpers.QuickDeleteOrg(orgName)
	//				})
	//
	//				When("the user specifies the desired broker", func() {
	//
	//					It("purges the service offering, asking for confirmation", func() {
	//						session := helpers.CFWithStdin(buffer, "purge-service-offering", broker1.ServiceName(), "-b", broker1.Name())
	//
	//						Expect(session).To(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//						Expect(session).To(Say("Really purge service offering %s from broker %s from Cloud Foundry?", broker1.ServiceName(), broker1.Name()))
	//						Expect(session).To(Say("Purging service %s...", broker1.ServiceName()))
	//						Expect(session).To(Say("OK"))
	//						Eventually(session).Should(Exit(0))
	//
	//						session = helpers.CF("marketplace")
	//						Expect(session).To(Say("OK"))
	//						Consistently(session).ShouldNot(Say(`%s.+%s`, broker1.ServiceName(), broker1.Name()))
	//						Expect(session).To(Say(`%s.+%s`, broker1.ServiceName(), broker2.Name()))
	//						Eventually(session).Should(Exit(0))
	//					})
	//				})
	//
	//				When("the user does not specify the desired broker", func() {
	//					It("does not purge the service offering", func() {
	//						session := helpers.CFWithStdin(buffer, "purge-service-offering", broker1.ServiceName())
	//
	//						Eventually(session.Err).Should(Say("Service '%s' is provided by multiple service brokers. Specify a broker by using the '-b' flag.", broker1.ServiceName()))
	//						Expect(session).To(Say("FAILED"))
	//
	//						Eventually(session).ShouldNot(Say("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database\\. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings\\. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances\\. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup\\."))
	//						Eventually(session).ShouldNot(Say("Purging service %s...", broker1.ServiceName()))
	//						Eventually(session).ShouldNot(Say("OK"))
	//						Eventually(session).Should(Exit(1))
	//					})
	//				})
	//			})
	//		})
	//	})

})
