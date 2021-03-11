package contacts

// THIS PACKAGE IS AUTOGENERATED DO NOT EDIT

import "github.com/beeekind/go-salesforce-sdk/types"

// Contact ...
//
// Represents a contact, which is a person associated with an account.
type Contact struct {
	// AcceptedEventRelations ...
	AcceptedEventRelations struct {
		Done      bool                     `json:"done"`
		Count     int                      `json:"count"`
		TotalSize int                      `json:"totalSize"`
		Records   []*AcceptedEventRelation `json:"records"`
	} `json:"AcceptedEventRelations"`
	// Account ...
	Account *Account `json:"Account"`
	// AccountContactRoles ...
	AccountContactRoles struct {
		Done      bool                  `json:"done"`
		Count     int                   `json:"count"`
		TotalSize int                   `json:"totalSize"`
		Records   []*AccountContactRole `json:"records"`
	} `json:"AccountContactRoles"`
	// AccountID ...
	//
	// ID of the account that’s the parent of this contact. We recommend that you update up to 50 contacts
	// simultaneously when changing the accounts on contacts enabled for a Customer Portal or partner portal. We also recommend
	// that you make this update after business hours.
	//
	// Properties:Create, Filter, Group, Nillable, Sort,
	// Update
	AccountID string `json:"AccountId"`
	// ActivityHistories ...
	ActivityHistories struct {
		Done      bool               `json:"done"`
		Count     int                `json:"count"`
		TotalSize int                `json:"totalSize"`
		Records   []*ActivityHistory `json:"records"`
	} `json:"ActivityHistories"`
	// Assets ...
	Assets struct {
		Done      bool     `json:"done"`
		Count     int      `json:"count"`
		TotalSize int      `json:"totalSize"`
		Records   []*Asset `json:"records"`
	} `json:"Assets"`
	// AssistantName ...
	//
	// The assistant’s name.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	AssistantName types.NullableString `json:"AssistantName"`
	// AssistantPhone ...
	//
	// The assistant’s telephone number.
	//
	// Properties:Create, Filter, Group, Nillable, Sort,
	// Update
	AssistantPhone types.NullableString `json:"AssistantPhone"`
	// AttachedContentDocuments ...
	AttachedContentDocuments struct {
		Done      bool                       `json:"done"`
		Count     int                        `json:"count"`
		TotalSize int                        `json:"totalSize"`
		Records   []*AttachedContentDocument `json:"records"`
	} `json:"AttachedContentDocuments"`
	// Attachments ...
	Attachments struct {
		Done      bool          `json:"done"`
		Count     int           `json:"count"`
		TotalSize int           `json:"totalSize"`
		Records   []*Attachment `json:"records"`
	} `json:"Attachments"`
	// AuthorizationFormConsents ...
	AuthorizationFormConsents struct {
		Done      bool                        `json:"done"`
		Count     int                         `json:"count"`
		TotalSize int                         `json:"totalSize"`
		Records   []*AuthorizationFormConsent `json:"records"`
	} `json:"AuthorizationFormConsents"`
	// Birthdate ...
	//
	// The contact’s birthdate.Filter criteria for report filters, list view filters, and SOQL queries
	// ignore the year portion of the Birthdate field. For example, this SOQL query returns contacts with birthdays
	// later in the year than today:SELECT Name, BirthdateFROM ContactWHERE Birthdate > TODAY
	//
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	Birthdate types.Date `json:"Birthdate"`
	// CampaignMembers ...
	CampaignMembers struct {
		Done      bool              `json:"done"`
		Count     int               `json:"count"`
		TotalSize int               `json:"totalSize"`
		Records   []*CampaignMember `json:"records"`
	} `json:"CampaignMembers"`
	// CaseContactRoles ...
	CaseContactRoles struct {
		Done      bool               `json:"done"`
		Count     int                `json:"count"`
		TotalSize int                `json:"totalSize"`
		Records   []*CaseContactRole `json:"records"`
	} `json:"CaseContactRoles"`
	// Cases ...
	Cases struct {
		Done      bool    `json:"done"`
		Count     int     `json:"count"`
		TotalSize int     `json:"totalSize"`
		Records   []*Case `json:"records"`
	} `json:"Cases"`
	// CleanStatus ...
	//
	// Indicates the record’s clean status as compared with Data.com. Values include: Matched, Different,
	// Acknowledged, NotFound, Inactive, Pending, SelectMatch, or Skipped.Several values for CleanStatus appear with
	// different labels on the contact record. Matched appears as In Sync Acknowledged appears as Reviewed Pending appears
	// as Not Compared
	//
	// Properties:Create, Filter, Group, Nillable, Restricted picklist, Sort, Update
	CleanStatus types.NullableString `json:"CleanStatus"`
	// CombinedAttachments ...
	CombinedAttachments struct {
		Done      bool                  `json:"done"`
		Count     int                   `json:"count"`
		TotalSize int                   `json:"totalSize"`
		Records   []*CombinedAttachment `json:"records"`
	} `json:"CombinedAttachments"`
	// CommSubscriptionConsents ...
	CommSubscriptionConsents struct {
		Done      bool                       `json:"done"`
		Count     int                        `json:"count"`
		TotalSize int                        `json:"totalSize"`
		Records   []*CommSubscriptionConsent `json:"records"`
	} `json:"CommSubscriptionConsents"`
	// ContactCleanInfos ...
	ContactCleanInfos struct {
		Done      bool                `json:"done"`
		Count     int                 `json:"count"`
		TotalSize int                 `json:"totalSize"`
		Records   []*ContactCleanInfo `json:"records"`
	} `json:"ContactCleanInfos"`
	// ContactRequests ...
	ContactRequests struct {
		Done      bool              `json:"done"`
		Count     int               `json:"count"`
		TotalSize int               `json:"totalSize"`
		Records   []*ContactRequest `json:"records"`
	} `json:"ContactRequests"`
	// ContentDocumentLinks ...
	ContentDocumentLinks struct {
		Done      bool                   `json:"done"`
		Count     int                    `json:"count"`
		TotalSize int                    `json:"totalSize"`
		Records   []*ContentDocumentLink `json:"records"`
	} `json:"ContentDocumentLinks"`
	// ContractContactRoles ...
	ContractContactRoles struct {
		Done      bool                   `json:"done"`
		Count     int                    `json:"count"`
		TotalSize int                    `json:"totalSize"`
		Records   []*ContractContactRole `json:"records"`
	} `json:"ContractContactRoles"`
	// ContractsSigned ...
	ContractsSigned struct {
		Done      bool        `json:"done"`
		Count     int         `json:"count"`
		TotalSize int         `json:"totalSize"`
		Records   []*Contract `json:"records"`
	} `json:"ContractsSigned"`
	// CreatedBy ...
	CreatedBy *User `json:"CreatedBy"`
	// CreatedByID ...
	CreatedByID string `json:"CreatedById"`
	// CreatedDate ...
	CreatedDate types.Datetime `json:"CreatedDate"`
	// CreditMemos ...
	CreditMemos struct {
		Done      bool          `json:"done"`
		Count     int           `json:"count"`
		TotalSize int           `json:"totalSize"`
		Records   []*CreditMemo `json:"records"`
	} `json:"CreditMemos"`
	// DeclinedEventRelations ...
	DeclinedEventRelations struct {
		Done      bool                     `json:"done"`
		Count     int                      `json:"count"`
		TotalSize int                      `json:"totalSize"`
		Records   []*DeclinedEventRelation `json:"records"`
	} `json:"DeclinedEventRelations"`
	// Department ...
	//
	// The contact’s department.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	Department types.NullableString `json:"Department"`
	// Description ...
	//
	// A description of the contact. Label is Contact Description up to 32 KB.
	//
	// Properties:Create,
	// Nillable, Update
	Description types.NullableString `json:"Description"`
	// DuplicateRecordItems ...
	DuplicateRecordItems struct {
		Done      bool                   `json:"done"`
		Count     int                    `json:"count"`
		TotalSize int                    `json:"totalSize"`
		Records   []*DuplicateRecordItem `json:"records"`
	} `json:"DuplicateRecordItems"`
	// Email ...
	//
	// The contact’s email address.
	//
	// Properties:Create, Filter, Group, idLookup, Nillable, Sort,
	// Update
	Email types.NullableString `json:"Email"`
	// EmailBouncedDate ...
	//
	// If bounce management is activated and an email sent to the contact bounces, the date and time of the
	// bounce.
	//
	// Properties:Create, Filter, Nillable, Sort, Update
	EmailBouncedDate types.Datetime `json:"EmailBouncedDate"`
	// EmailBouncedReason ...
	//
	// If bounce management is activated and an email sent to the contact bounces, the reason for the
	// bounce.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	EmailBouncedReason types.NullableString `json:"EmailBouncedReason"`
	// EmailMessageRelations ...
	EmailMessageRelations struct {
		Done      bool                    `json:"done"`
		Count     int                     `json:"count"`
		TotalSize int                     `json:"totalSize"`
		Records   []*EmailMessageRelation `json:"records"`
	} `json:"EmailMessageRelations"`
	// EmailStatuses ...
	EmailStatuses struct {
		Done      bool           `json:"done"`
		Count     int            `json:"count"`
		TotalSize int            `json:"totalSize"`
		Records   []*EmailStatus `json:"records"`
	} `json:"EmailStatuses"`
	// EventRelations ...
	EventRelations struct {
		Done      bool             `json:"done"`
		Count     int              `json:"count"`
		TotalSize int              `json:"totalSize"`
		Records   []*EventRelation `json:"records"`
	} `json:"EventRelations"`
	// Events ...
	Events struct {
		Done      bool     `json:"done"`
		Count     int      `json:"count"`
		TotalSize int      `json:"totalSize"`
		Records   []*Event `json:"records"`
	} `json:"Events"`
	// Fax ...
	//
	// The contact’s fax number. Label is Business Fax.
	//
	// Properties:Create, Filter, Group,
	// Nillable, Sort, Update
	Fax types.NullableString `json:"Fax"`
	// FeedSubscriptionsForEntity ...
	FeedSubscriptionsForEntity struct {
		Done      bool                  `json:"done"`
		Count     int                   `json:"count"`
		TotalSize int                   `json:"totalSize"`
		Records   []*EntitySubscription `json:"records"`
	} `json:"FeedSubscriptionsForEntity"`
	// Feeds ...
	Feeds struct {
		Done      bool           `json:"done"`
		Count     int            `json:"count"`
		TotalSize int            `json:"totalSize"`
		Records   []*ContactFeed `json:"records"`
	} `json:"Feeds"`
	// FirstName ...
	//
	// The contact’s first name up to 40 characters.
	//
	// Properties:Create, Filter, Group, Nillable,
	// Sort, Update
	FirstName types.NullableString `json:"FirstName"`
	// Histories ...
	Histories struct {
		Done      bool              `json:"done"`
		Count     int               `json:"count"`
		TotalSize int               `json:"totalSize"`
		Records   []*ContactHistory `json:"records"`
	} `json:"Histories"`
	// HomePhone ...
	//
	// The contact’s home telephone number.
	//
	// Properties:Create, Filter, Group, Nillable, Sort,
	// Update
	HomePhone types.NullableString `json:"HomePhone"`
	// ID ...
	ID string `json:"Id"`
	// Individual ...
	Individual *Individual `json:"Individual"`
	// IndividualID ...
	//
	// ID of the data privacy record associated with this contact. This field is available if Data Protection and
	// Privacy is enabled.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	IndividualID string `json:"IndividualId"`
	// Invoices ...
	Invoices struct {
		Done      bool       `json:"done"`
		Count     int        `json:"count"`
		TotalSize int        `json:"totalSize"`
		Records   []*Invoice `json:"records"`
	} `json:"Invoices"`
	// IsDeleted ...
	//
	// Indicates whether the object has been moved to the Recycle Bin (true) or not (false). Label is
	// Deleted.
	//
	// Properties:Defaulted on create, Filter
	IsDeleted bool `json:"IsDeleted"`
	// IsEmailBounced ...
	//
	// If bounce management is activated and an email is sent to a contact, indicates whether the email bounced
	// (true) or not (false).
	//
	// Properties:Defaulted on create, Filter, Group, Sort
	IsEmailBounced bool `json:"IsEmailBounced"`
	// Jigsaw ...
	//
	// References the company’s ID in Data.com. If an account has a value in this field, it means that the account
	// was imported from Data.com. If the field value is null, the account was not imported from Data.com. Maximum
	// size is 20 characters. Available in API version 22.0 and later. Label is Data.com Key. Important The Jigsaw
	// field is exposed in the API to support troubleshooting for import errors and reimporting of corrected data. Do
	// not modify this value.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	Jigsaw types.NullableString `json:"Jigsaw"`
	// JigsawContactID ...
	JigsawContactID types.NullableString `json:"JigsawContactId"`
	// Languages__c ...
	Languages__c types.NullableString `json:"Languages__c"`
	// LastActivityDate ...
	//
	// Value is the most recent of either: Due date of the most recent event logged against the record. Due date of
	// the most recently closed task associated with the record.
	//
	// Properties:Filter, Group, Nillable,
	// Sort
	LastActivityDate types.Date `json:"LastActivityDate"`
	// LastCURequestDate ...
	LastCURequestDate types.Datetime `json:"LastCURequestDate"`
	// LastCUUpdateDate ...
	LastCUUpdateDate types.Datetime `json:"LastCUUpdateDate"`
	// LastModifiedBy ...
	LastModifiedBy *User `json:"LastModifiedBy"`
	// LastModifiedByID ...
	LastModifiedByID string `json:"LastModifiedById"`
	// LastModifiedDate ...
	LastModifiedDate types.Datetime `json:"LastModifiedDate"`
	// LastName ...
	//
	// Required. Last name of the contact up to 80 characters.
	//
	// Properties:Create, Filter, Group, Sort,
	// Update
	LastName string `json:"LastName"`
	// LastReferencedDate ...
	//
	// The timestamp when the current user last accessed this record, a record related to this record, or a list
	// view.
	//
	// Properties:Filter, Nillable, Sort
	LastReferencedDate types.Datetime `json:"LastReferencedDate"`
	// LastViewedDate ...
	//
	// The timestamp when the current user last viewed this record or list view. If this value is null, the user
	// might have only accessed this record or list view (LastReferencedDate) but not viewed it.
	//
	//
	// Properties:Filter, Nillable, Sort
	LastViewedDate types.Datetime `json:"LastViewedDate"`
	// LeadSource ...
	//
	// The lead’s source.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	LeadSource types.NullableString `json:"LeadSource"`
	// Level__c ...
	Level__c types.NullableString `json:"Level__c"`
	// ListEmailIndividualRecipients ...
	ListEmailIndividualRecipients struct {
		Done      bool                            `json:"done"`
		Count     int                             `json:"count"`
		TotalSize int                             `json:"totalSize"`
		Records   []*ListEmailIndividualRecipient `json:"records"`
	} `json:"ListEmailIndividualRecipients"`
	// MailingAddress ...
	//
	// The compound form of the mailing address. Read-only. For details on compound address fields, see Address
	// Compound Fields.
	//
	// Properties:Filter, Nillable
	MailingAddress types.Address `json:"MailingAddress"`
	// MailingCity ...
	//
	// Mailing address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MailingCity types.NullableString `json:"MailingCity"`
	// MailingCountry ...
	//
	// Mailing address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MailingCountry types.NullableString `json:"MailingCountry"`
	// MailingGeocodeAccuracy ...
	//
	// Accuracy level of the geocode for the mailing address. For details on geolocation compound field, see
	// Compound Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Group, Nillable,
	// Restricted picklist, Sort, Update, Query, Restricted picklist, Nillable
	MailingGeocodeAccuracy types.NullableString `json:"MailingGeocodeAccuracy"`
	// MailingLatitude ...
	//
	// Used with MailingLongitude to specify the precise geolocation of a mailing address. Acceptable values
	// are numbers between –90 and 90 up to 15 decimal places. For details on geolocation compound fields, see
	// Compound Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Nillable, Sort, Update
	MailingLatitude types.NullableFloat64 `json:"MailingLatitude"`
	// MailingLongitude ...
	//
	// Used with MailingLatitude to specify the precise geolocation of a mailing address. Acceptable values are
	// numbers between –180 and 180 up to 15 decimal places. For details on geolocation compound fields, see Compound
	// Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Nillable, Sort, Update
	MailingLongitude types.NullableFloat64 `json:"MailingLongitude"`
	// MailingPostalCode ...
	//
	// Mailing address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MailingPostalCode types.NullableString `json:"MailingPostalCode"`
	// MailingState ...
	//
	// Mailing address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MailingState types.NullableString `json:"MailingState"`
	// MailingStreet ...
	//
	// Street address for mailing address.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MailingStreet types.NullableString `json:"MailingStreet"`
	// MasterRecord ...
	MasterRecord *Contact `json:"MasterRecord"`
	// MasterRecordID ...
	//
	// If this record was deleted as the result of a merge, this field contains the ID of the record that remains. If
	// this record was deleted for any other reason, or has not been deleted, the value is null.
	//
	//
	// Properties:Filter, Group, Nillable, Sort
	MasterRecordID string `json:"MasterRecordId"`
	// MessagingSessions ...
	MessagingSessions struct {
		Done      bool                `json:"done"`
		Count     int                 `json:"count"`
		TotalSize int                 `json:"totalSize"`
		Records   []*MessagingSession `json:"records"`
	} `json:"MessagingSessions"`
	// MobilePhone ...
	//
	// Contact’s mobile phone number.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	MobilePhone types.NullableString `json:"MobilePhone"`
	// Name ...
	//
	// Concatenation of FirstName, MiddleName, LastName, and Suffix up to 203 characters, including
	// whitespaces.
	//
	// Properties:Filter, Group, Sort
	Name string `json:"Name"`
	// Notes ...
	Notes struct {
		Done      bool    `json:"done"`
		Count     int     `json:"count"`
		TotalSize int     `json:"totalSize"`
		Records   []*Note `json:"records"`
	} `json:"Notes"`
	// NotesAndAttachments ...
	NotesAndAttachments struct {
		Done      bool                 `json:"done"`
		Count     int                  `json:"count"`
		TotalSize int                  `json:"totalSize"`
		Records   []*NoteAndAttachment `json:"records"`
	} `json:"NotesAndAttachments"`
	// OpenActivities ...
	OpenActivities struct {
		Done      bool            `json:"done"`
		Count     int             `json:"count"`
		TotalSize int             `json:"totalSize"`
		Records   []*OpenActivity `json:"records"`
	} `json:"OpenActivities"`
	// Opportunities ...
	Opportunities struct {
		Done      bool           `json:"done"`
		Count     int            `json:"count"`
		TotalSize int            `json:"totalSize"`
		Records   []*Opportunity `json:"records"`
	} `json:"Opportunities"`
	// OpportunityContactRoles ...
	OpportunityContactRoles struct {
		Done      bool                      `json:"done"`
		Count     int                       `json:"count"`
		TotalSize int                       `json:"totalSize"`
		Records   []*OpportunityContactRole `json:"records"`
	} `json:"OpportunityContactRoles"`
	// OtherAddress ...
	//
	// The compound form of the other address. Read-only. For details on compound address fields, see Address
	// Compound Fields.
	//
	// Properties:Filter, Nillable
	OtherAddress types.Address `json:"OtherAddress"`
	// OtherCity ...
	//
	// Alternate address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherCity types.NullableString `json:"OtherCity"`
	// OtherCountry ...
	//
	// Alternate address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherCountry types.NullableString `json:"OtherCountry"`
	// OtherGeocodeAccuracy ...
	//
	// Accuracy level of the geocode for the other address. For details on geolocation compound fields, see
	// Compound Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Group, Nillable,
	// Restricted picklist, Sort, Update
	OtherGeocodeAccuracy types.NullableString `json:"OtherGeocodeAccuracy"`
	// OtherLatitude ...
	//
	// Used with OtherLongitude to specify the precise geolocation of an alternate address. Acceptable values
	// are numbers between –90 and 90 up to 15 decimal places. For details on geolocation compound fields, see
	// Compound Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Nillable, Sort, Update
	OtherLatitude types.NullableFloat64 `json:"OtherLatitude"`
	// OtherLongitude ...
	//
	// Used with OtherLatitude to specify the precise geolocation of an alternate address. Acceptable values
	// are numbers between –180 and 180 up to 15 decimal places. For details on geolocation compound fields, see
	// Compound Field Considerations and Limitations.
	//
	// Properties:Create, Filter, Nillable, Sort, Update
	OtherLongitude types.NullableFloat64 `json:"OtherLongitude"`
	// OtherPhone ...
	//
	// Telephone for alternate address.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherPhone types.NullableString `json:"OtherPhone"`
	// OtherPostalCode ...
	//
	// Alternate address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherPostalCode types.NullableString `json:"OtherPostalCode"`
	// OtherState ...
	//
	// Alternate address details.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherState types.NullableString `json:"OtherState"`
	// OtherStreet ...
	//
	// Street for alternate address.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	OtherStreet types.NullableString `json:"OtherStreet"`
	// OutgoingEmailRelations ...
	OutgoingEmailRelations struct {
		Done      bool                     `json:"done"`
		Count     int                      `json:"count"`
		TotalSize int                      `json:"totalSize"`
		Records   []*OutgoingEmailRelation `json:"records"`
	} `json:"OutgoingEmailRelations"`
	// Owner ...
	Owner *User `json:"Owner"`
	// OwnerID ...
	//
	// The ID of the owner of the account associated with this contact.
	//
	// Properties:Create, Defaulted on
	// create, Filter, Group, Sort, Update
	OwnerID string `json:"OwnerId"`
	// PersonRecord ...
	PersonRecord struct {
		Done      bool                        `json:"done"`
		Count     int                         `json:"count"`
		TotalSize int                         `json:"totalSize"`
		Records   []*UserEmailPreferredPerson `json:"records"`
	} `json:"PersonRecord"`
	// Phone ...
	//
	// Telephone number for the contact. Label is Business Phone.
	//
	// Properties:Create, Filter, Group,
	// Nillable, Sort, Update
	Phone types.NullableString `json:"Phone"`
	// PhotoURL ...
	//
	// Path to be combined with the URL of a Salesforce instance (Example:
	// https://yourInstance.salesforce.com/) to generate a URL to request the social network profile image associated with the contact. Generated URL
	// returns an HTTP redirect (code 302) to the social network profile image for the contact. Empty if Social Accounts
	// and Contacts isn't enabled or if Social Accounts and Contacts is disabled for the requesting user.
	//
	//
	// Properties:Filter, Group, Nillable, Sort
	PhotoURL types.NullableString `json:"PhotoUrl"`
	// ProcessInstances ...
	ProcessInstances struct {
		Done      bool               `json:"done"`
		Count     int                `json:"count"`
		TotalSize int                `json:"totalSize"`
		Records   []*ProcessInstance `json:"records"`
	} `json:"ProcessInstances"`
	// ProcessSteps ...
	ProcessSteps struct {
		Done      bool                      `json:"done"`
		Count     int                       `json:"count"`
		TotalSize int                       `json:"totalSize"`
		Records   []*ProcessInstanceHistory `json:"records"`
	} `json:"ProcessSteps"`
	// RecordActionHistories ...
	RecordActionHistories struct {
		Done      bool                   `json:"done"`
		Count     int                    `json:"count"`
		TotalSize int                    `json:"totalSize"`
		Records   []*RecordActionHistory `json:"records"`
	} `json:"RecordActionHistories"`
	// RecordActions ...
	RecordActions struct {
		Done      bool            `json:"done"`
		Count     int             `json:"count"`
		TotalSize int             `json:"totalSize"`
		Records   []*RecordAction `json:"records"`
	} `json:"RecordActions"`
	// RecordAssociatedGroups ...
	RecordAssociatedGroups struct {
		Done      bool                        `json:"done"`
		Count     int                         `json:"count"`
		TotalSize int                         `json:"totalSize"`
		Records   []*CollaborationGroupRecord `json:"records"`
	} `json:"RecordAssociatedGroups"`
	// ReportsTo ...
	ReportsTo *Contact `json:"ReportsTo"`
	// ReportsToID ...
	//
	// This field doesn’t appear if IsPersonAccount is true.
	//
	// Properties:Create, Filter, Group,
	// Nillable, Sort, Update
	ReportsToID string `json:"ReportsToId"`
	// ReturnOrders ...
	ReturnOrders struct {
		Done      bool           `json:"done"`
		Count     int            `json:"count"`
		TotalSize int            `json:"totalSize"`
		Records   []*ReturnOrder `json:"records"`
	} `json:"ReturnOrders"`
	// Salutation ...
	//
	// Honorific abbreviation, word, or phrase to be used in front of name in greetings, such as Dr. or
	// Mrs.
	//
	// Properties:Create, Filter, Group, Nillable, Sort, Update
	Salutation types.NullableString `json:"Salutation"`
	// ServiceAppointments ...
	ServiceAppointments struct {
		Done      bool                  `json:"done"`
		Count     int                   `json:"count"`
		TotalSize int                   `json:"totalSize"`
		Records   []*ServiceAppointment `json:"records"`
	} `json:"ServiceAppointments"`
	// Shares ...
	Shares struct {
		Done      bool            `json:"done"`
		Count     int             `json:"count"`
		TotalSize int             `json:"totalSize"`
		Records   []*ContactShare `json:"records"`
	} `json:"Shares"`
	// SystemModstamp ...
	SystemModstamp types.Datetime `json:"SystemModstamp"`
	// Tasks ...
	Tasks struct {
		Done      bool    `json:"done"`
		Count     int     `json:"count"`
		TotalSize int     `json:"totalSize"`
		Records   []*Task `json:"records"`
	} `json:"Tasks"`
	// Title ...
	//
	// Title of the contact, such as CEO or Vice President.
	//
	// Properties:Create, Filter, Group,
	// Nillable, Sort, Update
	Title types.NullableString `json:"Title"`
	// TopicAssignments ...
	TopicAssignments struct {
		Done      bool               `json:"done"`
		Count     int                `json:"count"`
		TotalSize int                `json:"totalSize"`
		Records   []*TopicAssignment `json:"records"`
	} `json:"TopicAssignments"`
	// UndecidedEventRelations ...
	UndecidedEventRelations struct {
		Done      bool                      `json:"done"`
		Count     int                       `json:"count"`
		TotalSize int                       `json:"totalSize"`
		Records   []*UndecidedEventRelation `json:"records"`
	} `json:"UndecidedEventRelations"`
	// Users ...
	Users struct {
		Done      bool    `json:"done"`
		Count     int     `json:"count"`
		TotalSize int     `json:"totalSize"`
		Records   []*User `json:"records"`
	} `json:"Users"`
}
