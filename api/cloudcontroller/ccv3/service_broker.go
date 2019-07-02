package ccv3

import (
	"bytes"
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// ServiceBroker represents a Cloud Controller V3 Service Broker.
type ServiceBroker struct {
	// GUID is a unique service broker identifier.
	GUID string `json:"guid,omitempty"`
	// Name is the name of the service broker.
	Name string `json:"name"`
	// URL is the url of the service broker.
	URL string `json:"url"`
	// Credentials contains the credentials for authenticating with the service broker.
	Credentials ServiceBrokerCredentials `json:"credentials"`
	// This is the relationship for the space GUID
	Relationships *ServiceBrokerRelationships `json:"relationships,omitempty"`
}

// ServiceBrokerCredentials represents a data structure for the Credentials
// of V3 Service Broker.
type ServiceBrokerCredentials struct {
	// Type is the type of credentials for the service broker, e.g. "basic"
	Type constant.ServiceBrokerCredentialsType `json:"type"`
	// Data is the credentials data of the service broker of a particular type.
	Data ServiceBrokerCredentialsData `json:"data"`
}

// ServiceBrokerCredentialsData represents a data structure for the Credentials Data
// of V3 Service Broker Credentials.
type ServiceBrokerCredentialsData struct {
	// Username is the Basic Auth username for the service broker.
	Username string `json:"username"`
	// Password is the Basic Auth password for the service broker.
	Password string `json:"password"`
}

type ServiceBrokerRelationships struct {
	Space ServiceBrokerRelationshipsSpace `json:"space"`
}

type ServiceBrokerRelationshipsSpace struct {
	Data ServiceBrokerRelationshipsSpaceData `json:"data"`
}

type ServiceBrokerRelationshipsSpaceData struct {
	GUID string `json:"guid"`
}

// CreateServiceBroker registers a new service broker.
func (client *Client) CreateServiceBroker(name, username, password, brokerURL, spaceGUID string) (Warnings, error) {
	bodyBytes, err := json.Marshal(newServiceBroker(name, username, password, brokerURL, spaceGUID))
	if err != nil {
		return nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostServiceBrokerRequest,
		Body:        bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return nil, err
	}

	response := cloudcontroller.Response{}
	err = client.connection.Make(request, &response)

	return response.Warnings, err
}

// GetServiceBrokers lists service brokers.
func (client *Client) GetServiceBrokers() ([]ServiceBroker, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceBrokersRequest,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullList []ServiceBroker
	warnings, err := client.paginate(request, ServiceBroker{}, func(item interface{}) error {
		if serviceBroker, ok := item.(ServiceBroker); ok {
			fullList = append(fullList, serviceBroker)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceBroker{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullList, warnings, err
}

func newServiceBroker(name, username, password, brokerURL, spaceGUID string) ServiceBroker {
	serviceBroker := ServiceBroker{
		Name: name,
		URL:  brokerURL,
		Credentials: ServiceBrokerCredentials{
			Type: constant.BasicCredentials,
			Data: ServiceBrokerCredentialsData{
				Username: username,
				Password: password,
			},
		},
	}

	if spaceGUID != "" {
		serviceBroker.Relationships = &ServiceBrokerRelationships{
			Space: ServiceBrokerRelationshipsSpace{
				Data: ServiceBrokerRelationshipsSpaceData{
					GUID: spaceGUID,
				},
			},
		}
	}

	return serviceBroker
}
