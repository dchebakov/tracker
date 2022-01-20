package collector

type CollectorLog struct {
	CustomerID int64   `json:"customerID" validate:"required"`
	Timestamp  int64   `json:"timestamp"  validate:"required"`
	TagID      *int64  `json:"tagID"`
	UserID     *string `json:"userID"`
	RemoteIP   *string `json:"remoteIP"`
}
