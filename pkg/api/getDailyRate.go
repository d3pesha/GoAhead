package api

import (
	"GoAhead/pkg/model"
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (r CentralBankRoute) GetDailyRate() error {
	date := time.Now().Format("2006-01-02")

	xmlBody := fmt.Sprintf(`
		<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
		  <soap12:Body>
		    <GetCursOnDate xmlns="http://web.cbr.ru/">
		      <On_date>%s</On_date>
		    </GetCursOnDate>
		  </soap12:Body>
		</soap12:Envelope>
	`, date)

	resp, err := http.Post(url, "application/soap+xml; charset=utf-8", bytes.NewBufferString(xmlBody))
	if err != nil {
		return fmt.Errorf("error executing the request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading the response body: %s", err)
	}

	envelope := new(model.Envelope)

	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		return fmt.Errorf("error parsing XML: %v", err)
	}

	valutes := envelope.Body.GetCursOnDateResult.ValuteData.ValuteCursOnDate

	err = r.uc.GetDailyRate(context.Background(), valutes)
	if err != nil {
		return fmt.Errorf("error while outputting XML: %v", err)
	}

	return nil
}
