---
title: DevicePostureAttributes
page_id: schema-devicepostureattributes-326f6db4
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DevicePostureAttributes

```yaml
type: object
properties:
    attributes:
        type: object
        description: |
            Contains all the posture attributes assigned to a node.
            Attribute values can be strings, numbers or booleans.
        additionalProperties:
            x-additionalPropertiesName: Posture attributes
            anyOf:
                - type: string
                - type: number
                - type: boolean
    expiries:
        type: object
        description: |
            Contains the expiry time for each posture attribute, if set.
        additionalProperties:
            type: string
            format: date-time
            example: '2022-12-01T05:23:30Z'
            description: Expiry time for a given posture attribute.
example:
    attributes:
        custom:myScore: 80
        custom:diskEncryption: true
        custom:myAttribute: my_value
        node:os: linux
        node:osVersion: 5.19.0-42-generic
        node:tsReleaseTrack: stable
        node:tsVersion: 1.40.0
        node:tsAutoUpdate: false
        node:tsStateEncrypted: false
    expiries:
        custom:myScore: '2024-04-23T18:25:43.511Z'
```
