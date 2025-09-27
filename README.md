# arclog
arc-dps log uploader

## Roadmap
- [x] db uploads table
  - [x] columns: id, path, url, status, status_reason, statecode, created_at, updated_at
  - [x] up and down migrations
- [x] watcher
  - [x] write to uploads on file
  - [x] workers
  - [x] concurrent workers
- [x] http client
  - [x] handle ratelimit
  - [x] post to dps.report
  - [x] handle post options 
- [ ] upload command
  - [x] no flag - uploads all pending logs
  - [x] upload -watch (move watch command here)
  - [ ] -dir (upload only logs in dir)
  - [ ] -file (upload only given log)
- [ ] log command
  - [ ] add (add logs to db)
    - [ ] no flag - adds all logs
    - [ ] -file (add only file)
    - [ ] -dir (add only from dir)
  - [ ] list (list logs in db)
    - [ ] -f - filter opts (pending, uploaded, failed, skipped)
  - [ ] clear (deactivate logs active = 0)
    - [ ] -d (hard delete)
- [ ] add a way to exclude logs (example: dont post wvw logs)
- [ ] wvw live info
  - [ ] would need to parse evtc file
  - [ ] enemy count?