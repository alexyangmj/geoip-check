# geoip-check

Return geolocation info (country, city, timezone, lat, long) in bulk using MaxMind Lite DB and Gregory Oschwald's geoip2-golang library

The binary for MacOS (compiled on Sonoma 14.x) is included in this repository.

`geoip-check` is for Intel processor x86-64 based and `geoip-check-arm64` is for Apple Silicon based.

```
geoip-check v1.2
Last Update: 14 Apr 2024, Alex Yang (linkedin.com/in/4yang)

Usage for Single IP query:
   geoip-check [IPv4/v6] [Optional_Switch]

Optional_Switch for output format:
   0   Suppresss showing source IP
   I   Show only source IP, Country ISO Code
   C   Show only source IP, Country
   c   Show only source IP, City
   T   Show only source IP, Timezone
   L   Show only source IP, Latitude, Longitude
   Cc  Show only source IP, Country, City

Example:
   geoip-check 74.125.200.100
   geoip-check 2607:f8b0:4003:0c00:0000:0000:0000:006a
   geoip-check 74.125.200.101 Cc
   geoip-check 74.125.200.101 0C

Output:
   % geoip-check 74.125.200.100
   74.125.200.100,US,United States,,America/Chicago,37.751,-97.822

   % geoip-check 2607:f8b0:4003:c00::6a
   2607:f8b0:4003:c00::6a,US,United States,Tulsa,America/Chicago,36.16,-95.988

   % geoip-check 113.20.105.19 Cc
   113.20.105.19,Vietnam,Hanoi

   % geoip-check 74.125.200.102 0T
   America/Chicago

Usage for Bulk IP query:
   geoip-check [inputfile.txt] --> file extension must be .txt
   
Example:
   geoip-check input.txt
   geoip-check input.txt 0Cc
```

# Important:

`GeoLite2-City.mmdb` file must be present in your `$HOME` directory (that is: `/Users/<yourusername>` for MacOS  -OR-  `/home/<yourusername>` for most Linux distro)

For Linux's root user, your `$HOME` directory usually is `/root`

For Windows, you'll need to re-compile `geoip-check.go` file and modify the source path to windows `%UserProfile%` directory

# If you have any issue and need a little help

Please don't hesitate to DM me at **Linkedin** OR open an issue.

https://linkedin.com/in/4yang

# To contribute

Please make a PR to help improve this tool :)


