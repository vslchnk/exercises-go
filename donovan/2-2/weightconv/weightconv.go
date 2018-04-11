package weightconv

import "fmt"

type KG float64
type Pound float64

func (k KG) String() string    { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }
