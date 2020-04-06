package ccv3

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
	"code.cloudfoundry.org/cli/api/cloudcontroller/jsonry"
)

// ServiceOffering represents a Cloud Controller V3 Service Offering.
type ServiceOffering struct {
	// GUID is a unique service offering identifier.
	GUID string
	// Name is the name of the service offering.
	Name string
	// ServiceBrokerName is the name of the service broker
	ServiceBrokerName string `jsonry:"relationships.service_broker.data.name"`

	Metadata *Metadata
}

func (so *ServiceOffering) UnmarshalJSON(data []byte) error {
	return jsonry.Unmarshal(data, so)
}

// GetServiceOffering lists service offering with optional filters.
func (client *Client) GetServiceOfferings(query ...Query) ([]ServiceOffering, Warnings, error) {
	var resources []ServiceOffering

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetServiceOfferingsRequest,
		Query:        query,
		ResponseBody: ServiceOffering{},
		AppendToList: func(item interface{}) error {
			resources = append(resources, item.(ServiceOffering))
			return nil
		},
	})

	return resources, warnings, err
}

func (client *Client) GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (ServiceOffering, Warnings, error) {
	query := []Query{{Key: NameFilter, Values: []string{serviceOfferingName}}}
	if serviceBrokerName != "" {
		query = append(query, Query{Key: ServiceBrokerNamesFilter, Values: []string{serviceBrokerName}})
	}

	offerings, warnings, err := client.GetServiceOfferings(query...)
	if err != nil {
		return ServiceOffering{}, warnings, err
	}

	switch len(offerings) {
	case 0:
		return ServiceOffering{}, warnings, ccerror.ServiceOfferingNotFoundError{Name: serviceOfferingName}
	case 1:
		return offerings[0], warnings, nil
	default:
		return ServiceOffering{}, warnings, ccerror.ServiceOfferingNameAmbiguityError{
			Name:               serviceOfferingName,
			ServiceBrokerNames: extractBrokerNames(offerings),
		}
	}
}

func extractBrokerNames(offerings []ServiceOffering) (result []string) {
	for _, o := range offerings {
		result = append(result, o.ServiceBrokerName)
	}
	return
}
