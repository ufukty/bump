# bump

![Social card](assets/github-social-preview.png)

Smallest CLI utility that adds the next semantic version tag to the current commit in active branch in a git folder.

## Install

```sh
go install github.com/ufukty/bump
```

If you don't have the go compiler, installing it is too easy to skip [go.dev/dl](https://go.dev/dl)

## Usage

```sh
$ cd my-beautiful-git-project
# either of
$ bump fix
$ bump minor
$ bump major
```

## Suggestions

Don't forget pushing tags to GitHub next time you push.

```sh
git push origin --tags
```

## Contributions

Issues, PRs and Discussions are open and welcome.

## License

MIT
