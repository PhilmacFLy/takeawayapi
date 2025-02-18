package takeawayapi

import (
	"testing"
	"time"
)

func getTimeWithoutTimezone() time.Time {
	now := time.Now()
	const layout = "2006-01-02 15:04:05"
	now, _ = time.Parse(layout, now.Format(layout))
	return now
}

func TestParseTakeAwayTime(t *testing.T) {
	oldtime := getTimeWithoutTimezone()
	takeAwayTime := oldtime.Format(takeAwayTimeFormat)
	newtime, err := ParseTakeAwayTime(takeAwayTime)
	if err != nil {
		t.Fatalf(`ParseTakeAwayTime errored wit error: %v`, err)
	}
	if newtime != oldtime {
		t.Fatalf(`ParseTakeAwayTime returned wrong time: Oldtime: %v NewTime %v`, oldtime, newtime)
	}
}

func TestGetCurrentTime(t *testing.T) {
	tac := NewClient("de")
	nowtime := getTimeWithoutTimezone()
	r, err := tac.GetCurrentTime(DE, "O3QQ11PN", 1)
	if err != nil {
		t.Fatalf(`GetCurrentTime errored wit error: %v`, err)
	}
	gotTime := r.CurrentTime
	if gotTime.Before(nowtime) {
		t.Fatalf(`Time from Api Before Time: %v`, gotTime)
	}
	if gotTime.After(nowtime.Add(time.Minute * 5)) {
		t.Fatalf(`GetCurrentTime defiation to big: %v vs %v`, nowtime, gotTime)
	}
}

func TestGetRestaurants(t *testing.T) {
	tac := NewClient("de")
	postcodes := []string{"90461", "18147", "92431", "79111", "45897"}
	for _, postcode := range postcodes {
		r, err := tac.GetRestaurants(postcode, DE, "", "")
		if err != nil {
			t.Fatalf(`GetRestaurants for postcode: %v errored with error: %v`, postcode, err)
		}
		if len(r.Restaurants) == 0 {
			t.Fatalf(`GetRestaurants for postcotde: %v returned no restaurants`, postcode)
		}
		if r.City.PostCode != postcode {
			t.Fatalf(`GetRestaurants returned wrong postcode: Expected Postcode %v Got Postcode %v`, postcode, r.City.PostCode)
		}
	}
}

func TestGetCountriesData(t *testing.T) {
	tac := NewClient("de")
	countries, err := tac.GetCountriesData()
	if err != nil {
		t.Fatalf(`GetCountriesData errored with error: %v`, err)
	}
	if len(countries.CountryData) == 0 {
		t.Fatalf(`GetCountriesData returned no countries`)
	}
	if len(countries.Cs.CountryTranslations) == 0 {
		t.Fatalf(`GetCountriesData returned no translations`)
	}
}

func TestGetRestaurantData(t *testing.T) {
	tac := NewClient("de")
	restaurantId := "O3QQ11PN"
	restaurantData, err := tac.GetRestaurantData(restaurantId, "", DE, "", "", "")
	if err != nil {
		t.Fatalf(`GetRestaurantData errored with error: %v`, err)
	}
	if restaurantData.RestaurantID != restaurantId {
		t.Fatalf(`GetRestaurantData returned wrong restaurantId: Expected %v Got %v`, restaurantId, restaurantData.RestaurantID)
	}
	if restaurantData.Name == "" {
		t.Fatalf(`GetRestaurantData returned no name`)
	}
	if len(restaurantData.Menu.CategorieStruct.Categories) == 0 {
		t.Fatalf(`GetRestaurantData returned no categories`)
	}

}

func TestGetRestaurantCheckoutData(t *testing.T) {
	tac := NewClient("de")
	restaurantId := "O3QQ11PN"
	restaurantData, err := tac.GetRestaurantData(restaurantId, "", DE, "", "", "")
	if err != nil {
		t.Fatalf(`GetRestaurantData errored with error: %v`, err)
	}
	if restaurantData.RestaurantID != restaurantId {
		t.Fatalf(`GetRestaurantData returned wrong restaurantId: Expected %v Got %v`, restaurantId, restaurantData.RestaurantID)
	}
	if restaurantData.Name == "" {
		t.Fatalf(`GetRestaurantData returned no name`)
	}

}

func TestGetRestaurantReviews(t *testing.T) {
	tac := NewClient("de")
	restaurantId := "O3QQ11PN"
	reviews, err := tac.GetRestaurantReviews(restaurantId, 1)
	if err != nil {
		t.Fatalf(`GetRestaurantReviews errored with error: %v`, err)
	}
	if len(reviews) == 0 {
		t.Fatalf(`GetRestaurantReviews returned no reviews`)
	}
}
