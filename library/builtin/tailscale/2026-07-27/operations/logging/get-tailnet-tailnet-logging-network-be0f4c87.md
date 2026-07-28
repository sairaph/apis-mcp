---
title: List network flow logs
page_id: operation-get-tailnet-tailnet-logging-network-cd6bc0c1
path: operations/logging
description: |-
    List all network flow logs for a tailnet.

    OAuth Scope: `logs:network:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/logging/network
operation_ids:
    - listNetworkFlowLogs
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List network flow logs

`GET /tailnet/{tailnet}/logging/network`

Operation ID: `listNetworkFlowLogs`

List all network flow logs for a tailnet.

OAuth Scope: `logs:network:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/start'
- $ref: '#/components/parameters/end'
```

## Definition

```yaml
summary: List network flow logs
description: |
    List all network flow logs for a tailnet.

    OAuth Scope: `logs:network:read`.
operationId: listNetworkFlowLogs
tags:
    - Logging
responses:
    '200':
        description: Successful operation. The `logs` field contains an array of [NetworkFlowLog](#model/networkflowlog) objects.
        content:
            application/json:
                schema:
                    type: object
                    description: A structured response for all of a Tailnet's network flow logs over a period of time.
                    properties:
                        logs:
                            type: array
                            description: Matching log entries, ordered chronologically.
                            items:
                                $ref: '#/components/schemas/NetworkFlowLog'
    '400':
        description: Request has missing or invalid parameter(s).
        $ref: '#/components/responses/400'
    '403':
        description: User does not have sufficient access to view network flow logs.
        $ref: '#/components/responses/403'
    '404':
        description: Logging is not supported on this deployment of Tailscale.
        $ref: '#/components/responses/404'
    '502':
        description: The system was unable to communicate with logging server.
        $ref: '#/components/responses/502'
```
