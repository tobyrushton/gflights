package gflights

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/tobyrushton/gflights/internal/pb"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/proto"
)

func serialiseLocations(locations []string, locationType pb.LocationType) []*pb.Location {
	locationsRet := make([]*pb.Location, 0, len(locations))
	for _, l := range locations {
		locationsRet = append(locationsRet, &pb.Location{
			Type: locationType,
			Name: l,
		})
	}
	return locationsRet
}

func serialiseSearchFlight(
	date time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	stops Stops,
) *pb.SearchFlight {
	return &pb.SearchFlight{
		Date:         date.Format(time.DateOnly),
		SrcLocations: append(serialiseLocations(srcCities, pb.LocationType_CITY), serialiseLocations(srcAirports, pb.LocationType_AIRPORT)...),
		DstLocations: append(serialiseLocations(dstCities, pb.LocationType_CITY), serialiseLocations(dstAirports, pb.LocationType_AIRPORT)...),
		Stops:        pb.Stops(stops).Enum(),
	}
}

func serialiseSearchFlights(args Args) []*pb.SearchFlight {
	if args.Options.TripType == OneWay {
		return []*pb.SearchFlight{
			serialiseSearchFlight(args.DepartureDate, args.SrcCities, args.SrcAirports, args.DstCities, args.DstAirports, args.Options.Stops),
		}
	}
	return []*pb.SearchFlight{
		serialiseSearchFlight(args.DepartureDate, args.SrcCities, args.SrcAirports, args.DstCities, args.DstAirports, args.Options.Stops),
		serialiseSearchFlight(args.ReturnDate, args.DstCities, args.DstAirports, args.SrcCities, args.SrcAirports, args.Options.Stops),
	}
}

func serialiseFlights(flights []Flight) []*pb.Flight {
	serFlights := make([]*pb.Flight, 0, len(flights))

	for _, flight := range flights {
		serFlights = append(serFlights, &pb.Flight{
			DepCode:      flight.DepAirportCode,
			Date:         flight.DepTime.Format(time.DateOnly),
			DestCode:     flight.ArrAirportCode,
			AirlineCode:  flight.FlightCode.AirlineCode,
			FlightNumber: flight.FlightCode.FlightNumber,
		})
	}

	return serFlights
}

func serialiseBookingFlights(trip *TripSelection, src []string, dst []string) []*pb.BookingFlight {
	serBookingFlights := make([]*pb.BookingFlight, 0)

	for _, segment := range trip.Segments {
		serBookingFlights = append(serBookingFlights, &pb.BookingFlight{
			Date:         segment.Legs[0].DepTime.Format(time.DateOnly),
			Flights:      serialiseFlights(segment.Legs),
			SrcLocations: serialiseLocations(src, pb.LocationType_CITY),
			DstLocations: serialiseLocations(dst, pb.LocationType_CITY),
		})
	}

	return serBookingFlights
}

func serialiseTravelers(travelers Travelers) []pb.Traveler {
	travelersRet := make([]pb.Traveler, 0,
		travelers.Adults+travelers.Children+
			travelers.InfantsInSeat+travelers.InfantsOnLap)

	for i := 0; i < travelers.Adults; i++ {
		travelersRet = append(travelersRet, pb.Traveler_ADULT)
	}
	for i := 0; i < travelers.Children; i++ {
		travelersRet = append(travelersRet, pb.Traveler_CHILD)
	}
	for i := 0; i < travelers.InfantsInSeat; i++ {
		travelersRet = append(travelersRet, pb.Traveler_INFANT_IN_SEAT)
	}
	for i := 0; i < travelers.InfantsOnLap; i++ {
		travelersRet = append(travelersRet, pb.Traveler_INFANT_ON_LAP)
	}
	return travelersRet
}

func (s *Session) SerialiseURL(ctx context.Context, args Args) (string, error) {
	var err error

	if err = args.Validate(); err != nil {
		return "", err
	}

	args.SrcCities, err = s.abbrCities(ctx, args.SrcCities, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbreviated src cities: %v", err)
	}

	args.DstCities, err = s.abbrCities(ctx, args.DstCities, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbreviated dst cities: %v", err)
	}
	urlProto := &pb.SearchUrl{
		Flight:    serialiseSearchFlights(args),
		Travelers: serialiseTravelers(args.Options.Travelers),
		Class:     pb.Class(args.Options.Class),
		TripType:  pb.TripType(args.Options.TripType),
	}

	tfs, err := proto.Marshal(urlProto)
	if err != nil {
		return "", fmt.Errorf("error during url serialization: %s", err)
	}

	return "https://www.google.com/travel/flights/search" +
		"?tfs=" + base64.RawURLEncoding.EncodeToString(tfs) +
		"&curr=" + args.Options.Currency.String() +
		"&hl=" + args.Options.Lang.String(), nil
}

func (s *Session) SerialiseBookingURL(ctx context.Context, trip *TripSelection) (string, error) {
	if trip == nil {
		return "", fmt.Errorf("trip selection cannot be nil")
	}
	if len(trip.Segments) == 0 {
		return "", fmt.Errorf("trip selection must have at least one segment")
	}
	srcCity, err := s.abbrCities(
		ctx,
		[]string{trip.Segments[0].Legs[0].DepCity},
		language.English,
	)
	if err != nil {
		return "", fmt.Errorf("could not abbreviate source city: %v", err)
	}

	dstCity, err := s.abbrCities(
		ctx,
		[]string{trip.Segments[0].Legs[len(trip.Segments[0].Legs)-1].ArrCity},
		language.English,
	)
	if err != nil {
		return "", fmt.Errorf("could not abbreviate destination city: %v", err)
	}

	urlProto := &pb.BookingUrl{
		Flight:    serialiseBookingFlights(trip, srcCity, dstCity),
		Travelers: serialiseTravelers(trip.Travelers),
		Class:     pb.Class(trip.Class),
		TripType:  pb.TripType(trip.TripType),
	}

	tfs, err := proto.Marshal(urlProto)
	if err != nil {
		return "", fmt.Errorf("error during booking url serialization: %s", err)
	}

	return "https://www.google.com/travel/flights/booking" +
		"?tfs=" + base64.RawURLEncoding.EncodeToString(tfs), nil
}
