# buffalo-xbuild

This plugin is an experimental rebuilding of the `buffalo build` command. The goal is to eventually put this into Buffalo "core", but since building binaries is so important to Buffalo, I want to make sure we have all the kinks worked out before merging it into the main repo.

## Installation

```bash
$ go get -u -v github.com/gobuffalo/buffalo-xbuild
```

## Usage

The usage of this plugin is the exact same as `buffalo build`, it is a drop in replacement. The only difference is use the `buffalo xbuild` command instead of the `buffalo build` command.

```bash
$ buffalo xbuild
```
