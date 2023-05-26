/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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

// ProxyV1Service struct for ProxyV1Service
type ProxyV1Service struct {
	// The unique string that we created to identify the Service resource.
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. Supports UTF-8 characters. **This value should not have PII.**
	UniqueName *string `json:"unique_name,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
	ChatInstanceSid *string `json:"chat_instance_sid,omitempty"`
	// The URL we call when the interaction status changes.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// The default `ttl` value for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value.
	DefaultTtl              *int    `json:"default_ttl,omitempty"`
	NumberSelectionBehavior *string `json:"number_selection_behavior,omitempty"`
	GeoMatchLevel           *string `json:"geo_match_level,omitempty"`
	// The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
	InterceptCallbackUrl *string `json:"intercept_callback_url,omitempty"`
	// The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
	OutOfSessionCallbackUrl *string `json:"out_of_session_callback_url,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Service resource.
	Url *string `json:"url,omitempty"`
	// The URLs of resources related to the Service.
	Links *map[string]interface{} `json:"links,omitempty"`
}
