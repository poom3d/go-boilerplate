// GENERATED FILE: DO NOT EDIT!

package audiences

// Types used by the API.
// implements the service definition of Audience
type Audience struct {
	AudienceId     string `json:"audienceId,omitempty"`
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	EstimatedUsers int32  `json:"estimatedUsers,omitempty"`
}

// implements the service definition of Audiences
type Audiences struct {
	Data []Audience `json:"data,omitempty"`
}

// implements the service definition of AgeRange
type AgeRange struct {
	From int32  `json:"from,omitempty"`
	To   int32  `json:"to,omitempty"`
	Type string `json:"type,omitempty"`
}

// implements the service definition of Campaign
type Campaign struct {
	CampaignId string `json:"campaignId,omitempty"`
	LinkId     string `json:"linkId,omitempty"`
	Action     string `json:"action,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	Condition  string `json:"condition,omitempty"`
}

// implements the service definition of Demographic
type Demographic struct {
	Gender   string   `json:"gender,omitempty"`
	Age      AgeRange `json:"age,omitempty"`
	Location string   `json:"location,omitempty"`
	Email    string   `json:"email,omitempty"`
}

// implements the service definition of Interest
type Interest struct {
	Tag string `json:"tag,omitempty"`
}

// implements the service definition of Interests
type Interests struct {
	Data []Interest `json:"data,omitempty"`
}

// implements the service definition of UserStatus
type UserStatus struct {
	Status    string `json:"status,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Condition string `json:"condition,omitempty"`
}

// implements the service definition of UnauthorizedError
type UnauthorizedError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// implements the service definition of Error
type Error struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ListRecommendedAudiencesParameters holds parameters to ListRecommendedAudiences
type ListRecommendedAudiencesParameters struct {
	Limit int64 `json:"limit,omitempty"`
}

// ListRecommendedAudiencesResponses holds responses of ListRecommendedAudiences
type ListRecommendedAudiencesResponses struct {
	OK      *Audiences
	Code401 *UnauthorizedError
	Default *Error
}

// ListAudiencesParameters holds parameters to ListAudiences
type ListAudiencesParameters struct {
	Limit int64 `json:"limit,omitempty"`
}

// ListAudiencesResponses holds responses of ListAudiences
type ListAudiencesResponses struct {
	OK      *Audiences
	Code401 *UnauthorizedError
	Default *Error
}

// CreateAudienceParameters holds parameters to CreateAudience
type CreateAudienceParameters struct {
	Audience Audience `json:"audience,omitempty"`
}

// CreateAudienceResponses holds responses of CreateAudience
type CreateAudienceResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}

// ShowAudienceByIdParameters holds parameters to ShowAudienceById
type ShowAudienceByIdParameters struct {
	AudienceId string `json:"audienceId,omitempty"`
}

// ShowAudienceByIdResponses holds responses of ShowAudienceById
type ShowAudienceByIdResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}

// FilterAudienceByDemographicParameters holds parameters to FilterAudienceByDemographic
type FilterAudienceByDemographicParameters struct {
	AudienceId  string      `json:"audienceId,omitempty"`
	Demographic Demographic `json:"demographic,omitempty"`
}

// FilterAudienceByDemographicResponses holds responses of FilterAudienceByDemographic
type FilterAudienceByDemographicResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}

// FilterAudienceByInterestsParameters holds parameters to FilterAudienceByInterests
type FilterAudienceByInterestsParameters struct {
	AudienceId string    `json:"audienceId,omitempty"`
	Interests  Interests `json:"interests,omitempty"`
}

// FilterAudienceByInterestsResponses holds responses of FilterAudienceByInterests
type FilterAudienceByInterestsResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}

// FilterAudienceByUserStatusParameters holds parameters to FilterAudienceByUserStatus
type FilterAudienceByUserStatusParameters struct {
	AudienceId string     `json:"audienceId,omitempty"`
	Userstatus UserStatus `json:"userstatus,omitempty"`
}

// FilterAudienceByUserStatusResponses holds responses of FilterAudienceByUserStatus
type FilterAudienceByUserStatusResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}

// FilterAudienceByCampaignParameters holds parameters to FilterAudienceByCampaign
type FilterAudienceByCampaignParameters struct {
	AudienceId string   `json:"audienceId,omitempty"`
	Campaign   Campaign `json:"campaign,omitempty"`
}

// FilterAudienceByCampaignResponses holds responses of FilterAudienceByCampaign
type FilterAudienceByCampaignResponses struct {
	OK      *Audience
	Code401 *UnauthorizedError
	Default *Error
}
