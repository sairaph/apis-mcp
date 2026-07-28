---
title: Get log streaming status
page_id: operation-get-tailnet-tailnet-logging-logtype-stream-status-64bff5ee
path: operations/logging
description: |-
    Retrieve the log streaming status for the provided log type.

    OAuth Scope: `log_streaming:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/logging/{logType}/stream/status
operation_ids:
    - getLogStreamingStatus
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get log streaming status

`GET /tailnet/{tailnet}/logging/{logType}/stream/status`

Operation ID: `getLogStreamingStatus`

Retrieve the log streaming status for the provided log type.

OAuth Scope: `log_streaming:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/logType'
```

## Definition

```yaml
summary: Get log streaming status
description: |
    Retrieve the log streaming status for the provided log type.

    OAuth Scope: `log_streaming:read`.
operationId: getLogStreamingStatus
tags:
    - Logging
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/LogstreamEndpointPublishingStatus'
    '404':
        description: Log streaming has not been configured, this `logType` is not supported, or user does not have sufficient access to view log streaming status.
        $ref: '#/components/responses/404'
    '502':
        description: The system was unable to communicate with logging server.
        $ref: '#/components/responses/502'
```
