package basic_diffs

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	insights_dashboards "github.com/andrei-pokhila/insights-dashboards/gen/go/dashboards"
	"github.com/andrei-pokhila/insights-dashboards/src/storage"
)

//go:embed price_diff.sql
var priceDiffTemplate string

func GetCandles(r *insights_dashboards.BasicRequest) *insights_dashboards.FundingResponse {
	var (
		rawTmpl  bytes.Buffer
		response insights_dashboards.FundingResponse
	)

	tmpl, err := template.New("price_diff.sql").Parse(priceDiffTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&rawTmpl, nil)
	if err != nil {
		panic(err)
	}

	startTime := time.Unix(0, r.GetStart()*int64(time.Millisecond))
	endTime := time.Unix(0, r.GetEnd()*int64(time.Millisecond))
	e := strings.Replace(r.GetExchange().String(), "_", "-", -1)

	conn := storage.NewConnection()
	rows, err := conn.Query(rawTmpl.String(),
		clickhouse.Named("markets", r.GetMarkets()),
		clickhouse.Named("startTime", startTime),
		clickhouse.Named("endTime", endTime),
		clickhouse.Named("exchange", strings.ToLower(e)),
		clickhouse.Named("windowSize", r.GetWindowSize()),
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			point    insights_dashboards.FundingPoint
			datetime time.Time
		)

		if err := rows.Scan(&datetime, &point.Market, &point.FundingRate); err != nil {
			panic(err)
		}

		point.Timestamp = datetime.UnixNano() / int64(time.Millisecond)

		response.Points = append(response.Points, &point)
	}

	return &response
}
