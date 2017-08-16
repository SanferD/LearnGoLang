// convert between pound and kilogram
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

func PToK(p Pound) Kilogram { return Kilogram(p*0.453592) }
func KToP(k Kilogram) Pound { return Pound(k*2.20462) }