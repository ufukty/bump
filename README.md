# bump

![Social card](assets/github-social-preview.png)

Bump is the smallest and simplest CLI utility that bumps the version number of a project.

Bump accepts `major`, `minor`, `patch` commands as those labels are described in [Semantic Versioning 2.0.0](https://semver.org) with addition of `alpha` command. Commands are used to determine which label of latest version tag will be incremented.

Bump starts with `git describe` in background to learn the latest version, increments the correct label and calls `git tag` which assigns the version number.

## Install

```sh
go install github.com/ufukty/bump@latest
```

If you don't have the go compiler, installing it is too easy to skip [go.dev/dl](https://go.dev/dl)

## Usage

```sh
$ cd my-beautiful-git-project
# either of
$ bump major
$ bump minor
$ bump patch
$ bump alpha
```

### Commands

Bump follows `MAJOR`.`MINOR`.`PATCH` label rules and issues `v1.0.0`-like tags. Also, Bump uses a 4th label called `ALPHA`. This is for distinctively tagging sequent [alpha](https://github.com/ufukty/bump/issues/1) releases. 4th label is omitted for non-alpha versions for complience with [Semantic Versioning 2.0.0](https://semver.org). Bump reads and writes `v` prefixed version tags.

| $v_i$    | command      | $v_{i+1}$  | command desc.   |
| -------- | ------------ | ---------- | --------------- |
| `v1.2.3` | `bump major` | `v2.0.0`   | breaking change |
| `v1.2.3` | `bump minor` | `v1.3.0`   | new features    |
| `v1.2.3` | `bump patch` | `v1.2.4`   | bug fixes       |
| `v1.2.3` | `bump alpha` | `v1.2.3.1` | not for use     |

### Publishing tags

Don't forget pushing tags to GitHub next time you push.

```sh
git push origin --tags
```

## Contributions

Issues, PRs and Discussions are open and welcome.

## License

MIT
