module github.com/dialangproject/content

go 1.25.0

replace github.com/dialangproject/common => ../common

require (
	github.com/dialangproject/common v0.0.0-00010101000000-000000000000
	github.com/magiconair/properties v1.8.10
	github.com/pariz/gountries v0.1.6
	golang.org/x/text v0.29.0
)

require (
	github.com/lib/pq v1.10.9 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
