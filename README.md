# jarate - a system usage daemon-viewer-thingamajig

[![Go](https://github.com/stanekondrej/jarate/actions/workflows/go.yml/badge.svg)](https://github.com/stanekondrej/jarate/actions/workflows/go.yml)

_I'm not a script kiddie, Dad; I'm a sysadmin! Well, the difference being: one
is a job, and the other's mental sickness!_

## Why the name?

sniper tf2 jarate

## How do I use this?

There are two endpoints:

- `/` - this is the standard, _"better"_ one - it is basically the way to
estabilish a WebSocket connection with this daemon. The daemon then feeds the
client data over the connection in a certain interval (that you can configure)
- `/oneshot` - this one gives you data, only over HTTP. You don't have to
estabilish a WS connection with the server if you just want to get the data
once.

## How do I work with the data?

The response is formatted as JSON, with the following structure:

```jsonc
{
  "cpu": {
    "per_cpu": [
      14.432              // You probably only have a single CPU
    ],
  },
  "mem": {                // in bytes
    "used": 7000000000,   // 7_000_000_000 (7GB)
    "total": 16000000000
  }
}
```

## More customizations?

`jarate --help` should be enough to answer that question.

-----

`Be polite. Be efficient. Have a plan to monitor the performance of every piece
of metal you come across.`
