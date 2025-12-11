package gflights

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tobyrushton/gflights/internal/utils"
)

func (s *Session) getPriceGridReqData(ctx context.Context, args PriceGridArgs) (string, error) {
	if err := args.Validate(); err != nil {
		return "", err
	}
	serStartDepartRange := args.StartDepartureRange.Format("2006-01-02")
	serEndDepartRange := args.EndDepartureRange.Format("2006-01-02")
	serStartReturnRange := args.StartReturnRange.Format("2006-01-02")
	serEndReturnRange := args.EndReturnRange.Format("2006-01-02")

	rawData, err := s.getCalendarRawData(ctx, calendarArgs{
		RangeStartDate: args.StartDepartureRange,
		RangeEndDate:   args.StartReturnRange,
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
	suffix := fmt.Sprintf(`],null,null,null,1],[\"%s\",\"%s\"],[\"%s\",\"%s\"]]"]`,
		serStartDepartRange, serEndDepartRange, serStartReturnRange, serEndReturnRange)

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestPriceGrid(ctx context.Context, args PriceGridArgs) (*http.Response, error) {
	url := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetCalendarGrid?f.sid=4413567882004396298&bl=boq_travel-frontend-flights-ui_20251209.02_p0&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=1168350&rt=c"

	reqData, err := s.getPriceGridReqData(ctx, args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqData +
			`&at=AI2cvzaUCzY0WjuQdaTUyWMizEwC%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

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

func (s *Session) GetPriceGrid(ctx context.Context, args PriceGridArgs) ([]SimpleOffer, error) {
	resp, err := s.doRequestPriceGrid(ctx, args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	offers := make([]SimpleOffer, 0)

	body := bufio.NewReader(resp.Body)
	utils.SkipPrefix(body)

	for {
		utils.ReadLine(body) // skip line
		bytesToDecode, err := utils.GetInnerBytes(body)
		if err != nil {
			return offers, nil
		}
		offers_, _ := getPriceCalendarSection(bytesToDecode)
		offers = append(offers, offers_...)
	}
}
