package takeawayapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// For all structs I tried to make the names as clear as possible, so you can easily understand what they are for.
// If you have any idea what the other fields are for, please let me know. I will update the structs accordingly.

// Country struct is used to store the information about a country.

type CountryCode int

// Result from the getcountrydata api
// [('NL', '1'), ('DE', '2'), ('BE', '3'), ('AT', '5'), ('CH', '6'), ('LU', '10'), ('PL', '180'), ('PT', '181'), ('VN', '239')]
const (
	NL CountryCode = 1
	DE CountryCode = 2
	BE CountryCode = 3
	AT CountryCode = 5
	CH CountryCode = 6
	LU CountryCode = 10
	PL CountryCode = 180
	PT CountryCode = 181
	VN CountryCode = 239
)

type apiError struct {
	Nok struct {
		Error struct {
			ErrorID   int    `json:"errorid"`
			ErrorText string `json:"errortext"`
		} `json:"error"`
	} `json:"nok"`
}

type countriesResponse struct {
	AvailableCountries AvailableCountries `json:"av"`
}

type AvailableCountries struct {
	CountryData []Country `json:"cd"`
	Cs          struct {
		CountryTranslations []CountryTranslation `json:"ct"`
	} `json:"cs"`
	Em  []interface{} `json:"em"`
	API struct {
		Rd  int `json:"rd"`
		Rdc int `json:"rdc"`
	} `json:"api"`
}

type CountryTranslation struct {
	Ci string `json:"ci"`
	Tr struct {
		Nl string `json:"nl"`
		De string `json:"de"`
		Fr string `json:"fr"`
		En string `json:"en"`
		Sv string `json:"sv"`
		Da string `json:"da"`
	} `json:"tr"`
	Im string `json:"im"`
	Sc struct {
		St []struct {
			Si string `json:"si"`
			Tr struct {
				En string `json:"en"`
				Nl string `json:"nl"`
				Da string `json:"da"`
				Fr string `json:"fr"`
				De string `json:"de"`
				Vi string `json:"vi"`
				Sk string `json:"sk"`
				Pl string `json:"pl"`
				Pt string `json:"pt"`
			} `json:"tr"`
		} `json:"st"`
	} `json:"sc"`
}

type Country struct {
	CountryA2           string `json:"cy"`
	Domain              string `json:"nm"`
	APIDomain           string `json:"su"`
	P1                  string `json:"p1"`
	P2                  string `json:"p2"`
	P3                  string `json:"p3,omitempty"`
	Gse                 string `json:"gse"`
	Cre                 string `json:"cre"`
	E1                  string `json:"e1"`
	E2                  string `json:"e2"`
	Pse                 string `json:"pse"`
	Psw                 any    `json:"psw"`
	Taa                 any    `json:"taa"`
	Name                string `json:"tw"`
	Emailadress         string `json:"se"`
	LogoURL             string `json:"lo"`
	FlagIconURL         string `json:"fl"`
	SmallLogoURL        string `json:"hl"`
	Ic                  string `json:"ic"`
	InternalCountryCode CountryCode
	Sc                  string `json:"sc"`
	Pie                 int    `json:"pie"`
	Erp                 struct {
		Pm []int `json:"pm"`
	} `json:"erp"`
	Lye int    `json:"lye"`
	Lyv int    `json:"lyv"`
	Lyn string `json:"lyn"`
	Ls  struct {
		La []string `json:"la"`
	} `json:"ls"`
	Si                      string `json:"si"`
	CountryNameTranslations struct {
		Nl string `json:"nl"`
		De string `json:"de"`
		Fr string `json:"fr"`
		En string `json:"en"`
		Vi string `json:"vi"`
		Pl string `json:"pl"`
		Pt string `json:"pt"`
		Lu string `json:"lu"`
		Bg string `json:"bg"`
		Ro string `json:"ro"`
		It string `json:"it"`
		Da string `json:"da"`
		No string `json:"no"`
	} `json:"cn"`
	Ac struct {
		Gp struct {
			Sid int    `json:"sid"`
			Ak  string `json:"ak"`
			Pb  string `json:"pb"`
		} `json:"gp"`
		Nm struct {
			Sid int    `json:"sid"`
			Ul  string `json:"ul"`
			Ak  string `json:"ak"`
		} `json:"nm"`
		As int `json:"as"`
	} `json:"ac"`
	Mv  any    `json:"mv"`
	Tip bool   `json:"tip"`
	Dcr string `json:"dcr"`
	Mor string `json:"mor"`
}

type LocationData struct {
	PostCode      string `json:"pc"`
	CountryA2     string `json:"cy"`
	TeriotoryName string `json:"tn"`
	Province      string `json:"pr"`
}

type GeoLocationDataResponse struct {
	LocationData LocationData `json:"ld"`
}

type restaurantsResponseOuter struct {
	RestaurantsResponse RestaurantsResponse `json:"rs"`
}

type RestaurantsResponse struct {
	City struct {
		Ps       string `json:"ps"`
		CityName string `json:"pt"`
		PostCode string `json:"ptd"`
	} `json:"cp"`
	Restaurants    []Restaurant `json:"rt"`
	CurrentTime    time.Time
	CurrentTimeStr string `json:"ct"`
	Unx            int    `json:"unx"`
	Wd             string `json:"wd"`
}

type Restaurant struct {
	ID         string `json:"id"`
	Pcid       string `json:"pcid"`
	Name       string `json:"nm"`
	Branchname string `json:"bn"`
	Op         string `json:"op"`
	Hd         string `json:"hd"`
	Dm         struct {
		Ah any `json:"ah"`
		Dl struct {
			Op int    `json:"op"`
			Oh string `json:"oh"`
		} `json:"dl"`
		Pu struct {
			Op int    `json:"op"`
			Oh string `json:"oh"`
		} `json:"pu"`
	} `json:"dm"`
	Tip           int `json:"tip"`
	New           int `json:"new"`
	Ply           int `json:"ply"`
	EstimatedTime any `json:"est"`
	Eta           struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"eta"`
	Ft             string `json:"ft"`
	Ck             string `json:"ck"`
	Ds             string `json:"ds"`
	Logo           string `json:"lo"`
	CloudinaryLogo string `json:"cloudinaryLogo"`
	Cs             struct {
		Ct [][]string `json:"ct"`
	} `json:"cs"`
	Sr struct {
		S1  int     `json:"s1"`
		S2  int     `json:"s2"`
		S3  any     `json:"s3"`
		S4  float64 `json:"s4"`
		S5  any     `json:"s5"`
		S6  any     `json:"s6"`
		S7  any     `json:"s7"`
		S8  any     `json:"s8"`
		S9  any     `json:"s9"`
		S12 any     `json:"s12"`
	} `json:"sr"`
	Nt           string `json:"nt"`
	ChainPenalty bool   `json:"chain_penalty"`
	Sc           struct {
		Si []string `json:"si"`
	} `json:"sc"`
	Pd struct {
		Pz struct {
			Num1 string `json:"1"`
			Num2 string `json:"2"`
			Num3 string `json:"3"`
		} `json:"pz"`
	} `json:"pd,omitempty"`
	Pm struct {
		Me []struct {
			Mi string `json:"mi"`
		} `json:"me"`
	} `json:"pm"`
	Dc struct {
		Ma string `json:"ma"`
		Co []struct {
			Fr string `json:"fr"`
			To string `json:"to"`
			Ct string `json:"ct"`
		} `json:"co"`
		Ddf []interface{} `json:"ddf"`
	} `json:"dc"`
	Rv      string  `json:"rv"`
	Rvd     string  `json:"rvd"`
	Address Address `json:"ad"`
	Bd      string  `json:"bd"`
}

type Address struct {
	Street      string `json:"st"`
	Housenumber string `json:"hn"`
	Postcode    string `json:"pc"`
	Town        string `json:"tn"`
	City        string `json:"ci"`
	Latitude    string `json:"lt"`
	Longitude   string `json:"ln"`
}

type restaurantDataResponse struct {
	RestaurantData RestaurantData `json:"rd"`
}

type RestaurantData struct {
	Name             string `json:"nm"`
	Branch           string `json:"bn"`
	RestaurantID     string `json:"ri"`
	TelephoneNumbers struct {
		No1 string `json:"no1"`
	} `json:"tel"`
	HeaderImageURL   string  `json:"mh"`
	CloudinaryHeader string  `json:"cloudinaryHeader"`
	Tr               string  `json:"tr"`
	Pne              string  `json:"pne"`
	Rci              string  `json:"rci"`
	Ad               Address `json:"ad"`
	Oo               struct {
		Lu             string `json:"lu"`
		CloudinaryLogo string `json:"cloudinaryLogo"`
		Nt             string `json:"nt"`
		Sl             string `json:"sl"`
		Rv             string `json:"rv"`
		Rvd            string `json:"rvd"`
		Bd             string `json:"bd"`
		Cim            int    `json:"cim"`
		Ft             string `json:"ft"`
		Eba            bool   `json:"eba"`
	} `json:"oo"`
	Sco   string      `json:"sco"`
	Ddf   string      `json:"ddf"`
	Smid  interface{} `json:"smid"`
	Pty   string      `json:"pty"`
	Pro   int         `json:"pro"`
	Cph   string      `json:"cph"`
	Murl  string      `json:"murl"`
	Rte   int         `json:"rte"`
	Legal struct {
		Owner string  `json:"own"`
		Vat   string  `json:"vat"`
		Tcr   string  `json:"tcr"`
		Crn   string  `json:"crn"`
		Adr   Address `json:"adr"`
	} `json:"lgl"`
	Ck string `json:"ck"`
	Ds int    `json:"ds"`
	Op string `json:"op"`
	Ac string `json:"ac"`
	Dm struct {
		Ah string `json:"ah"`
		Dl struct {
			Op  int    `json:"op"`
			Mpt string `json:"mpt"`
			Eta struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"eta"`
			Oh string `json:"oh"`
		} `json:"dl"`
		Pu struct {
			Op  int    `json:"op"`
			Mpt string `json:"mpt"`
			Oh  string `json:"oh"`
		} `json:"pu"`
	} `json:"dm"`
	Ply int `json:"ply"`
	Pd  struct {
		Pz struct {
			Num1 string `json:"1"`
			Num2 string `json:"2"`
			Num3 string `json:"3"`
		} `json:"pz"`
	} `json:"pd"`
	Pm struct {
		Me []struct {
			Mi string `json:"mi"`
			Mt string `json:"mt"`
			Mf string `json:"mf"`
		} `json:"me"`
	} `json:"pm"`
	DeliveryTimes ServiceTimes `json:"dt"`
	PickupTimes   ServiceTimes `json:"pt"`
	Dc            struct {
		Ma string `json:"ma"`
	} `json:"dc"`
	DeliveryData struct {
		Da []struct {
			Postcodes struct {
				PostCodesArray []string `json:"pp"`
			} `json:"pc"`
			Ma    string `json:"ma"`
			Costs []struct {
				Fr string `json:"fr"`
				To string `json:"to"`
				Ct string `json:"ct"`
			} `json:"co"`
		} `json:"da"`
	} `json:"dd"`
	Menu struct {
		CategorieStruct struct {
			Categories []struct {
				ID            string        `json:"id"`
				Name          string        `json:"nm"`
				Description   string        `json:"ds"`
				Cti           string        `json:"cti"`
				Ot            []interface{} `json:"ot"`
				ProductStruct struct {
					Products Products `json:"pr"`
				} `json:"ps"`
				CloudinaryChain string `json:"cloudinaryChain,omitempty"`
			} `json:"ct"`
		} `json:"cs"`
	} `json:"mc"`
	Rt struct {
		Cr  float64 `json:"cr"`
		Prr float64 `json:"prr"`
	} `json:"rt"`
	CurrentTimeStr string `json:"ct"`
	CurrentTime    time.Time
	Wd             string `json:"wd"`
	Ce             int    `json:"ce"`
}

type ServiceTimes struct {
	CurrentTimeStr string `json:"ct"`
	CurrentTime    time.Time
	Td             Times `json:"td"`
	Tm             Times `json:"tm"`
}

type Times struct {
	Times []struct {
		StartTimeStr string `json:"st"`
		StartTime    time.Time
		EndTimeStr   string `json:"et"`
		EndTime      time.Time
	} `json:"ti"`
}

type Products []Product

func (p *Products) UnmarshalJSON(data []byte) error {
	fmt.Printf("Raw JSON for 'pr': %s\n", string(data)) // Debug: Print the raw data

	// Try to unmarshal into a slice (expected case)
	var products []Product
	if err := json.Unmarshal(data, &products); err == nil {
		*p = products
		fmt.Println("Successfully unmarshaled 'pr' as an array")
		return nil
	} else {
		fmt.Printf("Failed to unmarshal as array: %v\n", err)
	}

	// If unmarshaling into a slice fails, try unmarshaling a single object
	var singleProduct Product
	if err := json.Unmarshal(data, &singleProduct); err == nil {
		*p = []Product{singleProduct} // Wrap single product into a slice
		fmt.Println("Successfully unmarshaled 'pr' as a single object")
		return nil
	} else {
		fmt.Printf("Failed to unmarshal as single object: %v\n", err)
	}

	// If both fail, return an error with the problematic data
	return fmt.Errorf("failed to unmarshal 'pr': data=%s", string(data))
}

type Product struct {
	ID                string `json:"id"`
	Name              string `json:"nm"`
	Description       string `json:"ds,omitempty"`
	Ah                string `json:"ah,omitempty"`
	PickupCost        string `json:"pc"` // Should be string instead of float64
	DeliveryCost      string `json:"tc"` // Should be string instead of float64
	Pu                string `json:"pu,omitempty"`
	CloudinaryProduct string `json:"cloudinaryProduct,omitempty"`
	Xfm               int    `json:"xfm"`
	Fai               Fai    `json:"fai,omitempty"`
	SideItems         struct {
		SideDishes []SideDish `json:"sd"`
	} `json:"ss,omitempty"`
}

// Fai struct with custom UnmarshalJSON for All, Add, and Xtr fields
type Fai struct {
	All CustomAll `json:"all"`
	Add CustomAdd `json:"add,omitempty"`
	Xtr CustomXtr `json:"xtr,omitempty"`
	Nut string    `json:"nut,omitempty"`
}

// Custom type for "all" field to handle both array and object
type CustomAll []string

func (a *CustomAll) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an array (expected case)
	var allStrings []string
	if err := json.Unmarshal(data, &allStrings); err == nil {
		*a = allStrings
		return nil
	}

	// If not an array, try to unmarshal as an object with "id" field
	var allObject struct {
		ID []string `json:"id"`
	}
	if err := json.Unmarshal(data, &allObject); err == nil {
		*a = allObject.ID
		return nil
	}

	return fmt.Errorf("failed to unmarshal 'all': data=%s", string(data))
}

// Custom type for "add" field to handle both array and object
type CustomAdd struct {
	IDs []string `json:"id,omitempty"`
}

func (a *CustomAdd) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object with "id" field
	var addObject struct {
		ID []string `json:"id"`
	}
	if err := json.Unmarshal(data, &addObject); err == nil {
		a.IDs = addObject.ID
		return nil
	}

	// Try to unmarshal as an empty array (valid case)
	var addArray []string
	if err := json.Unmarshal(data, &addArray); err == nil {
		a.IDs = addArray
		return nil
	}

	return fmt.Errorf("failed to unmarshal 'add': data=%s", string(data))
}

// Custom type for "xtr" field to handle both array and object
type CustomXtr struct {
	Extras map[string]string
}

func (x *CustomXtr) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object (expected case)
	var extrasMap map[string]string
	if err := json.Unmarshal(data, &extrasMap); err == nil {
		x.Extras = extrasMap
		return nil
	}

	// If it's an empty array, treat it as an empty map
	var emptyArray []interface{}
	if err := json.Unmarshal(data, &emptyArray); err == nil {
		x.Extras = make(map[string]string) // Empty map instead of array
		return nil
	}

	return fmt.Errorf("failed to unmarshal 'xtr': data=%s", string(data))
}

type SideDish struct {
	Nm string `json:"nm"`
	Cc struct {
		Ch []struct {
			ID           string `json:"id"`
			Name         string `json:"nm"`
			PickupCost   string `json:"pc"`
			DeliveryCost string `json:"tc"`
			Xfm          int    `json:"xfm"`
		} `json:"ch"`
	} `json:"cc"`
	Tp string `json:"tp"`
}

type reviewsResponse struct {
	ReviewStruct struct {
		Reviews []Review `json:"rv"`
	} `json:"rr"`
}

type Review struct {
	Name    string `json:"nm"`
	TimeStr string `json:"ti"`
	Time    time.Time
	Remark  string `json:"rm"`
	Kw      string `json:"kw"`
	Be      string `json:"be"`
	Dm      string `json:"dm"`
	Zo      string `json:"zo"`
	Ny      string `json:"ny"`
}

type currentTimeResponseOuter struct {
	CurrentTimeResponse CurrentTimeResponse `json:"st"`
}

type CurrentTimeResponse struct {
	CurrentTimeStr string `json:"ct"`
	CurrentTime    time.Time
	Rs             int    `json:"rs"`
	Wd             string `json:"wd"`
}
