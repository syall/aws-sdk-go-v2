# v1.7.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.3 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.2 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.6.1 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2023-11-16)

* **Feature**: Amazon Pinpoint now offers additional operations as part of version 2 of the SMS and voice APIs. This release includes 26 new APIs to create and manage phone number registrations, add verified destination numbers, and request sender IDs.

# v1.5.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2023-09-18)

* **Announcement**: [BREAKFIX] Change in MaxResults datatype from value to pointer type in cognito-sync service.
* **Feature**: Adds several endpoint ruleset changes across all models: smaller rulesets, removed non-unique regional endpoints, fixes FIPS and DualStack endpoints, and make region not required in SDK::Endpoint. Additional breakfix to cognito-sync field.

# v1.2.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2023-08-01)

* No change notes available for this release.

# v1.2.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.14 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.13 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.12 (2023-06-15)

* No change notes available for this release.

# v1.1.11 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.10 (2023-05-04)

* No change notes available for this release.

# v1.1.9 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.8 (2023-04-10)

* No change notes available for this release.

# v1.1.7 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.6 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.5 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.4 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.1.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.1.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.0.20 (2022-12-23)

* No change notes available for this release.

# v1.0.19 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.18 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.17 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.16 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.15 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.14 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.13 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.12 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.11 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.10 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.9 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.8 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.7 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.6 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.5 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.4 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.3 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.2 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2022-04-14)

* No change notes available for this release.

# v1.0.0 (2022-03-31)

* **Release**: New AWS service client module
* **Feature**: Amazon Pinpoint now offers a version 2.0 suite of SMS and voice APIs, providing increased control over sending and configuration. This release is a new SDK for sending SMS and voice messages called PinpointSMSVoiceV2.

