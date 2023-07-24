# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [1.2.0] - 2023-07-21

### Added
- Added `-merge` flag (default to `false`) to merge branch.

### Changed
- Changed `-to` flag to `-destination`
- Changed `-from` flag to `-source`

### Fixed
- Fixed warning message for `-push` and `-force` flag


## [1.1.0] - 2023-07-21

### Added
- [samber/oops](https://github.com/samber/oops) for better stack trace.

### Changed
- `workingDir`, `branchFrom`, `branchTo`, and `projects` now defaults to empty string.
- Throw error if parse flag failed.
- Removing `sync.WaitGroup{}`

### Fixed

- Fixed flow error.
- Fixed GitHub Actions version not found.


## [1.0.1] - 2023-07-03

### Fixed

- Fixed path error.

### Changed

- Updated [README.md](README.md)


## [1.0.0] - 2023-07-03

### Added
- Initial release.


## [0.1.6] - 2023-07-03

### Fixed

- Test fix auto release.


## [0.1.5] - 2023-07-03

### Fixed

- Test fix auto release.


## [0.1.4] - 2023-07-03

### Fixed

- Test fix auto release.


## [0.1.3] - 2023-07-03

### Added

- Changelog file.


## [0.1.2] - 2023-07-03

### Fixed

- Fixed test build.


[1.2.0]: https://github.com/shadowbane/branch-changer/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/shadowbane/branch-changer/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/shadowbane/branch-changer/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/shadowbane/branch-changer/compare/v0.1.6...v1.0.0
[0.1.6]: https://github.com/shadowbane/branch-changer/compare/v0.1.5...v0.1.6
[0.1.5]: https://github.com/shadowbane/branch-changer/compare/v0.1.4...v0.1.5
[0.1.4]: https://github.com/shadowbane/branch-changer/compare/v0.1.3...v0.1.4
[0.1.3]: https://github.com/shadowbane/branch-changer/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/shadowbane/branch-changer/releases/tag/v0.1.2
