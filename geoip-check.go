package main

import (
	"fmt"
	"log"
	"net"
    "os"
    "path/filepath"
    "bufio"
    "strings"
    "github.com/oschwald/geoip2-golang"
)

func main() {
    Banner := "geoip-check v1.0\n"
    Banner = Banner + "Last Update: 10 Apr 2024, Alex Yang (linkedin.com/in/4yang)\n\n"
    Banner = Banner + "Usage: geoip-check [IPv4/v6]\n"
    Banner = Banner + "Usage: geoip-check inputfile.txt (file extension must be .txt)\n"
    
    isFile := false
    
    defer func() {
        if r := recover(); r != nil { 
            fmt.Println(Banner) 
        }
    }()
    
    inputIP := os.Args[1]
    if !strings.Contains(inputIP, ".") && !strings.Contains(inputIP, ":") {
        fmt.Println(Banner)
        return
    }
    
    if strings.Count(inputIP, ".") != 3 && strings.Count(inputIP, ":") != 7 {
        if !strings.HasSuffix(inputIP, ".txt") { 
            fmt.Println(Banner)
            return
        } else {
            isFile = true
        }
    }
    
    homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

    filePath := filepath.Join(homeDir, "GeoLite2-City.mmdb")
    
	db, err := geoip2.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    
    errline := 1
    
    if isFile {
        file, err := os.Open(inputIP)
 
	   if err != nil {
		  fmt.Println("failed opening file: %s", err)
            return
	   }
 
	   scanner := bufio.NewScanner(file)
	   scanner.Split(bufio.ScanLines)
	   
   	    for scanner.Scan() {
            txtlines := scanner.Text()
            if len(txtlines) == 0 { continue }
            if strings.Count(txtlines, ".") != 3 && strings.Count(txtlines, ":") != 7 {
                fmt.Println("Error in line: [", errline, "] IP: [", txtlines, "] - check if IP address is in correct IPv4/v6 format!")
                errline++
                continue
            }
            ip := net.ParseIP(txtlines)
            record, err := db.City(ip)
            if err != nil {
                fmt.Println("Error in line: [", errline, "] IP: [", txtlines ,"] - check if IP address is in correct IPv4/v6 format!")
                errline++
                continue
            }
            fmt.Printf("%v,", txtlines)
            fmt.Printf("%v,", record.Country.IsoCode)
            fmt.Printf("%v,", record.Country.Names["en"])
            fmt.Printf("%v,", record.City.Names["en"])
            fmt.Printf("%v,", record.Location.TimeZone)
            fmt.Printf("%v,", record.Location.Latitude)
            fmt.Printf("%v\n", record.Location.Longitude)
            errline++
        }
        file.Close()
    } else {
        ip := net.ParseIP(inputIP)
        record, err := db.City(ip)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%v,", inputIP)
        fmt.Printf("%v,", record.Country.IsoCode)
        fmt.Printf("%v,", record.Country.Names["en"])
        fmt.Printf("%v,", record.City.Names["en"])
        fmt.Printf("%v,", record.Location.TimeZone)
        fmt.Printf("%v,", record.Location.Latitude)
        fmt.Printf("%v\n", record.Location.Longitude)
    }
}

func IsIpv4Net(host string) bool {
   return net.ParseIP(host) != nil
}

func IsIpv6Net(host string) bool {
   return net.ParseIP(host) != nil
}