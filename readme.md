## Introduction

> Work In Progress!

[![Build Status](https://travis-ci.org/frozzare/alfred.svg?branch=master)](https://travis-ci.org/frozzare/alfred) [![Go Report Card](https://goreportcard.com/badge/github.com/frozzare/alfred)](https://goreportcard.com/report/github.com/frozzare/alfred)

Alfred is a development environment for applications using Docker. It's not a replacement for [Docker Compose](https://docs.docker.com/compose/).

Check out the [examples](https://github.com/frozzare/alfred/tree/master/examples).

## Setup

Before using Alfred you have to point all `*.dev` domains to localhost. 

## Proxy container

Alfred requires a proxy container like [nginx](https://github.com/jwilder/nginx-proxy) or [Caddy](https://github.com/frozzare/caddy-proxy) to proxy domains to the right container.

The first time you use Alfred to start the proxy container, the type flag is optional to switch proxy container. Default proxy container is [Caddy](https://github.com/frozzare/caddy-proxy).

The proxy containers will only bind port 80 and not 443 (yet, pull request?).

```sh
alfred proxy start [--type=nginx]
```

You can stop the proxy container with `alfred proxy stop` or delete the Docker container.

## Application container

You're application don't need a configuration file but it's recommended since the default configuration may not work for you're application. If you just have HTML site it's easy. Checkout the [HTML example](https://github.com/frozzare/alfred/tree/master/examples/html).

```json
{
    "path": "./public"
}
```

If you need some more advanced or running PHP or some other language you need to configure which image is used and maybe environment variables.

```json
{
    "image": "isotopab/php:7.0-apache",
    "env": [
        "SITEPATH=/var/www/html/public"
    ]
}
```

When you ready with your configuration you just run `alfred start` in the same directory as the `alfred.json` exists in, if no config file exists in will use the default configuration which are configured for a HTML site.

## Configuration

The real configuration used for starting a application container can be view running `alfred config`. This shows the [HTML example](https://github.com/frozzare/alfred/tree/master/examples/html) configuration. 

```json
{
  "env": [
    "VIRTUAL_HOST=html.dev",
    "VIRTUAL_PORT=2015"
  ],
  "image": "joshix/caddy",
  "host": "html.dev",
  "links": [],
  "port": 2015,
  "path": "/u/go/src/github.com/frozzare/alfred/examples/html/public:/var/www/html:ro"
}
```

All values can be modified with `alfred.json`. The host value is based on your folder if no config value exists.

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)