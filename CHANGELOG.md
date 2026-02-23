# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),

## [0.1.4] - 2026-02-23

### Added

- Added a new `random` package with `Choice` and `Sample` functions.

## [0.1.3] - 2026-02-23

### Changed

- Rename from _probatigo_ to _vero_

## [0.1.2] - 2026-02-23

### Removed

- Internal function `AlmostEqualTime` and its tests;
- Packages `require` and `assert`.

### Changed

- Logic of `AlmostEqualTime` moved into `AssertAlmostEqualTime`;

- One package `check` instead of packages `require` and `assert`.

## [0.1.1] - 2026-02-23

### Changed

- Fixed the condition in `assert.AlmostEqualTime` and `require.AlmostEqualTime`: previously, when the times were equal, the checker failed due to an inverted boolean flag.

### Added

- Added tests for the public checker.

## [0.1.0] - 2026-02-22

### Added

- GitHub Actions for every push (`.github/worflows/main.yml`) and for release only (`.github/worflows/release.yml`);

- New public checker functions: `check.assert.AlmostEqualTime` and `check.require.AlmostEqualTime`.
These functions compare two time values with a specified precision and either fail immediately or perform a soft assertion.

- Internal function `AlmostEqualTime` - contains the comparison logic used by the test public functions.

- Internal function `ShowDiffAlmostEqualTime` - displays the two time values used in `AlmostEqualTime`.

- New packages: `check.assert` and `check.require`.