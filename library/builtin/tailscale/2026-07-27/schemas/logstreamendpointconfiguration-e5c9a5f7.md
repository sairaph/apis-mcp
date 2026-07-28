---
title: LogstreamEndpointConfiguration
page_id: schema-logstreamendpointconfiguration-e5c9a5f7
path: schemas
description: The current configuration of a log streaming endpoint.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# LogstreamEndpointConfiguration

The current configuration of a log streaming endpoint.

```yaml
type: object
description: The current configuration of a log streaming endpoint.
example:
    logType: configuration
    destinationType: elastic
    url: http://100.71.134.73:80/config-log-datastream
    user: myusername
properties:
    logType:
        description: The type of log that is streamed to this endpoint.
        readOnly: true
        $ref: '#/components/schemas/LogType'
    destinationType:
        type: string
        enum:
            - splunk
            - elastic
            - panther
            - cribl
            - datadog
            - axiom
            - s3
        description: The type of system to which logs are being streamed.
        example: elastic
    url:
        type: string
        description: The URL to which log streams are being posted. If the DestinationType is `s3`, the URL may be (and often is) empty to use the official Amazon S3 endpoint.
        example: http://100.71.134.73:80/config-log-datastream
    user:
        type: string
        description: The username with which log streams to this endpoint are authenticated.
        example: myusername
    uploadPeriodMinutes:
        type: integer
        description: An optional number of minutes to wait in between uploading new logs. If the quantity of logs does not fit within a single upload, multiple uploads will be made.
        maximum: 1440
        example: 5
    compressionFormat:
        type: string
        enum:
            - zstd
            - gzip
            - none
        description: The compression algorithm with which to compress logs. `none` disables compression. Defaults to `none`.
        example: zstd
    token:
        type: string
        description: The token/password with which log streams to this endpoint should be authenticated.
        writeOnly: true
        example: mytoken
    s3Bucket:
        type: string
        description: The S3 bucket name. Required if the destinationType is `s3`.
        example: mycompany-mybucket
    s3Region:
        type: string
        description: The region in which the S3 bucket is located. Required if the destinationType is `s3`.
        example: us-east-1
    s3KeyPrefix:
        type: string
        description: An optional S3 key prefix to prepend to the auto-generated S3 key name.
    s3AuthenticationType:
        type: string
        enum:
            - accesskey
            - rolearn
        description: What type of authentication to use for S3. Required if the destinationType is `s3`. Tailscale recommends using `rolearn`. See [Amazon documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_common-scenarios_third-party.html).
    s3AccessKeyId:
        type: string
        description: The S3 access key ID. Required if the destinationType is `s3` and `authenticationType` is `accesskey`.
    s3SecretAccessKey:
        type: string
        description: The S3 secret access key. Required if the destinationType is `s3` and `authenticationType` is `accesskey`.
        writeOnly: true
    s3RoleArn:
        type: string
        description: The Role ARN that Tailscale should supply to AWS when authenticating using role-based authentication. Required if the destinationType is `s3` and `authenticationType` is `rolearn`.
    s3ExternalId:
        type: string
        description: The AWS external id that Tailscale supplies when authenticating using role-based authentication. Populated if the destinationType is `s3` and `authenticationType` is `rolearn`. This corresponds to the `externalId` field obtained from [tailnet/{tailnet}/aws-external-id](#tag/logging/POST/tailnet/{tailnet}/aws-external-id).
        readOnly: true
    gcsBucket:
        type: string
        description: The GCS bucket name. Required if the destinationType is `gcs`.
    gcsKeyPrefix:
        type: string
        description: An optional GCS key prefix to append to the GCS bucket name.
    gcsScopes:
        type: array
        items:
            type: string
        description: The GCS scopes needed to be able to write to the GCS bucket.
    gcsCredentials:
        type: string
        description: The JSON workload identity credentials from GCS needed for accessing the GCS account.
```
