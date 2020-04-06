package v7action_test

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"errors"

	"code.cloudfoundry.org/cli/actor/actionerror"
	. "code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/actor/v7action/v7actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Offering Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v7actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v7actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil, nil, nil)
	})

	Describe("GetServiceOfferingByNameAndBroker", func() {
		It("delegates to the client method", func() {
			fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerReturns(
				ccv3.ServiceOffering{Name: "fakeServiceOfferingName"},
				ccv3.Warnings{"a warning"},
				nil,
			)

			serviceOffering, warnings, executionError := actor.GetServiceOfferingByNameAndBroker("fakeServiceOfferingName", "fakeServiceBrokerName")

			Expect(serviceOffering.Name).To(Equal("fakeServiceOfferingName"))
			Expect(warnings).To(ConsistOf("a warning"))
			Expect(executionError).NotTo(HaveOccurred())

			Expect(fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerCallCount()).To(Equal(1))
			actualServiceOfferingName, actualServiceBrokerName := fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerArgsForCall(0)
			Expect(actualServiceOfferingName).To(Equal("fakeServiceOfferingName"))
			Expect(actualServiceBrokerName).To(Equal("fakeServiceBrokerName"))
		})

		DescribeTable(
			"when the client returns an error ",
			func(clientError, expectedError error) {
				fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerReturns(ccv3.ServiceOffering{}, ccv3.Warnings{"a warning"}, clientError)

				_, warnings, err := actor.GetServiceOfferingByNameAndBroker("fake-service-offering", "fake-service-broker")
				Expect(err).To(MatchError(expectedError))
				Expect(warnings).To(ConsistOf("a warning"))
			},
			Entry(
				"ServiceOfferingNotFoundError",
				ccerror.ServiceOfferingNotFoundError{Name: "foo", ServiceBrokerName: "bar"},
				actionerror.ServiceNotFoundError{Name: "foo", Broker: "bar"},
			),
			Entry(
				"ServiceOfferingNameAmbiguityError",
				ccerror.ServiceOfferingNameAmbiguityError{
					Name:               "foo",
					ServiceBrokerNames: []string{"bar", "baz"},
				},
				actionerror.DuplicateServiceError{
					Name:           "foo",
					ServiceBrokers: []string{"bar", "baz"},
				},
			),
			Entry(
				"other error",
				errors.New("boom"),
				errors.New("boom"),
			),
		)
	})

	Describe("PurgeServiceOfferingByNameAndBroker", func() {
		BeforeEach(func() {
			fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerReturns(ccv3.ServiceOffering{}, ccv3.Warnings{"a warning"}, nil)
		})

		It("gets the service offering guid", func() {
			warnings, err := actor.PurgeServiceOfferingByNameAndBroker("fake-service-offering", "fake-service-broker")
			Expect(err).NotTo(HaveOccurred())
			Expect(warnings).To(ConsistOf("a warning"))

			Expect(fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerCallCount()).To(Equal(1))
			actualOffering, actualBroker := fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerArgsForCall(0)
			Expect(actualOffering).To(Equal("fake-service-offering"))
			Expect(actualBroker).To(Equal("fake-service-broker"))
		})

		DescribeTable(
			"when the client returns an error ",
			func(clientError, expectedError error) {
				fakeCloudControllerClient.GetServiceOfferingByNameAndBrokerReturns(ccv3.ServiceOffering{}, ccv3.Warnings{"a warning"}, clientError)

				warnings, err := actor.PurgeServiceOfferingByNameAndBroker("fake-service-offering", "fake-service-broker")
				Expect(err).To(MatchError(expectedError))
				Expect(warnings).To(ConsistOf("a warning"))
			},
			Entry(
				"ServiceOfferingNotFoundError",
				ccerror.ServiceOfferingNotFoundError{Name: "foo", ServiceBrokerName: "bar"},
				actionerror.ServiceNotFoundError{Name: "foo", Broker: "bar"},
			),
			Entry(
				"ServiceOfferingNameAmbiguityError",
				ccerror.ServiceOfferingNameAmbiguityError{
					Name:               "foo",
					ServiceBrokerNames: []string{"bar", "baz"},
				},
				actionerror.DuplicateServiceError{
					Name:           "foo",
					ServiceBrokers: []string{"bar", "baz"},
				},
			),
			Entry(
				"other error",
				errors.New("boom"),
				errors.New("boom"),
			),
		)
	})
})
