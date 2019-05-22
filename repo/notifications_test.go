package repo_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/evenfound/even-go/repo"
)

func TestDataIsUnmarshalable(t *testing.T) {
	for _, n := range createNotificationExamples() {
		var (
			actual  = make(map[string]interface{})
			err     error
			payload []byte
		)
		if payload, err = n.Data(); err != nil {
			t.Error(err)
		}
		if err = json.Unmarshal(payload, &actual); err != nil {
			t.Errorf("failed unmarshaling '%s': %s\n", n.GetType(), err)
			continue
		}
	}
}

func TestWebsocketDataIsUnmarshalable(t *testing.T) {
	for _, n := range createNotificationExamples() {
		var (
			actual  = make(map[string]interface{})
			err     error
			payload []byte
		)
		if payload, err = n.WebsocketData(); err != nil {
			t.Error(err)
		}
		if err = json.Unmarshal(payload, &actual); err != nil {
			t.Errorf("failed unmarshaling '%s': %s\n", n.GetType(), err)
			continue
		}
		if _, ok := actual["notification"]; !ok {
			t.Errorf("missing 'notification' JSON key in marshalled payload of %s", n.GetType())
		}
	}
}

// TestNotificationMarshalling ensures that the new Notification marshal format is
// functioning properly. This applies to notifications which have been marshaled in
// the datastore with json.Marshal(Notification{}). Some notifications have been
// marshaled in the datastore with json.Marshal(Notification{}.NotifierData), and
// TestLegacyNotificationMarshalling covers those cases.
func TestNotificationMarshalling(t *testing.T) {
	for _, n := range createNotificationExamples() {
		var (
			expected = repo.NewNotification(n, time.Now(), false)
			actual   = &repo.Notification{}
		)
		data, err := json.Marshal(expected)
		if err != nil {
			t.Errorf("failed marshaling '%s': %s\n", expected.GetType(), err)
			continue
		}
		if err := json.Unmarshal(data, actual); err != nil {
			t.Errorf("failed unmarshaling '%s': %s\n", expected.GetType(), err)
			continue
		}

		if actual.GetType() != expected.GetType() {
			t.Error("Expected notification to match types, but did not")
			t.Errorf("Expected: %s\n", expected.GetType())
			t.Errorf("Actual: %s\n", actual.GetType())
		}
		if !reflect.DeepEqual(actual.NotifierData, expected.NotifierData) {
			t.Error("Expected notifier data to match, but did not")
			t.Errorf("Expected: %+v\n", expected.NotifierData)
			t.Errorf("Actual: %+v\n", actual.NotifierData)
		}
	}
}

// TestLegacyNotificationMarshalling ensures that the legacy Notification marshaling is
// functioning properly. This applies to notifications which have been marshaled in
// the datastore with json.Marshal(Notification{}.NotifierData).
func TestLegacyNotificationMarshalling(t *testing.T) {
	for _, n := range createLegacyNotificationExamples() {
		var (
			actual = &repo.Notification{}
		)
		data, err := json.Marshal(n)
		if err != nil {
			t.Errorf("failed marshaling '%s': %s\n", n.GetType(), err)
			continue
		}
		if err := json.Unmarshal(data, actual); err != nil {
			t.Errorf("failed unmarshaling '%s': %s\n", n.GetType(), err)
			continue
		}

		if actual.GetID() != n.GetID() {
			t.Error("Expected notification to match ID, but did not")
			t.Errorf("Expected: %s\n", n.GetID())
			t.Errorf("Actual: %s\n", actual.GetID())
		}
		if actual.GetType() != n.GetType() {
			t.Error("Expected notification to match types, but did not")
			t.Errorf("Expected: %s\n", n.GetType())
			t.Errorf("Actual: %s\n", actual.GetType())
		}
		if !reflect.DeepEqual(actual.NotifierData, n) {
			t.Error("Expected notifier data to match, but did not")
			t.Errorf("Expected: %+v\n", n)
			t.Errorf("Actual: %+v\n", actual.NotifierData)
		}
	}
}

func createNotificationExamples() []repo.Notifier {
	return append([]repo.Notifier{
		{
			ID:     "disputeNotificationID",
			Type:   repo.NotifierTypeModeratorDisputeExpiry,
			CaseID: repo.NewNotificationID(),
		},
		{
			ID:      "buyerDisputeTimeoutID",
			Type:    repo.NotifierTypeBuyerDisputeTimeout,
			OrderID: repo.NewNotificationID(),
		},
		{
			ID:      "buyerDisputeExpiryID",
			Type:    repo.NotifierTypeBuyerDisputeExpiry,
			OrderID: repo.NewNotificationID(),
		},
		{
			ID:      "saleAgingID",
			Type:    repo.NotifierTypeVendorDisputeTimeout,
			OrderID: repo.NewNotificationID(),
		},
		{
			ID:      "vendorFinalizedPayment",
			Type:    repo.NotifierTypeVendorFinalizedPayment,
			OrderID: repo.NewNotificationID(),
		},
	},
		createLegacyNotificationExamples()...)
}

func createLegacyNotificationExamples() []repo.Notifier {
	return []repo.Notifier{
		{
			ID:      "orderCompletionID",
			Type:    repo.NotifierTypeCompletionNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "disputeAcceptedID",
			Type:    repo.NotifierTypeDisputeAcceptedNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "disputeCloseID",
			Type:    repo.NotifierTypeDisputeCloseNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "disputeOpenID",
			Type:    repo.NotifierTypeDisputeOpenNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "disputeUpdateID",
			Type:    repo.NotifierTypeDisputeUpdateNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:     "followID",
			Type:   repo.NotifierTypeFollowNotification,
			PeerId: repo.NewNotificationID(),
		},
		{
			ID:      "fulfillmentID",
			Type:    repo.NotifierTypeFulfillmentNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:     "moderatorAddID",
			Type:   repo.NotifierTypeModeratorAddNotification,
			PeerId: repo.NewNotificationID(),
		},
		{
			ID:     "moderatorRemoveID",
			Type:   repo.NotifierTypeModeratorRemoveNotification,
			PeerId: repo.NewNotificationID(),
		},
		{
			ID:      "orderCancelID",
			Type:    repo.NotifierTypeOrderCancelNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "orderConfirmID",
			Type:    repo.NotifierTypeOrderConfirmationNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "orderDeclinedID",
			Type:    repo.NotifierTypeOrderDeclinedNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "orderNotificationID",
			Type:    repo.NotifierTypeOrderNewNotification,
			BuyerID: repo.NewNotificationID(),
		},
		{
			ID:      "paymentID",
			Type:    repo.NotifierTypePaymentNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "processingErrorID",
			Type:    repo.NotifierTypeProcessingErrorNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:      "refundID",
			Type:    repo.NotifierTypeRefundNotification,
			OrderId: repo.NewNotificationID(),
		},
		{
			ID:     "unfollowID",
			Type:   repo.NotifierTypeUnfollowNotification,
			PeerId: repo.NewNotificationID(),
		},
	}
}

func TestNotificationSatisfiesNotifierInterface(t *testing.T) {
	notifier := repo.VendorDisputeTimeout{
		ID:      "saleAgingID",
		Type:    repo.NotifierTypeVendorDisputeTimeout,
		OrderID: repo.NewNotificationID(),
	}
	var _ repo.Notifier = repo.NewNotification(notifier, time.Now(), false)
}
