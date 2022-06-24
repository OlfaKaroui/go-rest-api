package conversation

import (
	"net/http"
	"time"

	messagebird "github.com/messagebird/go-rest-api/v7"
)

type WebhookCreateRequest struct {
	ChannelID string         `json:"channelId"`
	Events    []WebhookEvent `json:"events"`
	URL       string         `json:"url"`
}

type WebhookUpdateRequest struct {
	Events []WebhookEvent `json:"events,omitempty"`
	URL    string         `json:"url,omitempty"`
	Status WebhookStatus  `json:"status,omitempty"`
}

type WebhookList struct {
	Offset     int
	Limit      int
	Count      int
	TotalCount int
	Items      []*Webhook
}

type Webhook struct {
	ID              string
	ChannelID       string
	Events          []WebhookEvent
	URL             string
	Status          WebhookStatus
	CreatedDatetime *time.Time
	UpdatedDatetime *time.Time
}

type WebhookEvent string

const (
	WebhookEventConversationCreated WebhookEvent = "conversation.created"
	WebhookEventConversationUpdated WebhookEvent = "conversation.updated"
	WebhookEventMessageCreated      WebhookEvent = "message.created"
	WebhookEventMessageUpdated      WebhookEvent = "message.updated"
)

// WebhookStatus indicates what state a Webhook is in.
// At the moment there are only 2 statuses; enabled or disabled.
type WebhookStatus string

const (
	// WebhookStatusEnabled indictates that the webhook is enabled.
	WebhookStatusEnabled WebhookStatus = "enabled"
	// WebhookStatusDisabled indictates that the webhook is disabled.
	WebhookStatusDisabled WebhookStatus = "disabled"
)

// CreateWebhook registers a webhook that is invoked when something interesting
// happens.
func CreateWebhook(c *messagebird.Client, req *WebhookCreateRequest) (*Webhook, error) {
	webhook := &Webhook{}
	if err := request(c, webhook, http.MethodPost, webhooksPath, req); err != nil {
		return nil, err
	}

	return webhook, nil
}

// DeleteWebhook ensures an existing webhook is deleted and no longer
// triggered. If the error is nil, the deletion was successful.
func DeleteWebhook(c *messagebird.Client, id string) error {
	return request(c, nil, http.MethodDelete, webhooksPath+"/"+id, nil)
}

// ListWebhooks gets a collection of webhooks. Pagination can be set in options.
func ListWebhooks(c *messagebird.Client, options *ListRequestOptions) (*WebhookList, error) {
	query := paginationQuery(options)

	webhookList := &WebhookList{}
	if err := request(c, webhookList, http.MethodGet, webhooksPath+"?"+query, nil); err != nil {
		return nil, err
	}

	return webhookList, nil
}

// ReadWebhook gets a single webhook based on its ID.
func ReadWebhook(c *messagebird.Client, id string) (*Webhook, error) {
	webhook := &Webhook{}
	if err := request(c, webhook, http.MethodGet, webhooksPath+"/"+id, nil); err != nil {
		return nil, err
	}

	return webhook, nil
}

// UpdateWebhook updates a single webhook based on its ID with any values set in WebhookUpdateRequest.
// Do not set any values that should not be updated.
func UpdateWebhook(c *messagebird.Client, id string, req *WebhookUpdateRequest) (*Webhook, error) {
	webhook := &Webhook{}
	if err := request(c, webhook, http.MethodPatch, webhooksPath+"/"+id, req); err != nil {
		return nil, err
	}

	return webhook, nil
}
