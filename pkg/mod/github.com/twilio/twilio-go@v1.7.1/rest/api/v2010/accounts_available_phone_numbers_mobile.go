/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListAvailablePhoneNumberMobile'
type ListAvailablePhoneNumberMobileParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
	AreaCode *int `json:"AreaCode,omitempty"`
	// The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
	Contains *string `json:"Contains,omitempty"`
	// Whether the phone numbers can receive text messages. Can be: `true` or `false`.
	SmsEnabled *bool `json:"SmsEnabled,omitempty"`
	// Whether the phone numbers can receive MMS messages. Can be: `true` or `false`.
	MmsEnabled *bool `json:"MmsEnabled,omitempty"`
	// Whether the phone numbers can receive calls. Can be: `true` or `false`.
	VoiceEnabled *bool `json:"VoiceEnabled,omitempty"`
	// Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeAllAddressRequired *bool `json:"ExcludeAllAddressRequired,omitempty"`
	// Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeLocalAddressRequired *bool `json:"ExcludeLocalAddressRequired,omitempty"`
	// Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeForeignAddressRequired *bool `json:"ExcludeForeignAddressRequired,omitempty"`
	// Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
	Beta *bool `json:"Beta,omitempty"`
	// Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
	NearNumber *string `json:"NearNumber,omitempty"`
	// Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada.
	NearLatLong *string `json:"NearLatLong,omitempty"`
	// The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada.
	Distance *int `json:"Distance,omitempty"`
	// Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
	InPostalCode *string `json:"InPostalCode,omitempty"`
	// Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
	InRegion *string `json:"InRegion,omitempty"`
	// Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada.
	InRateCenter *string `json:"InRateCenter,omitempty"`
	// Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
	InLata *string `json:"InLata,omitempty"`
	// Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
	InLocality *string `json:"InLocality,omitempty"`
	// Whether the phone numbers can receive faxes. Can be: `true` or `false`.
	FaxEnabled *bool `json:"FaxEnabled,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAvailablePhoneNumberMobileParams) SetPathAccountSid(PathAccountSid string) *ListAvailablePhoneNumberMobileParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetAreaCode(AreaCode int) *ListAvailablePhoneNumberMobileParams {
	params.AreaCode = &AreaCode
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetContains(Contains string) *ListAvailablePhoneNumberMobileParams {
	params.Contains = &Contains
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetSmsEnabled(SmsEnabled bool) *ListAvailablePhoneNumberMobileParams {
	params.SmsEnabled = &SmsEnabled
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetMmsEnabled(MmsEnabled bool) *ListAvailablePhoneNumberMobileParams {
	params.MmsEnabled = &MmsEnabled
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetVoiceEnabled(VoiceEnabled bool) *ListAvailablePhoneNumberMobileParams {
	params.VoiceEnabled = &VoiceEnabled
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetExcludeAllAddressRequired(ExcludeAllAddressRequired bool) *ListAvailablePhoneNumberMobileParams {
	params.ExcludeAllAddressRequired = &ExcludeAllAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetExcludeLocalAddressRequired(ExcludeLocalAddressRequired bool) *ListAvailablePhoneNumberMobileParams {
	params.ExcludeLocalAddressRequired = &ExcludeLocalAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetExcludeForeignAddressRequired(ExcludeForeignAddressRequired bool) *ListAvailablePhoneNumberMobileParams {
	params.ExcludeForeignAddressRequired = &ExcludeForeignAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetBeta(Beta bool) *ListAvailablePhoneNumberMobileParams {
	params.Beta = &Beta
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetNearNumber(NearNumber string) *ListAvailablePhoneNumberMobileParams {
	params.NearNumber = &NearNumber
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetNearLatLong(NearLatLong string) *ListAvailablePhoneNumberMobileParams {
	params.NearLatLong = &NearLatLong
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetDistance(Distance int) *ListAvailablePhoneNumberMobileParams {
	params.Distance = &Distance
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetInPostalCode(InPostalCode string) *ListAvailablePhoneNumberMobileParams {
	params.InPostalCode = &InPostalCode
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetInRegion(InRegion string) *ListAvailablePhoneNumberMobileParams {
	params.InRegion = &InRegion
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetInRateCenter(InRateCenter string) *ListAvailablePhoneNumberMobileParams {
	params.InRateCenter = &InRateCenter
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetInLata(InLata string) *ListAvailablePhoneNumberMobileParams {
	params.InLata = &InLata
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetInLocality(InLocality string) *ListAvailablePhoneNumberMobileParams {
	params.InLocality = &InLocality
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetFaxEnabled(FaxEnabled bool) *ListAvailablePhoneNumberMobileParams {
	params.FaxEnabled = &FaxEnabled
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetPageSize(PageSize int) *ListAvailablePhoneNumberMobileParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAvailablePhoneNumberMobileParams) SetLimit(Limit int) *ListAvailablePhoneNumberMobileParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AvailablePhoneNumberMobile records from the API. Request is executed immediately.
func (c *ApiService) PageAvailablePhoneNumberMobile(CountryCode string, params *ListAvailablePhoneNumberMobileParams, pageToken, pageNumber string) (*ListAvailablePhoneNumberMobileResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Mobile.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AreaCode != nil {
		data.Set("AreaCode", fmt.Sprint(*params.AreaCode))
	}
	if params != nil && params.Contains != nil {
		data.Set("Contains", *params.Contains)
	}
	if params != nil && params.SmsEnabled != nil {
		data.Set("SmsEnabled", fmt.Sprint(*params.SmsEnabled))
	}
	if params != nil && params.MmsEnabled != nil {
		data.Set("MmsEnabled", fmt.Sprint(*params.MmsEnabled))
	}
	if params != nil && params.VoiceEnabled != nil {
		data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
	}
	if params != nil && params.ExcludeAllAddressRequired != nil {
		data.Set("ExcludeAllAddressRequired", fmt.Sprint(*params.ExcludeAllAddressRequired))
	}
	if params != nil && params.ExcludeLocalAddressRequired != nil {
		data.Set("ExcludeLocalAddressRequired", fmt.Sprint(*params.ExcludeLocalAddressRequired))
	}
	if params != nil && params.ExcludeForeignAddressRequired != nil {
		data.Set("ExcludeForeignAddressRequired", fmt.Sprint(*params.ExcludeForeignAddressRequired))
	}
	if params != nil && params.Beta != nil {
		data.Set("Beta", fmt.Sprint(*params.Beta))
	}
	if params != nil && params.NearNumber != nil {
		data.Set("NearNumber", *params.NearNumber)
	}
	if params != nil && params.NearLatLong != nil {
		data.Set("NearLatLong", *params.NearLatLong)
	}
	if params != nil && params.Distance != nil {
		data.Set("Distance", fmt.Sprint(*params.Distance))
	}
	if params != nil && params.InPostalCode != nil {
		data.Set("InPostalCode", *params.InPostalCode)
	}
	if params != nil && params.InRegion != nil {
		data.Set("InRegion", *params.InRegion)
	}
	if params != nil && params.InRateCenter != nil {
		data.Set("InRateCenter", *params.InRateCenter)
	}
	if params != nil && params.InLata != nil {
		data.Set("InLata", *params.InLata)
	}
	if params != nil && params.InLocality != nil {
		data.Set("InLocality", *params.InLocality)
	}
	if params != nil && params.FaxEnabled != nil {
		data.Set("FaxEnabled", fmt.Sprint(*params.FaxEnabled))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAvailablePhoneNumberMobileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AvailablePhoneNumberMobile records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAvailablePhoneNumberMobile(CountryCode string, params *ListAvailablePhoneNumberMobileParams) ([]ApiV2010AvailablePhoneNumberMobile, error) {
	response, errors := c.StreamAvailablePhoneNumberMobile(CountryCode, params)

	records := make([]ApiV2010AvailablePhoneNumberMobile, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AvailablePhoneNumberMobile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAvailablePhoneNumberMobile(CountryCode string, params *ListAvailablePhoneNumberMobileParams) (chan ApiV2010AvailablePhoneNumberMobile, chan error) {
	if params == nil {
		params = &ListAvailablePhoneNumberMobileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010AvailablePhoneNumberMobile, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAvailablePhoneNumberMobile(CountryCode, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAvailablePhoneNumberMobile(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAvailablePhoneNumberMobile(response *ListAvailablePhoneNumberMobileResponse, params *ListAvailablePhoneNumberMobileParams, recordChannel chan ApiV2010AvailablePhoneNumberMobile, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AvailablePhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAvailablePhoneNumberMobileResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAvailablePhoneNumberMobileResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAvailablePhoneNumberMobileResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAvailablePhoneNumberMobileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}