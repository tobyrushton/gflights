package gflights

import (
	"fmt"
	"time"
	"unicode"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

type Travelers struct {
	Adults        int
	Children      int
	InfantsInSeat int
	InfantsOnLap  int
}

// Traveler rules:
//   - Max 9 total travelers per search
//   - At least 1 adult per search
//   - One adult for every 2 infants
//   - One adult per infant on lap
func (t *Travelers) Validate() error {
	if t.Adults < 1 {
		return fmt.Errorf("at least one adult traveler is required")
	}
	if t.Adults+t.Children+t.InfantsInSeat+t.InfantsOnLap > 9 {
		return fmt.Errorf("maximum of 9 travelers allowed per search")
	}
	if t.InfantsOnLap > t.Adults {
		return fmt.Errorf("each infant on lap must be accompanied by an adult")
	}
	if float64(t.InfantsOnLap+t.InfantsInSeat)/2 > float64(t.Adults) {
		return fmt.Errorf("one adult is required for every two infants")
	}
	if t.Children < 0 || t.InfantsInSeat < 0 || t.InfantsOnLap < 0 {
		return fmt.Errorf("number of children and infants cannot be negative")
	}
	return nil
}

// Class represents the travel class for the flight
type Class int64

const (
	Economy Class = iota + 1
	PremiumEconomy
	Business
	First
)

// TripType represents the type of trip: one-way or round-trip
type TripType int64

const (
	RoundTrip TripType = iota + 1
	OneWay
)

// Stops represents the number of stops for the flight
type Stops int64

const (
	NonStop Stops = iota + 1
	OneStop
	TwoPlusStops
	AnyStops
)

type Options struct {
	Travelers Travelers
	Class     Class
	TripType  TripType
	Stops     Stops
	Lang      language.Tag
	Currency  currency.Unit
}

func DefaultOptions() Options {
	return Options{
		Travelers: Travelers{Adults: 1},
		Class:     Economy,
		TripType:  OneWay,
		Stops:     AnyStops,
		Lang:      language.English,
		Currency:  currency.USD,
	}
}

func validateDate(a, b time.Time) error {
	if a.IsZero() {
		return fmt.Errorf("departure date is required")
	}
	if a.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("departure date cannot be in the past")
	}
	if !b.IsZero() && b.Before(a) {
		return fmt.Errorf("return date cannot be before departure date")
	}
	return nil
}

func isAirportCode(code string) bool {
	if len(code) != 3 {
		return false
	}
	for _, r := range code {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func validateLocations(srcCities, srcAirports, dstCities, dstAirports []string) error {
	if len(srcCities)+len(srcAirports) == 0 {
		return fmt.Errorf("at least one source city or airport is required")
	}
	if len(dstCities)+len(dstAirports) == 0 {
		return fmt.Errorf("at least one destination city or airport is required")
	}

	for _, s := range srcAirports {
		if !isAirportCode(s) {
			return fmt.Errorf("src airport '%s' is not an airport code", s)
		}
	}

	for _, d := range dstAirports {
		if !isAirportCode(d) {
			return fmt.Errorf("dst airport '%s' is not an airport code", d)
		}
	}

	return nil
}

type Args struct {
	DepartureDate, ReturnDate                      time.Time
	SrcCities, SrcAirports, DstCities, DstAirports []string
	Options                                        Options
}

func (a *Args) Validate() error {
	if err := validateDate(a.DepartureDate, a.ReturnDate); err != nil {
		return err
	}
	if err := validateLocations(a.SrcCities, a.SrcAirports, a.DstCities, a.DstAirports); err != nil {
		return err
	}
	if err := a.Options.Travelers.Validate(); err != nil {
		return err
	}
	return nil
}

type FlightCode struct {
	AirlineCode  string // airline code
	FlightNumber string // flight number
}

// Flight describes a single, one-way flight.
type Flight struct {
	DepAirportCode string        // departure airport code
	DepAirportName string        // departure airport name
	DepCity        string        // departure city
	ArrAirportName string        // arrival airport name
	ArrAirportCode string        // arrival airport code
	ArrCity        string        // arrival city
	DepTime        time.Time     // departure time
	ArrTime        time.Time     // arrival time
	Duration       time.Duration // duration of the flight
	Airplane       string        // airplane model
	FlightCode     FlightCode    // flight code
	Unknown        []any         // it contains all unknown data which are parsed from the Google Flights API
	AirlineName    string        // airline name
	Legroom        string        // legroom in the airplane seats
}

type SimpleOffer struct {
	StartDate  time.Time // start date of the offer
	ReturnDate time.Time // return date of the offer
	Price      float64   // price of the offer
}

type OutboundOffer struct {
	SimpleOffer

	Flight         []Flight
	SrcAirportCode string        // code of the airport where the trip starts
	DstAirportCode string        // destination airport
	SrcCity        string        // source city
	DstCity        string        // destination city
	Duration       time.Duration // total duration of the trip

	s     *Session
	token string
	args  Args
}

type PriceRange struct {
	Min float64 // minimum price
	Max float64 // maximum price
}

type ReturnOffer struct {
	Flight []Flight
	Price  float64

	token string
}

type FlightSegment struct {
	Legs []Flight // Contains the legs of the segment
}

type TripSelection struct {
	Segments  []FlightSegment // Contains the different segments of the flight (e.g., outbound and return)
	Price     float64
	Class     Class
	TripType  TripType
	Travelers Travelers
}
