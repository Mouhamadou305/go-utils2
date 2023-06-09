package types

import (
	"github.com/Mouhamadou305/go-utils2/pkg/api/models"
)

type RegistrationData models.Integration

type AdditionalSubscriptionData struct {
	SubscriptionID string `json:"subscriptionID"`
}

type EventUpdate struct {
	KeptnEvent     models.KeptnContextExtendedCE
	MetaData       EventUpdateMetaData
	SubscriptionID string // optional
}

type EventUpdateMetaData struct {
	Subject string
}

type EventSenderKeyType struct{}

var EventSenderKey = EventSenderKeyType{}

type EventSender func(ce models.KeptnContextExtendedCE) error
