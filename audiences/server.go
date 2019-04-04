// GENERATED FILE: DO NOT EDIT!

package audiences

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func intValue(s string) (v int64) {
	v, _ = strconv.ParseInt(s, 10, 64)
	return v
}

// This package-global variable holds the user-written Provider for API services.
// See the Provider interface for details.
var provider Provider

// These handlers serve API methods.

// Handler
//
func HandleListRecommendedAudiences(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &ListRecommendedAudiencesParameters{}
	// get request fields in path and query parameters
	// instantiate the responses structure
	responses := &ListRecommendedAudiencesResponses{}
	// call the service provider
	err = provider.ListRecommendedAudiences(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleListAudiences(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &ListAudiencesParameters{}
	// get request fields in path and query parameters
	// instantiate the responses structure
	responses := &ListAudiencesResponses{}
	// call the service provider
	err = provider.ListAudiences(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleCreateAudience(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &CreateAudienceParameters{}
	err = c.BindJSON(&parameters.Audience)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// get request fields in path and query parameters
	// instantiate the responses structure
	responses := &CreateAudienceResponses{}
	// call the service provider
	err = provider.CreateAudience(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleShowAudienceById(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &ShowAudienceByIdParameters{}
	// get request fields in path and query parameters
	// name:"audienceId" type:"string" position:PATH nativeType:"string" fieldName:"AudienceId" parameterName:"audienceId" serialize:true
	value := c.Param("audienceId")
	parameters.AudienceId = value
	// instantiate the responses structure
	responses := &ShowAudienceByIdResponses{}
	// call the service provider
	err = provider.ShowAudienceById(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleFilterAudienceByDemographic(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &FilterAudienceByDemographicParameters{}
	err = c.BindJSON(&parameters.Demographic)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// get request fields in path and query parameters
	// name:"audienceId" type:"string" position:PATH nativeType:"string" fieldName:"AudienceId" parameterName:"audienceId" serialize:true
	value := c.Param("audienceId")
	parameters.AudienceId = value
	// instantiate the responses structure
	responses := &FilterAudienceByDemographicResponses{}
	// call the service provider
	err = provider.FilterAudienceByDemographic(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleFilterAudienceByInterests(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &FilterAudienceByInterestsParameters{}
	err = c.BindJSON(&parameters.Interests)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// get request fields in path and query parameters
	// name:"audienceId" type:"string" position:PATH nativeType:"string" fieldName:"AudienceId" parameterName:"audienceId" serialize:true
	value := c.Param("audienceId")
	parameters.AudienceId = value
	// instantiate the responses structure
	responses := &FilterAudienceByInterestsResponses{}
	// call the service provider
	err = provider.FilterAudienceByInterests(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleFilterAudienceByUserStatus(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &FilterAudienceByUserStatusParameters{}
	err = c.BindJSON(&parameters.Userstatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// get request fields in path and query parameters
	// name:"audienceId" type:"string" position:PATH nativeType:"string" fieldName:"AudienceId" parameterName:"audienceId" serialize:true
	value := c.Param("audienceId")
	parameters.AudienceId = value
	// instantiate the responses structure
	responses := &FilterAudienceByUserStatusResponses{}
	// call the service provider
	err = provider.FilterAudienceByUserStatus(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Handler
//
func HandleFilterAudienceByCampaign(c *gin.Context) {
	var err error
	// instantiate the parameters structure
	parameters := &FilterAudienceByCampaignParameters{}
	err = c.BindJSON(&parameters.Campaign)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// get request fields in path and query parameters
	// name:"audienceId" type:"string" position:PATH nativeType:"string" fieldName:"AudienceId" parameterName:"audienceId" serialize:true
	value := c.Param("audienceId")
	parameters.AudienceId = value
	// instantiate the responses structure
	responses := &FilterAudienceByCampaignResponses{}
	// call the service provider
	err = provider.FilterAudienceByCampaign(parameters, responses)
	if err == nil {
		if responses.OK != nil {
			// write the normal response
			c.JSON(200, responses.OK)
			return
		}
		if responses.Default != nil {
			// write the error response
			c.JSON(int(responses.Default.Code), responses.Default)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// Initialize the API service.
func Initialize(p Provider, engine *gin.Engine) {
	provider = p
	// method.Path=/audiences/recommended
	engine.GET("/audiences/recommended", HandleListRecommendedAudiences)
	// method.Path=/audiences
	engine.GET("/audiences", HandleListAudiences)
	// method.Path=/audiences
	engine.POST("/audiences", HandleCreateAudience)
	// method.Path=/audience/{audienceId}
	engine.GET("/audience/:audienceId", HandleShowAudienceById)
	// method.Path=/audiences/{audienceId}/demographic
	engine.POST("/audiences/:audienceId/demographic", HandleFilterAudienceByDemographic)
	// method.Path=/audiences/{audienceId}/interests
	engine.POST("/audiences/:audienceId/interests", HandleFilterAudienceByInterests)
	// method.Path=/audiences/{audienceId}/behaviours/userstatus
	engine.POST("/audiences/:audienceId/behaviours/userstatus", HandleFilterAudienceByUserStatus)
	// method.Path=/audiences/{audienceId}/behaviours/campaign
	engine.POST("/audiences/:audienceId/behaviours/campaign", HandleFilterAudienceByCampaign)
}
