# Changes

## [0.39.0](https://www.github.com/googleapis/google-api-go-client/compare/v0.38.0...v0.39.0) (2021-02-04)


### Features

* **all:** auto-regenerate discovery clients , refs [#855](https://www.github.com/googleapis/google-api-go-client/issues/855) [#854](https://www.github.com/googleapis/google-api-go-client/issues/854) [#853](https://www.github.com/googleapis/google-api-go-client/issues/853) [#851](https://www.github.com/googleapis/google-api-go-client/issues/851) [#850](https://www.github.com/googleapis/google-api-go-client/issues/850) [#848](https://www.github.com/googleapis/google-api-go-client/issues/848)


### Bug Fixes

* **transport:** expand OS environment variables in cert provider command ([#852](https://www.github.com/googleapis/google-api-go-client/issues/852)) ([be6c56a](https://www.github.com/googleapis/google-api-go-client/commit/be6c56a1948a57eb0300613a70ef608330ca36e0))

## [0.38.0](https://www.github.com/googleapis/google-api-go-client/compare/v0.37.0...v0.38.0) (2021-01-29)


### Features

* **all:** auto-regenerate discovery clients , refs [#846](https://www.github.com/googleapis/google-api-go-client/issues/846) [#845](https://www.github.com/googleapis/google-api-go-client/issues/845) [#844](https://www.github.com/googleapis/google-api-go-client/issues/844) [#840](https://www.github.com/googleapis/google-api-go-client/issues/840)


### Bug Fixes

* **internal:** don't self-sign JWT when an endpoint provided ([#847](https://www.github.com/googleapis/google-api-go-client/issues/847)) ([55f262c](https://www.github.com/googleapis/google-api-go-client/commit/55f262c3a4e8d287ceeeee844b0d174299acc439))
* **internal:** don't use self-signed JWT with impersonation ([#788](https://www.github.com/googleapis/google-api-go-client/issues/788)) ([1dc7dac](https://www.github.com/googleapis/google-api-go-client/commit/1dc7dacd54b4b93f5465b71f2ee8c27e59630454))

## [0.37.0](https://www.github.com/googleapis/google-api-go-client/compare/v0.36.0...v0.37.0) (2021-01-25)


### Features

* **all:** auto-regenerate discovery clients , refs [#839](https://www.github.com/googleapis/google-api-go-client/issues/839) [#838](https://www.github.com/googleapis/google-api-go-client/issues/838) [#836](https://www.github.com/googleapis/google-api-go-client/issues/836) [#834](https://www.github.com/googleapis/google-api-go-client/issues/834) [#833](https://www.github.com/googleapis/google-api-go-client/issues/833) [#831](https://www.github.com/googleapis/google-api-go-client/issues/831) [#830](https://www.github.com/googleapis/google-api-go-client/issues/830) [#829](https://www.github.com/googleapis/google-api-go-client/issues/829) [#828](https://www.github.com/googleapis/google-api-go-client/issues/828) [#827](https://www.github.com/googleapis/google-api-go-client/issues/827) [#825](https://www.github.com/googleapis/google-api-go-client/issues/825) [#824](https://www.github.com/googleapis/google-api-go-client/issues/824) [#823](https://www.github.com/googleapis/google-api-go-client/issues/823) [#822](https://www.github.com/googleapis/google-api-go-client/issues/822) [#820](https://www.github.com/googleapis/google-api-go-client/issues/820) [#819](https://www.github.com/googleapis/google-api-go-client/issues/819) [#817](https://www.github.com/googleapis/google-api-go-client/issues/817) [#816](https://www.github.com/googleapis/google-api-go-client/issues/816) [#812](https://www.github.com/googleapis/google-api-go-client/issues/812) [#811](https://www.github.com/googleapis/google-api-go-client/issues/811) [#810](https://www.github.com/googleapis/google-api-go-client/issues/810) [#809](https://www.github.com/googleapis/google-api-go-client/issues/809) [#807](https://www.github.com/googleapis/google-api-go-client/issues/807) [#806](https://www.github.com/googleapis/google-api-go-client/issues/806) [#805](https://www.github.com/googleapis/google-api-go-client/issues/805) [#803](https://www.github.com/googleapis/google-api-go-client/issues/803) [#800](https://www.github.com/googleapis/google-api-go-client/issues/800) [#799](https://www.github.com/googleapis/google-api-go-client/issues/799) [#793](https://www.github.com/googleapis/google-api-go-client/issues/793) [#792](https://www.github.com/googleapis/google-api-go-client/issues/792) [#786](https://www.github.com/googleapis/google-api-go-client/issues/786) [#784](https://www.github.com/googleapis/google-api-go-client/issues/784) [#782](https://www.github.com/googleapis/google-api-go-client/issues/782) [#779](https://www.github.com/googleapis/google-api-go-client/issues/779) [#771](https://www.github.com/googleapis/google-api-go-client/issues/771) [#770](https://www.github.com/googleapis/google-api-go-client/issues/770) [#768](https://www.github.com/googleapis/google-api-go-client/issues/768)
* **transport/bytestream:** Add Close method for shutdown ([#787](https://www.github.com/googleapis/google-api-go-client/issues/787)) ([96bfd87](https://www.github.com/googleapis/google-api-go-client/commit/96bfd877fbc5869c0a4d87de4daeb1f76aaca79d)), refs [#775](https://www.github.com/googleapis/google-api-go-client/issues/775)


### Bug Fixes

* **all:** use CheckResponse for media downloads ([#773](https://www.github.com/googleapis/google-api-go-client/issues/773)) ([39cbab0](https://www.github.com/googleapis/google-api-go-client/commit/39cbab06d28f1d017bfc016c6735f6f45c51c90e)), refs [#752](https://www.github.com/googleapis/google-api-go-client/issues/752)
* **compute:** don't tigger linter for field named Deprecated ([#774](https://www.github.com/googleapis/google-api-go-client/issues/774)) ([d2bc921](https://www.github.com/googleapis/google-api-go-client/commit/d2bc921f997425bc267d8e4845286b0d67bbe1ef)), refs [#767](https://www.github.com/googleapis/google-api-go-client/issues/767)
* don't use markdown style links ([#789](https://www.github.com/googleapis/google-api-go-client/issues/789)) ([09ddacb](https://www.github.com/googleapis/google-api-go-client/commit/09ddacba9c3b45798fa309d3719638c754ec69a1)), refs [#712](https://www.github.com/googleapis/google-api-go-client/issues/712)
* **transport/grpc:** check Compute Engine environment for DirectPath ([#781](https://www.github.com/googleapis/google-api-go-client/issues/781)) ([89287b6](https://www.github.com/googleapis/google-api-go-client/commit/89287b68a240f818e9ae70a6395b1d72e21ee236))

## [0.36.0](https://www.github.com/googleapis/google-api-go-client/compare/v0.35.0...v0.36.0) (2020-12-03)


### Features

* **all:** auto-regenerate discovery clients , refs [#766](https://www.github.com/googleapis/google-api-go-client/issues/766) [#762](https://www.github.com/googleapis/google-api-go-client/issues/762) [#758](https://www.github.com/googleapis/google-api-go-client/issues/758) [#760](https://www.github.com/googleapis/google-api-go-client/issues/760) [#757](https://www.github.com/googleapis/google-api-go-client/issues/757) [#756](https://www.github.com/googleapis/google-api-go-client/issues/756) [#754](https://www.github.com/googleapis/google-api-go-client/issues/754) [#753](https://www.github.com/googleapis/google-api-go-client/issues/753) [#749](https://www.github.com/googleapis/google-api-go-client/issues/749) [#747](https://www.github.com/googleapis/google-api-go-client/issues/747) [#744](https://www.github.com/googleapis/google-api-go-client/issues/744)
* **internaloption:** add better support for self-signed JWT ([#738](https://www.github.com/googleapis/google-api-go-client/issues/738)) ([1a7550f](https://www.github.com/googleapis/google-api-go-client/commit/1a7550f9546052997806ff7ea9bcba55326bdb16))
* **transport:** Add default certificate caching support ([#721](https://www.github.com/googleapis/google-api-go-client/issues/721)) ([caa4d89](https://www.github.com/googleapis/google-api-go-client/commit/caa4d89fd452600f9911cfa0945500566cf9c72e))


### Bug Fixes

* **google-api-go-generator:** add patch for compute mtls endpoint ([#761](https://www.github.com/googleapis/google-api-go-client/issues/761)) ([445fe0b](https://www.github.com/googleapis/google-api-go-client/commit/445fe0be627de9769c5b252e77858c78682a6b8c))

## [0.35.0](https://www.github.com/googleapis/google-api-go-client/compare/v0.34.0...v0.35.0) (2020-11-06)


### Features

* **all:** auto-regenerate discovery clients , refs [#743](https://www.github.com/googleapis/google-api-go-client/issues/743) [#741](https://www.github.com/googleapis/google-api-go-client/issues/741) [#739](https://www.github.com/googleapis/google-api-go-client/issues/739) [#737](https://www.github.com/googleapis/google-api-go-client/issues/737) [#735](https://www.github.com/googleapis/google-api-go-client/issues/735) [#733](https://www.github.com/googleapis/google-api-go-client/issues/733) [#730](https://www.github.com/googleapis/google-api-go-client/issues/730) [#729](https://www.github.com/googleapis/google-api-go-client/issues/729) [#724](https://www.github.com/googleapis/google-api-go-client/issues/724)
* **internaloption:** add EnableDirectPath internaloption ([#732](https://www.github.com/googleapis/google-api-go-client/issues/732)) ([baf33b2](https://www.github.com/googleapis/google-api-go-client/commit/baf33b2baf0a1a5459e7b4ff193fa47117829169))

## v0.34.0

- transport:
  - Fix mergeEndpoint logic to support default endpoints without scheme.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.


## v0.33.0

- idtoken:
  - Add an example of setting HTTP auth header.
- internal:
  - Refactor service-account.json into a testdata folder.
- transport:
  - Add device certificate Authentication support to GRPC.
  - Support `GOOGLE_API_USE_CLIENT_CERTIFICATE` and
    `GOOGLE_API_USE_MTLS_ENDPOINT` environment variables to conform with
    [AIP-4114](https://google.aip.dev/auth/4114).
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.32.0

- option:
  - Add experimental ImpersonateCredentials option for impersonating a service
    account.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.31.0

- cloudcommerceprocurement:
  - Regenerate `cloudcommerceprocurement` v1.
- Updated dependencies.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.30.0

### Changes

- idtoken:
  - Fix flaky ecdsa test.
  - Fix some typos in the docs.
  - Fix `WithCredentialsJSON` not working with `NewClient`.
  - Speed up tests.
- internal:
  - Remove the install of staticcheck.
  - Automate dependency updates with Renovate.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.29.0

### Changes

- Various updates to autogenerated clients.
- transport: internal bug fixes for mTLS.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.28.0

### Changes

- gensupport:
  - Retry the initial request for a media upload in the storage library only.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.27.0

### Changes

- gensupport:
  - Expand retryable errors to include wrapped errors and transient network
    failures.
  - Add retry for the initial request in a resumable or multipart upload.
- transport/http:
  - Don't reuse a base transport between clients. This fixes a race condition in
    defaultBaseTransport.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.26.0

### Changes

- idtoken:
  - Populate Claims map.
- transport/http:
  - Update default HTTP transport settings to use a larger value for
    MaxIdleConnsPerHost. This improves performance in the storage client.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.25.0

### Announcements

[goolgeapis/google-api-go-client](https://github.com/googleapis/google-api-go-client)
has moved its source of truth to GitHub and is no longer a mirror. This means
that our contributing process has changed a bit. We will now be conducting all
code reviews on GitHub which means we now accept Pull Requests! If you have a
version of the codebase previously checked out you may wish to update your git
remote to point to GitHub.

### Changes

- all:
  - Updated instructions in CONTRIBUTING.md for pull requests.
- idtoken:
  - Validate now checks to see if the token is expired.
- sheets:
  - Update ExtendedValue Fields to be pointer types.
- support/bunder:
  - Fix a deadlock that could when handler limit was set to one.
- transport:
  - Allow `GOOGLE_API_USE_MTLS` overriding the mTLS endpoint behavior for the
    HTTP client.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.24.0

### Changes

- googleapi:
  - Return more details with errors.
- sqladmin
  - Make StorageAutoResize a pointer type for v1.
- transport/http:
  - When provided, use the TokenSource from options for NewTransport. This fixes
    a bug in idtoken.NewClient where the wrong TokenSource was being used for
    authentication.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.23.0

### Changes

- apigee:
  - Re-enable generation of this client.
- compute:
  - Make Id a on ExternalVpnGateway a pointer type.
- idtoken:
  - Add new package to support making requests with and validating Google ID
    tokens.
- slides:
  - Make int values of Range optional.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.22.0

### Depreciation Notice

- Package `google.golang.org/api/sql/v1beta4` has been deprecated as it was
  generated under the wrong name. This package will be removed in a future
  release. Please migrate to: `google.golang.org/api/sqladmin/v1beta4`.

### Changes

- Apigee client has temporarily been disabled.

- Updated custom search example to be in line with new API.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.21.0

- Disabled automatic switching to *.mtls.googleapis.com endpoints.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.20.0

- WithGRPCConnectionPool is a no-op for some APIs.

- correctly report Go version of runtime.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.19.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.18.0

- Add the WithClientCertSource option for mTLS (client TLS certificates), currently only supported for HTTP clients.

- Allow host:port for endpoint overrides, rather than requiring the full base URL (for google.golang.org/api clients).

- Make DialPool work with WithGRPCConn plus non-zero pool size [googleapis/google-cloud-go#1780](https://github.com/googleapis/google-cloud-go/issues/1780)

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.17.0

- Revert sqladmin package name back from sql to sqladmin. (#448)

- Various updates to autogenerated clients.

Internal:

- transport/grpc: add internal WithDialPool option for GAPIC clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.16.0

- Increase the default chunk size for uploads (e.g., for the storage package) to 16 MB.

- transport:
  - Automatically populate QuotaProject from the "quota_project_id" field in the JSON credentials file.
  - Add grpc.DialPool, which opens multiple grpc.ClientConns based on WithGRPCConnectionPool option.

- Added a check to prevent mixed calls to Add and AddWait.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.15.0

- monitoring/v3:
  - Rename Service to MService; revert APIService to Service.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.14.0

- Fix for setting custom HTTP headers in the absence of UserAgent.

- Add a client option for disabling telemetry such as OpenCensus.

- Performance improvements to google.golang.org/api/support/bundler.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.13.0

- Changes to how media path redirection is handled in generated code.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.12.0

- Various updates to autogenerated clients.

## v0.11.0

- Various updates to autogenerated clients.

- Module information now indicates go 1.11 as oldest supported version.  As of
  October 1, 2019 versions 1.9 and 1.10 are no longer supported.

- Removed the following APIs which are no longer available via the discovery
  service: dfareporting/v2.8, prediction/*.

- The internal gensupport library has been relocated to the more idiomatic
  path internal/gensupport.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.10.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.9.0

- Small fix to chunking retry logic such that each chunk has its own retry
  deadline, instead of unbounded retries.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.8.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.7.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.6.0

- Add support for GCP DirectPath.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.5.0

- Better support for google.api.HttpBody.
- Support for google.api.HttpBody in the healthcare API.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.4.0

- Includes a re-pin of opencensus, greatly reducing the transitive
dependency list.
- Deletes photoslibrary/v1. The photoslibrary team hopes to fully support Go in
the near future, but this autogenerated library is ready to be sunset. If you
rely on this client, please vendor this library at v0.3.2.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.3.2

This patch releases re-builds the go.sum. This was not possible in the
previous release.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.3.1

This patch release removes github.com/golang/lint from the transitive
dependency list, resolving `go get -u` problems.

_Please note_: this release intentionally has a broken go.sum. Please use v0.3.2.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.3.0

go.mod modifications, including removal of go 1.12 statement and update of
opencensus dependency.

_Please note_: the release version is not indicative of an individual client's
stability or version.

## v0.2.0

General improvements.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

## v0.1.0

Initial release along with Go module support.

_Please note:_ the release version is not indicative of an individual client's
stability or version.
