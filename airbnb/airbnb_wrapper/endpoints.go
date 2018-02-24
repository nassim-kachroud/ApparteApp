package airbnbWrapper

const (
	Authorize = "AUTHORIZE"
	// Search              = "SEARCH"
	// GetReviews          = "GET_REVIEWS"
	ViewUserInfo    = "VIEW_USER_INFO"
	ViewListingInfo = "VIEW_LISTING_INFO"
	// CreateMessageThread = "CREATE_MESSAGE_THREAD"
	// GetMessages         = "GET_MESSAGES"
	GetUserInfo         = "GET_USER_INFO"
	GetCalendarInfo     = "GET_CALENDAR_INFO"
	GetTransactionsInfo = "GET_TRANSACTIONS_INFO"
)

var Endpoints = map[string]string{
	Authorize: "https://api.airbnb.com/v1/authorize",
	// Search:              "https://api.airbnb.com/v2/search_results?",
	// GetReviews:          "https://api.airbnb.com/v2/reviews?",
	ViewUserInfo: "https://api.airbnb.com/v2/users",
	// ViewListingInfo:     "https://api.airbnb.com/v2/listings",
	// CreateMessageThread: "https://api.airbnb.com/v1/threads/create",
	// GetMessages:         "https://api.airbnb.com/v1/threads",
	GetUserInfo:         "https://api.airbnb.com/v1/account/active",
	GetCalendarInfo:     "https://api.airbnb.com/v2/batch",
	GetTransactionsInfo: "https://api.airbnb.com/v2/batch",
}
