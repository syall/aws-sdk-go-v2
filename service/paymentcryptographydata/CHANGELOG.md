# v1.5.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.5 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.4 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.4.3 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2023-08-31)

* **Feature**: Make KeyCheckValue field optional when using asymmetric keys as Key Check Values typically only apply to symmetric keys

# v1.1.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2023-08-01)

* No change notes available for this release.

# v1.1.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.4 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.3 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.2 (2023-06-15)

* No change notes available for this release.

# v1.0.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2023-06-08)

* **Release**: New AWS service client module
* **Feature**: Initial release of AWS Payment Cryptography DataPlane Plane service for performing cryptographic operations typically used during card payment processing.

