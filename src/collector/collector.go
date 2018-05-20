package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"math/rand"
	"strconv"
)

const NAMESPACE = "benchmark"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Exporter struct {

	error        prometheus.Gauge
	totalScrapes prometheus.Counter
	scrapeErrors *prometheus.CounterVec
	nByte int
	nMetrics int
}

func NewExporter(nByte, nMetrics int) *Exporter{
	return &Exporter{
		nByte: nByte,
		nMetrics: nMetrics,
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: NAMESPACE,
			Subsystem: "",
			Name:      "scrapes_total",
			Help:      "Total number of times jmx was scraped for metrics.",
		}),
		scrapeErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: NAMESPACE,
			Subsystem: "",
			Name:      "scrape_errors_total",
			Help:      "Total number of times an error occurred scraping a jmx.",
		}, []string{"collector"}),
		error: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: NAMESPACE,
			Subsystem: "",
			Name:      "last_scrape_error",
			Help:      "Whether the last scrape of metrics from jmx resulted in an error (1 for error, 0 for success).",
		}),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})

	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()

	e.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	e.scrape(ch, e.nByte, e.nMetrics)
	ch <- e.totalScrapes
	ch <- e.error
	e.scrapeErrors.Collect(ch)
}

func (e *Exporter) scrape(ch chan<- prometheus.Metric, nByte, nMetric int){

	e.totalScrapes.Inc()

	var benchamarkDesc *prometheus.Desc

	for i:=0 ; i<nMetric; i++ {

		help, name := e.randSeq(nByte), e.randSeq(nByte)
		benchamarkDesc = prometheus.NewDesc(
			prometheus.BuildFQName(NAMESPACE, "", name),
			help,
			[]string{"id"}, nil,
		)
		scrapeTime := time.Now()
		ch <- prometheus.MustNewConstMetric(benchamarkDesc, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), strconv.Itoa(i))
	}
}

func (e *Exporter) randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


