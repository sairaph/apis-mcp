---
title: Get log streaming configuration
page_id: operation-get-tailnet-tailnet-logging-logtype-stream-05269e14
path: operations/logging
description: |-
    Retrieve the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/logging/{logType}/stream
operation_ids:
    - getLogStreamingConfiguration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get log streaming configuration

`GET /tailnet/{tailnet}/logging/{logType}/stream`

Operation ID: `getLogStreamingConfiguration`

Retrieve the log streaming configuration for the provided log type.

OAuth Scope: `log_streaming:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/logType'
```

## Definition

```yaml
summary: Get log streaming configuration
description: |
    Retrieve the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming:read`.
operationId: getLogStreamingConfiguration
tags:
    - Logging
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/LogstreamEndpointConfiguration'
    '404':
        description: Log streaming has not been configured, this `logType` is not supported, or user does not have sufficient access to view log streaming configuration.
        $ref: '#/components/responses/404'
```
