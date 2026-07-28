---
title: KeyCapabilities
page_id: schema-keycapabilities-44de49fc
path: schemas
description: '`capabilities` is a mapping of resources to permissible actions.'
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# KeyCapabilities

`capabilities` is a mapping of resources to permissible actions.

```yaml
type: object
description: |
    `capabilities` is a mapping of resources to permissible actions.
properties:
    devices:
        type: object
        description: |
            `devices` specifies the key's permissions over devices.
            This field is only populated for auth keys.
        properties:
            create:
                type: object
                description: |
                    `create` specifies the key's permissions when creating devices.
                properties:
                    reusable:
                        description: |
                            reusable for auth keys only; reusable auth keys can be used multiple times to register different devices.
                            Learn more about reusable auth keys at https://tailscale.com/kb/1085/#types-of-auth-keys.
                        type: boolean
                        example: true
                    ephemeral:
                        description: |
                            ephemeral for auth keys only; ephemeral keys are used to connect and then clean up short-lived devices.
                            Learn about ephemeral nodes at https://tailscale.com/kb/1111/.
                        type: boolean
                        example: false
                    preauthorized:
                        description: |
                            preauthorized for auth keys only; these are also referred to as "pre-approved" keys. 'true' means that devices
                            registered with this key won't require additional approval from a tailnet admin.
                            Learn about device approval at https://tailscale.com/kb/1099/.
                        type: boolean
                        example: true
                    tags:
                        description: |
                            tags are the tags that will be set on devices registered with this key.
                            Learn about tags at https://tailscale.com/kb/1068/.

                            Whether tags are required or optional depends on the owner of the auth key:
                            - When creating an auth key owned by the tailnet (using OAuth), it must have tags. The auth tags specified for that new auth key must exactly match the tags that are on the OAuth client used to create that auth key (or they must be tags that are owned by the tags that are on the OAuth client used to create the auth key).
                            - When creating an auth key owned by a user (using a user's access token), tags are optional.
                        type: array
                        items:
                            type: string
                        example:
                            - tag:example
```
