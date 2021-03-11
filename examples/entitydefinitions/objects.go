package entitydefinitions

// THIS PACKAGE IS AUTOGENERATED DO NOT EDIT

import "github.com/beeekind/go-salesforce-sdk/types"

// EntityDefinition ...
type EntityDefinition struct {
	// ID ...
	ID string `json:"Id"`
	// DurableID ...
	DurableID string `json:"DurableId"`
	// LastModifiedDate ...
	LastModifiedDate types.Datetime `json:"LastModifiedDate"`
	// LastModifiedByID ...
	LastModifiedByID string `json:"LastModifiedById"`
	// QualifiedAPIName ...
	QualifiedAPIName string `json:"QualifiedApiName"`
	// NamespacePrefix ...
	NamespacePrefix string `json:"NamespacePrefix"`
	// DeveloperName ...
	DeveloperName string `json:"DeveloperName"`
	// MasterLabel ...
	MasterLabel string `json:"MasterLabel"`
	// Label ...
	Label string `json:"Label"`
	// PluralLabel ...
	PluralLabel string `json:"PluralLabel"`
	// DefaultCompactLayoutID ...
	DefaultCompactLayoutID string `json:"DefaultCompactLayoutId"`
	// IsCustomizable ...
	IsCustomizable bool `json:"IsCustomizable"`
	// IsApexTriggerable ...
	IsApexTriggerable bool `json:"IsApexTriggerable"`
	// IsWorkflowEnabled ...
	IsWorkflowEnabled bool `json:"IsWorkflowEnabled"`
	// IsProcessEnabled ...
	IsProcessEnabled bool `json:"IsProcessEnabled"`
	// IsCompactLayoutable ...
	IsCompactLayoutable bool `json:"IsCompactLayoutable"`
	// DeploymentStatus ...
	DeploymentStatus string `json:"DeploymentStatus"`
	// KeyPrefix ...
	KeyPrefix string `json:"KeyPrefix"`
	// IsCustomSetting ...
	IsCustomSetting bool `json:"IsCustomSetting"`
	// IsDeprecatedAndHidden ...
	IsDeprecatedAndHidden bool `json:"IsDeprecatedAndHidden"`
	// IsReplicateable ...
	IsReplicateable bool `json:"IsReplicateable"`
	// IsRetrieveable ...
	IsRetrieveable bool `json:"IsRetrieveable"`
	// IsSearchLayoutable ...
	IsSearchLayoutable bool `json:"IsSearchLayoutable"`
	// IsSearchable ...
	IsSearchable bool `json:"IsSearchable"`
	// IsTriggerable ...
	IsTriggerable bool `json:"IsTriggerable"`
	// IsIDEnabled ...
	IsIDEnabled bool `json:"IsIdEnabled"`
	// IsEverCreatable ...
	IsEverCreatable bool `json:"IsEverCreatable"`
	// IsEverUpdatable ...
	IsEverUpdatable bool `json:"IsEverUpdatable"`
	// IsEverDeletable ...
	IsEverDeletable bool `json:"IsEverDeletable"`
	// IsFeedEnabled ...
	IsFeedEnabled bool `json:"IsFeedEnabled"`
	// IsQueryable ...
	IsQueryable bool `json:"IsQueryable"`
	// IsMruEnabled ...
	IsMruEnabled bool `json:"IsMruEnabled"`
	// DetailURL ...
	DetailURL string `json:"DetailUrl"`
	// EditURL ...
	EditURL string `json:"EditUrl"`
	// NewURL ...
	NewURL string `json:"NewUrl"`
	// EditDefinitionURL ...
	EditDefinitionURL string `json:"EditDefinitionUrl"`
	// HelpSettingPageName ...
	HelpSettingPageName string `json:"HelpSettingPageName"`
	// HelpSettingPageURL ...
	HelpSettingPageURL string `json:"HelpSettingPageUrl"`
	// RunningUserEntityAccessID ...
	RunningUserEntityAccessID string `json:"RunningUserEntityAccessId"`
	// PublisherID ...
	PublisherID string `json:"PublisherId"`
	// IsLayoutable ...
	IsLayoutable bool `json:"IsLayoutable"`
	// RecordTypesSupported ...
	RecordTypesSupported string `json:"RecordTypesSupported"`
	// InternalSharingModel ...
	InternalSharingModel string `json:"InternalSharingModel"`
	// ExternalSharingModel ...
	ExternalSharingModel string `json:"ExternalSharingModel"`
	// HasSubtypes ...
	HasSubtypes bool `json:"HasSubtypes"`
	// IsSubtype ...
	IsSubtype bool `json:"IsSubtype"`
	// IsAutoActivityCaptureEnabled ...
	IsAutoActivityCaptureEnabled bool `json:"IsAutoActivityCaptureEnabled"`
	// IsInterface ...
	IsInterface bool `json:"IsInterface"`
	// ImplementsInterfaces ...
	ImplementsInterfaces string `json:"ImplementsInterfaces"`
	// ImplementedBy ...
	ImplementedBy string `json:"ImplementedBy"`
	// ExtendsInterfaces ...
	ExtendsInterfaces string `json:"ExtendsInterfaces"`
	// ExtendedBy ...
	ExtendedBy string `json:"ExtendedBy"`
	// DefaultImplementation ...
	DefaultImplementation string `json:"DefaultImplementation"`
	// Description ...
	Description string `json:"Description"`
}