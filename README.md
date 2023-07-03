```bash
██████  ██████   █████  ███    ██  ██████ ██   ██      ██████ ██   ██  █████  ███    ██  ██████  ███████ ██████
██   ██ ██   ██ ██   ██ ████   ██ ██      ██   ██     ██      ██   ██ ██   ██ ████   ██ ██       ██      ██   ██
██████  ██████  ███████ ██ ██  ██ ██      ███████     ██      ███████ ███████ ██ ██  ██ ██   ███ █████   ██████
██   ██ ██   ██ ██   ██ ██  ██ ██ ██      ██   ██     ██      ██   ██ ██   ██ ██  ██ ██ ██    ██ ██      ██   ██
██████  ██   ██ ██   ██ ██   ████  ██████ ██   ██      ██████ ██   ██ ██   ██ ██   ████  ██████  ███████ ██   ██
#  
#  Version:      v1.0.1
#  Maintainer:   Adli I. Ifkar <adly.shadowbane@gmail.com>
```

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=flat-square)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/license-MIT-blue?style=flat-square&label=License)](LICENSE)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/shadowbane/branch-changer/build.yml?logo=githubactions&style=flat-square&label=Build%20Status)
![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/shadowbane/branch-changer?include_prereleases&style=flat-square&logo=git&label=Latest%20Tag)
![GitHub all releases](https://img.shields.io/github/downloads/shadowbane/branch-changer/total?style=flat-square&label=Total%20Downloads)


# Branch Changer

This is a simple script to change branch of multiple git repositories at once.

## Background
Well, I made this app 'cause I'm a little bit bored and annoyed that I need to quickly change branch of multiple git repositories at once. So, I made this script to do that.
You could easily run a bash script, but like I said, I'm bored and I want to make something useful.

## Usage

Download from [release page](https://github.com/shadowbane/branch-changer/releases/latest) and run it from console.

### Linux & Intel MacOS
```bash
$ wget -c https://github.com/shadowbane/branch-changer/releases/latest/download/branch-changer-linux-amd64 -o /usr/bin/branch-changer
$ sudo chmod +x /usr/bin/branch-changer
```

### Apple Silicon MacOS
```bash
$ wget -c https://github.com/shadowbane/branch-changer/releases/latest/download/branch-changer-darwin-arm64 -o /usr/bin/branch-changer
$ sudo chmod +x /usr/bin/branch-changer
```

### Windows
Sadly, I don't have Windows machine to test this script. But, I think it should work on Windows 10/11 with WSL2.

### Parameters
```bash
Usage of ./branch-changer:
  -failedonly
        Load failed projects only
  -force
        Force switch branch (stash changes)
  -from string
        Branch to merge from
  -projects string
        Projects to merge
  -push
        Push newly create branch to remote
  -to string
        Branch to merge to
  -workdir string
        Working directory
```
- `-failedonly` : Load failed projects only
- `-force` : Force switch branch (stash changes before changing branch)
- `-from` : Branch to merge from (or, source branch)
- `-to`: Branch to merge to (or, destination branch)
- `-projects` : Projects to merge
- `-push` : Push newly create branch to remote
- `-workdir` : Working directory. All the projects must live in this directory.

### Running failed projects only
If you're like me, and need to switch several branch at once, sometimes you will face a problem where some of the projects failed to switch branch.
The script will store all failed projects in `failed.txt` file in the working directory.
Also, you could use `-failedonly` parameter to load only failed projects.

```bash
branch-changer --workdir /path/to/working/directory --failedonly
```

## To Do
- Add tests

# License
MIT License
