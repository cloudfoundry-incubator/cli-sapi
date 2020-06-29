package resources_test

import (
	"encoding/json"

	. "code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("service plan resource", func() {
	DescribeTable(
		"Unmarshaling",
		func(servicePlan ServicePlan, serialized string) {
			var parsed ServicePlan
			Expect(json.Unmarshal([]byte(serialized), &parsed)).NotTo(HaveOccurred())
			Expect(parsed).To(Equal(servicePlan))
		},
		Entry(
			"basic",
			ServicePlan{
				GUID:                "fake-service-plan-guid",
				Name:                "fake-service-plan-name",
				ServiceOfferingGUID: "fake-service-offering-guid",
			},
			`{
				"guid": "fake-service-plan-guid",
				"name": "fake-service-plan-name",
				"relationships": {
					"service_offering": {
						"data": {
							"guid": "fake-service-offering-guid"
						}
					}
				}
			}`,
		),
		Entry(
			"detailed",
			ServicePlan{
				GUID:                "fake-service-plan-guid",
				Name:                "fake-service-plan-name",
				ServiceOfferingGUID: "fake-service-offering-guid",
				Description:         "fake-description",
				Available:           true,
				VisibilityType:      "public",
				Free:                true,
				SpaceGUID:           "fake-space-guid",
				Metadata: &Metadata{
					Labels: map[string]types.NullString{
						"foo": types.NewNullString("bar"),
						"baz": types.NewNullString(),
					},
				},
			},
			`{
				"guid": "fake-service-plan-guid",
				"name": "fake-service-plan-name",
				"description": "fake-description",
				"available": true,
				"visibility_type": "public",
				"free": true,
				"metadata": {
					"labels": {
						"foo": "bar",
						"baz": null
					}
				},
				"relationships": {
					"service_offering": {
						"data": {
							"guid": "fake-service-offering-guid"
						}
					},
					"space": {
						"data": {
							"guid": "fake-space-guid"
						}
					}
				}
			}`,
		),
	)
})
