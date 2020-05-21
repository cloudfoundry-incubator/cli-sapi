package resources

import (
	"code.cloudfoundry.org/cli/types"
	"code.cloudfoundry.org/jsonry"
)

type Domain struct {
	GUID             string         `json:"guid,omitempty"`
	Name             string         `json:"name"`
	Internal         types.NullBool `json:"internal,omitempty"`
	OrganizationGUID string         `jsonry:"relationships.organization.data.guid,omitempty"`
	RouterGroup      string         `jsonry:"router_group.guid,omitempty"`
	Protocols        []string       `jsonry:"supported_protocols,omitempty"`

	// Metadata is used for custom tagging of API resources
	Metadata *Metadata `json:"metadata,omitempty"`
}

func (d Domain) MarshalJSON() ([]byte, error) {
	return jsonry.Marshal(d)
}

func (d *Domain) UnmarshalJSON(data []byte) error {
	return jsonry.Unmarshal(data, d)
}

func (d Domain) Shared() bool {
	return d.OrganizationGUID == ""
}
