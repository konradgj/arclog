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
- [ ] http client
  - [ ] handle ratelimit
  - [x] post to dps.report
  - [x] handle post options 
- [ ] upload command
  - [ ] upload -dir (upload from dir maybe recursive?)
  - [x] upload -watch (move watch command here)
- [ ] analisis commands (list?)
  - [ ] list uploads per status
  - [ ] list uploaded with link to follow
  - [ ] list if log succes or not?
- [ ] wvw live info
  - [ ] would need to parse evtc file
  - [ ] enemy count?