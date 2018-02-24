package airbnbWrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ricardo-ch/apparte-app/airbnb/airbnb_wrapper/response"
	"github.com/ricardo-ch/testairgo/airgo/net"
)

// GetCalendarQueryTemplate ...
const (
	GetCalendarQueryTemplate = `{
		"operations":[
		   {  
			  "method":"GET",
			  "path":"/calendar_days",
			  "query":{  
				 "start_date":"%s",
				 "listing_id":"%s",
				 "_format":"host_calendar",
				 "end_date":"%s"
			  }
		   }
		],
		"_transaction":false
	 }`
)

// GetTransactionsQueryTemplate ...
const (
	GetTransactionsQueryTemplate = `{
		"operations":[
		   {  
			  "method":"GET",
			  "path":"/reservations",
			  "query":{  
				 "start_date":"%s",
				 "host_id":"%s",
				 "_format":"for_mobile_list",
				 "_order":"start_date"
			  }
		   }
		],
		"_transaction":false
	 }`
)

// AccessToken ...
type AccessToken struct {
	Token string `json:"access_token"`
}

// Airgo ...
type Airgo struct {
	ApiKey      string
	AccessToken AccessToken
}

// NewAPI ...
func NewAPI() *Airgo {
	return &Airgo{}
}

func (a *Airgo) authorize(params map[string]string) (AccessToken, error) {
	var token AccessToken
	client := http.Client{}
	data, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", Endpoints[Authorize], strings.NewReader(string(data)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}
	json.Unmarshal(b, &token)
	return token, nil
}

// Login check the user and password in AirBnb
// based on  http://airbnbapi.org/#login-by-email
func (a *Airgo) Login(params map[string]string) (AccessToken, error) {
	params[ClientID] = a.ApiKey
	params[GrantType] = GrantTypePassword
	token, err := a.authorize(params)

	return token, err
}

// GetUserInfo ...
func (a *Airgo) GetUserInfo(token AccessToken, params *url.Values) (response.UserInfoResponse, error) {
	var response response.UserInfoResponse
	baseURL, err := url.Parse(Endpoints[GetUserInfo])
	if err != nil {
		return response, err
	}

	headers := http.Header{}
	headers.Set("X-Airbnb-OAuth-Token", token.Token)

	params.Add(ClientID, a.ApiKey)
	baseURL.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseURL.String(), headers)
	json.Unmarshal(b, &response)
	return response, nil
}

// ViewUserInfo ...
func (a *Airgo) ViewUserInfo(userID string, params *url.Values) (response.ViewUserInfoResponse, error) {
	var ui response.ViewUserInfoResponse
	u := fmt.Sprintf("%s/%s?", Endpoints[ViewUserInfo], userID)

	baseURL, err := url.Parse(u)
	if err != nil {
		return ui, err
	}
	params.Add(ClientID, a.ApiKey)
	baseURL.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseURL.String(), headers)
	json.Unmarshal(b, &ui)
	return ui, nil
}

// GetAllListings ...
func (a *Airgo) GetAllListings(params *url.Values) ([]response.ViewListingInfoResponse, error) {
	var lis []response.ViewListingInfoResponse
	baseURL, err := url.Parse(Endpoints[ViewListingInfo])
	if err != nil {
		return lis, err
	}

	params.Add(ClientID, a.ApiKey)

	baseURL.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseURL.String(), headers)
	if err != nil {
		return lis, err
	}
	json.Unmarshal(b, &lis)
	return lis, nil
}

// GetCalendar ...
func (a *Airgo) GetCalendar(token AccessToken, params map[string]string) ([]response.CalendarDay, error) {
	var calendarResponse response.GetCalendarInfoResponse
	client := http.Client{}

	postData := fmt.Sprintf(GetCalendarQueryTemplate, params[StartDate], params[ListingId], params[EndDate])

	req, err := http.NewRequest("POST", Endpoints[GetCalendarInfo], strings.NewReader(postData))
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Airbnb-OAuth-Token", token.Token)

	resp, err := client.Do(req)
	if err != nil {
		return []response.CalendarDay{}, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []response.CalendarDay{}, err
	}
	json.Unmarshal(b, &calendarResponse)
	return calendarResponse.Operations[0].CalendarOperationsReponse.CalendarDays, nil
}

// GetTransactions ...
func (a *Airgo) GetTransactions(token AccessToken, params map[string]string) ([]response.ReservationInfo, error) {
	var getTransactionsInfoResponse response.GetTransactionsInfoResponse
	client := http.Client{}

	postData := fmt.Sprintf(GetTransactionsQueryTemplate, params[StartDate], params[UserID])

	req, err := http.NewRequest("POST", Endpoints[GetTransactionsInfo], strings.NewReader(postData))
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Airbnb-OAuth-Token", token.Token)

	resp, err := client.Do(req)
	if err != nil {
		return []response.ReservationInfo{}, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []response.ReservationInfo{}, err
	}
	json.Unmarshal(b, &getTransactionsInfoResponse)
	return calendarResponse.Operations[0].CalendarOperationsReponse.CalendarDays, nil
}

// func (a *Airgo) LoginFB(params *url.Values) (AccessToken, error) {
// 	params.Add(ClientID, a.ApiKey)
// 	params.Add(AssertionType, AssertionTypeFacebook)
// 	params.Add(PreventAccountCreation, PreventAccountCreationTrue)
// 	token, err := a.authorize(params)
// 	return token, err
// }

// func (a *Airgo) LoginGoogle(params *url.Values) (AccessToken, error) {
// 	params.Add(ClientID, a.ApiKey)
// 	params.Add(AssertionType, AssertionTypeGoogle)
// 	params.Add(PreventAccountCreation, PreventAccountCreationTrue)
// 	token, err := a.authorize(params)
// 	return token, err
// }

// // ListingSearch
// func (a *Airgo) ListingSearch(params *url.Values) (response.ListingSearchResponse, error) {
// 	var response response.ListingSearchResponse
// 	baseUrl, err := url.Parse(Endpoints[Search])
// 	if err != nil {
// 		return response, err
// 	}

// 	params.Add(ClientID, a.ApiKey)
// 	baseUrl.RawQuery = params.Encode()

// 	client := net.HttpClient{}
// 	headers := http.Header{}
// 	b, err := client.Get(baseUrl.String(), headers)
// 	json.Unmarshal(b, &response)
// 	return response, nil
// }

// // GetReviews Returns reviews for a given listing.
// func (a *Airgo) GetReviews(params *url.Values) (response.ReviewResponse, error) {
// 	var reviews response.ReviewResponse
// 	baseUrl, err := url.Parse(Endpoints[GetReviews])
// 	if err != nil {
// 		return reviews, err
// 	}
// 	params.Add(ClientID, a.ApiKey)
// 	baseUrl.RawQuery = params.Encode()

// 	client := net.HttpClient{}
// 	headers := http.Header{}
// 	b, err := client.Get(baseUrl.String(), headers)
// 	json.Unmarshal(b, &reviews)
// 	return reviews, nil
// }

// func (a *Airgo) CreateMessageThread(token AccessToken, params *url.Values) (response.CreateThreadResponse, error) {
// 	var response response.CreateThreadResponse
// 	headers := http.Header{}
// 	headers.Set("X-Airbnb-OAuth-Token", token.Token)
// 	params.Add(ClientID, a.ApiKey)

// 	client := net.HttpClient{}
// 	b, err := client.Post(Endpoints[CreateMessageThread], *params, headers)

// 	if err != nil {
// 		return response, err
// 	}
// 	json.Unmarshal(b, &response)
// 	return response, nil
// }

// func (a *Airgo) GetMessages(token AccessToken, params *url.Values) {

// }

// func (a *Airgo) GetUserId(token AccessToken, params *url.Values) (int, error) {
// 	var id int
// 	baseUrl, err := url.Parse(Endpoints[GetUserInfo])
// 	if err != nil {
// 		return response, err
// 	}

// 	headers := http.Header{}
// 	headers.Set("X-Airbnb-OAuth-Token", token.Token)

// 	params.Add(ClientID, a.ApiKey)
// 	baseUrl.RawQuery = params.Encode()

// 	client := net.HttpClient{}
// 	b, err := client.Get(baseUrl.String(), headers)
// 	json.Unmarshal(b, &response)
// 	return response, nil
// }
