# golang-template

[![codecov](https://codecov.io/gh/yiranzai/golang-template/branch/master/graph/badge.svg)](https://codecov.io/gh/yiranzai/golang-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/golang-template)](https://goreportcard.com/report/github.com/yiranzai/golang-template)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/golang-template/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/golang-template?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/golang-template/badges/users.svg)](https://www.codetriage.com/yiranzai/golang-template)
[![Release](https://img.shields.io/github/release/yiranzai/golang-template.svg?style=flat-square)](https://github.com/yiranzai/golang-template/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-template.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-template?ref=badge_shield)

Golang project general template for me.

## 目录

______________________________________________________________________

<!--ts-->
   * [golang-template](#golang-template)
      * [目录](#目录)
      * [Setup](#setup)
         * [Setup by Command](#setup-by-command)
         * [Setup on GitHub](#setup-on-github)
      * [Usage](#usage)
         * [Install](#install)
         * [Test](#test)
         * [<a href="https://pre-commit.com/" rel="nofollow">Pre-commit</a>](#pre-commit)
      * [Github Workflows](#github-workflows)
         * [Golang Test And Coverage](#golang-test-and-coverage)
         * [<a href="https://github.com/pantheon-systems/autotag">Autotag</a>](#autotag)
         * [<a href="https://github.com/goreleaser/goreleaser-action">Goreleaser</a>](#goreleaser)
         * [<a href="https://github.com/yiranzai/github-markdown-toc">Github Markdown TOC</a>](#github-markdown-toc)
      * [License](#license)

<!-- Added by: runner, at: Sat Apr 10 09:08:32 UTC 2021 -->

<!--te-->

______________________________________________________________________

## Setup

### Setup by Command

1. `git clone https://github.com/yiranzai/golang-template your_awesome_project`
1. Replace all strings `yiranzai/golang-template` in this repository to `your_github_username/your_awesome_project`
1. Replace all strings `golang-template` in this repository to `your_awesome_project`

### Setup on GitHub

Click "Use this template" button on GitHub project page.

## Usage

Something.

### Install

```sh
go get github.com/yiranzai/golang-template
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

### [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc)

This [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc) fork for [@ekalinin](https://github.com/ekalinin)'s [Github Markdown TOC](https://github.com/ekalinin/github-markdown-toc).

I Added flags to support for more features.

See [ekalinin/github-markdown-toc#110](https://github.com/ekalinin/github-markdown-toc/issues/110) and [ekalinin/github-markdown-toc#115](https://github.com/ekalinin/github-markdown-toc/pull/115)

```ini
--all Find all Markdown files for non-hidden folders
--auto Ignore ts/te tags, Automatically at the end/head of the file
--head The TOC is generated in the header of the file, requires --auto
```

## License

This project is licensed under the MIT License.
See the [LICENSE](./LICENSE) file
for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-template.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-template?ref=badge_large)
