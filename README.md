# service

**`Highly opinionated Go service.`** 🐻

## What

This is a complete over-engineered Golang service template, which strongly relies on the building blocks of my (also) [**`highly opinionated Go backend kit`**](https://github.com/neoxelox/kit).

> This template is already being used by some companies on production, handling non-trivial volumes of traffic. However, use it at your own "risk", this is a personal project and support is not the main intention.

## Features

This repository packs so many features that I am unable to sit down and list them all : ). Feel free to take a look at the code.

## Before Getting Started

### Development

Make a copy of `envs/dev/.env.example` to `envs/dev/.env` and fill the variables. `envs/dev/.env` is an ignored file so you won't have the chance to commit it with potential production values 😉.

Also, to setup the environment, follow these steps:

1. Install dependencies: `pip install -r scripts/requirements.txt`
2. Install tools `inv tool.install --include "dev*"`

Run `inv help` for further commands and `inv <command> --help` for their usage.

### Production

To make this environment work you will have to make the following changes:

1. Make a copy of `envs/prod/.env.example` to `envs/prod/.env` (which is ignored) and fill the variables.

2. Create 2 [Cloudflare Tunnels](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/), one for exterior user accessible services `ext`, and the other for interior company-only accessible services and tools `int`.

3. Download both Tunnels certificates to `envs/prod/certs/` (which is ignored).

4. Fill both Tunnels configuration files `envs/prod/cloudflared-ext.yaml & cloudflared-int.yaml` with the Tunnel ID, credentials path and hostnames.

5. Change the `scripts/tasks.py` `build & deploy` tasks to tag and push your service images to your registry of preference (it should be a private registry...).

6) Finally, fill the `envs/prod/docker-compose.yaml` so the service containers point to your own registry images.

## Contribute

Feel free to contribute to this project : ) .

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT) - read the [LICENSE](LICENSE) file for details.
