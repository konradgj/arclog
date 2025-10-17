# arclog

`arclog` is a small command-line tool for managing arcdps combat logs and uploading them to [dps.report](https://dps.report).

It provides a local SQLite database tracking logs, a watcher to monitor for new logs, and an uploader that posts logs to dps.report.

## Features

- Store and query logs in a local DB.
  - Add files or directories of logs to the DB.
- Watch a log directory and automatically add & upload new logs.
- Upload pending logs to dps.report with per-file status tracking.
- List uploads by various filters.

## Quick install (from source)

You can build it locally with the Go toolchain.

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
- Flags available (set):
  - `-p` set log path (default is arcdps default)
  - `-t` set dps.report usertoken

#### log
- `arclog log add /path/to/file.zevtc /path/to/dir` — add files/dirs to the DB (supports multiple paths)
- `arclog log list` 
- Flags available (list):
  - `-s` list logs filtered by upload status (pending, uploading, uploaded, failed, skipped)
  - `-p` list logs filtered by relative path
    - relative path supports wildcards: eg. deimos% or "Deimos (17154)"
  - `-d` list logs by date
  - `--from` list logs from (>=) date
  - `--to` list logs to (<=) date

#### watch
- `arclog watch` — watch the configured log directory and add them to db

#### upload
- `arclog upload` — upload all pending logs
- `arclog upload -p /path/to/logs` — upload logs found at the given paths
- `arclog upload -s failed` — attempt uploads for logs with a specific status
- `arclog upload -w` — monitor log directory and upload as logs are created
- Flags available: 
  - `-p` paths to upload from (supports multiple paths)
  - `-s` upload logs with specific status
  - `-w` watch for new files and add to upload queue
  - `-a` to upload anonymously
  - `-d` to enable detailed WvW logs

For help on any command run:

```bash
arclog <command> -h
```

## Roadmap (short)

- Add exclude rules (e.g. skip WvW logs)
- (maybe) Add parsing for WvW live info
