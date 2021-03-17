package main

import (
	"client/conf"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var config conf.Conf
var parserErr error
var protocol conf.Protocol

func init() {
	// var confFile string
	// flag.StringVar(&confFile, "c", os.Args[1], "config file")
	// flag.Parse()

	// config, parserErr = conf.ConfParser(confFile, &protocol)
	// if parserErr != nil {
	// 	log.Fatalf("parser config failed:", parserErr.Error())
	// }

	prometheus.MustRegister(uptimeGauge)
}

func recordMetrics() {
	go func() {
		for {
			uptimeGauge.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	uptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "echo_client",
		Name:      "uptime_seconds",
		Help:      "How long the system has been up",
	})
)

type Tracer struct {
	start        time.Time
	dns          time.Time
	tlsHandshake time.Time
	connect      time.Time
}

func (t *Tracer) DNSStart(dsi httptrace.DNSStartInfo) {
	t.dns = time.Now()
}

func (t *Tracer) DNSDone(ddi httptrace.DNSDoneInfo) {
	fmt.Printf("DNS Done: %v\n", time.Since(t.dns))
}

func (t *Tracer) TLSHandshakeStart() {
	t.tlsHandshake = time.Now()
}

func (t *Tracer) TLSHandshakeDone(cs tls.ConnectionState, err error) {
	fmt.Printf("TLS HandShake: %v\n", time.Since(t.tlsHandshake))
}

func (t *Tracer) ConnectStart(network, addr string) {
	t.connect = time.Now()
}

func (t *Tracer) ConnectDone(network, addr string, err error) {
	fmt.Printf("Connect Time: %v\n", time.Since(t.connect))
}

func (t *Tracer) GetFirstResponseByte() {
	fmt.Printf("Time from start to first byte: %v\n", time.Since(t.start))
}

func main() {
	//recordMetrics()
	//start := time.Now().Second()
	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":5000", nil)
	//started := time.Now().Second()
	//uptimeGauge.Set(float64(started - start))

	server := httptest.NewServer(http.HandlerFunc(http.NotFound))
	defer server.Close()

	tracer := &Tracer{}
	trace := &httptrace.ClientTrace{
		DNSStart:             tracer.DNSStart,
		DNSDone:              tracer.DNSDone,
		ConnectStart:         tracer.ConnectStart,
		ConnectDone:          tracer.ConnectDone,
		TLSHandshakeStart:    tracer.TLSHandshakeStart,
		TLSHandshakeDone:     tracer.TLSHandshakeDone,
		GotFirstResponseByte: tracer.GetFirstResponseByte,
	}

	req, _ := http.NewRequest("GET", "http://10.244.76.17:5000", nil)
	tracer.start = time.Now()
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total time: %v\n", time.Since(tracer.start))
}
