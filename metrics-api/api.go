package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmot"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Metrics struct {
	Ok          bool `header:"Ok"`
	Cpu         int  `header:"Cpu"`
	CpuBurst    int  `header:"CpuBurst"`
	Memory      int  `header:"Memory"`
	MemoryBurst int  `header:"MemoryBurst"`
	Count       int  `header:"Count"`
}

type Message struct {
	Message string `header:"Message"`
}

var m = Metrics{
	Ok:          true,
	Cpu:         0,
	CpuBurst:    0,
	Memory:      0,
	MemoryBurst: 0,
	Count:       0,
}

// @Summary Ping
// @Produce json
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// @Summary Retrieve metrics
// @Produce  json
// @Router /metrics [get]
func getMetrics(c *gin.Context) {
	m.Cpu = 20 + m.CpuBurst + rand.Intn(10)
	m.Memory = 20 + m.MemoryBurst + rand.Intn(10)
	m.Count++

	if err := c.ShouldBindHeader(&m); err != nil {
		c.JSON(200, err)
	}

	c.JSON(200, gin.H{
		"Ok":     m.Ok,
		"Cpu":    m.Cpu,
		"Memory": m.Memory,
		"Count":  m.Count,
	})
}

// @Summary Produce error
// @Produce json
// @Router /error [get]
func makeError(c *gin.Context) {
	m.CpuBurst = 70
	m.MemoryBurst = 70
	m.Ok = false
	c.JSON(200, gin.H{
		"message": "stay calm, it happens",
	})
}

// @Summary Repair the error
// @Produce json
// @Router /repair [get]
func repair(c *gin.Context) {
	m.CpuBurst = 0
	m.MemoryBurst = 0
	m.Ok = true
	c.JSON(200, gin.H{
		"message": "error repaired, hooray!",
	})
}

// @Summary Sleep for N seconds
// @Produce json
// @Router /sleep [get]
func sleep(c *gin.Context) {
	durationString := c.Param("duration")
	durationInt, err := strconv.Atoi(durationString)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Cannot convert to integer.",
		})
		return
	}
	time.Sleep(time.Duration(durationInt) * time.Second)
	message := "Slept for " + durationString + " seconds."
	c.String(200, message)
}

// @Summary Connect to a web page
// @Produce json
// @Router /connect [get]
func connect(c *gin.Context) {
	// Retrieve the original context
	ctx := c.Request.Context()

	// Declare span - global
	spanGlobal, ctx := opentracing.StartSpanFromContext(ctx, "global")

	// Retrieve page
	page := c.Query("page")
	start := time.Now()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", page, nil)
	resp, err := client.Do(req)

	defer resp.Body.Close()
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		c.JSON(resp.StatusCode, gin.H{
			"message": err,
		})
		return
	}
	message := "Connecting to " + page + " spent " + elapsed.Truncate(time.Millisecond).String()
	c.String(resp.StatusCode, message)

	// End span - global
	spanGlobal.Finish()
}

// @Summary Connect to two web pages
// @Produce json
// @Router /connect2 [get]
func connect2(c *gin.Context) {
	// Retrieve the original context
	ctx := c.Request.Context()

	// Use Elastic APM library
	opentracing.SetGlobalTracer(apmot.New())

	// Declare span - global
	spanGlobal, ctx := opentracing.StartSpanFromContext(ctx, "global")
	//spanGlobal, ctx := apm.StartSpan(ctx, "connect2", "custom")

	// Retrieve pages
	page1 := c.Query("page1")
	page2 := c.Query("page2")

	// Declare span - page1
	spanPage1, _ := opentracing.StartSpanFromContext(ctx, "page1")
	//spanPage1, ctx := apm.StartSpan(ctx, "page1", "custom")

	// Connect to first page
	start1 := time.Now()
	client1 := &http.Client{}
	req1, _ := http.NewRequest("GET", page1, nil)

	// Serialize the wire, injecting context into headers
	opentracing.GlobalTracer().Inject(
		spanPage1.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req1.Header))
	resp1, err := client1.Do(req1)

	defer resp1.Body.Close()
	end1 := time.Now()
	elapsed1 := end1.Sub(start1)
	if err != nil {
		c.JSON(resp1.StatusCode, gin.H{
			"message": err,
		})
		return
	}
	message := "Connecting to " + page1 + " spent " + elapsed1.Truncate(time.Millisecond).String()

	// End span - page1
	spanPage1.Finish()
	//spanPage1.End()

	// Declare span - page2
	spanPage2, _ := opentracing.StartSpanFromContext(ctx, "page2")
	//spanPage2, ctx := apm.StartSpan(ctx, "page2", "custom")

	// Connect to second page
	start2 := time.Now()
	client2 := &http.Client{}
	req2, _ := http.NewRequest("GET", page2, nil)
	// Serialize the wire
	opentracing.GlobalTracer().Inject(
		spanPage1.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req2.Header))
	resp2, err := client2.Do(req2)

	defer resp2.Body.Close()
	end2 := time.Now()
	elapsed2 := end2.Sub(start2)
	if err != nil {
		c.JSON(resp2.StatusCode, gin.H{
			"message": err,
		})
		return
	}
	message += "\nConnecting to " + page2 + " spent " + elapsed2.Truncate(time.Millisecond).String()

	// End span - page2
	spanPage2.Finish()
	//spanPage2.End()

	// Return results
	c.String(200, message)

	// End span - global
	spanGlobal.Finish()
	//spanGlobal.End()
}
