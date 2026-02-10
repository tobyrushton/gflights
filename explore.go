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
)

func (s *Session) getExploreRawData(ctx context.Context, args ExploreArgs) (string, error) {
	serSrcs, err := s.serialiseFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Options.Lang)
	if err != nil {
		return "", err
	}

	serDepartureDate := args.DepartureDate.Format("2006-01-02")
	serReturnDate := args.ReturnDate.Format("2006-01-02")

	serTravelers := serialiseFlightTravelers(args.Options.Travelers)
	serStops := serialiseFlightStop(args.Options.Stops)

	rawData := ""

	rawData += fmt.Sprintf(`[[%f,%f],[%f,%f]],null,[null,null,%d,null,[],%d,%s,null,null,null,null,null,null,`,
		args.Coordinates.NorthLat,
		args.Coordinates.EastLng,
		args.Coordinates.SouthLat,
		args.Coordinates.WestLng,
		args.Options.TripType,
		args.Options.Class,
		serTravelers,
	)

	rawData += fmt.Sprintf(`[[[[%s]],[],null,%s,null,null,\"%s\"]`,
		serSrcs, serStops, serDepartureDate)

	if args.Options.TripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[],[[%s]],null,%s,null,null,\"%s\"]`,
			serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func (s *Session) getExploreReqData(ctx context.Context, args ExploreArgs) (string, error) {
	rawData, err := s.getExploreRawData(ctx, args)
	if err != nil {
		return "", err
	}

	prefix := `[null,"[[],`
	suffix := `],null,null,null,1],null,1,null,0,null,1,[694,867],4]"]`

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

// parseFlightData extracts flight information from the Google API response
func parseFlightData(jsonData []byte) ([]ExploreOffer, error) {
	var rawData []any
	if err := json.Unmarshal(jsonData, &rawData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Navigate to the flights array
	// Structure: [0][4][0] contains the array of flight data
	if len(rawData) < 5 {
		return nil, fmt.Errorf("unexpected data structure: insufficient top-level elements")
	}

	mainArray, ok := rawData[4].([]any)
	if !ok || len(mainArray) == 0 {
		return nil, fmt.Errorf("could not find main array at index 4")
	}

	flightsArray, ok := mainArray[0].([]any)
	if !ok {
		return nil, fmt.Errorf("could not find flights array")
	}

	var flights []ExploreOffer

	for _, item := range flightsArray {
		flightData, ok := item.([]any)
		if !ok || len(flightData) < 7 {
			continue
		}

		// Extract city ID (index 0)
		cityID, _ := flightData[0].(string)

		// Extract price from index 1
		// Format: [[null, price], ...]
		var price int
		if priceArray, ok := flightData[1].([]any); ok && len(priceArray) > 0 {
			if priceData, ok := priceArray[0].([]any); ok && len(priceData) > 1 {
				if priceFloat, ok := priceData[1].(float64); ok {
					price = int(priceFloat)
				}
			}
		}

		// Extract airline and airport info from index 6
		// Format: [airline_code, airline_name, stops, total_price, null, airport_code, ...]
		if airlineInfo, ok := flightData[6].([]any); ok && len(airlineInfo) > 5 {
			airlineCode, _ := airlineInfo[0].(string)
			airlineName, _ := airlineInfo[1].(string)
			airportCode, _ := airlineInfo[5].(string)

			// Extract number of stops (index 2)
			stops := 0
			if stopsFloat, ok := airlineInfo[2].(float64); ok {
				stops = int(stopsFloat)
			}

			// Check if it's a multi-carrier flight
			isMulti := airlineCode == "multi"

			flight := ExploreOffer{
				CityID:         cityID,
				AirportCode:    airportCode,
				Price:          price,
				Airline:        airlineName,
				AirlineCode:    airlineCode,
				IsMultiCarrier: isMulti,
				Stops:          stops,
			}

			// Only add flights with valid data
			if flight.AirportCode != "" && flight.Price > 0 {
				flights = append(flights, flight)
			}
		}
	}

	return flights, nil
}

func (s *Session) doExplore(ctx context.Context, args ExploreArgs) (*http.Response, error) {
	url := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetExploreDestinations?f.sid=-6055481678483262293&bl=boq_travel-frontend-flights-ui_20260206.02_p0&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=256256&rt=c"

	reqData, err := s.getExploreReqData(ctx, args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqData +
			`&at=AAuQa1oq5qIkgkQ2nG9vQZFTgSME%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header["cookie"] = s.cookies
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","US","%s",1,null,[-120],null,[[48764689,47907128,48676280,48710756,48627726,48480739,48593234,48707380]],1,[]]`, args.Options.Currency))

	return s.client.Do(req)
}

func (s *Session) GetExplore(ctx context.Context, args ExploreArgs) ([]ExploreOffer, error) {
	resp, err := s.doExplore(ctx, args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	utils.SkipPrefix(body)

	finalOffers := []ExploreOffer{}

	for {
		utils.ReadLine(body) // skip line
		bytesToDecode, err := utils.GetInnerBytes(body)
		if err != nil {
			if err.Error() == "EOF" {
				return finalOffers, nil
			}
			return nil, err
		}
		offers, _ := parseFlightData(bytesToDecode)
		finalOffers = append(finalOffers, offers...)
	}
}
