---
title: List devices hosting a Service
page_id: operation-get-tailnet-tailnet-services-servicename-devices-0e4ffeaf
path: operations/services
description: |-
    List all devices that are hosting the specified Service.

    OAuth Scope: `services`, `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}/devices
operation_ids:
    - listServiceHosts
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List devices hosting a Service

`GET /tailnet/{tailnet}/services/{serviceName}/devices`

Operation ID: `listServiceHosts`

List all devices that are hosting the specified Service.

OAuth Scope: `services`, `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
```

## Definition

```yaml
summary: List devices hosting a Service
description: |
    List all devices that are hosting the specified Service.

    OAuth Scope: `services`, `devices:core`.
operationId: listServiceHosts
tags:
    - Services
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        hosts:
                            type: array
                            items:
                                $ref: '#/components/schemas/ServiceHostInfo'
    '400':
        description: Invalid parameters or no permission to Services.
        $ref: '#/components/responses/400'
    '403':
        description: Access to the Service or devices is forbidden.
        $ref: '#/components/responses/403'
    '404':
        description: Service not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
