package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"strconv"
	"net/http"
	//"fmt"
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
	page := c.Query("page")
	//fmt.Println(page)
	//resp, err := http.Get(page)
	start := time.Now()
	_, err := http.Get(page)
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	// time.Sleep(time.Duration(durationInt) * time.Second)
	message := "Connecting to " + page + " spent " + elapsed.Truncate(time.Millisecond).String()
	c.String(200, message)
}

