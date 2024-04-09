# geoip-check
Return geolocation info (country, city, timezone, lat, long) in bulk using MaxMind Lite DB and Gregory Oschwald's geoip2-golang library

The binary for MacOS (compiled on Sonoma 14.x) is included in this repository.

Usage: ./geoip-check [IPv4/v6]

Usage: ./geoip-check inputfile.txt (file extension must be .txt)

Important:
GeoLite2-City.mmdb file must be present in your $HOME directory (that is: /Users/<yourusername> for MacOS  -OR-  /home/<yourusername> for most Linux distro)
For Linux's root user, your $HOME directory usually is /root
For Windows, you'll need to re-compile geoip-check.go file and modify the source path to windows %%UserProfile% directory

If you have any issue and need a little help, don't hesitate to DM me at Linkedin OR open an issue.
https://linkedin.com/in/4yang
