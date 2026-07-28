---
title: Set custom device posture attributes
page_id: operation-post-device-deviceid-attributes-attributekey-4b7a58c3
path: operations/devices
description: |-
    Create or update a custom posture attribute on the specified device.
    User-managed attributes must be in the custom namespace,
    which is indicated by prefixing the attribute key with `custom:`.

    OAuth Scope: `devices:posture_attributes`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/attributes/{attributeKey}
operation_ids:
    - setCustomDevicePostureAttributes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set custom device posture attributes

`POST /device/{deviceId}/attributes/{attributeKey}`

Operation ID: `setCustomDevicePostureAttributes`

Create or update a custom posture attribute on the specified device.
User-managed attributes must be in the custom namespace,
which is indicated by prefixing the attribute key with `custom:`.

OAuth Scope: `devices:posture_attributes`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
- $ref: '#/components/parameters/attributeKey'
```

## Definition

```yaml
summary: Set custom device posture attributes
description: |
    Create or update a custom posture attribute on the specified device.
    User-managed attributes must be in the custom namespace,
    which is indicated by prefixing the attribute key with `custom:`.

    OAuth Scope: `devices:posture_attributes`.
operationId: setCustomDevicePostureAttributes
tags:
    - Devices
requestBody:
    required: true
    content:
        application/json:
            schema:
                type: object
                properties:
                    value:
                        anyOf:
                            - type: string
                            - type: number
                            - type: boolean
                        example: my_value
                        description: |
                            A value can be either a string, number or boolean.

                            A string value can have a maximum length of 50 characters,
                            and can only contain letters, numbers, underscores, and periods.

                            A number value is an integer and must be a JSON safe number (up to 2^53 - 1).
                    expiry:
                        type: string
                        format: date-time
                        example: '2022-12-01T05:23:30Z'
                        description: |
                            An optional expiry time for a given posture attribute. If set, Tailscale
                            will automatically remove the attribute within a few minutes after the specified
                            time.
                    comment:
                        type: string
                        maxLength: 200
                        description: |
                            An optional comment indicating a reason why an attribute is set,
                            which will be added to the audit log.
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DevicePostureAttributes'
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
