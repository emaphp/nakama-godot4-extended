# Nakama Godot 4 Extended #

## In construction!!! ##

## About ##

This is a reimplementation of [Nakama Godot Demo](https://github.com/heroiclabs/nakama-godot-demo/) but for Godot 4. The original project has been modified to also include module development using the Go programming language.

## Requirements ##

Your system needs these tools installed:

 - The Go Programming Language (>=1.24)
 - Docker
 - Docker Compose Plugin

## Setup ##

### Build the server modules ###

This steps builds the server modules. Compiled modules are stored in `nakama-server/nakama/modules`.

```
 cd nakama-server/
 make build
```

### Run the server  ###

This step initializes a PostgreSQL container and then runs Nakama. Nakama configuration can be modified in `nakama-server/nakama/config.yml`.

```
 cd nakama-server/
 docker compose up -d postgres nakama
```

### Import Godot project ###

[TODO]

## Licenses ##

This project is dual-licensed:

- The source code is available under the Apache 2.0 license.
- Art assets (images, audio files) are [CC-By-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/). You should attribute them to Heroic Labs (https://heroiclabs.com/).
