package v7action_test

import (
	"errors"

	. "code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/actor/v7action/v7actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Broker Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v7actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v7actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil, nil, nil)
	})

	Describe("GetServiceBrokers", func() {
		var (
			serviceBrokers []ServiceBroker
			warnings       Warnings
			executionError error
		)

		JustBeforeEach(func() {
			serviceBrokers, warnings, executionError = actor.GetServiceBrokers()
		})

		When("the cloud controller request is successful", func() {
			When("the cloud controller returns service brokers", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetServiceBrokersReturns([]ccv3.ServiceBroker{
						{
							GUID: "service-broker-guid-1",
							Name: "service-broker-1",
							URL:  "service-broker-url-1",
						},
						{
							GUID: "service-broker-guid-2",
							Name: "service-broker-2",
							URL:  "service-broker-url-2",
						},
					}, ccv3.Warnings{"some-service-broker-warning"}, nil)
				})

				It("returns the service brokers and warnings", func() {
					Expect(executionError).NotTo(HaveOccurred())

					Expect(serviceBrokers).To(ConsistOf(
						ServiceBroker{Name: "service-broker-1", GUID: "service-broker-guid-1", URL: "service-broker-url-1"},
						ServiceBroker{Name: "service-broker-2", GUID: "service-broker-guid-2", URL: "service-broker-url-2"},
					))
					Expect(warnings).To(ConsistOf("some-service-broker-warning"))
					Expect(fakeCloudControllerClient.GetServiceBrokersCallCount()).To(Equal(1))
				})
			})
		})

		When("the cloud controller returns an error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetServiceBrokersReturns(
					nil,
					ccv3.Warnings{"some-service-broker-warning"},
					errors.New("no service broker"))
			})

			It("returns an error and warnings", func() {
				Expect(executionError).To(MatchError("no service broker"))
				Expect(warnings).To(ConsistOf("some-service-broker-warning"))
			})
		})
	})

	Describe("GetServiceBrokerByName", func() {
		var (
			ccv3ServiceBrokers []ccv3.ServiceBroker
			serviceBroker      ServiceBroker

			serviceBroker1Name string
			serviceBroker1Guid string

			serviceBrokerNotTheOneYouWant string
			notTheBrokerYouAreLookingFor  string

			warnings   Warnings
			executeErr error
		)

		BeforeEach(func() {
			ccv3ServiceBrokers = []ccv3.ServiceBroker{
				{Name: serviceBrokerNotTheOneYouWant, GUID: notTheBrokerYouAreLookingFor},
				{Name: serviceBroker1Name, GUID: serviceBroker1Guid},
			}
		})

		JustBeforeEach(func() {
			serviceBroker, warnings, executeErr = actor.GetServiceBrokerByName(serviceBroker1Name)
		})

		When("the API layer call is successful", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetServiceBrokersReturns(
					ccv3ServiceBrokers,
					ccv3.Warnings{"some-service-broker-warning"},
					nil,
				)
			})

			It("returns back the serviceBrokers and warnings", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(fakeCloudControllerClient.GetServiceBrokersCallCount()).To(Equal(1))

				Expect(serviceBroker).To(Equal(
					ServiceBroker{Name: serviceBroker1Name, GUID: serviceBroker1Guid},
				))
				Expect(warnings).To(ConsistOf("some-service-broker-warning"))

			})
		})

		When("when the API layer call returns an error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetServiceBrokersReturns(
					[]ccv3.ServiceBroker{},
					ccv3.Warnings{"some-service-broker-warning"},
					errors.New("list-error"),
				)
			})

			It("returns the error and prints warnings", func() {
				Expect(executeErr).To(MatchError("list-error"))
				Expect(warnings).To(ConsistOf("some-service-broker-warning"))
				Expect(serviceBroker).To(Equal(ServiceBroker{}))

				Expect(fakeCloudControllerClient.GetServiceBrokersCallCount()).To(Equal(1))
			})
		})
	})

	Describe("CreateServiceBroker", func() {
		const (
			name      = "name"
			url       = "url"
			username  = "username"
			password  = "password"
			spaceGUID = "space-guid"
		)

		var (
			warnings       Warnings
			executionError error
		)

		JustBeforeEach(func() {
			warnings, executionError = actor.CreateServiceBroker(name, username, password, url, spaceGUID)
		})

		When("the client request is successful", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.CreateServiceBrokerReturns(ccv3.Warnings{"some-creation-warning"}, nil)
			})

			It("succeeds and returns warnings", func() {
				Expect(executionError).NotTo(HaveOccurred())

				Expect(warnings).To(ConsistOf("some-creation-warning"))
			})

			It("passes the service broker credentials to the client", func() {
				Expect(fakeCloudControllerClient.CreateServiceBrokerCallCount()).To(Equal(1))
				n, u, p, l, s := fakeCloudControllerClient.CreateServiceBrokerArgsForCall(0)
				Expect(n).To(Equal(name))
				Expect(u).To(Equal(username))
				Expect(p).To(Equal(password))
				Expect(l).To(Equal(url))
				Expect(s).To(Equal(spaceGUID))
			})
		})

		When("the client returns an error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.CreateServiceBrokerReturns(ccv3.Warnings{"some-other-warning"}, errors.New("invalid broker"))
			})

			It("fails and returns warnings", func() {
				Expect(executionError).To(MatchError("invalid broker"))

				Expect(warnings).To(ConsistOf("some-other-warning"))
			})
		})
	})

	Describe("DeleteServiceBroker", func() {
		var (
			serviceBrokerGUID = "some-service-broker-guid"
			warnings          Warnings
			executionError    error
		)

		JustBeforeEach(func() {
			warnings, executionError = actor.DeleteServiceBroker(serviceBrokerGUID)
		})

		When("the client request is successful", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.DeleteServiceBrokerReturns(ccv3.Warnings{"some-deletion-warning"}, nil)
			})

			It("succeeds and returns warnings", func() {
				Expect(executionError).NotTo(HaveOccurred())

				Expect(warnings).To(ConsistOf("some-deletion-warning"))
			})

			It("passes the service broker credentials to the client", func() {
				Expect(fakeCloudControllerClient.DeleteServiceBrokerCallCount()).To(Equal(1))
				actualServiceBrokerGUID := fakeCloudControllerClient.DeleteServiceBrokerArgsForCall(0)
				Expect(actualServiceBrokerGUID).To(Equal(serviceBrokerGUID))
			})
		})

		When("the client returns an error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.DeleteServiceBrokerReturns(ccv3.Warnings{"some-other-warning"}, errors.New("invalid broker"))
			})

			It("fails and returns warnings", func() {
				Expect(executionError).To(MatchError("invalid broker"))

				Expect(warnings).To(ConsistOf("some-other-warning"))
			})
		})
	})
})
