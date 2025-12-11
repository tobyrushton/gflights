package gflights

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/tobyrushton/gflights/internal/utils"
)

func (s *Session) getPriceGraphReqData(ctx context.Context, args PriceGraphArgs) (string, error) {
	serializedRangeStartDate := args.RangeStartDate.Format("2006-01-02")
	serializedRangeEndDate := args.RangeEndDate.Format("2006-01-02")

	rawData, err := s.getCalendarRawData(ctx, calendarArgs{
		RangeStartDate: args.RangeStartDate,
		RangeEndDate:   args.RangeEndDate,
		SrcCities:      args.SrcCities,
		SrcAirports:    args.SrcAirports,
		DstCities:      args.DstCities,
		DstAirports:    args.DstAirports,
		Options:        args.Options,
	})
	if err != nil {
		return "", err
	}

	prefix := `[null,"[null,`
	suffix := fmt.Sprintf(`],null,null,null,1],[\"%s\",\"%s\"]]"]`,
		serializedRangeStartDate, serializedRangeEndDate)
	if args.Options.TripType == RoundTrip {
		suffix = fmt.Sprintf(`],null,null,null,1,null,null,null,null,null,[]],[\"%s\",\"%s\"],null,[%d,%d]]"]`,
			serializedRangeStartDate, serializedRangeEndDate, args.TripLength, args.TripLength)
	}

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestPriceGraph(ctx context.Context, args PriceGraphArgs) (*http.Response, error) {
	url := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetCalendarGraph?f.sid=-8920707734915550076&bl=boq_travel-frontend-ui_20230627.07_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=261464&rt=c"

	reqDate, err := s.getPriceGraphReqData(ctx, args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
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

func (s *Session) GetPriceGraph(ctx context.Context, args PriceGraphArgs) ([]SimpleOffer, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	offers := []SimpleOffer{}

	resp, err := s.doRequestPriceGraph(ctx, args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	utils.SkipPrefix(body)

	for {
		utils.ReadLine(body) // skip line
		bytesToDecode, err := utils.GetInnerBytes(body)
		if err != nil {
			sort.Slice(offers, func(i, j int) bool {
				return offers[i].DepartureDate.Before(offers[j].DepartureDate)
			})
			return offers, nil
		}

		offers_, _ := getPriceCalendarSection(bytesToDecode)
		offers = append(offers, offers_...)
	}
}
