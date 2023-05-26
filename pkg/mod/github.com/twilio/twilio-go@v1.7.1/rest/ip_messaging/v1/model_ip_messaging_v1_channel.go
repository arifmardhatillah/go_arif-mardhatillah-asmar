/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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

// IpMessagingV1Channel struct for IpMessagingV1Channel
type IpMessagingV1Channel struct {
	Sid           *string                 `json:"sid,omitempty"`
	AccountSid    *string                 `json:"account_sid,omitempty"`
	ServiceSid    *string                 `json:"service_sid,omitempty"`
	FriendlyName  *string                 `json:"friendly_name,omitempty"`
	UniqueName    *string                 `json:"unique_name,omitempty"`
	Attributes    *string                 `json:"attributes,omitempty"`
	Type          *string                 `json:"type,omitempty"`
	DateCreated   *time.Time              `json:"date_created,omitempty"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	CreatedBy     *string                 `json:"created_by,omitempty"`
	MembersCount  *int                    `json:"members_count,omitempty"`
	MessagesCount *int                    `json:"messages_count,omitempty"`
	Url           *string                 `json:"url,omitempty"`
	Links         *map[string]interface{} `json:"links,omitempty"`
}
