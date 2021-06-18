package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"regexp"
	"time"
)

type checkTcp struct {
	network            string
	timeout            float64
	status             string
	tls                bool
	noCheckCertificate bool
}

func dial(c checkTcp) (net.Conn, error) {
	var timeout time.Duration
	if c.timeout > 0 {
		timeout = time.Duration(c.timeout * float64(time.Second))
	} else {
		timeout = time.Duration(5 * int64(time.Second))
	}
	d := &net.Dialer{Timeout: timeout}
	if c.tls {
		return tls.DialWithDialer(d, "tcp", c.network, &tls.Config{
			InsecureSkipVerify: c.noCheckCertificate,
		})
	} else {
		return d.Dial("tcp", c.network)
	}
}

func performTCPChecks(c checkTcp, ch chan string) {

	start := time.Now()
	conn, err := dial(c)
	//conn, err := net.DialTimeout("tcp", c.network, timeout)
	elapsedSeconds := float64(time.Since(start)) / float64(time.Second)
	if err != nil {
		//fmt.Println("Error:", err)
		e := checkErr(err)
		c.status = fmt.Sprintf("CRITICAL - %s %s %f", e, c.network, elapsedSeconds)
		ch <- c.status
		//fmt.Println("Error:", c.network, c.status)
	} else {
		defer conn.Close()
		c.status = fmt.Sprintf("OK - %s %f", c.network, elapsedSeconds)
		ch <- c.status
	}
}

func checkErr(err error) string {

	if err == nil {
		return "OK"
	} else if match, _ := regexp.MatchString(".*lookup.*", err.Error()); match {
		return "Unknown Host"
	} else if match, _ := regexp.MatchString(".*refuse.*", err.Error()); match {
		return "Connection refused"
	} else if match, _ := regexp.MatchString(".*timeout.*", err.Error()); match {
		return "Connection Timeout"
	} else if match, _ := regexp.MatchString(".*authority.*", err.Error()); match {
		return "Certificate signed by unknown authority"
	} else {
		return "Unkonow error"
	}
}
