# nightfall_cli
![Nightfall_CLI](https://cdn.nightfall.ai/nightfall-dark-logo-tm.png "nightfall_cli")
### NightfallCLI - a code review tool that protects you from committing sensitive information

NightfallCLI scans your code for secrets or sensitive information. It’s intended to be used as a part of your CI to simplify the development process, improve your 
security, and ensure you never accidentally leak secrets or other sensitive information via an accidental commit.

## Supported Services
### GithubActions
[NightfallDLP](https://github.com/nightfallai/nightfall_dlp_action)

## Additional Configuration
Aside from which detectors to scan on, you can add additional fields to your config, `./nightfall/config.json`, to ignore tokens and files as well as specify which files to exclusively scan on.
### Token Exclusion
To ignore specific tokens, you can add the `tokenExclusionList` field to your config. The `tokenExclusionList` is a list of strings, and it mutes findings that match any of the given regex patterns.

Here's an example use case:

```tokenExclusionList: ["NF-gGpblN9cXW2ENcDBapUNaw3bPZMgcABs", "^127\\."]```

In the example above, findings with the API token `NF-gGpblN9cXW2ENcDBapUNaw3bPZMgcABs` as well as local IP addresses starting with `127.` would not be reported. For more information on how we match tokens, take a look at [regexp](https://golang.org/pkg/regexp/).
### File Inclusion/Exclusion
To omit files from being scanned, you can add the `fileExclusionList` field to your config. In addition, to only scan on certain files, add the `fileInclusionList` to the config.

Here's an example use case:
```
    fileExclusionList: ["*/tests/*"],
    fileInclusionList: ["*.go", "*.json"]
```
In the example, we are ignoring all file paths with a `tests` subdirectory, and only scanning on `go` and `json` files.
Note: we are using [gobwas/glob](https://github.com/gobwas/glob) to match file path patterns. Unlike the token regex matching, file paths must be completely matched by the given pattern. e.g. If `tests` is a subdirectory, it will not be matched by `tests/*`, which is only a partial match.
