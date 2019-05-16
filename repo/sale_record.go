package repo

import (
	"strings"
	"time"

	"github.com/evenfound/even-go/pb"
)

var (
	VendorDisputeTimeout_lastInterval = time.Duration(45*24) * time.Hour
)

// SaleRecord represents a one-to-one relationship with records
// in the SQL datastore
type SaleRecord struct {
	Contract                     *pb.RicardianContract
	OrderID                      string
	OrderState                   pb.OrderState
	Timestamp                    time.Time
	LastDisputeTimeoutNotifiedAt time.Time
}

// SupportsTimedEscrowRelease indicates whether the underlying AcceptedCurrency supports
// a time-bassed release behavior.
// TODO: Express this from the wallet-interface instead
func (r *SaleRecord) SupportsTimedEscrowRelease() bool {
	if r.Contract.BuyerOrder != nil &&
		r.Contract.BuyerOrder.Payment != nil {
		switch strings.ToUpper(r.Contract.BuyerOrder.Payment.Coin) {
		case "BTC":
			return true
		case "TBTC":
			return true
		case "BCH":
			return true
		case "TBCH":
			return true
		case "LTC":
			return true
		case "TLTC":
			return true
		case "ZEC":
			return false
		case "TZEC":
			return false
		}
	}
	return false
}

// IsDisputeable returns whether the Sale is in a state that it can be disputed with a
// third-party moderator
func (r *SaleRecord) IsDisputeable() bool {
	if r.IsModeratedContract() {
		switch r.OrderState {
		case pb.OrderState_PARTIALLY_FULFILLED, pb.OrderState_FULFILLED:
			return true
		}
	}
	return false
}

// IsModeratedContract indicates whether the SaleRecord has a contract which includes
// a third-party moderator
func (r *SaleRecord) IsModeratedContract() bool {
	return r.Contract != nil && r.Contract.BuyerOrder != nil && r.Contract.BuyerOrder.Payment != nil && r.Contract.BuyerOrder.Payment.Method == pb.Order_Payment_MODERATED
}

// BuildVendorDisputeTimeoutLastNotification returns a Notification that alerts a SaleRecord
// is more than 45 days old and already expired
func (r *SaleRecord) BuildVendorDisputeTimeoutLastNotification(createdAt time.Time) *Notification {
	notification := &VendorDisputeTimeout{
		ID:        NewNotificationID(),
		ExpiresIn: uint(0),
		OrderID:   r.OrderID,
		Thumbnail: Thumbnail{},
		Type:      NotifierTypeVendorDisputeTimeout,
	}
	if len(r.Contract.VendorListings) > 0 && len(r.Contract.VendorListings[0].Item.Images) > 0 {
		notification.Thumbnail = Thumbnail{
			Tiny:  r.Contract.VendorListings[0].Item.Images[0].Tiny,
			Small: r.Contract.VendorListings[0].Item.Images[0].Small,
		}
	}
	return NewNotification(notification, createdAt, false)
}
