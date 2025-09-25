# arclog
arc-dps log uploader

## Roadmap
- [ ] db uploads table
 - [ ] columns: id, path, url, status, status_reason, statecode, created_at, updated_at
 - [ ] up and down migrations
- [ ] write to uploads on file
 - [ ] workers
 - [ ] concurrent workers
 - [ ] handle ratelimit
- [ ] usertoken command
 - [ ] user fav browser to get user token and set in config
- [ ] http client
 - [ ] post to dps.report
 - [ ] handle post options
- [ ] wvw live info
  - [ ] would need to parse evtc file
  - [ ] enemy count?