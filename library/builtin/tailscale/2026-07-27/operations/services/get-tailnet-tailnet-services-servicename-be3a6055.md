---
title: Get a Service
page_id: operation-get-tailnet-tailnet-services-servicename-ad07f85e
path: operations/services
description: |-
    Retrieve the details for the specified Service.

    OAuth Scope: `services:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}
operation_ids:
    - getService
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a Service

`GET /tailnet/{tailnet}/services/{serviceName}`

Operation ID: `getService`

Retrieve the details for the specified Service.

OAuth Scope: `services:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
```

## Definition

```yaml
summary: Get a Service
description: |
    Retrieve the details for the specified Service.

    OAuth Scope: `services:read`.
operationId: getService
tags:
    - Services
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/VIPServiceInfo'
    '400':
        $ref: '#/components/responses/400'
    '403':
        description: Access to the Service is forbidden.
        $ref: '#/components/responses/403'
    '404':
        description: Service not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
