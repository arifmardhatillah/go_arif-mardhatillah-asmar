/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ServerlessV1FunctionVersion struct for ServerlessV1FunctionVersion
type ServerlessV1FunctionVersion struct {
	// The unique string that we created to identify the Function Version resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Function Version resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Function Version resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the Function resource that is the parent of the Function Version resource.
	FunctionSid *string `json:"function_sid,omitempty"`
	// The URL-friendly string by which the Function Version resource can be referenced. It can be a maximum of 255 characters. All paths begin with a forward slash ('/'). If a Function Version creation request is submitted with a path not containing a leading slash, the path will automatically be prepended with one.
	Path       *string `json:"path,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	// The date and time in GMT when the Function Version resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Function Version resource.
	Url   *string                 `json:"url,omitempty"`
	Links *map[string]interface{} `json:"links,omitempty"`
}
