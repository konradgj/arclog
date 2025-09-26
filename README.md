# arclog
arc-dps log uploader

## Roadmap
- [x] db uploads table
  - [x] columns: id, path, url, status, status_reason, statecode, created_at, updated_at
  - [x] up and down migrations
- [ ] watcher
  - [x] write to uploads on file
  - [ ] workers
  - [ ] concurrent workers
- [ ] usertoken command
  - [ ] user fav browser to get user token and set in config
- [ ] http client
  - [ ] handle ratelimit
  - [ ] post to dps.report
  - [ ] handle post options
- [ ] wvw live info
  - [ ] would need to parse evtc file
  - [ ] enemy count?