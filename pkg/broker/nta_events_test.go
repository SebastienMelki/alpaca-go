package broker_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sebastienmelki/alpaca-go/pkg/broker"
)

// Verbatim sandbox lines from GET /v1/events/nta, captured 2026-11-03.
const (
	ntaJNLC  = `{"id":"0ee89bcf-a004-4ed0-b0c5-af44e057f185","system_date":"2026-09-01","settle_date":"2026-09-01","entry_type":"JNLC","description":"DT269010605 deposit","net_amount":40,"symbol":"","qty":0,"status":"executed","currency":"USD","price":null,"per_share_amount":null,"account_id":"517dc9b6-1e23-4ce0-98b5-25bce762fa56","at":"2026-09-01T08:23:35.748617Z","event_ulid":"01M1E122J43G4TFNW408G8MP6A","event_id":36373146}`
	ntaOPEXP = `{"id":"43008efa-531c-4a15-a422-5c54c31868e3","system_date":"2026-09-04","settle_date":"2026-09-04","entry_type":"OPEXP","description":"Options Expiry","net_amount":0,"symbol":"AAPL260904C00327500","qty":-1,"status":"executed","currency":"USD","group_id":"dccb8a7c-026e-45dd-bd17-b2655c2c5095","cusip":"AAPL260904C00327500","price":null,"per_share_amount":null,"account_id":"517dc9b6-1e23-4ce0-98b5-25bce762fa56","at":"2026-09-05T01:00:08.541437Z","event_ulid":"01M1QH8YYXV3YHYBCVGXAMWW5N","event_id":36416561}`
	ntaSWP   = `{"id":"afd85e00-d5e1-424b-a39a-ec09a9a15f21","system_date":"2026-09-01","settle_date":"2026-09-01","entry_type":"SWP","description":"sweep.transactions.id: b873784a-27fc-40a5-bb30-37caaebe689f","net_amount":-2082.9,"symbol":"SWEEPFDIC","qty":2082.9,"status":"executed","currency":"USD","entry_sub_type":"DP","cusip":"SWEEPFDIC","price":"1","per_share_amount":null,"account_id":"2af0afa3-e3c1-461d-aeff-6be838d56992","at":"2026-09-01T19:45:08.694857Z","event_ulid":"01M1F8212PYD4JX63MRWXC2CW1","event_id":36377309}`
)

func writeSSE(w http.ResponseWriter, lines ...string) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ": welcome to the Alpaca events\n\n")
	for _, line := range lines {
		fmt.Fprintf(w, "data: %s\n\n", line)
	}
	w.(http.Flusher).Flush()
}

func TestSubscribeNTAEvents_DecodesSandboxEvents(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/events/nta" {
			t.Errorf("path = %q, want /v1/events/nta", r.URL.Path)
		}
		gotQuery = r.URL.Query()
		writeSSE(w, ntaJNLC, ntaOPEXP, ntaSWP)
	}))
	defer srv.Close()

	client := broker.NewClient("key", "secret",
		broker.WithBaseURL(srv.URL), broker.WithDiscardUnknownFields(true))

	stream, err := client.V1.SubscribeNTAEvents(context.Background(),
		&broker.SubscribeNTAEventsRequest{IncludePreprocessing: true})
	if err != nil {
		t.Fatalf("SubscribeNTAEvents: %v", err)
	}
	defer stream.Close()

	var events []*broker.NTAEvent
	for range 3 {
		ev := &broker.NTAEvent{}
		if !stream.Next(ev) {
			t.Fatalf("Next returned false after %d events: %v", len(events), stream.Err())
		}
		events = append(events, ev)
	}

	if got := gotQuery.Get("include_preprocessing"); got != "true" {
		t.Errorf("include_preprocessing = %q, want true", got)
	}
	if got := gotQuery.Get("since_ulid"); got != "" {
		t.Errorf("since_ulid = %q, want empty on first connect", got)
	}

	if got := events[0].GetId(); got != "0ee89bcf-a004-4ed0-b0c5-af44e057f185" {
		t.Errorf("event 1 Id = %q", got)
	}
	if got := events[0].GetEntryType(); got != "JNLC" {
		t.Errorf("event 1 EntryType = %q", got)
	}
	if got := events[0].GetNetAmount(); got != 40 {
		t.Errorf("event 1 NetAmount = %v, want 40", got)
	}
	if got := events[0].GetPrice(); got != "" {
		t.Errorf("event 1 Price = %q, want empty for null", got)
	}
	if got := events[0].GetEventUlid(); got != "01M1E122J43G4TFNW408G8MP6A" {
		t.Errorf("event 1 EventUlid = %q", got)
	}

	if got := events[1].GetGroupId(); got != "dccb8a7c-026e-45dd-bd17-b2655c2c5095" {
		t.Errorf("event 2 GroupId = %q", got)
	}
	if got := events[1].GetQty(); got != -1 {
		t.Errorf("event 2 Qty = %v, want -1", got)
	}
	if got := events[1].GetAt(); got != "2026-09-05T01:00:08.541437Z" {
		t.Errorf("event 2 At = %q", got)
	}

	if got := events[2].GetPrice(); got != "1" {
		t.Errorf("event 3 Price = %q, want \"1\"", got)
	}
	if got := events[2].GetNetAmount(); got != -2082.9 {
		t.Errorf("event 3 NetAmount = %v, want -2082.9", got)
	}
	if got := events[2].GetQty(); got != 2082.9 {
		t.Errorf("event 3 Qty = %v, want 2082.9", got)
	}
}

func TestSubscribeNTAEvents_SendsSinceULID(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query()
		writeSSE(w, ntaSWP)
	}))
	defer srv.Close()

	client := broker.NewClient("key", "secret",
		broker.WithBaseURL(srv.URL), broker.WithDiscardUnknownFields(true))

	stream, err := client.V1.SubscribeNTAEvents(context.Background(), &broker.SubscribeNTAEventsRequest{
		SinceUlid:            "01M1QH8YYXV3YHYBCVGXAMWW5N",
		IncludePreprocessing: true,
	})
	if err != nil {
		t.Fatalf("SubscribeNTAEvents: %v", err)
	}
	defer stream.Close()

	ev := &broker.NTAEvent{}
	if !stream.Next(ev) {
		t.Fatalf("Next returned false: %v", stream.Err())
	}

	if got := gotQuery.Get("since_ulid"); got != "01M1QH8YYXV3YHYBCVGXAMWW5N" {
		t.Errorf("since_ulid = %q", got)
	}
	if got := gotQuery.Get("include_preprocessing"); got != "true" {
		t.Errorf("include_preprocessing = %q, want true", got)
	}
	if got := ev.GetEventUlid(); got != "01M1F8212PYD4JX63MRWXC2CW1" {
		t.Errorf("EventUlid = %q", got)
	}
}
