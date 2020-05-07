package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	//"go.elastic.co/apm"
	"fmt"
	"go.elastic.co/apm/module/apmot"
	"strconv"
	"time"
)

func sleep(c *gin.Context) {
	// Retrieve the original context
	ctx := c.Request.Context()

	// Use Elastic APM library
	opentracing.SetGlobalTracer(apmot.New())

	// DEBUG: print all request headers
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}

	// Deserialize the wire, extracting context from headers
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}

	// Start span
	globalSpan, ctx := opentracing.StartSpanFromContext(
		ctx,
		"sleep",
		ext.RPCServerOption(wireContext))

	// Sleep for N seconds, send span evey second
	durationString := c.Param("duration")
	durationInt, err := strconv.Atoi(durationString)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Cannot convert to integer.",
		})
		return
	}
	for i := 0; i < durationInt; i++ {
		spanSleepSecond, _ := opentracing.StartSpanFromContext(ctx, "sleepNo"+strconv.Itoa(i+1))
		time.Sleep(time.Second)
		spanSleepSecond.Finish()
	}
	// time.Sleep(time.Duration(durationInt) * time.Second)
	message := "Slept for " + durationString + " seconds."
	c.String(200, message)

	// End span
	globalSpan.Finish()
}
