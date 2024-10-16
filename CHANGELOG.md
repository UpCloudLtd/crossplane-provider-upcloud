# Changelog

All notable changes to this project will be documented in this file.
See updating [Changelog example here](https://keepachangelog.com/en/1.0.0/)

## [Unreleased]

## [0.0.7] - 2024-10-16

### Changed

- **BREAKING** provider: API root group of all CRDs is now `upcloud.com`, previously `upcloud.io`
- go module is now `github.com/UpCloudLtd/crossplane-provider-upcloud`, previously `github.com/UpCloudLtd/provider-upcloud`
- HTTP request header `User-Agent` is now configured to `crossplane-provider-upcloud/{tag}`

[Unreleased]: https://github.com/UpCloudLtd/crossplane-provider-upcloud/compare/v0.0.7...HEAD
[0.0.7]: https://github.com/UpCloudLtd/crossplane-provider-upcloud/compare/v0.0.6...v0.0.7