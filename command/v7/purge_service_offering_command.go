package v7

import (
	"code.cloudfoundry.org/cli/command/flag"
)

//type PurgeServiceOfferingActor interface {
//	PurgeServiceOffering(service v2action.Service) (v2action.Warnings, error)
//	GetServiceByNameAndBrokerName(serviceName, brokerName string) (v2action.Service, v2action.Warnings, error)
//}

type PurgeServiceOfferingCommand struct {
	BaseCommand

	RequiredArgs    flag.Service `positional-args:"yes"`
	ServiceBroker   string       `short:"b" description:"Purge a service from a particular service broker. Required when service name is ambiguous"`
	Force           bool         `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}  `usage:"CF_NAME purge-service-offering SERVICE [-b BROKER] [-f]\n\nWARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup."`
	relatedCommands interface{}  `related_commands:"marketplace, purge-service-instance, service-brokers"`
}

func (cmd PurgeServiceOfferingCommand) Execute(args []string) error {
	if err := cmd.SharedActor.CheckTarget(false, false); err != nil {
		return err
	}

	warnings, err := cmd.Actor.PurgeServiceOfferingByNameAndBroker(cmd.RequiredArgs.ServiceOffering, cmd.ServiceBroker)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()

	//cmd.UI.DisplayText("WARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup.\n")
	//
	//if !cmd.Force {
	//	var promptMessage string
	//	if cmd.ServiceBroker != "" {
	//		promptMessage = "Really purge service offering {{.ServiceOffering}} from broker {{.ServiceBroker}} from Cloud Foundry?"
	//	} else {
	//		promptMessage = "Really purge service offering {{.ServiceOffering}} from Cloud Foundry?"
	//	}
	//
	//	purgeServiceOffering, promptErr := cmd.UI.DisplayBoolPrompt(false, promptMessage, map[string]interface{}{
	//		"ServiceOffering": cmd.RequiredArgs.ServiceOffering,
	//		"ServiceBroker":   cmd.ServiceBroker,
	//	})
	//	if promptErr != nil {
	//		return promptErr
	//	}
	//
	//	if !purgeServiceOffering {
	//		cmd.UI.DisplayText("Purge service offering cancelled")
	//		return nil
	//	}
	//}
	//
	//cmd.UI.DisplayText("Purging service {{.ServiceOffering}}...", map[string]interface{}{
	//	"ServiceOffering": cmd.RequiredArgs.ServiceOffering,
	//})

	return nil
}
