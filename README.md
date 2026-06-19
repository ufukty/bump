# bump

![Social card](assets/github-social-preview.png)

Bump is the smallest and simplest CLI utility that bumps the version number of a project.

Bump accepts `major`, `minor`, `patch` commands as those labels are described in [Semantic Versioning 2.0.0](https://semver.org) with addition of `alpha` command. Commands are used to determine which label of latest version tag will be incremented.

Bump starts with `git describe` in background to learn the latest version, increments the correct label and calls `git tag` which assigns the version number.

## Install

Use Go compiler ([go.dev/dl](https://go.dev/dl)) to download and install Bump:

```sh
go install github.com/ufukty/bump@latest
```

Check if everything works:

```sh
bump help
```

## Issuing version tags

```sh
$ cd my-beautiful-git-project
```

### Release versions

Bump issues release versions tags in the form of $v\text{MAJOR}.\text{MINOR}.\text{PATCH}$. Issuing the next version is easy once the label of increment is decided. Landing on $v1.0.0$ requires the `--force` flag. See [Accidental Backwards-Compatibility Promise Prevention](#accidental-backwards-compatibility-promise-prevention).

| Command      | $V_i$    | $V_{i+1}$ | When to use?    |
| ------------ | -------- | --------- | --------------- |
| `bump major` | $v1.2.3$ | $v2.0.0$  | Breaking change |
| `bump minor` | $v1.2.3$ | $v1.3.0$  | New feature     |
| `bump patch` | $v1.2.3$ | $v1.2.4$  | Bug fixe        |

### Pre-release versions

Issuing pre-release tags is done using alpha-tracks. An alpha-track is a series of $v\text{TARGET-alpha}.N$ style tags. The track target is set at initialization and follows the same guidelines as the release version tags. The numeric identifier starts at 1 and incremented at each track iteration. Finalizing an alpha-track issues the $v{TARGET}$. Landing on and targeting $v1.0.0$ requires the `--force` flag. See [Accidental Backwards-Compatibility Promise Prevention](#accidental-backwards-compatibility-promise-prevention).

#### Initiating a new alpha-track

Provingi the alpha command a label name (eg. `major`, `minor`, `patch`) Bump calculates the target version and issues a tag for the first alpha version. You must be on a release version to initiate a new alpha-track.

| Command            | $V_i$    | $V_{i+1}$               |
| ------------------ | -------- | ----------------------- |
| `bump alpha major` | $v1.2.3$ | $v2.0.0\text{-alpha.}1$ |
| `bump alpha minor` | $v1.2.3$ | $v1.3.0\text{-alpha.}1$ |
| `bump alpha patch` | $v1.2.3$ | $v1.2.4\text{-alpha.}1$ |

#### Iterating the current alpha-track

Without additional argument alpha command iterates the track and assigns a tag for the next pre-release version.

| Command      | $V_i$                   | $V_{i+1}$               |
| ------------ | ----------------------- | ----------------------- |
| `bump alpha` | $v2.0.0\text{-alpha.}5$ | $v2.0.0\text{-alpha.}6$ |
| `bump alpha` | $v1.3.0\text{-alpha.}5$ | $v1.3.0\text{-alpha.}6$ |
| `bump alpha` | $v1.2.4\text{-alpha.}5$ | $v1.2.4\text{-alpha.}6$ |

#### Finalizing the current alpha-track

Finalizing issues a version tag with the alpha-track's target, this time without the alpha suffix.

| Command               | $V_i$                   | $V_{i+1}$ |
| --------------------- | ----------------------- | --------- |
| `bump alpha finalize` | $v2.0.0\text{-alpha.}6$ | $v2.0.0$  |
| `bump alpha finalize` | $v1.3.0\text{-alpha.}6$ | $v1.3.0$  |
| `bump alpha finalize` | $v1.2.4\text{-alpha.}6$ | $v1.2.4$  |

## Accidental backwards-compatibility promise prevention

Among many communities the v1 is expected to mark the beginning of backwards-compatibility promise in a software project. [SemVer2.0.0](https://semver.org) is a specification that explicitly states similar expectation. [0ver](https://0ver.org) is another one that suggests against ever reaching to v1. To protect developers from using commands mistakenly and suggesting such promises to their communities unconsciously Bump rejects issuing (or even targeting) the $v1.0.0$ version tag. Developers need to use the `--force` flag to proceed. Such as:

```sh
bump major --force
bump alpha major --force
bump alpha finalize --force
```

### Publishing tags

Don't forget pushing tags to GitHub next time you push.

```sh
git push origin --tags
```

## Contributions

Issues, PRs and Discussions are open and welcome.

## License

MIT
