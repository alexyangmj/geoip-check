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
    Banner := "geoip-check v1.1c\n"
    Banner = Banner + "Last Update: 12 Apr 2024, Alex Yang (linkedin.com/in/4yang)\n\n"
    Banner = Banner + "Usage for Single IP query:\n"
    Banner = Banner + "   geoip-check [IPv4/v6] [Optional_Switch]\n\n"
    Banner = Banner + "Optional_Switch for output format:\n"
    Banner = Banner + "0   Suppresss showing source IP\n"
    Banner = Banner + "I   Show only source IP, Country ISO Code\n"
    Banner = Banner + "C   Show only source IP, Country\n"
    Banner = Banner + "c   Show only source IP, City\n"
    Banner = Banner + "T   Show only source IP, Timezone\n"
    Banner = Banner + "L   Show only source IP, Latitude, Longitude\n"
    Banner = Banner + "Cc  Show only source IP, Country, City\n\n"
    Banner = Banner + "Example:\n"
    Banner = Banner + "   geoip-check 74.125.200.100\n"
    Banner = Banner + "   geoip-check 2607:f8b0:4003:0c00:0000:0000:0000:006a\n"
    Banner = Banner + "   geoip-check 74.125.200.101 Cc\n"
    Banner = Banner + "   geoip-check 74.125.200.101 0C\n\n"
    Banner = Banner + "Usage for Bulk IP query:\n"
    Banner = Banner + "   geoip-check [inputfile.txt] --> file extension must be .txt\n\n"
    Banner = Banner + "Example:\n"
    Banner = Banner + "   geoip-check input.txt\n"
    Banner = Banner + "   geoip-check input.txt 0Cc\n\n"
    
    var inputIP     string
    var isFile      bool   = false
    var Switch      string = "NIL"

    defer func() {
        if r := recover(); r != nil {
            fmt.Println(Banner) 
        }
    }()

    if len(os.Args) == 1  { 
        fmt.Println(Banner)
        return
    } 
    
    if len(os.Args) == 2 { inputIP = os.Args[1] } 

    if len(os.Args) == 3 { 
        inputIP = os.Args[1]
        Switch  = os.Args[2] 
    }
    
    if !strings.Contains(inputIP, ".") && !strings.Contains(inputIP, ":") {
        fmt.Println(Banner)
        return
    }
    
    if strings.Count(inputIP, ".") != 3 && strings.Count(inputIP, ":") < 4 {
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
            if strings.Count(txtlines, ".") != 3 && strings.Count(txtlines, ":") < 4 {
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
            switch Switch {
            case "0":
                fmt.Printf("%v,", record.Country.IsoCode)
                fmt.Printf("%v,", record.Country.Names["en"])
                fmt.Printf("%v,", record.City.Names["en"])
                fmt.Printf("%v,", record.Location.TimeZone)
                fmt.Printf("%v,", record.Location.Latitude)
                fmt.Printf("%v\n", record.Location.Longitude)    
            case "I":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v\n", record.Country.IsoCode)
            case "C":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v\n", record.Country.Names["en"])
            case "c":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v\n", record.City.Names["en"])
            case "T":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v\n", record.Location.TimeZone)
            case "L":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v,", record.Location.Latitude)
                fmt.Printf("%v\n", record.Location.Longitude)
            case "Cc":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v,", record.Country.Names["en"])
                fmt.Printf("%v\n", record.City.Names["en"])
            case "0C":
                fmt.Printf("%v\n", record.Country.Names["en"])
            case "0I":
                fmt.Printf("%v\n", record.Country.IsoCode)                
            case "0c":
                fmt.Printf("%v\n", record.City.Names["en"])
            case "0T":
                fmt.Printf("%v\n", record.Location.TimeZone)
            case "0L":
                fmt.Printf("%v,", record.Location.Latitude)
                fmt.Printf("%v\n", record.Location.Longitude)
            case "0Cc":
                fmt.Printf("%v,", record.Country.IsoCode)
                fmt.Printf("%v\n", record.City.Names["en"])
            case "NIL":
                fmt.Printf("%v,", txtlines)
                fmt.Printf("%v,", record.Country.IsoCode)
                fmt.Printf("%v,", record.Country.Names["en"])
                fmt.Printf("%v,", record.City.Names["en"])
                fmt.Printf("%v,", record.Location.TimeZone)
                fmt.Printf("%v,", record.Location.Latitude)
                fmt.Printf("%v\n", record.Location.Longitude)
            default:
                fmt.Println("Unrecognized switch!")
            }
            errline++
        }
        file.Close()
    } else {
        ip := net.ParseIP(inputIP)
        record, err := db.City(ip)
        if err != nil {
            log.Fatal(err)
        }

        switch Switch {
        case "0":
            fmt.Printf("%v,", record.Country.IsoCode)
            fmt.Printf("%v,", record.Country.Names["en"])
            fmt.Printf("%v,", record.City.Names["en"])
            fmt.Printf("%v,", record.Location.TimeZone)
            fmt.Printf("%v,", record.Location.Latitude)
            fmt.Printf("%v\n", record.Location.Longitude)    
        case "I":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v\n", record.Country.IsoCode)
        case "C":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v\n", record.Country.Names["en"])
        case "c":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v\n", record.City.Names["en"])
        case "T":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v\n", record.Location.TimeZone)
        case "L":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v,", record.Location.Latitude)
            fmt.Printf("%v\n", record.Location.Longitude)
        case "Cc":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v,", record.Country.Names["en"])
            fmt.Printf("%v\n", record.City.Names["en"])
        case "0C":
            fmt.Printf("%v\n", record.Country.Names["en"])
        case "0I":
            fmt.Printf("%v\n", record.Country.IsoCode)
        case "0c":
            fmt.Printf("%v\n", record.City.Names["en"])
        case "0T":
            fmt.Printf("%v\n", record.Location.TimeZone)
        case "0L":
            fmt.Printf("%v,", record.Location.Latitude)
            fmt.Printf("%v\n", record.Location.Longitude)
        case "0Cc":
            fmt.Printf("%v,", record.Country.Names["en"])
            fmt.Printf("%v\n", record.City.Names["en"])
        case "NIL":
            fmt.Printf("%v,", inputIP)
            fmt.Printf("%v,", record.Country.IsoCode)
            fmt.Printf("%v,", record.Country.Names["en"])
            fmt.Printf("%v,", record.City.Names["en"])
            fmt.Printf("%v,", record.Location.TimeZone)
            fmt.Printf("%v,", record.Location.Latitude)
            fmt.Printf("%v\n", record.Location.Longitude)
        default:
            fmt.Println("Unrecognized switch!")
        }
    }
}

func IsIpv4Net(host string) bool {
   return net.ParseIP(host) != nil
}

func IsIpv6Net(host string) bool {
   return net.ParseIP(host) != nil
}