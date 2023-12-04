# go-ratelimit
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go RateLimit is a package that allows you to rate limit request into your golang applications. This package implements the token bucket algorithm and implements it per client/ip address

## Installation
Get Go RateLimit package on your project:

```bash
go get github.com/ihsanardanto-djoin/go-ratelimit
```

## Usage
This packages provides a middleware using Popular Go Framework (Echo, Chi, Gin) which can be added as a global middleware or as a single route.

```go
// in server file or anywhere middleware should be registered
rl := goratelimit.NewRateLimiter(2, 6)         // Create a new rate limiter
e.Use(echowritelimit.RateLimitMiddleware(rl))
```

```go
// in route file or anywhere route should be registered
rl := goratelimit.NewRateLimiter(2, 6)         // Create a new rate limiter
router.Echo.GET("api/v1/posts", handler, echowritelimit.RateLimitMiddleware(rl))
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/MarketingPipeline/README-Quotes/blob/main/LICENSE) file for details.

## Contributors
<a href="https://github.com/ihsanardanto-djoin/go-ratelimit/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ihsanardanto-djoin/go-ratelimit" />
</a>
