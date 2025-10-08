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
- [x] upload command
  - [x] no flag - uploads all pending logs
  - [x] upload -watch (move watch command here)
  - [x] -p (upload files in given path)
  - [x] -s (upload with specific status)
- [ ] log command
  - [x] add (add logs to db)
    - [x] add files or dirs (can add multiple paths i commands)
  - [x] list (list logs in db)
    - [x] -f - filter opts (pending, uploaded, failed, skipped)
    - [ ] -d (filter by date -d date, -b before -a after)
  - [ ] clear (deactivate logs active = 0)
    - [ ] -d (hard delete)
- [ ] add a way to exclude logs (example: dont post wvw logs)
- [ ] wvw live info
  - [ ] would need to parse evtc file
  - [ ] enemy count?