---
name: ue-test-patterns
description: Patterns for writing and running UE 5.7 automation tests
---
## UE 5.7 Test Execution
- Headless: `UnrealEditor-Cmd <project> -ExecCmds="Automation RunTests <filter>;Quit" -nullrhi -unattended -nosplash -ABSLOG=<path>`
- Parse "Test Completed." lines: `Result={Success|Fail}` and `Path={}`
- Failure details between `BeginEvents` / `EndEvents` markers
- Exit codes unreliable on macOS — always parse logs
- Add `sleep 1` after test for log flush on macOS

## Test Naming
- `ProjectName.Unit.*` — headless unit tests
- `ProjectName.Visual.*` — GPU visual tests
- `ProjectName.Integration.*` — integration tests
