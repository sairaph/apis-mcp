---
title: Get a posture integration
page_id: operation-get-posture-integrations-id-1629f33d
path: operations/deviceposture
description: |-
    Gets the posture integration identified by `{id}`.

    OAuth Scope: `feature_settings:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /posture/integrations/{id}
operation_ids:
    - getPostureIntegration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a posture integration

`GET /posture/integrations/{id}`

Operation ID: `getPostureIntegration`

Gets the posture integration identified by `{id}`.

OAuth Scope: `feature_settings:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/id'
```

## Definition

```yaml
summary: Get a posture integration
description: |
    Gets the posture integration identified by `{id}`.

    OAuth Scope: `feature_settings:read`.
operationId: getPostureIntegration
tags:
    - DevicePosture
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/PostureIntegration'
    '404':
        description: Posture integration not found, or user is not authorized to read it.
        $ref: '#/components/responses/404'
```
