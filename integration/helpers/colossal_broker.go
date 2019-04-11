package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

type obj map[string]interface{}
type ColossalBroker struct {
	Name       string
	AppsDomain string
	Path       string
	Config     obj
}

func NewColossalBroker(name string, domain string) *ColossalBroker {
	services := []obj{}

	for i := 0; i < 30; i++ {
		name := "service-" + strconv.Itoa(i)

		plans := []obj{}
		for j := 0; j < 30; j++ {
			name := "plan-" + strconv.Itoa(j)
			plans = append(plans, obj{
				"name":           name,
				"id":             name + "-guid",
				"description":    "Description for fake plan " + name,
				"max_storage_tb": 5,
				"metadata": obj{
					"cost": 0,
					"bullets": []obj{
						{"content": "Shared fake server"},
						{"content": "5 TB storage"},
					},
				},
				"schemas": obj{},
			})
		}

		services = append(services, obj{
			"name":                  name,
			"id":                    name + "-guid",
			"description":           "fake " + name,
			"tags":                  []string{"no-sql", "relational"},
			"requires":              []string{},
			"max_db_per_node":       5,
			"instances_retrievable": true,
			"bindings_retrievable":  true,
			"bindable":              true,
			"metadata": obj{
				"provider": obj{name: "the name"},
				"listing": obj{
					"imageUrl":        "http://catgifpage.com/cat.gif",
					"blurb":           "fake broker that is fake",
					"longDescription": "A long time ago, in a galaxy far far away...",
				},
				"displayName":      "The Fake Broker",
				"documentationUrl": "http://documentation.url",
				"shareable":        true,
			},
			"dashboard_client": obj{
				"id":           "sso-" + name,
				"secret":       "sso-secret" + name,
				"redirect_uri": "http://example.com",
			},
			"plan_updateable": true,
			"plans":           plans,
		})
	}

	return &ColossalBroker{
		Name:       name,
		AppsDomain: domain,
		Path:       NewAssets().ServiceBroker,
		Config: obj{
			"behaviors": obj{
				"catalog": obj{
					"sleep_seconds": 0,
					"status":        200,
					"body": obj{
						"services": services,
					},
				},
			},

			"service_instances":                   obj{},
			"service_bindings":                    obj{},
			"max_fetch_service_instance_requests": 1,
		},
	}
}

func CreateColossalBroker(domain string) *ColossalBroker {
	name := NewServiceBrokerName()
	broker := NewColossalBroker(name, domain)

	broker.
		Push().
		Configure().
		Create()

	return broker
}

func (broker *ColossalBroker) ToJSON() string {
	bytes, err := json.Marshal(broker.Config)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile("/tmp/colossal-broker-config.json", bytes, 0644)
	Expect(err).ToNot(HaveOccurred())

	return string(bytes)
}

func (broker *ColossalBroker) Push() *ColossalBroker {
	Eventually(CF(
		"push", broker.Name,
		"--no-start",
		"-m", DefaultMemoryLimit,
		"-p", broker.Path,
		"--no-route",
	)).Should(Exit(0))

	Eventually(CF(
		"map-route",
		broker.Name,
		broker.AppsDomain,
		"--hostname", broker.Name,
	)).Should(Exit(0))

	Eventually(CF("start", broker.Name)).Should(Exit(0))

	return broker
}

func (broker *ColossalBroker) Configure() *ColossalBroker {
	uri := fmt.Sprintf("http://%s.%s%s", broker.Name, broker.AppsDomain, "/config")
	config := broker.ToJSON()
	fmt.Printf("CONFIG :: %s\n", config)
	body := strings.NewReader(config)
	req, err := http.NewRequest("POST", uri, body)
	Expect(err).ToNot(HaveOccurred())
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	Expect(err).ToNot(HaveOccurred())

	fullBody, err := ioutil.ReadAll(resp.Body)
	Expect(err).ToNot(HaveOccurred())
	fmt.Printf("RESPONSE :: %s\n", fullBody)

	defer resp.Body.Close()

	return broker
}

func (broker *ColossalBroker) Create() *ColossalBroker {
	appURI := fmt.Sprintf("http://%s.%s", broker.Name, broker.AppsDomain)
	Eventually(CF("create-service-broker", broker.Name, "username", "password", appURI)).Should(Exit(0))
	Eventually(CF("service-brokers")).Should(And(Exit(0), Say(broker.Name)))
	return broker
}
