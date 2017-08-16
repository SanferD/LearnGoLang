// convert between foot and meter
package heightconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

func FToM(f Foot) Meter { return Meter(f*0.3048) }
func MToF(m Meter) Foot { return Foot(m*3.28084) }
