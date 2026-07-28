---
title: Update a posture integration
page_id: operation-patch-posture-integrations-id-e13a6cd0
path: operations/deviceposture
description: |-
    Updates the posture integration identified by `{id}`. You may omit the `clientSecret` from your request to retain the previously configured `clientSecret`.

    OAuth Scope: `feature_settings`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /posture/integrations/{id}
operation_ids:
    - updatePostureIntegration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update a posture integration

`PATCH /posture/integrations/{id}`

Operation ID: `updatePostureIntegration`

Updates the posture integration identified by `{id}`. You may omit the `clientSecret` from your request to retain the previously configured `clientSecret`.

OAuth Scope: `feature_settings`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/id'
```

## Definition

```yaml
summary: Update a posture integration
description: |
    Updates the posture integration identified by `{id}`. You may omit the `clientSecret` from your request to retain the previously configured `clientSecret`.

    OAuth Scope: `feature_settings`.
operationId: updatePostureIntegration
tags:
    - DevicePosture
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/PostureIntegration'
            example:
                cloudId: global
                clientId: 93013672-b00c-4344-80ca-7ecf74f9dce1
                tenantId: d1ae389b-5207-43a2-afca-2de6b03ac7e3
                clientSecret: as32598d#@%asdf
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/PostureIntegration'
    '403':
        description: User does not have sufficient access to update this posture integration.
        $ref: '#/components/responses/403'
    '404':
        description: Posture integration not found.
        $ref: '#/components/responses/404'
```
