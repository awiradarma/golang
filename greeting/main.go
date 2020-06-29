package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	m := os.Getenv("MESSAGE")
	c := os.Getenv("COLLECTOR")
	remote := os.Getenv("REMOTE")
	h, _ := os.Hostname()
	if m == "" {
		m = "hello"
	}
	if c == "" {
		c = "localhost:6831"
	}

	tracer, reporter, closer, err := ginopentracing.InitTracing(fmt.Sprintf("greeting-api-go-gin:%s::", m), c, ginopentracing.WithEnableInfoLog(false))
	if err != nil {
		panic("unable to init tracing")
	}
	defer closer.Close()
	defer reporter.Close()
	opentracing.SetGlobalTracer(tracer)

	t := ginopentracing.OpenTracer([]byte("api-request-"))

	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	r.Use(t)
	p.Use(r)
	r.GET("/greetings", func(c *gin.Context) {
		var span opentracing.Span
		if gin.IsDebugging() {
			logHttpHeaders(c)
			logGinContext(c)
		}

		if cspan, ok := c.Get("tracing-context"); ok {
			span = ginopentracing.StartSpanWithParent(cspan.(opentracing.Span).Context(), "greeting", c.Request.Method, c.Request.URL.Path)
		} else {
			span = ginopentracing.StartSpanWithHeader(&c.Request.Header, "greeting", c.Request.Method, c.Request.URL.Path)
		}
		defer span.Finish()

		remoteResponse := ""
		if remote != "" {
			httpClient := &http.Client{}
			httpReq, _ := http.NewRequest("GET", remote, nil)

			// Transmit the span's TraceContext as HTTP headers on our
			// outbound request.
			opentracing.GlobalTracer().Inject(
				span.(opentracing.Span).Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(httpReq.Header))

			resp, err := httpClient.Do(httpReq)
			if err != nil {
				log.Fatal("Error reading response. ", err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("Error reading body. ", err)
			}

			remoteResponse = fmt.Sprintf(" -> %s", body)
		}

		c.String(200, m+" ("+h+")"+remoteResponse)
	})
	r.Run()
}

func logHttpHeaders(c *gin.Context) {
	for k, vals := range c.Request.Header {
		log.Printf("%s", k)
		for _, v := range vals {
			log.Printf("\t%s", v)
		}
		log.Printf("\n")
	}
}

func logGinContext(c *gin.Context) {
	for k, vals := range c.Keys {
		log.Println(k, vals)
	}
}
