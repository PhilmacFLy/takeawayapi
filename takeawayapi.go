package takeawayapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const takeAwayTimeFormat = "2006-01-02 15:04:05"

const takeAwayURL = "https://%s.citymeal.com/android/android.php"
const takeAwayPassword = "4ndro1d"

var defaultParams = map[string]string{
	"language":      "de",
	"version":       "5.7",
	"systemVersion": "24",
	"appVersion":    "4.15.3.2",
}

func ParseTakeAwayTime(takeAwayTime string) (time.Time, error) {
	return time.Parse(takeAwayTimeFormat, takeAwayTime)
}

// Client represents the Takeaway API client
type TakeAwayClient struct {
	BaseURL  string
	Language string
	HTTP     *http.Client
	Headers  map[string]string
}

// CallFunction makes a request to the API, processes the response, and unmarshals it into resultStruct
func (tac *TakeAwayClient) sendRequest(function string, resultStruct any, params ...interface{}) error {
	// Generate MD5 checksum
	hash := md5.New()
	paramStrings := []string{function}

	// Convert parameters to strings
	for _, param := range params {
		paramStrings = append(paramStrings, fmt.Sprintf("%v", param))
	}
	paramStrings = append(paramStrings, takeAwayPassword) // Append password

	// Compute MD5 hash
	hash.Write([]byte(strings.Join(paramStrings, "")))
	md5sum := hex.EncodeToString(hash.Sum(nil))

	// Prepare request parameters
	data := url.Values{}
	data.Set("var1", function)
	for i, param := range params {
		data.Set(fmt.Sprintf("var%d", i+2), fmt.Sprintf("%v", param))
	}
	data.Set("var0", md5sum)

	// Add default parameters
	for key, value := range defaultParams {
		data.Set(key, value)
	}

	// Create request
	req, err := http.NewRequest("POST", tac.BaseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Add custom headers
	for key, value := range tac.Headers {
		req.Header.Set(key, value)
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check if the response contains an error
	var apiError apiError
	if err := json.Unmarshal(body, &apiError); err == nil && apiError.Nok.Error.ErrorID != 0 {
		return fmt.Errorf("API error: %v", apiError.Nok.Error)
	}

	// Unmarshal into the provided success struct
	if err := json.Unmarshal(body, resultStruct); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return nil
}

// NewClient initializes a new API client with the given language.
func NewClient(language string) *TakeAwayClient {
	baseURL := fmt.Sprintf(takeAwayURL, language)
	return &TakeAwayClient{
		BaseURL:  baseURL,
		Language: language,
		HTTP:     &http.Client{Timeout: time.Second * 10},
		Headers:  map[string]string{},
	}
}

// NewClientWithHTTPClient initializes a new API client with the given language and HTTP client.
func NewClientWithHTTPClient(language string, httpClient *http.Client) *TakeAwayClient {
	baseURL := fmt.Sprintf(takeAwayURL, language)
	return &TakeAwayClient{
		BaseURL:  baseURL,
		Language: language,
		HTTP:     httpClient,
		Headers:  map[string]string{},
	}
}

// SetHeader sets a header for the client
func (tac *TakeAwayClient) SetHeader(key, value string) {
	tac.Headers[key] = value
}

// SetHeaders sets multiple headers for the client
func (tac *TakeAwayClient) SetHeaders(headers map[string]string) {
	tac.Headers = headers
}

// AppendHeaders appends a header to the client
func (tac *TakeAwayClient) AppendHeaders(headers map[string]string) {
	for key, value := range headers {
		tac.Headers[key] = value
	}
}

// GetCurrentTime returns the current time from the API
func (tac *TakeAwayClient) GetCurrentTime(cc CountryCode, RestaurantID string, OrderingMode int) (CurrentTimeResponse, error) {
	function := "getcurrenttime"
	var currentTimeResponse currentTimeResponseOuter
	err := tac.sendRequest(function, &currentTimeResponse, cc, RestaurantID, OrderingMode)
	if err != nil {
		return CurrentTimeResponse{}, fmt.Errorf("error sending %s request: %v", function, err)
	}
	currentTimeResponse.CurrentTimeResponse.CurrentTime, err = ParseTakeAwayTime(currentTimeResponse.CurrentTimeResponse.CurrentTimeStr)
	if err != nil {
		return CurrentTimeResponse{}, fmt.Errorf("error parsing current time: %v", err)
	}
	return currentTimeResponse.CurrentTimeResponse, nil
}

// GetRestaurants returns a list of restaurants for the given postal code or coordinates
func (tac *TakeAwayClient) GetRestaurants(postalCode string, cc CountryCode, latitude string, longitude string) (RestaurantsResponse, error) {
	function := "getrestaurants"
	var restaurantsResponse restaurantsResponseOuter
	err := tac.sendRequest(function, &restaurantsResponse, postalCode, cc, latitude, longitude, tac.Language)
	if err != nil {
		return RestaurantsResponse{}, fmt.Errorf("error sending %s request: %v", function, err)
	}
	restaurantsResponse.RestaurantsResponse.CurrentTime, err = ParseTakeAwayTime(restaurantsResponse.RestaurantsResponse.CurrentTimeStr)
	if err != nil {
		return RestaurantsResponse{}, fmt.Errorf("error parsing current time: %v", err)
	}
	return restaurantsResponse.RestaurantsResponse, nil
}

// GetCountriesData returns a list of available countries
func (tac *TakeAwayClient) GetCountriesData() (AvailableCountries, error) {
	function := "getcountriesdata"
	var countriesResponse countriesResponse
	err := tac.sendRequest(function, &countriesResponse)
	if err != nil {
		return AvailableCountries{}, fmt.Errorf("error sending %s request: %v", function, err)
	}
	return countriesResponse.AvailableCountries, nil
}

func (tac *TakeAwayClient) getRestaurantData(function string, restaurantId string, postcode string, cc CountryCode, latitude string, longitude string, clientID string) (RestaurantData, error) {
	var restaurantDataResponse restaurantDataResponse
	err := tac.sendRequest(function, &restaurantDataResponse, restaurantId, cc, postcode, latitude, longitude, clientID)
	if err != nil {
		return RestaurantData{}, fmt.Errorf("error sending %s request: %v", function, err)
	}
	restaurantDataResponse.RestaurantData.CurrentTime, err = ParseTakeAwayTime(restaurantDataResponse.RestaurantData.CurrentTimeStr)
	if err != nil {
		return RestaurantData{}, fmt.Errorf("error parsing current time: %v", err)
	}
	for i := range restaurantDataResponse.RestaurantData.DeliveryTimes.Td.Times {
		restaurantDataResponse.RestaurantData.DeliveryTimes.Td.Times[i].StartTime, err = ParseTakeAwayTime(restaurantDataResponse.RestaurantData.DeliveryTimes.Td.Times[i].StartTimeStr)
		if err != nil {
			return RestaurantData{}, fmt.Errorf("error parsing delivery StartTimeStr time: %v", err)
		}
		restaurantDataResponse.RestaurantData.DeliveryTimes.Td.Times[i].EndTime, err = ParseTakeAwayTime(restaurantDataResponse.RestaurantData.DeliveryTimes.Td.Times[i].EndTimeStr)
		if err != nil {
			return RestaurantData{}, fmt.Errorf("error parsing delivery EndTime time: %v", err)
		}
	}
	for i := range restaurantDataResponse.RestaurantData.DeliveryTimes.Tm.Times {
		restaurantDataResponse.RestaurantData.DeliveryTimes.Tm.Times[i].StartTime, err = ParseTakeAwayTime(restaurantDataResponse.RestaurantData.DeliveryTimes.Tm.Times[i].StartTimeStr)
		if err != nil {
			return RestaurantData{}, fmt.Errorf("error parsing delivery StartTimeStr time: %v", err)
		}
		restaurantDataResponse.RestaurantData.DeliveryTimes.Tm.Times[i].EndTime, err = ParseTakeAwayTime(restaurantDataResponse.RestaurantData.DeliveryTimes.Tm.Times[i].EndTimeStr)
		if err != nil {
			return RestaurantData{}, fmt.Errorf("error parsing delivery EndTime time: %v", err)
		}
	}
	return restaurantDataResponse.RestaurantData, nil
}

// GetRestaurantData returns data for a specific restaurant including all menu items
func (tac *TakeAwayClient) GetRestaurantData(restaurantId string, postcode string, cc CountryCode, latitude string, longitude string, clientID string) (RestaurantData, error) {
	function := "getrestaurantdata"
	return tac.getRestaurantData(function, restaurantId, postcode, cc, latitude, longitude, clientID)
}

// GetRestaurantCheckoutData returns data for a specific restaurant without menu items
func (tac *TakeAwayClient) GetRestaurantCheckoutData(restaurantId string, postcode string, cc CountryCode, latitude string, longitude string, clientID string) (RestaurantData, error) {
	function := "getrestaurantcheckoutdata"
	return tac.getRestaurantData(function, restaurantId, postcode, cc, latitude, longitude, clientID)
}

// GetRestaurantReviews returns reviews for a specific restaurant
func (tac *TakeAwayClient) GetRestaurantReviews(restaurantID string, page int) ([]Review, error) {
	function := "restaurantreviews"
	var restaurantReviewsResponse reviewsResponse
	err := tac.sendRequest(function, &restaurantReviewsResponse, restaurantID, page)
	if err != nil {
		return []Review{}, fmt.Errorf("error sending %s request: %v", function, err)
	}
	return restaurantReviewsResponse.ReviewStruct.Reviews, nil
}
