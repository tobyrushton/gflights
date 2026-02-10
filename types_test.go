package gflights_test

import (
	"testing"
	"time"

	"github.com/tobyrushton/gflights"
)

type withValidate interface {
	Validate() error
}

var wrongAirportCode = "wrong"

func testValidate(t *testing.T, args withValidate, wantErr string) {
	gotErr := args.Validate()

	if gotErr == nil && wantErr != "" {
		t.Fatalf("expected error '%s', got nil", wantErr)
	} else if gotErr != nil && gotErr.Error() != wantErr {
		t.Fatalf("expected error '%s', got '%s'", wantErr, gotErr.Error())
	}
}

func TestValidateArgs(t *testing.T) {
	args := &gflights.Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "at least one destination city or airport is required")

	args = &gflights.Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "at least one source city or airport is required")

	args = &gflights.Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "src airport 'wrong' is not an airport code")

	args = &gflights.Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "dst airport 'wrong' is not an airport code")

	args = &gflights.Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		DepartureDate: time.Now().AddDate(0, 0, 3),
		ReturnDate:    time.Now().AddDate(0, 0, 1),
	}
	testValidate(t, args, "return date cannot be before departure date")

	args = &gflights.Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		DepartureDate: time.Now().AddDate(0, 0, -1),
		ReturnDate:    time.Now().AddDate(0, 0, 1),
	}
	testValidate(t, args, "departure date cannot be in the past")
}

func TestValidateTravelers(t *testing.T) {
	args := &gflights.Travelers{}
	testValidate(t, args, "at least one adult traveler is required")

	args = &gflights.Travelers{
		Adults:   1,
		Children: -1,
	}

	testValidate(t, args, "number of children and infants cannot be negative")

	args = &gflights.Travelers{
		Adults:       1,
		Children:     1,
		InfantsOnLap: 2,
	}
	testValidate(t, args, "each infant on lap must be accompanied by an adult")

	args = &gflights.Travelers{
		Adults:        1,
		Children:      1,
		InfantsOnLap:  1,
		InfantsInSeat: 2,
	}
	testValidate(t, args, "one adult is required for every two infants")
}

func TestValidatePriceGraphArgs(t *testing.T) {
	args := &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 1),
		RangeEndDate:   time.Now().AddDate(0, 0, 10),
		TripLength:     5,
	}
	testValidate(t, args, "at least one destination city or airport is required")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 1),
		RangeEndDate:   time.Now().AddDate(0, 0, 10),
		TripLength:     5,
	}
	testValidate(t, args, "at least one source city or airport is required")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 10),
		RangeEndDate:   time.Now().AddDate(0, 0, 1),
		TripLength:     5,
	}
	testValidate(t, args, "range end date cannot be before range start date")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, -1),
		RangeEndDate:   time.Now().AddDate(0, 0, 10),
		TripLength:     5,
	}
	testValidate(t, args, "range start date cannot be in the past")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 1),
		RangeEndDate:   time.Now().AddDate(0, 0, 10),
		TripLength:     0,
	}
	testValidate(t, args, "trip length must be at least 1 day")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 1),
		RangeEndDate:   time.Now().AddDate(0, 0, 1),
		TripLength:     15,
	}
	testValidate(t, args, "range start date cannot be the same as range end date")

	args = &gflights.PriceGraphArgs{
		SrcCities:      []string{"abc"},
		SrcAirports:    []string{},
		DstCities:      []string{"abc"},
		DstAirports:    []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 1),
		RangeEndDate:   time.Now().AddDate(0, 0, 163),
		TripLength:     15,
	}
	testValidate(t, args, "number of days between dates is larger than 161, 162")
}

func TestValidateExploreArgs(t *testing.T) {
	// Test: no source locations
	args := &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "at least one source city or airport is required")

	// Test: invalid airport code - too short
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{},
		SrcAirports:   []string{"LH"},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "src airport 'LH' is not an airport code")

	// Test: invalid airport code - too long
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{},
		SrcAirports:   []string{"LHRX"},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "src airport 'LHRX' is not an airport code")

	// Test: invalid airport code - lowercase
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{},
		SrcAirports:   []string{"lhr"},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "src airport 'lhr' is not an airport code")

	// Test: north latitude out of bounds (> 90)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 91.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: north latitude out of bounds (< -90)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: -91.0,
			SouthLat: -95.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: south latitude out of bounds (> 90)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 0.0,
			SouthLat: 91.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: south latitude out of bounds (< -90)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 0.0,
			SouthLat: -91.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: east longitude out of bounds (> 180)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  181.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: east longitude out of bounds (< -180)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  -181.0,
			WestLng:  -185.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: west longitude out of bounds (> 180)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  100.0,
			WestLng:  181.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: west longitude out of bounds (< -180)
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  -181.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "coordinates are out of bounds")

	// Test: north latitude equals south latitude
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 10.0,
			SouthLat: 10.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "north latitude must be greater than south latitude")

	// Test: north latitude less than south latitude
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 5.0,
			SouthLat: 10.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "north latitude must be greater than south latitude")

	// Test: east longitude equals west longitude
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  100.0,
			WestLng:  100.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "east longitude must be greater than west longitude")

	// Test: east longitude less than west longitude
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  92.0,
			WestLng:  141.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "east longitude must be greater than west longitude")

	// Test: valid with boundary coordinates
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcAirports:   []string{"LHR"},
		SrcCities:     []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 90.0,
			SouthLat: -90.0,
			EastLng:  180.0,
			WestLng:  -180.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid with only cities
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London", "Paris"},
		SrcAirports:   []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid with only airports
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{},
		SrcAirports:   []string{"LHR", "LGW", "STN"},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid with cities and airports
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcCities:     []string{"London"},
		SrcAirports:   []string{"CDG"},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid Southeast Asia coordinates
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcAirports:   []string{"LHR"},
		SrcCities:     []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 28.0,
			SouthLat: -11.0,
			EastLng:  141.0,
			WestLng:  92.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid very small bounding box
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcAirports:   []string{"LHR"},
		SrcCities:     []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 51.5074,
			SouthLat: 51.5073,
			EastLng:  -0.1277,
			WestLng:  -0.1278,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid crossing equator
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcAirports:   []string{"LHR"},
		SrcCities:     []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 10.0,
			SouthLat: -10.0,
			EastLng:  50.0,
			WestLng:  20.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")

	// Test: valid crossing prime meridian
	args = &gflights.ExploreArgs{
		DepartureDate: time.Now().AddDate(0, 0, 7),
		ReturnDate:    time.Now().AddDate(0, 0, 14),
		SrcAirports:   []string{"LHR"},
		SrcCities:     []string{},
		Coordinates: gflights.ExploreCoordinates{
			NorthLat: 50.0,
			SouthLat: 40.0,
			EastLng:  10.0,
			WestLng:  -10.0,
		},
		Options: gflights.Options{Travelers: gflights.Travelers{Adults: 1}},
	}
	testValidate(t, args, "")
}
