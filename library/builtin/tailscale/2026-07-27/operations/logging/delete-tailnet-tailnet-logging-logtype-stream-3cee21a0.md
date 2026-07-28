---
title: Disable log streaming
page_id: operation-delete-tailnet-tailnet-logging-logtype-stream-2d8ddef1
path: operations/logging
description: |-
    Delete the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /tailnet/{tailnet}/logging/{logType}/stream
operation_ids:
    - disableLogStreaming
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Disable log streaming

`DELETE /tailnet/{tailnet}/logging/{logType}/stream`

Operation ID: `disableLogStreaming`

Delete the log streaming configuration for the provided log type.

OAuth Scope: `log_streaming`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/logType'
```

## Definition

```yaml
summary: Disable log streaming
description: |
    Delete the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming`.
operationId: disableLogStreaming
tags:
    - Logging
responses:
    '200':
        description: Successful operation.
    '403':
        description: User does not have sufficient access to update log streaming configuration.
        $ref: '#/components/responses/403'
    '404':
        description: Log streaming has not been configured, this `logType` is not supported, or user does not have sufficient access to view log streaming configuration.
        $ref: '#/components/responses/404'
```
