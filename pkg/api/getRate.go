package api

import (
	"GoAhead/pkg/model"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"io"
	"net/http"
	"time"
)

const url = "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx"

func (r CentralBankRoute) GetCursByValue(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	val := c.Query("val")

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
		c.String(http.StatusInternalServerError, "Error executing the request: %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.String(resp.StatusCode, "HTTP Error: %s", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf(c, "Error reading the response body: %s", err)
	}

	envelope := new(model.Envelope)

	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing XML: %v", err)
		return
	}

	valutes := envelope.Body.GetCursOnDateResult.ValuteData.ValuteCursOnDate

	response, err := r.uc.GetCursByValue(c, valutes, val)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while outputting XML: %v", err)
	}

	c.JSON(http.StatusOK, response)
}
