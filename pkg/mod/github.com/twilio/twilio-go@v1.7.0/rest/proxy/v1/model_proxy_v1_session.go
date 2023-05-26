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

// ProxyV1Session struct for ProxyV1Session
type ProxyV1Session struct {
	// The unique string that we created to identify the Session resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/proxy/api/service) the session is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Session resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session started.
	DateStarted *time.Time `json:"date_started,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session ended.
	DateEnded *time.Time `json:"date_ended,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session last had an interaction.
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
	DateExpiry *time.Time `json:"date_expiry,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. Supports UTF-8 characters. **This value should not have PII.**
	UniqueName *string `json:"unique_name,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The reason the Session ended.
	ClosedReason *string `json:"closed_reason,omitempty"`
	// The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
	Ttl  *int    `json:"ttl,omitempty"`
	Mode *string `json:"mode,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Session resource.
	Url *string `json:"url,omitempty"`
	// The URLs of resources related to the Session.
	Links *map[string]interface{} `json:"links,omitempty"`
}