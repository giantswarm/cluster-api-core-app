# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.22-gs5] - 2021-11-24

### Fixed

- Prepare to run multiple apps (CRD Webhook, Dynamic Webhook and Controller).

### Changed

- Align CRD webhook name with `cluster-api-app`.

## [0.3.22-gs4] - 2021-08-16

## [0.3.22-gs3] - 2021-08-16

## [0.3.22-gs2] - 2021-08-13

### Changed

- Don't publish app to the app collection.
- Use giantswarm/config for config management.

## [0.3.22-gs1] - 2021-08-02

### Changed

- Updated cluster-api to v0.3.22.

## [0.3.19-gs1] - 2021-06-17

### Changed

- Updated cluster-api to v0.3.19.

## [0.3.14-gs3] - 2021-06-01

### Removed

- Remove kube-rbac-proxy for the metrics endpoint.

### Fixed

- Fixed label selector for webhook and manager services.

## [0.3.14-gs2] - 2021-05-27

## [0.3.14-gs1] - 2021-05-12

## [0.0.2] - 2021-03-18

### Added

- Set `appVersion` explicitly.

## [0.0.1] - 2021-03-03

### Added

- Add initial app implementation.


[Unreleased]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.22-gs5...HEAD
[0.3.22-gs5]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.22-gs4...v0.3.22-gs5
[0.3.22-gs4]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.22-gs3...v0.3.22-gs4
[0.3.22-gs3]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.22-gs2...v0.3.22-gs3
[0.3.22-gs2]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.22-gs1...v0.3.22-gs2
[0.3.22-gs1]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.19-gs1...v0.3.22-gs1
[0.3.19-gs1]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.14-gs3...v0.3.19-gs1
[0.3.14-gs3]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.14-gs2...v0.3.14-gs3
[0.3.14-gs2]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.3.14-gs1...v0.3.14-gs2
[0.3.14-gs1]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.0.2...v0.3.14-gs1
[0.0.2]: https://github.com/giantswarm/cluster-api-core-app/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/cluster-api-core-app/releases/tag/v0.0.1
