# SIAK REST

A proxy server to convert your SIAK-NG response to JSON, instead of HTML.

Docs: `/swagger/index.html`, [Global Instance](https://siak-rest.up.railway.app/swagger/index.html)

## Running

You will need [golang](https://go.dev/). There is no docker compose script (yet), so you can run it by running

```
$ go run cmd/http-server/main.go
```

## Idea

SIAK-NG identifies its user with two cookies: `siakng_cc` and `Mojavi`. These two cookies are unique to your session.
This proxy simply forwards these cookies from the `X-Siakng-Cc` and `X-Mojavi` header and forward it as a cookie to
SIAK-NG server, and once the response is gathered by the proxy, the proxy simply parses the HTML and bakes the JSON
response.
