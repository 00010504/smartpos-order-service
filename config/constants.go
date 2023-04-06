package config

const (
	ConsumerGroupID = "order_service"

	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

const (
	OrderStatusDraft    = "7069e210-7d2e-4a12-9160-3ef82f18ef4d"
	OrderStatusPayed    = "d3bde6a2-532c-4f08-811f-0385e804c885"
	OrderStatusPostpone = "15c0d291-45e5-4077-89d7-9b365e65cfed"
)

const (
	OrderDiscountTypeNone       = "9a2aa8fe-806e-44d7-8c9d-575fa67ebefd"
	OrderDiscountTypePercentage = "1fe92aa8-2a61-4bf1-b907-182b497584ad"
	OrderDiscountTypeAmount     = "9fb3ada6-a73b-4b81-9295-5c1605e54552"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

const (
	DD_MM_YYYY          = "02-01-2006"
	DD_MM_YYYY_HH_MM    = "02-01-2006 15:04"
	DD_MM_YYYY_HH_MM_SS = "02-01-2006 15:04:05"
)

var ()
