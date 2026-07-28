---
title: List configuration audit logs
page_id: operation-get-tailnet-tailnet-logging-configuration-258fdb1b
path: operations/logging
description: |-
    List all configuration audit logs for a tailnet.

    OAuth Scope: `logs:configuration:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/logging/configuration
operation_ids:
    - listConfigurationAuditLogs
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List configuration audit logs

`GET /tailnet/{tailnet}/logging/configuration`

Operation ID: `listConfigurationAuditLogs`

List all configuration audit logs for a tailnet.

OAuth Scope: `logs:configuration:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/start'
- $ref: '#/components/parameters/end'
- $ref: '#/components/parameters/actor'
- $ref: '#/components/parameters/target'
- $ref: '#/components/parameters/event'
```

## Definition

```yaml
summary: List configuration audit logs
description: |
    List all configuration audit logs for a tailnet.

    OAuth Scope: `logs:configuration:read`.
operationId: listConfigurationAuditLogs
tags:
    - Logging
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    description: A structured response for all of a Tailnet's audit logs over a period of time.
                    properties:
                        version:
                            type: string
                            description: Version of audit logs response.
                            example: '1.1'
                        tailnet:
                            type: string
                            example: example.com
                            description: The tailnet on which the logged configuration changes were made.
                        logs:
                            type: array
                            description: Matching log entries, ordered chronologically.
                            items:
                                $ref: '#/components/schemas/ConfigurationAuditLog'
    '400':
        description: Request has missing or invalid parameter(s).
        $ref: '#/components/responses/400'
    '403':
        description: User does not have sufficient access to view configuration audit logs.
        $ref: '#/components/responses/403'
    '404':
        description: Logging is not supported on this deployment of Tailscale.
        $ref: '#/components/responses/404'
```
