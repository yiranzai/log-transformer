# log-transformer

[![codecov](https://codecov.io/gh/yiranzai/log-transformer/branch/master/graph/badge.svg)](https://codecov.io/gh/yiranzai/log-transformer)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/log-transformer)](https://goreportcard.com/report/github.com/yiranzai/log-transformer)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/log-transformer/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/log-transformer?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/log-transformer/badges/users.svg)](https://www.codetriage.com/yiranzai/log-transformer)
[![Release](https://img.shields.io/github/release/yiranzai/log-transformer.svg?style=flat-square)](https://github.com/yiranzai/log-transformer/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Flog-transformer.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Flog-transformer?ref=badge_shield)

Parse Log Write And Transform To Other

## 目录

______________________________________________________________________

<!--ts-->

- [log-transformer](#log-transformer)
  - [目录](#%E7%9B%AE%E5%BD%95)
  - [Usage](#usage)
    - [Install](#install)
    - [Run](#run)
    - [Test](#test)
    - [<a href="https://pre-commit.com/" rel="nofollow">Pre-commit</a>](#pre-commit)
  - [Github Workflows](#github-workflows)
    - [Golang Test And Coverage](#golang-test-and-coverage)
    - [<a href="https://github.com/pantheon-systems/autotag">Autotag</a>](#autotag)
    - [<a href="https://github.com/goreleaser/goreleaser-action">Goreleaser</a>](#goreleaser)
    - [<a href="https://github.com/yiranzai/github-markdown-toc">Github Markdown TOC</a>](#github-markdown-toc)
  - [License](#license)

<!-- Added by: runner, at: Sat Feb 12 07:49:54 UTC 2022 -->

<!--te-->

______________________________________________________________________

## Usage

Something.

### Install

```shell
go get github.com/yiranzai/log-transformer
```

### Run

```shell
$ cp conf/config.example.yaml conf/config.yaml

$ log-transformer
or
$ log-transformer -f conf/config.yaml
```

### Test

```sh
go get gotest.tools/v3
```

### [Pre-commit](https://pre-commit.com/)

check or fix code style.

see [.pre-commit-config.yaml](.pre-commit-config.yaml)

e.g:

- go fmt
- golines
- go mod tiny
- go vet

Install the [pre-commit](https://pre-commit.com/)

```sh
pip install pre-commit
pre-commit install
vim .pre-commit-config.yaml
```

## Github Workflows

This repo used some workflows

### Golang Test And Coverage

Golang Test And Coverage upload to [Codecov](https://codecov.io)

### [Autotag](https://github.com/pantheon-systems/autotag)

Automatically increment version tags to a git repo based on commit messages.

### [Goreleaser](https://github.com/goreleaser/goreleaser-action)

GitHub Action for GoReleaser

#### [ghaction-import-gpg](https://github.com/crazy-max/ghaction-import-gpg)

GitHub Action to easily import a GPG key.

[New Repository secret](https://github.com/yiranzai/golang-project-template/settings/secrets/actions/new)

add `YOUR_PRIVATE_KEY` and `PASSPHRASE` secrets.

### [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc)

This [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc) fork
for [@ekalinin](https://github.com/ekalinin)'s [Github Markdown TOC](https://github.com/ekalinin/github-markdown-toc).

I Added flags to support for more features.

See [ekalinin/github-markdown-toc#110](https://github.com/ekalinin/github-markdown-toc/issues/110)
and [ekalinin/github-markdown-toc#115](https://github.com/ekalinin/github-markdown-toc/pull/115)

```ini
--all Find all Markdown files for non-hidden folders
--auto Ignore ts/te tags, Automatically at the end/head of the file
--head The TOC is generated in the header of the file, requires --auto
```

## License

This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Flog-transformer.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Flog-transformer?ref=badge_large)
