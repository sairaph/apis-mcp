---
title: List all Services
page_id: operation-get-tailnet-tailnet-services-cdf86689
path: operations/services
description: |-
    List all Services configured for the tailnet. This includes all Services in the "advertised"
    tab of the Services page in the Tailscale admin console.

    OAuth Scope: `services:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/services
operation_ids:
    - listServices
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List all Services

`GET /tailnet/{tailnet}/services`

Operation ID: `listServices`

List all Services configured for the tailnet. This includes all Services in the "advertised"
tab of the Services page in the Tailscale admin console.

OAuth Scope: `services:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List all Services
description: |
    List all Services configured for the tailnet. This includes all Services in the "advertised"
    tab of the Services page in the Tailscale admin console.

    OAuth Scope: `services:read`.
operationId: listServices
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
                        vipServices:
                            type: array
                            items:
                                $ref: '#/components/schemas/VIPServiceInfo'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
