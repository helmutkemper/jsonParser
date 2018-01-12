package jsonParser

import (
  "encoding/json"
  "fmt"
)

func ExampleParserInterface() {
  var punchInData interface{}
  s := `{
          "data": "01092017",
          "hora": "0900",
          "longitude": "-48.465146",
          "latitude": "-27.428847",
          "provider": "gps",
          "offline": "false"
        }`

  if err := json.Unmarshal([]byte(s), &punchInData); err != nil {
    panic(err)
  }

  fmt.Printf( "data: %v\n", Parser( punchInData, "data", int(0) ) )
  fmt.Printf( "hora: %v\n", Parser( punchInData, "hora", int(0) ) )
  fmt.Printf( "offline: %v\n", Parser( punchInData, "offline", false ) )
  fmt.Printf( "longitude: %v\n", Parser( punchInData, "longitude", float64(0) ) )
  fmt.Printf( "latitude: %v\n", Parser( punchInData, "latitude", float64(0) ) )
  fmt.Printf( "provider: %v\n\n\n", Parser( punchInData, "provider", "") )

  // Output:
  // data: 1092017
  // hora: 900
  // offline: false
  // longitude: -48.465146
  // latitude: -27.428847
  // provider: gps

}
