---
title: Create a posture integration
page_id: operation-post-tailnet-tailnet-posture-integrations-19df399e
path: operations/deviceposture
description: |-
    Create a posture integration, returning the resulting [PostureIntegration](#model/postureintegration).

    Must include `provider` and `clientSecret`.

    Currently, only one integration for each provider is supported.

    OAuth Scope: `feature_settings`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/posture/integrations
operation_ids:
    - createPostureIntegration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create a posture integration

`POST /tailnet/{tailnet}/posture/integrations`

Operation ID: `createPostureIntegration`

Create a posture integration, returning the resulting [PostureIntegration](#model/postureintegration).

Must include `provider` and `clientSecret`.

Currently, only one integration for each provider is supported.

OAuth Scope: `feature_settings`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create a posture integration
description: |
    Create a posture integration, returning the resulting [PostureIntegration](#model/postureintegration).

    Must include `provider` and `clientSecret`.

    Currently, only one integration for each provider is supported.

    OAuth Scope: `feature_settings`.
operationId: createPostureIntegration
tags:
    - DevicePosture
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/PostureIntegration'
            example:
                provider: intune
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
        description: User does not have sufficient access to create posture integrations.
        $ref: '#/components/responses/403'
    '409':
        description: A posture integration for the same provider already exists.
        $ref: '#/components/responses/409'
```
