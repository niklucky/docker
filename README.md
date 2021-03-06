# Docker Spawner 

This is a very specific package for creating and managing containers remotely over HTTP.
For security purposes, a layer was written that allows you to greatly narrow the control options.

In order not to open the Docker control to the outside and severely limit the possibilities.

## Features

* Very limited HTTP API to control docker containers: Create, Stop, List
* You can't created any container for any image. Only what set in setup
* You can get remote info about machine to consider balance loading
* System prune 

## Example

In folder `examples`

You can just run `make build` to build binary or `make` to run example in watch mode (reflex).
## Setup

Simple example is in `examples/main.go` with basic setup.
In short: 
* Init `spawner.New()` with passing basic options: `AutoDelete`, `Image` (docker image will be used to spawn)
* Set in *Environement variables* or in `.env` file `PORT`, `AUTH_KEY` to start HTTP-server

## API

### Create container based on Image

`POST /containers`

```json
{
  "start": true, // optional. By default true
  "env": ["MY_ENV=Testing", "ENV_VAR2=123"],
  "image": "postgres:13", // Optional. 1st Image from setup will be used
}
```

### List of running containers

`GET /containers`

```json
{
  "data": [{}], // List of containers created by spawner
  "info": {}, // System info will be provided if you pass ?info=true
  "error": null
}
```

### Stop container

`PUT /containers/:id/stop`

💥 Danger! `id` is optional so if you will not pass - all containers created by spawner will be stopped

### Deleting container

`DELETE /containers/:id`


## Tests

Run `make test`
## TODO

* Support for multiple nodes
* System info 

> Use at your own risk