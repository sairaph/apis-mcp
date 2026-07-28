---
title: Delete a posture integration
page_id: operation-delete-posture-integrations-id-56af6874
path: operations/deviceposture
description: |-
    Delete a specific posture integration.

    OAuth Scope: `feature_settings`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /posture/integrations/{id}
operation_ids:
    - deletePostureIntegration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a posture integration

`DELETE /posture/integrations/{id}`

Operation ID: `deletePostureIntegration`

Delete a specific posture integration.

OAuth Scope: `feature_settings`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/id'
```

## Definition

```yaml
summary: Delete a posture integration
description: |
    Delete a specific posture integration.

    OAuth Scope: `feature_settings`.
operationId: deletePostureIntegration
tags:
    - DevicePosture
responses:
    '200':
        description: Successful operation.
    '403':
        description: User does not have sufficient access to delete this posture integration.
        $ref: '#/components/responses/403'
    '404':
        description: Posture integration not found.
        $ref: '#/components/responses/404'
```
