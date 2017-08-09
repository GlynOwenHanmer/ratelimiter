package PocketMediaLimiter

import "time"

// frequency represents the rate at which something repeats in the units of number-of-events per second
type frequency float64

// period returns the period Duration for one cycle at a given frequency.
func (f frequency) period() time.Duration  {
	return time.Duration(1000000000.0 / float64(f))
}
