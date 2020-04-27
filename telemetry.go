package assembled

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptrace"
	"time"
)

type events struct {
	DNSStart             int64 `json:"dns_start,string,omitempty"`
	DNSDone              int64 `json:"dns_done,string,omitempty"`
	ConnectStart         int64 `json:"connect_start,string,omitempty"`
	ConnectDone          int64 `json:"connect_done,string,omitempty"`
	TLSHandshakeStart    int64 `json:"tls_start,string,omitempty"`
	TLSHandshakeDone     int64 `json:"tls_done,string,omitempty"`
	GotConn              int64 `json:"got_conn,string,omitempty"`
	GotFirstResponseByte int64 `json:"got_byte,string,omitempty"`
	DecodeBodyDone       int64 `json:"decode_body_done,string,omitempty"`
}

type timing struct {
	Start      time.Time `json:"start"`
	Events     events    `json:"events"`
	StatusCode int       `json:"code"`
	ID         string    `json:"id"`
}

func (t *timing) D() int64 {
	return int64(time.Now().UTC().Sub(t.Start) / time.Millisecond)
}

func withTelemetry(ctx context.Context, c *Client, r *http.Request) (*http.Request, func(*http.Response)) {
	select {
	case metric := <-c.metrics:
		if metric != nil {
			payload, err := json.Marshal(&metric)
			if err == nil {
				r.Header.Set("Client-Telemetry", string(payload))
			}
		}
	default:
		// pass
	}

	t := &timing{Start: time.Now().UTC()}
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) { t.Events.DNSStart = t.D() },
		DNSDone:  func(_ httptrace.DNSDoneInfo) { t.Events.DNSDone = t.D() },
		ConnectStart: func(_, _ string) {
			if t.Events.DNSDone == 0 {
				// connecting to IP
				t.Events.DNSDone = t.D()
			}
		},
		ConnectDone: func(net, addr string, err error) {
			t.Events.ConnectDone = t.D()
		},
		GotConn:              func(_ httptrace.GotConnInfo) { t.Events.GotConn = t.D() },
		GotFirstResponseByte: func() { t.Events.GotFirstResponseByte = t.D() },
		TLSHandshakeStart:    func() { t.Events.TLSHandshakeStart = t.D() },
		TLSHandshakeDone:     func(_ tls.ConnectionState, _ error) { t.Events.TLSHandshakeDone = t.D() },
	}

	return r.WithContext(httptrace.WithClientTrace(ctx, trace)), func(resp *http.Response) {
		if resp == nil {
			return
		}
		id := resp.Header.Get("Request-Id")
		if len(id) == 0 {
			return
		}

		t.ID = id
		t.Events.DecodeBodyDone = t.D()

		select {
		case c.metrics <- t:
		default:
		}
	}
}
