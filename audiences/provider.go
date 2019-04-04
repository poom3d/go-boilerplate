// GENERATED FILE: DO NOT EDIT!

package audiences

// To create a server, first write a class that implements this interface.
// Then pass an instance of it to Initialize().
type Provider interface {

	//
	ListRecommendedAudiences(parameters *ListRecommendedAudiencesParameters, responses *ListRecommendedAudiencesResponses) (err error)

	//
	ListAudiences(parameters *ListAudiencesParameters, responses *ListAudiencesResponses) (err error)

	//
	CreateAudience(parameters *CreateAudienceParameters, responses *CreateAudienceResponses) (err error)

	//
	ShowAudienceById(parameters *ShowAudienceByIdParameters, responses *ShowAudienceByIdResponses) (err error)

	//
	FilterAudienceByDemographic(parameters *FilterAudienceByDemographicParameters, responses *FilterAudienceByDemographicResponses) (err error)

	//
	FilterAudienceByInterests(parameters *FilterAudienceByInterestsParameters, responses *FilterAudienceByInterestsResponses) (err error)

	//
	FilterAudienceByUserStatus(parameters *FilterAudienceByUserStatusParameters, responses *FilterAudienceByUserStatusResponses) (err error)

	//
	FilterAudienceByCampaign(parameters *FilterAudienceByCampaignParameters, responses *FilterAudienceByCampaignResponses) (err error)
}
