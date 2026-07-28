---
title: Delete a Service
page_id: operation-delete-tailnet-tailnet-services-servicename-35ae48da
path: operations/services
description: |-
    Delete the specified Service from the tailnet.

    OAuth Scope: `services`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}
operation_ids:
    - deleteService
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a Service

`DELETE /tailnet/{tailnet}/services/{serviceName}`

Operation ID: `deleteService`

Delete the specified Service from the tailnet.

OAuth Scope: `services`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
```

## Definition

```yaml
summary: Delete a Service
description: |
    Delete the specified Service from the tailnet.

    OAuth Scope: `services`.
operationId: deleteService
tags:
    - Services
responses:
    '200':
        description: Successful operation.
    '400':
        $ref: '#/components/responses/400'
    '403':
        description: Access to delete the Service is forbidden.
        $ref: '#/components/responses/403'
    '404':
        description: Service not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
