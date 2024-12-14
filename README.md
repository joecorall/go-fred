# fred_go_toolkit
FRED GoLang Toolkit for Federal Reserve Economic Data

![build status](https://github.com/joecorall/go-fred/actions/workflows/lint-test.yml/badge.svg) [![GoDoc](https://godoc.org/github.com/nswekosk/fred_go_toolkit?status.svg)](https://godoc.org/github.com/nswekosk/fred_go_toolkit)

This is a GoLang toolkit for working with the Federal Reserve Economic Data. FRED contains frequently updated US macro and regional economic time series at annual, quarterly, monthly, weekly, and daily frequencies. FRED aggregates economic data from a variety of sources - most of which are US government agencies. The economic time series in FRED contains observation or measurement periods associated with data values. For instance, the US unemployment rate for the month of January, 1990 was 5.4 percent and for the month of January, 2000 was 4.0 percent.

## Installation

   Setup Go environment.
   Run 'go get github.com/nswekosk/fred_go_toolkit'

## Get a FRED API key

Sign up for a Fred API key: [http://api.stlouisfed.org/api_key.html](http://api.stlouisfed.org/api_key.html)

## Usage Example

```go
import (
    fred "github.com/joecorall/go-fred"
)

fredConfig := fred.FredConfig{
    APIKey: 'api-key',
    // File types and log files are optional.
    // You can use local constants fred.FileTypeJSON or fred.FileTypeXML.
    // If no file type is specified, the default type is fred.FileTypeXML.
    FileType: fred.FileTypeJSON,
    LogFile: "log.log",
}

client, err := fred.CreateFredClient(fredConfig)
if err != nil {
  // handle error
}

// params will be your input parameters for the API call.
// The value is of type interface{} because it can take an int, string, or boolean.
// All types will be converted to string format for the API request.
// In this example, that will send a request to the url: https://api.stlouisfed.org/fred/category?category_id=125&api_key=apiKey&file_type=json
params := make(map[string]interface{})
params["category_id"] = 125
fc, err := client.GetRecord(params)
if err != nil {
  // handle error
}
```
