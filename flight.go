package gflights

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tobyrushton/gflights/internal/utils"
	"golang.org/x/text/language"
)

const (
	flightAirportConst rune = '0'
	flightCityConst    rune = '5'
)

type rawDataArgs struct {
	Args

	OutboundFlights []Flight
}

type requestFlightArgs struct {
	rawDataArgs

	FlightsSession string
}

type offer struct {
	Flight         []Flight
	StartDate      time.Time
	ReturnDate     time.Time
	Duration       time.Duration
	SrcAirportCode string
	DstAirportCode string
	SrcCity        string
	DstCity        string
	Price          float64
}

func serialiseFlightStop(stops Stops) string {
	switch stops {
	case NonStop:
		return "1"
	case OneStop:
		return "2"
	case TwoPlusStops:
		return "3"
	}
	return "0"
}

func (s *Session) serialiseFlightLocations(ctx context.Context, cities []string, airports []string, Lang language.Tag) (string, error) {
	abbrCities, err := s.abbrCities(ctx, cities, Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbrivated city names: %v", err)
	}

	serialised := ""

	for _, l := range airports {
		serialised += fmt.Sprintf(`[\"%s\",%c],`, l, flightAirportConst)
	}
	for _, l := range abbrCities {
		serialised += fmt.Sprintf(`[\"%s\",%c],`, l, flightCityConst)
	}

	return serialised[:len(serialised)-1], nil
}

func serialiseFlightTravelers(travelers Travelers) string {
	return fmt.Sprintf(
		`[%d,%d,%d,%d]`,
		travelers.Adults,
		travelers.Children,
		travelers.InfantsOnLap,
		travelers.InfantsInSeat,
	)
}

func serialiseFlightRoute(flights []Flight) string {
	if len(flights) == 0 {
		return ""
	}

	serFlights := ""

	for _, flight := range flights {
		serFlights += fmt.Sprintf(
			`[\"%s\",\"%s\",\"%s\",null,\"%s\",\"%s\"],`,
			flight.DepAirportCode,
			flight.DepTime.Format("2006-01-02"),
			flight.ArrAirportCode,
			flight.FlightCode.AirlineCode,
			flight.FlightCode.FlightNumber,
		)
	}
	return serFlights[:len(serFlights)-1]
}

func (s *Session) getRawData(ctx context.Context, args rawDataArgs) (string, error) {
	serSrcs, err := s.serialiseFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not serialize src flight src locations: %v", err)
	}
	serDsts, err := s.serialiseFlightLocations(ctx, args.DstCities, args.DstAirports, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not serialize src flight dst locations: %v", err)
	}

	serDate := args.DepartureDate.Format("2006-01-02")
	serReturnDate := args.ReturnDate.Format("2006-01-02")

	serAdults := serialiseFlightTravelers(args.Options.Travelers)
	serStops := serialiseFlightStop(args.Options.Stops)
	outboundFlights := serialiseFlightRoute(args.OutboundFlights)

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%d,%s,null,null,null,null,null,null,[`,
		args.Options.TripType, args.Options.Class, serAdults)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[%s],[],[],null,null,[],3]`,
		serSrcs, serDsts, serStops, serDate, outboundFlights)

	if args.Options.TripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
			serDsts, serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func (s *Session) getFlightReqData(ctx context.Context, args requestFlightArgs) (string, error) {
	rawData, err := s.getRawData(ctx, args.rawDataArgs)
	if err != nil {
		return "", err
	}

	prefix := `[null,"[[],`
	if args.FlightsSession != "" {
		prefix = fmt.Sprintf(`[null, "[[null, %s],`, args.FlightsSession)
	}
	suffix := `],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestFlights(ctx context.Context, args requestFlightArgs) (*http.Response, error) {
	url := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	reqData, err := s.getFlightReqData(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("could not get request data: %v", err)
	}

	jsonBody := []byte(
		`f.req=` + reqData +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to do GetShoppingResults request: %v", err)
	}
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header["cookie"] = s.cookies
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","US","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, args.Options.Currency)) // language, location, Currency

	return s.client.Do(req)
}

func getFlightsDuration(flights []Flight) time.Duration {
	return flights[len(flights)-1].ArrTime.Sub(flights[0].DepTime)
}

func flightSchema(
	flight *Flight,
	depYear, depMonth, depDay, depHours, depMinutes,
	arrYear, arrMonth, arrDay, arrHours, arrMinutes,
	duration *float64,
	flightNoPart1, flightNoPart2 *string,
) *[]any {
	return &[]any{
		&flight.Unknown[0],
		&flight.Unknown[1],
		&flight.Unknown[2],
		&flight.DepAirportCode,
		&flight.DepAirportName,
		&flight.ArrAirportName,
		&flight.ArrAirportCode,
		&flight.Unknown[3],
		&[]any{&depHours, &depMinutes},
		&flight.Unknown[4],
		&[]any{&arrHours, &arrMinutes},
		&duration,
		&flight.Unknown[5],
		&flight.Unknown[6],
		&flight.Unknown[7],
		&flight.Unknown[8],
		&flight.Unknown[9],
		&flight.Airplane,
		&flight.Unknown[10],
		&flight.Unknown[11],
		&[]any{&depYear, &depMonth, &depDay},
		&[]any{&arrYear, &arrMonth, &arrDay},
		&[]any{&flightNoPart1, &flightNoPart2, nil, &flight.AirlineName},
		&flight.Unknown[12],
		&flight.Unknown[13],
		&flight.Unknown[14],
		&flight.Unknown[15],
		&flight.Unknown[16],
		&flight.Unknown[17],
		&flight.Unknown[18],
		&flight.Legroom,
		&flight.Unknown[19],
	}
}

func getFlights(rawFlights []json.RawMessage) ([]Flight, error) {
	flights := []Flight{}
	for _, rawFlight := range rawFlights {
		flight := Flight{}
		flight.Unknown = make([]any, 20)
		var depHours, depMinutes, arrHours, arrMinutes, duration,
			depYear, depMonth, depDay, arrYear, arrMonth, arrDay float64
		var airlineCode, flightNumber string
		if err := json.Unmarshal(rawFlight, flightSchema(
			&flight,
			&depYear, &depMonth, &depDay, &depHours, &depMinutes,
			&arrYear, &arrMonth, &arrDay, &arrHours, &arrMinutes,
			&duration,
			&airlineCode, &flightNumber,
		)); err != nil {
			return nil, err
		}
		depCity, depLocation := iataLocation(flight.DepAirportCode)
		arrCity, arrLocation := iataLocation(flight.ArrAirportCode)
		flight.DepCity = depCity
		flight.DepTime = time.Date(int(depYear), time.Month(depMonth), int(depDay), int(depHours), int(depMinutes), 0, 0, depLocation)
		flight.ArrCity = arrCity
		flight.ArrTime = time.Date(int(arrYear), time.Month(arrMonth), int(arrDay), int(arrHours), int(arrMinutes), 0, 0, arrLocation)
		parsedDuration, _ := time.ParseDuration(fmt.Sprintf("%dm", int(duration)))
		flight.Duration = parsedDuration
		flight.FlightCode = FlightCode{
			AirlineCode:  airlineCode,
			FlightNumber: flightNumber,
		}
		flights = append(flights, flight)
	}
	return flights, nil
}

func offerSchema(rawFlights *[]json.RawMessage, price *float64) *[]any {
	return &[]any{&[]any{nil, nil, rawFlights}, &[]any{&[]any{nil, price}}}
}

func getSubsectionOffers(rawOffers []json.RawMessage, returnDate time.Time) ([]offer, error) {
	offers := []offer{}
	for _, rawOffer := range rawOffers {
		offer := offer{}
		rawFlights := []json.RawMessage{}

		if string(rawOffer) == "null" {
			continue
		}

		if err := json.Unmarshal(rawOffer, offerSchema(&rawFlights, &offer.Price)); err != nil {
			continue
		}

		flights, err := getFlights(rawFlights)
		if err != nil || len(flights) == 0 {
			continue
		}

		offer.Flight = flights

		offer.StartDate = flights[0].DepTime
		offer.ReturnDate = returnDate.UTC()

		offer.Duration = getFlightsDuration(flights)

		offer.SrcAirportCode = flights[0].DepAirportCode
		offer.DstAirportCode = flights[len(flights)-1].ArrAirportCode

		offer.SrcCity = flights[0].DepCity
		offer.DstCity = flights[len(flights)-1].ArrCity

		offers = append(offers, offer)
	}
	return offers, nil
}

func sectionOffersSchema(rawOffers1, rawOffers2 *[]json.RawMessage, priceRange *PriceRange) *[]any {
	return &[]any{nil, nil, &[]any{rawOffers1}, &[]any{rawOffers2}, nil, &[]any{nil, nil, nil, nil,
		&[]any{nil, &priceRange.Min}, &[]any{nil, &priceRange.Max}}}
}

func getSectionOffers(bytesToDecode []byte, returnDate time.Time) ([]offer, *PriceRange, error) {
	rawOffers1 := []json.RawMessage{}
	rawOffers2 := []json.RawMessage{}

	priceRange := PriceRange{}

	if err := json.Unmarshal(bytesToDecode, sectionOffersSchema(&rawOffers1, &rawOffers2, &priceRange)); err != nil {
		return nil, nil, err
	}

	allOffers := []offer{}
	offers1, err := getSubsectionOffers(rawOffers1, returnDate)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers1...)
	offers2, err := getSubsectionOffers(rawOffers2, returnDate)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers2...)

	return allOffers, &priceRange, nil
}

func getToken(bytes []byte) string {
	var data []any
	if err := json.Unmarshal(bytes, &data); err != nil || len(data) < 2 {
		return ""
	}

	innerData, ok := data[0].([]any)
	if !ok || len(innerData) < 5 {
		return ""
	}
	token, ok := innerData[4].(string)
	if !ok {
		return ""
	}

	return token
}

func (s *Session) doGetOffers(ctx context.Context, args requestFlightArgs) ([]offer, *PriceRange, string, error) {
	finalOffers := []offer{}
	var finalPriceRange *PriceRange

	resp, err := s.doRequestFlights(ctx, args)
	if err != nil {
		return nil, nil, "", err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	utils.SkipPrefix(body)
	token := ""

	for {
		utils.ReadLine(body) // skip line
		bytesToDecode, err := utils.GetInnerBytes(body)
		if err != nil {
			return finalOffers, finalPriceRange, token, nil
		}

		if token == "" {
			token = getToken(bytesToDecode)
		}

		offers, priceRange, _ := getSectionOffers(bytesToDecode, args.ReturnDate)
		finalOffers = append(finalOffers, offers...)
		if priceRange != nil {
			finalPriceRange = priceRange
		}
	}
}

func (s *Session) GetOutboundOffers(ctx context.Context, args Args) ([]OutboundOffer, *PriceRange, error) {
	if err := args.Validate(); err != nil {
		return nil, nil, err
	}

	offers, priceRange, token, err := s.doGetOffers(ctx, requestFlightArgs{
		rawDataArgs: rawDataArgs{
			Args: args,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	finalOffers := make([]OutboundOffer, len(offers))
	for i, offer := range offers {
		finalOffers[i] = OutboundOffer{
			SimpleOffer: SimpleOffer{
				DepartureDate: offer.StartDate,
				Price:         offer.Price,
			},
			Flight:         offer.Flight,
			SrcAirportCode: offer.SrcAirportCode,
			DstAirportCode: offer.DstAirportCode,
			SrcCity:        offer.SrcCity,
			DstCity:        offer.DstCity,
			s:              s,
			token:          token,
			args:           args,
		}
	}

	return finalOffers, priceRange, nil
}
