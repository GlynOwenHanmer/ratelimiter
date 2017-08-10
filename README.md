# ratelimiter

ratelimiter provides a simple token-bucket based rate limiter.

The components developed in this package can be used to limit a wide variety of events such as network requests to be done, access to the disk, or to space out in time the sample readings on a process that gathers data from other systems. 

## Usage ##
To import this package into your project,
```go
import "github.com/GlynOwenHanmer/ratelimiter"
```

Create a new limiter by creating a Frequency object and configuring the limiter with the Frequency object and a maximum burst rate.
```go
// rate := ratelimiter.Frequency(5)
rate, err := ratelimiter.NewFrequency(10, time.Second * 2)
if err != nil {
	log.Fatal(err)
}
burst := uint64(3)
limiter, err := ratelimiter.NewLimiter(rate, burst)
if err != nil {
	log.Fatal(err)
	return
}
for i := 0; i < 200 ; i++ {
	if limiter.Allow() {
		// Invoke your event here.
		fmt.Println("Event invoked.")
	}
	time.Sleep(time.Millisecond * 5)
}
```
