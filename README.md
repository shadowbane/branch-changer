```bash
██████  ██████   █████  ███    ██  ██████ ██   ██      ██████ ██   ██  █████  ███    ██  ██████  ███████ ██████
██   ██ ██   ██ ██   ██ ████   ██ ██      ██   ██     ██      ██   ██ ██   ██ ████   ██ ██       ██      ██   ██
██████  ██████  ███████ ██ ██  ██ ██      ███████     ██      ███████ ███████ ██ ██  ██ ██   ███ █████   ██████
██   ██ ██   ██ ██   ██ ██  ██ ██ ██      ██   ██     ██      ██   ██ ██   ██ ██  ██ ██ ██    ██ ██      ██   ██
██████  ██   ██ ██   ██ ██   ████  ██████ ██   ██      ██████ ██   ██ ██   ██ ██   ████  ██████  ███████ ██   ██
#  
#  Version:      development
#  Maintainer:   Adli I. Ifkar <adly.shadowbane@gmail.com>
```

# Branch Changer

This is a simple script to change branch of multiple git repositories at once.

## Background
Well, I made this app 'cause I'm a little bit bored and annoyed that I need to quickly change branch of multiple git repositories at once. So, I made this script to do that.
You could easily run a bash script, but like I said, I'm bored and I want to make something useful.

## Usage

### Linux / MacOS Intel Silicon


### MacOS Apple Silicon


### Windows
Sadly, I don't have Windows machine to test this script. But, I think it should work on Windows 10 with WSL2.

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

# License
MIT License
