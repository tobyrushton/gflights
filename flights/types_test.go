package flights

import (
	"testing"
	"time"
)

type withValidate interface {
	Validate() error
}

var wrongAirportCode = "wrong"

func testValidate(t *testing.T, args withValidate, wantErr string) {
	gotErr := args.Validate()

	if gotErr == nil && wantErr != "" {
		t.Fatalf("expected error '%s', got nil", wantErr)
	} else if gotErr.Error() != wantErr {
		t.Fatalf("expected error '%s', got '%s'", wantErr, gotErr.Error())
	}
}

func TestValidateArgs(t *testing.T) {
	args := &Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "at least one destination city or airport is required")

	args = &Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "at least one source city or airport is required")

	args = &Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "src airport 'wrong' is not an airport code")

	args = &Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}, DepartureDate: time.Now().AddDate(0, 0, 1)}
	testValidate(t, args, "dst airport 'wrong' is not an airport code")

	args = &Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		DepartureDate: time.Now().AddDate(0, 0, 3),
		ReturnDate:    time.Now().AddDate(0, 0, 1),
	}
	testValidate(t, args, "return date cannot be before departure date")

	args = &Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		DepartureDate: time.Now().AddDate(0, 0, -1),
		ReturnDate:    time.Now().AddDate(0, 0, 1),
	}
	testValidate(t, args, "departure date cannot be in the past")
}

func TestValidateTravelers(t *testing.T) {
	args := &Travelers{}
	testValidate(t, args, "at least one adult traveler is required")

	args = &Travelers{
		Adults:   1,
		Children: -1,
	}

	testValidate(t, args, "number of children and infants cannot be negative")

	args = &Travelers{
		Adults:       1,
		Children:     1,
		InfantsOnLap: 2,
	}
	testValidate(t, args, "each infant on lap must be accompanied by an adult")

	args = &Travelers{
		Adults:        1,
		Children:      1,
		InfantsOnLap:  1,
		InfantsInSeat: 2,
	}
	testValidate(t, args, "one adult is required for every two infants")
}
