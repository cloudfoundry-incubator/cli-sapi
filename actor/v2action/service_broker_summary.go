package v2action

// ServiceBrokerSummary represents a summary of a service broker and its service instances.
type ServiceBrokerSummary struct {
	ServiceBroker
	Services []ServiceSummary
}
