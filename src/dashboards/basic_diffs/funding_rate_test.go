package basic_diffs_test

import (
	"strings"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	insights_dashboards "github.com/andrei-pokhila/insights-dashboards/gen/go/dashboards"
	basicDiffs "github.com/andrei-pokhila/insights-dashboards/src/dashboards/basic_diffs"
	"github.com/andrei-pokhila/insights-dashboards/src/storage"
)

func TestFundingRate(t *testing.T) {
	conn := storage.NewConnection()
	query := basicDiffs.GetFundingQuery()

	var exchange insights_dashboards.Exchange
	exchange = insights_dashboards.Exchange_DYDX_V4
	granularity := "1 hour"

	request := insights_dashboards.BasicRequest{
		Exchange:    exchange,
		Markets:     []string{"BTC-USD"},
		Start:       1714521600000,
		End:         1714607999000,
		Granularity: &granularity,
		WindowSize:  72,
	}

	const (
		E     = "funding"
		TABLE = "exchanges_events"
	)

	startTime := time.Unix(0, request.GetStart()*int64(time.Millisecond))
	endTime := time.Unix(0, request.GetEnd()*int64(time.Millisecond))
	e := strings.Replace(request.GetExchange().String(), "_", "-", -1)
	t.Log(e)

	rows, err := conn.Query(query,
		clickhouse.Named("markets", request.GetMarkets()),
		clickhouse.Named("startTime", startTime),
		clickhouse.Named("endTime", endTime),
		clickhouse.Named("exchange", strings.ToLower(e)),
		clickhouse.Named("windowSize", request.GetWindowSize()),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			point    insights_dashboards.FundingPoint
			datetime time.Time
		)

		if err := rows.Scan(&datetime, &point.Market, &point.FundingRate); err != nil {
			t.Fatal(err)
		}

		point.Timestamp = datetime.UnixNano() / int64(time.Millisecond)

		t.Log(point.String())
	}
}
