package response

type Photo struct {
	Caption     string `json:"caption"`
	ID          int    `json:"id"`
	Large       string `json:"large"`
	LargeCover  string `json:"large_cover"`
	Medium      string `json:"medium"`
	MiniSquare  string `json:"mini_square"`
	Picture     string `json:"picture"`
	Small       string `json:"small"`
	SortOrder   int    `json:"sort_order"`
	Thumbnail   string `json:"thumbnail"`
	XLarge      string `json:"x_large"`
	XLargeCover string `json:"x_large_cover"`
	XMedium     string `json:"x_medium"`
	XSmall      string `json:"x_small"`
	XlPicture   string `json:"xl_picture"`
	XxLarge     string `json:"xx_large"`
}

type ViewListingInfoResponse struct {
	Listing struct {
		Access                        string      `json:"access"`
		AdditionalHouseRules          string      `json:"additional_house_rules"`
		Address                       string      `json:"address"`
		Amenities                     []string    `json:"amenities"`
		AmenitiesIds                  []int       `json:"amenities_ids"`
		Bathrooms                     int         `json:"bathrooms"`
		BedType                       string      `json:"bed_type"`
		BedTypeCategory               string      `json:"bed_type_category"`
		Bedrooms                      int         `json:"bedrooms"`
		Beds                          int         `json:"beds"`
		CalendarUpdatedAt             string      `json:"calendar_updated_at"`
		CancelPolicy                  int         `json:"cancel_policy"`
		CancelPolicyShortStr          string      `json:"cancel_policy_short_str"`
		CancellationPolicy            string      `json:"cancellation_policy"`
		CheckInTime                   interface{} `json:"check_in_time"`
		CheckInTimeEnd                string      `json:"check_in_time_end"`
		CheckInTimeEndsAt             interface{} `json:"check_in_time_ends_at"`
		CheckInTimeStart              string      `json:"check_in_time_start"`
		CheckOutTime                  interface{} `json:"check_out_time"`
		City                          string      `json:"city"`
		CleaningFeeNative             int         `json:"cleaning_fee_native"`
		CollectionIds                 interface{} `json:"collection_ids"`
		CommercialHostInfo            interface{} `json:"commercial_host_info"`
		Country                       string      `json:"country"`
		CountryCode                   string      `json:"country_code"`
		CurrencySymbolLeft            string      `json:"currency_symbol_left"`
		CurrencySymbolRight           interface{} `json:"currency_symbol_right"`
		Description                   string      `json:"description"`
		DescriptionLocale             string      `json:"description_locale"`
		ExperiencesOffered            string      `json:"experiences_offered"`
		ExtraUserInfo                 interface{} `json:"extra_user_info"`
		ExtrasPriceNative             int         `json:"extras_price_native"`
		ForceMobileLegalModal         bool        `json:"force_mobile_legal_modal"`
		GuestsIncluded                int         `json:"guests_included"`
		HasAgreedToLegalTerms         interface{} `json:"has_agreed_to_legal_terms"`
		HasAvailability               bool        `json:"has_availability"`
		HasDoubleBlindReviews         bool        `json:"has_double_blind_reviews"`
		HasLicense                    bool        `json:"has_license"`
		HasViewedCleaning             interface{} `json:"has_viewed_cleaning"`
		HasViewedIbPerfDashboardPanel interface{} `json:"has_viewed_ib_perf_dashboard_panel"`
		HasViewedTerms                bool        `json:"has_viewed_terms"`
		Hosts                         []User      `json:"hosts"`
		HouseRules                    string      `json:"house_rules"`
		ID                            int         `json:"id"`
		InBuilding                    bool        `json:"in_building"`
		InTotoArea                    bool        `json:"in_toto_area"`
		InstantBookWelcomeMessage     interface{} `json:"instant_book_welcome_message"`
		InstantBookable               bool        `json:"instant_bookable"`
		Interaction                   string      `json:"interaction"`
		IsBusinessTravelReady         bool        `json:"is_business_travel_ready"`
		IsLocationExact               bool        `json:"is_location_exact"`
		JurisdictionNames             string      `json:"jurisdiction_names"`
		JurisdictionRolloutNames      string      `json:"jurisdiction_rollout_names"`
		Language                      string      `json:"language"`
		Lat                           float64     `json:"lat"`
		License                       interface{} `json:"license"`
		ListingCleaningFeeNative      int         `json:"listing_cleaning_fee_native"`
		ListingMonthlyPriceNative     int         `json:"listing_monthly_price_native"`
		ListingNativeCurrency         string      `json:"listing_native_currency"`
		ListingOccupancyInfo          struct {
			ShowOccupancyMessage bool `json:"show_occupancy_message"`
		} `json:"listing_occupancy_info"`
		ListingPriceForExtraPersonNative int          `json:"listing_price_for_extra_person_native"`
		ListingSecurityDepositNative     int          `json:"listing_security_deposit_native"`
		ListingWeekendPriceNative        interface{}  `json:"listing_weekend_price_native"`
		ListingWeeklyPriceNative         int          `json:"listing_weekly_price_native"`
		Lng                              float64      `json:"lng"`
		Locale                           string       `json:"locale"`
		LocalizedCity                    string       `json:"localized_city"`
		MapImageURL                      string       `json:"map_image_url"`
		Market                           string       `json:"market"`
		MaxNights                        int          `json:"max_nights"`
		MaxNightsInputValue              int          `json:"max_nights_input_value"`
		MediumURL                        string       `json:"medium_url"`
		MinNights                        int          `json:"min_nights"`
		MinNightsInputValue              int          `json:"min_nights_input_value"`
		MonthlyPriceFactor               float64      `json:"monthly_price_factor"`
		MonthlyPriceNative               int          `json:"monthly_price_native"`
		Name                             string       `json:"name"`
		NativeCurrency                   string       `json:"native_currency"`
		Neighborhood                     string       `json:"neighborhood"`
		NeighborhoodOverview             string       `json:"neighborhood_overview"`
		Notes                            string       `json:"notes"`
		PersonCapacity                   int          `json:"person_capacity"`
		Photos                           []Photo      `json:"photos"`
		PictureCaptions                  []string     `json:"picture_captions"`
		PictureCount                     int          `json:"picture_count"`
		PictureURL                       string       `json:"picture_url"`
		PictureUrls                      []string     `json:"picture_urls"`
		Price                            int          `json:"price"`
		PriceForExtraPersonNative        int          `json:"price_for_extra_person_native"`
		PriceFormatted                   string       `json:"price_formatted"`
		PriceNative                      int          `json:"price_native"`
		PrimaryHost                      User         `json:"primary_host"`
		PropertyType                     string       `json:"property_type"`
		PropertyTypeID                   int          `json:"property_type_id"`
		PublicAddress                    string       `json:"public_address"`
		RecentReview                     RecentReview `json:"recent_review"`
		RequireGuestPhoneVerification    bool         `json:"require_guest_phone_verification"`
		RequireGuestProfilePicture       bool         `json:"require_guest_profile_picture"`
		RequiresLicense                  bool         `json:"requires_license"`
		ReviewRatingAccuracy             int          `json:"review_rating_accuracy"`
		ReviewRatingCheckin              int          `json:"review_rating_checkin"`
		ReviewRatingCleanliness          int          `json:"review_rating_cleanliness"`
		ReviewRatingCommunication        int          `json:"review_rating_communication"`
		ReviewRatingLocation             int          `json:"review_rating_location"`
		ReviewRatingValue                int          `json:"review_rating_value"`
		ReviewsCount                     int          `json:"reviews_count"`
		RoomType                         string       `json:"room_type"`
		RoomTypeCategory                 string       `json:"room_type_category"`
		SecurityDepositFormatted         string       `json:"security_deposit_formatted"`
		SecurityDepositNative            int          `json:"security_deposit_native"`
		SecurityPriceNative              int          `json:"security_price_native"`
		SmartLocation                    string       `json:"smart_location"`
		Space                            string       `json:"space"`
		SpecialOffer                     interface{}  `json:"special_offer"`
		SquareFeet                       interface{}  `json:"square_feet"`
		StarRating                       float64      `json:"star_rating"`
		State                            string       `json:"state"`
		Summary                          string       `json:"summary"`
		ThumbnailURL                     string       `json:"thumbnail_url"`
		ThumbnailUrls                    []string     `json:"thumbnail_urls"`
		TimeZoneName                     string       `json:"time_zone_name"`
		TotoOptIn                        interface{}  `json:"toto_opt_in"`
		Transit                          string       `json:"transit"`
		User                             struct {
			User User `json:"user"`
		} `json:"user"`
		UserID            int         `json:"user_id"`
		WeeklyPriceFactor int         `json:"weekly_price_factor"`
		WeeklyPriceNative int         `json:"weekly_price_native"`
		WirelessInfo      interface{} `json:"wireless_info"`
		XlPictureURL      string      `json:"xl_picture_url"`
		XlPictureUrls     []string    `json:"xl_picture_urls"`
		Zipcode           string      `json:"zipcode"`
	} `json:"listing"`
	Metadata struct{} `json:"metadata"`
}

type Price struct {
	Amount int `json:"native_price"`
}

type Guest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	FullName  string `json:"full_name"`
}

type Reservation struct {
	ID                      string `json:"confirmation_code"`
	Guest                   Guest  `json:"guest"`
	HostPayoutAmount        string `json:"host_payout_formatted"`
	HostRoundedPayoutAmount int    `json:"payout_price_in_host_currency"`
	NbNights                int    `json:"nights"`
	NbGuests                int    `json:"number_of_guests"`
	StartDate               string `json:"start_date"`
	Status                  string `json:"status"`
	StatusFormatted         string `json:"status_string"`
}

type ExternalCalendar struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type CalendarDay struct {
	Available        bool              `json:"available"`
	Date             string            `json:"date"`
	ListingID        int               `json:"listing_id"`
	Reason           *string           `json:"reason"`
	Type             string            `json:"type"`
	SubType          *string           `json:"subtype"`
	Reservation      *Reservation      `json:"reservation"`
	Price            Price             `json:"price"`
	ExternalCalendar *ExternalCalendar `json:"external_calendar"`
}

type CalendarOperationsReponse struct {
	CalendarDays []CalendarDay `json:"calendar_days"`
}

type Operation struct {
	CalendarOperationsReponse CalendarOperationsReponse `json:"response"`
}

type GetCalendarInfoResponse struct {
	Operations []Operation `json:"operations"`
}
