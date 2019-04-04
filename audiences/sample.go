package audiences

// To create a server, first write a class that implements this interface.
// Then pass an instance of it to Initialize().
type SampleProvider struct {
}

//
func (s *SampleProvider) ListRecommendedAudiences(parameters *ListRecommendedAudiencesParameters, responses *ListRecommendedAudiencesResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) ListAudiences(parameters *ListAudiencesParameters, responses *ListAudiencesResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) CreateAudience(parameters *CreateAudienceParameters, responses *CreateAudienceResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) ShowAudienceById(parameters *ShowAudienceByIdParameters, responses *ShowAudienceByIdResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) FilterAudienceByDemographic(parameters *FilterAudienceByDemographicParameters, responses *FilterAudienceByDemographicResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) FilterAudienceByInterests(parameters *FilterAudienceByInterestsParameters, responses *FilterAudienceByInterestsResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) FilterAudienceByUserStatus(parameters *FilterAudienceByUserStatusParameters, responses *FilterAudienceByUserStatusResponses) (err error) {
	return nil
}

//
func (s *SampleProvider) FilterAudienceByCampaign(parameters *FilterAudienceByCampaignParameters, responses *FilterAudienceByCampaignResponses) (err error) {
	return nil
}
