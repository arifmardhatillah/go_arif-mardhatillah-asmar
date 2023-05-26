/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
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

// AutopilotV1FieldValue struct for AutopilotV1FieldValue
type AutopilotV1FieldValue struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the FieldValue resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Field Type associated with the Field Value.
	FieldTypeSid *string `json:"field_type_sid,omitempty"`
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
	Language *string `json:"language,omitempty"`
	// The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource.
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The unique string that we created to identify the FieldValue resource.
	Sid *string `json:"sid,omitempty"`
	// The Field Value data.
	Value *string `json:"value,omitempty"`
	// The absolute URL of the FieldValue resource.
	Url *string `json:"url,omitempty"`
	// The word for which the field value is a synonym of.
	SynonymOf *string `json:"synonym_of,omitempty"`
}
