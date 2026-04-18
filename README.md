# ArvanCloud Status CLI

A terminal UI that polls the [ArvanCloud status page](https://status.arvancloud.com)
and shows live service health and recent incidents. Built with
[Bubble Tea v2](https://charm.land/bubbletea/v2).

```
╭──────────────────────────────────────────────────────────────────────╮
│ ArvanCloud  all systems operational                                  │
│  ● Cloud Server                                                      │
│      ● Forough (Iran-Tehran/Central 2)                               │
│      ● Simin (Iran-Tehran/Central 1)                                 │
│  ● CDN                                                               │
│      ● DNS                                                           │
│      ● CDN                                                           │
╰──────────────────────────────────────────────────────────────────────╯
╭──────────────────────────────────────────────────────────────────────╮
│ Incidents (50 shown)                                                 │
│ ▸ ● [MINOR]     Foreign Connectivity Issue   resolved · 2d ago       │
│   ● [MINOR]     DNS Resolution Issue         resolved · 3d ago       │
│   ● [SCHEDULED] LCT Network Maintenance      resolved · 1w ago       │
╰──────────────────────────────────────────────────────────────────────╯
↑/k up • ↓/j down • tab switch pane • enter detail • r refresh • q quit
```

## Requirements

- Go **1.26** or newer

## Install

```sh
go install github.com/parham-alvani/arvanstatus-cli@latest
```

…or build from source:

```sh
git clone https://github.com/parham-alvani/arvanstatus-cli
cd arvanstatus-cli
go build -o arvanstatus-cli .
./arvanstatus-cli
```

## Usage

Just run the binary — no flags, no config file:

```sh
arvanstatus-cli
```

The top pane shows the full ArvanCloud service tree colored by status. The
bottom pane lists the most recent incidents from the public StatusPal API.
Both panes refresh automatically every 60 seconds.

### Keybindings

| Key       | Action                          |
|-----------|---------------------------------|
| `↑` / `k` | move cursor up                  |
| `↓` / `j` | move cursor down                |
| `tab`     | switch focused pane             |
| `enter`   | open incident detail            |
| `esc`     | close detail                    |
| `r`       | refresh now                     |
| `?`       | toggle full help                |
| `q`, `^C` | quit                            |

### Status colors

| Color  | Meaning                       |
|--------|-------------------------------|
| green  | operational / resolved        |
| yellow | minor incident / monitoring   |
| orange | identified                    |
| red    | major incident / investigating|
| blue   | scheduled maintenance         |

## Data source

ArvanCloud's status page is powered by [StatusPal](https://www.statuspal.io).
This tool calls two unauthenticated JSON endpoints:

- `GET https://statuspal.io/api/v2/status_pages/arvancloud/summary`
- `GET https://statuspal.io/api/v2/status_pages/arvancloud/incidents`

Timestamps are returned in `Asia/Tehran` local time without offset and are
parsed accordingly.
