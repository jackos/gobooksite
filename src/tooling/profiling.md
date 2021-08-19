
## github.com/pkg/profile
Start a cpu profile, save it to working directory
---
```go
defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
```
---
