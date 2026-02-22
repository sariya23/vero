# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),

## [0.1.0] - 2026-02-22

### Added

- New public checker functions: `check.assert.AlmostEqualTime` and `check.require.AlmostEqualTime`.
These functions compare two time values with a specified precision and either fail immediately or perform a soft assertion.

- Internal function `AlmostEqualTime` - contains the comparison logic used by the test public functions.

- Internal function `ShowDiffAlmostEqualTime` - displays the two time values used in `AlmostEqualTime`.

- New packages: `check.assert` and `check.require`.