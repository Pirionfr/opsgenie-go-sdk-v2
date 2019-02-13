package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"net/url"
	"strconv"
)

type ListAlertNotesRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Offset          string
	Direction       RequestDirection
	Order           Order
	Limit           uint32
	params          string
}

func (r ListAlertNotesRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r ListAlertNotesRequest) ResourcePath() string {
	return "/v2/alerts/" + r.IdentifierValue + "/notes" + r.setIdentifierToListAlertNotesParams(r)
}

func (r ListAlertNotesRequest) Method() string {
	return "GET"
}

func (r ListAlertNotesRequest) setIdentifierToListAlertNotesParams(request ListAlertNotesRequest) string {

	params := url.Values{}

	if request.IdentifierType == ALERTID {
		params.Add("identifierType", "id")
	}

	if request.IdentifierType == ALIAS {
		params.Add("identifierType", "alias")
	}

	if request.IdentifierType == TINYID {
		params.Add("identifierType", "tiny")
	}

	if request.Offset != "" {
		params.Add("offset", request.Offset)
	}

	if request.Order == Asc {
		params.Add("order", "asc")
	}
	if request.Order == Desc {
		params.Add("order", "desc")
	}

	if request.Direction == NEXT {
		params.Add("direction", "next")
	}
	if request.Direction == PREV {
		params.Add("direction", "prev")
	}

	if request.Limit != 0 {
		params.Add("limit", strconv.Itoa(int(request.Limit)))
	}

	if len(params) != 0 {
		request.params = "?" + params.Encode()
	} else {
		request.params = ""
	}

	return request.params

}