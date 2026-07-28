---
title: Set log streaming configuration
page_id: operation-put-tailnet-tailnet-logging-logtype-stream-1096e577
path: operations/logging
description: |-
    Set the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming`. `device_invites` and `policy_file` are also required if streaming to a [private endpoint](https://tailscale.com/kb/1255/log-streaming#private-endpoints).
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PUT
api_endpoints:
    - /tailnet/{tailnet}/logging/{logType}/stream
operation_ids:
    - setLogStreamingConfiguration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set log streaming configuration

`PUT /tailnet/{tailnet}/logging/{logType}/stream`

Operation ID: `setLogStreamingConfiguration`

Set the log streaming configuration for the provided log type.

OAuth Scope: `log_streaming`. `device_invites` and `policy_file` are also required if streaming to a [private endpoint](https://tailscale.com/kb/1255/log-streaming#private-endpoints).

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/logType'
```

## Definition

```yaml
summary: Set log streaming configuration
description: |
    Set the log streaming configuration for the provided log type.

    OAuth Scope: `log_streaming`. `device_invites` and `policy_file` are also required if streaming to a [private endpoint](https://tailscale.com/kb/1255/log-streaming#private-endpoints).
operationId: setLogStreamingConfiguration
tags:
    - Logging
requestBody:
    description: |
        The [LogstreamEndpointConfiguration](#model/logstreamendpointconfiguration) to set.
        `logType` is specified in the request URL rather than the body.
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/LogstreamEndpointConfiguration'
            example:
                destinationType: elastic
                url: http://100.71.134.73:80/config-log-datastream
                user: myusername
                token: mytoken
responses:
    '200':
        description: Successful operation.
    '400':
        description: Request has missing or invalid parameter(s).
        $ref: '#/components/responses/400'
    '403':
        description: User does not have sufficient access to update log streaming configuration.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found, this `logType` is not supported, or user does not have sufficient access to view log streaming configuration.
        $ref: '#/components/responses/404'
```
