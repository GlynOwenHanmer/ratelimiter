package PocketMediaLimiter

// frequency represents the rate at which something repeats in the units of number-of-events per second
type frequency float64

// period returns the period, in seconds, for one cycle at a given frequency.
func (f frequency) period() float64  {
	return 1.0 / float64(f)
}
