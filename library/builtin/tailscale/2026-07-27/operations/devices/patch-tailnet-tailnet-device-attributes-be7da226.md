---
title: Batch update custom device posture attributes
page_id: operation-patch-tailnet-tailnet-device-attributes-8732ade9
path: operations/devices
description: |-
    Batch updates posture attributes across devices in a tailnet.

    This endpoint uses [JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396) semantics.
    Specifying `null` for an attribute will delete that attribute.

    Attributes must be in the `custom:` namespace.

    OAuth Scope: `devices:posture_attributes`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /tailnet/{tailnet}/device-attributes
operation_ids:
    - batchUpdateCustomDevicePostureAttributes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Batch update custom device posture attributes

`PATCH /tailnet/{tailnet}/device-attributes`

Operation ID: `batchUpdateCustomDevicePostureAttributes`

Batch updates posture attributes across devices in a tailnet.

This endpoint uses [JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396) semantics.
Specifying `null` for an attribute will delete that attribute.

Attributes must be in the `custom:` namespace.

OAuth Scope: `devices:posture_attributes`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Batch update custom device posture attributes
description: |
    Batch updates posture attributes across devices in a tailnet.

    This endpoint uses [JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396) semantics.
    Specifying `null` for an attribute will delete that attribute.

    Attributes must be in the `custom:` namespace.

    OAuth Scope: `devices:posture_attributes`.
operationId: batchUpdateCustomDevicePostureAttributes
tags:
    - Devices
requestBody:
    required: true
    content:
        application/json:
            schema:
                type: object
                properties:
                    nodes:
                        type: object
                        additionalProperties:
                            type: object
                            additionalProperties:
                                anyOf:
                                    - type: object
                                      properties:
                                        value:
                                            anyOf:
                                                - type: string
                                                - type: number
                                                - type: boolean
                                            description: |
                                                Attribute value.
                                        expiry:
                                            type: string
                                            format: date-time
                                            description: |
                                                Optional expiry time for the attribute.
                                                If set, Tailscale automatically removes the attribute within
                                                a few minutes after the specified time.
                                      required:
                                        - value
                                    - type: 'null'
                                description: |
                                    Attribute configuration with value and optional expiry, or `null` to delete the attribute.
                            description: |
                                A mapping of attribute name → attribute configuration for this device.
                        description: |
                            A mapping of deviceId → posture attributes.
                    comment:
                        type: string
                        maxLength: 200
                        description: |
                            An optional comment indicating why attributes are being set,
                            which will be added to the audit log.
                example:
                    nodes:
                        nPM2KNuedB21DEVEL:
                            custom:myattr:
                                value: my_value
                        nPpz3VEKzX11DEVEL:
                            custom:flag:
                                value: true
                                expiry: '2025-09-19T15:00:00Z'
                    comment: bulk posture attribute update
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: 'null'
    '400':
        description: An invalid request payload was sent
        $ref: '#/components/responses/400'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
