---
title: Delete custom device posture attributes
page_id: operation-delete-device-deviceid-attributes-attributekey-fda5a6dc
path: operations/devices
description: |-
    Delete a posture attribute from the specified device.
    This is only applicable to user-managed posture attributes in the custom namespace,
    which is indicated by prefixing the attribute key with `custom:`.

    OAuth Scope: `devices:posture_attributes`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /device/{deviceId}/attributes/{attributeKey}
operation_ids:
    - deleteCustomDevicePostureAttributes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete custom device posture attributes

`DELETE /device/{deviceId}/attributes/{attributeKey}`

Operation ID: `deleteCustomDevicePostureAttributes`

Delete a posture attribute from the specified device.
This is only applicable to user-managed posture attributes in the custom namespace,
which is indicated by prefixing the attribute key with `custom:`.

OAuth Scope: `devices:posture_attributes`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
- $ref: '#/components/parameters/attributeKey'
```

## Definition

```yaml
summary: Delete custom device posture attributes
description: |
    Delete a posture attribute from the specified device.
    This is only applicable to user-managed posture attributes in the custom namespace,
    which is indicated by prefixing the attribute key with `custom:`.

    OAuth Scope: `devices:posture_attributes`.
operationId: deleteCustomDevicePostureAttributes
tags:
    - Devices
responses:
    '200':
        description: Successful operation.
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '429':
        $ref: '#/components/responses/429'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
