package gflights

import (
	"bufio"
	"io"

	"github.com/tobyrushton/gflights/internal/utils"
)

func decodeMessage(body io.ReadCloser, out any) error {
	b := bufio.NewReader(body)
	utils.SkipPrefix(b)

	for {
		utils.ReadLine(b)
		bytesToDecode, err := utils.GetInnerBytes(b)
		if err != nil { // EOF
			return nil
		}

		switch out.(type) {
		case *[]SimpleOffer:
			err := decodeSimpleOffer(bytesToDecode, out.(*[]SimpleOffer))
			if err != nil {
				return err
			}
		case *[]OutboundOffer:
			break
		}
	}
}

func decodeSimpleOffer(bytesToDecode []byte, out *[]SimpleOffer) error {
	offers, _ := getPriceCalendarSection(bytesToDecode)
	*out = append(*out, offers...)
	return nil
}
