# rebrandly-go-sdk

A Go SDK for the Rebrandly API


# Authenticating

Create a new API key from the Rebrandly dashboard, and export as the environmental variable `REBRANDLY_API_KEY`

```bash

export REBRANDLY_API_KEY="YOUR_API_KEY"
```

# Status

Right now only basic endpoints are baked into the SDK, but the framework is in place that makes it extremely easy to add other endpoints.

If there is an endpoint you need, just ping me (Kris Nova) or just open a PR.

# Parameters

Some of the SDK functions accept `params` which maps to the accepted to parameters in the Rebrandly documentation [here](https://developers.rebrandly.com/docs)