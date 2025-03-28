# jarate - a system usage daemon-viewer-thingamajig

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
    "overall": 10.1231,
    "per_core": [
      10.192,
      14.832,
      83.1928392,
      19.41203122
    ],
    "freq": 2600 // in MHz
  },
  "mem": { // in bytes
    "used": 7000000000, // 7_000_000_000 (7GB)
    "total": 16000000000
  }
}
```

## More customizations?

`jarate --help` should be enough to answer that question.

-----

`Be polite. Be efficient. Have a plan to monitor the performance of every piece
of metal you come across.`
