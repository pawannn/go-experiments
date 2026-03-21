package header

import "fmt"

var rn = []byte("\r\n")

var errMalfunctionedHeaderLine error = fmt.Errorf("Malfunctioned header line")
var errMalfunctionedHeaders error = fmt.Errorf("Malfunctioned headers")
