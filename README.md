# zcs

`zcs` is a CLI tool and Golang client to fetch real time data from ZCS Azzurro inverters. `zcs` fetches inverter data from the ZCS cloud platform at https://www.zcsazzurroportal.com.

To get your API credentials, install the [Azzurro Operators](https://apps.apple.com/it/app/azzurro-operators/id6447290301?l=en-GB) app, and create a support request.

## Quick Start

```shell
$ zcs azzurro fetch-realtime \
    --client-id "JohnDoe" \
    --api-auth "Zcs a0CHc2e9Y88h" \
    --thing-id "ZMG8982372HCCB76"

Power importing: 900.00
Power exporting: 0.00
Power generating: 0.00
Power consuming: 900.00
Battery level: 20%
Last update: 2024-10-22 20:12:27 +0000 UTC
```

You can also set `client-id` and `api-auth` using environment variables:

```shell
export ZCS_CLIENT_ID="JohnDoe"
export ZCS_CLIENT_API_AUTH="Zcs a0CHc2e9Y88h"

zcs azzurro fetch-realtime \
    --thing-id "ZMG8982372HCCB76"
```

Or the config file `~/.zcs`:

```shell
$ cat ~/.zcs

client_id: JohnDoe
api_auth: "Zcs a0CHc2e9Y88h"

zcs azzurro fetch-realtime \
    --thing-id ZMG8982372HCCB76
```
