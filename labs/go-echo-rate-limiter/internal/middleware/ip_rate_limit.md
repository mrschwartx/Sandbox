# IP-Based Rate Limit

IP-Base Rate Limit is a technique to restrict the humber of requests coming from a single IP address within a certain period. It's suitable for:

- Public APIs without login (anonymous access)
- Preventing light DDoS attacks
- Avoiding abuse from scrapping bots

## Example Rule:

Maximum 100 requests per minute per IP

With this rule, anonymous users cannot send more than 100 requests per minute - exess requests will be rejected with a 429 Too Many Requests status.

