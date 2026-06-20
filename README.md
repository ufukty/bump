# bump

![Social card](assets/github-social-preview.png)

Bump is the smallest, simplest CLI utility for bumping a project's version number.

Bump accepts the `major`, `minor`, and `patch` commands — named after the version components defined in [Semantic Versioning 2.0.0](https://semver.org) — along with an additional `alpha` command. The command determines which component of the latest version tag is incremented.

Bump starts by running `git describe` in the background to learn the latest version, increments the appropriate component, and calls `git tag` to assign the new version number.

## Install

Use the Go compiler ([go.dev/dl](https://go.dev/dl)) to download and install Bump:

```sh
go install github.com/ufukty/bump@latest
```

Check that everything works:

```sh
bump help
```

## Issuing version tags

```sh
$ cd my-beautiful-git-project
```

### Release versions

Bump issues release version tags in the form $v\text{MAJOR}.\text{MINOR}.\text{PATCH}$. Issuing the next version is straightforward once you've decided which component to increment. Landing on $v1.0.0$ requires the `--force` flag; see [Accidental Backwards-Compatibility Promise Prevention](#accidental-backwards-compatibility-promise-prevention).

| Command      | $V_i$    | $V_{i+1}$ | When to use     |
| ------------ | -------- | --------- | --------------- |
| `bump major` | $v1.2.3$ | $v2.0.0$  | Breaking change |
| `bump minor` | $v1.2.3$ | $v1.3.0$  | New feature     |
| `bump patch` | $v1.2.3$ | $v1.2.4$  | Bug fix         |

### Pre-release versions

Pre-release tags are issued using alpha-tracks. An alpha-track is a series of $v\text{TARGET-alpha}.N$-style tags. The track's target is set at initiation and follows the same rules as release version tags. The numeric identifier starts at 1 and is incremented on each iteration of the track. Finalizing an alpha-track issues the final $v\text{TARGET}$ tag. Targeting or landing on $v1.0.0$ requires the `--force` flag; see [Accidental Backwards-Compatibility Promise Prevention](#accidental-backwards-compatibility-promise-prevention).

#### Initiating a new alpha-track

Given a component to bump (e.g., `major`, `minor`, or `patch`), Bump calculates the target version and issues the first alpha tag. You must be on a release version to initiate a new alpha-track.

| Command            | $V_i$    | $V_{i+1}$               |
| ------------------ | -------- | ----------------------- |
| `bump alpha major` | $v1.2.3$ | $v2.0.0\text{-alpha.}1$ |
| `bump alpha minor` | $v1.2.3$ | $v1.3.0\text{-alpha.}1$ |
| `bump alpha patch` | $v1.2.3$ | $v1.2.4\text{-alpha.}1$ |

#### Iterating the current alpha-track

Without an additional argument, the alpha command iterates the track and issues the next pre-release tag.

| Command      | $V_i$                   | $V_{i+1}$               |
| ------------ | ----------------------- | ----------------------- |
| `bump alpha` | $v2.0.0\text{-alpha.}5$ | $v2.0.0\text{-alpha.}6$ |
| `bump alpha` | $v1.3.0\text{-alpha.}5$ | $v1.3.0\text{-alpha.}6$ |
| `bump alpha` | $v1.2.4\text{-alpha.}5$ | $v1.2.4\text{-alpha.}6$ |

#### Finalizing the current alpha-track

Finalizing issues a version tag carrying the alpha-track's target, this time without the alpha suffix.

| Command               | $V_i$                   | $V_{i+1}$ |
| --------------------- | ----------------------- | --------- |
| `bump alpha finalize` | $v2.0.0\text{-alpha.}6$ | $v2.0.0$  |
| `bump alpha finalize` | $v1.3.0\text{-alpha.}6$ | $v1.3.0$  |
| `bump alpha finalize` | $v1.2.4\text{-alpha.}6$ | $v1.2.4$  |

## Accidental backwards-compatibility promise prevention

In many communities, v1 is expected to mark the start of a project's backwards-compatibility guarantees. [SemVer 2.0.0](https://semver.org) is a specification that makes a similar expectation explicit; [0ver](https://0ver.org) is another convention, one that argues against ever reaching v1 at all. To keep developers from issuing these commands by mistake and signaling such a promise to their communities unintentionally Bump refuses to issue (or even target) the $v1.0.0$ tag. You must pass the `--force` flag to proceed. For example:

```sh
bump major --force
bump alpha major --force
bump alpha finalize --force
```

## Publishing tags

Don't forget to push your tags the next time you push; Git doesn't include them by default.

```sh
git push origin --tags
```

## Contributions

Issues, PRs, and Discussions are open and welcome.

## License

MIT
