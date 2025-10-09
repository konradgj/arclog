# arclog

`arclog` is a small command-line tool for managing arcdps combat logs and uploading them to [dps.report](https://dps.report).

It provides a local SQLite-backed database of found logs, a watcher to monitor your log directory, and an uploader that posts logs to dps.report while handling rate limits and retry/state tracking.

## Features

- Store and query logs in a local DB.
- Add files or directories of logs to the DB.
- Watch a log directory and automatically add & upload new logs.
- Upload pending logs to dps.report with per-file status tracking.
- List uploads by various filters.

## Quick install (from source)

This project is written in Go. You can build it locally with the Go toolchain.

1. Clone the repository:

```bash
git clone https://github.com/konradgj/arclog.git
cd arclog
```

2. Build the binary (recommended):

```bash
go build -o arclog ./
```

3. Or install into your Go bin dir:

```bash
go install ./
```

Make sure your $GOBIN or $GOPATH/bin is on your PATH if you used `go install`.

## Usage

The CLI binary exposes a few top-level commands: `config`, `log`, `watch`, and `upload`.

General form:

```bash
arclog [--debug] <command> [subcommand] [flags]
```

Global flags:
- `--debug` — enable development logging output

### Commands and examples

#### config
- `arclog config show` — print current configuration
- `arclog config set` — set saved config values
- Flags available:
  - `-p` set log path (default is arcdps default)
  - `-t` set dps.report usertoken

#### log
- `arclog log add /path/to/file.evtx /path/to/dir` — add files/dirs to the DB (supports multiple paths)
- `arclog log list -s pending` 
- Flags available:
  - `-s` list logs filtered by upload status (pending, uploading, uploaded, failed, skipped)
  - `-p` list logs filtered by relative path

#### watch
- `arclog watch` — watch the configured log directory and process new files
- Flags available: 
  - `-a` to upload anonymously
  - `-d` to enable detailed WvW logs

#### upload
- `arclog upload` — upload all pending logs
- `arclog upload -p /path/to/logs` — upload logs found at the given paths
- `arclog upload -s failed` — attempt uploads for logs with a specific status
- `arclog upload -w` — monitor log directory and upload as logs are created
- Flags available: 
  - `-a` to upload anonymously
  - `-d` to enable detailed WvW logs

For help on any command run:

```bash
arclog <command> -h
```

## Configuration

The CLI stores a small config (app dir `arclog`) which currently supports:

- logpath — the path to watch for logs
- usertoken — an optional user token used when uploading

Set them with `arclog config set` and view with `arclog config show`.

## Roadmap (short)

- Improve log filtering (date ranges)
- Add exclude rules (e.g. skip WvW logs)
- Add richer parsing for WvW live info
