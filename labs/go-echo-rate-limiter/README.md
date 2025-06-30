# Example Rate Limit

This repository demonstrates how to implement in-memory rate limiting in Echo framework:

- ✅ IP-based rate limiting
- ✅ User-based rate limiting using JWT
- ✅ Custom rate limit per endpoint

> No Redis required. Simple and educational implementation.

## Basic Theory

**Rate limiting** is a technique used to control the number of requests a client can make to a server in a given time window. It protects your application from abuse, DDoS attacks, and accidental overloads.

There are several common strategies:

### 🔹 IP-Based Rate Limiting

Limits the number of requests coming from a single IP address.
Useful for public APIs and anonymous traffic.

**Example Rule:**
`Allow max 100 requests per IP per minute`

### 🔹 User-Based Rate Limiting

Limits requests based on authenticated user identity, usually extracted from a JWT token.
Enables fine-grained control per user, especially useful in multi-tenant systems or subscription-based APIs.

**Example Rule:**
`Free users: 100 requests/day`
`Premium users: 10,000 requests/day`

### 🔹 Endpoint-Based Rate Limiting

Applies different limits to different API routes.
Useful when some endpoints are more sensitive or resource-intensive (e.g., `/login`, `/checkout`).

**Example Rule:**
`/login → max 5 requests/min`
`/products → max 200 requests/min`
---

## In-Memory Storage

In this example, all rate limits are stored in memory using Go’s native data structures (`map`, `sync.Map`, `time.Ticker`, etc).
This approach is:

- ✅ Simple
- ✅ Fast
- ⚠️ Not suitable for horizontally scaled environments (use Redis in that case)

> Case go to `./internal/middleware/`

## References
