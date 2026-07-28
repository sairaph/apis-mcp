---
title: Create or get AWS external id
page_id: operation-post-tailnet-tailnet-aws-external-id-30d443e9
path: operations/logging
description: |-
    Get an AWS external id to use for streaming tailnet logs to S3 using role-based authentication,
    creating a new one for this tailnet when necessary.

    OAuth Scope: `log_streaming`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/aws-external-id
operation_ids:
    - getAwsExternalId
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create or get AWS external id

`POST /tailnet/{tailnet}/aws-external-id`

Operation ID: `getAwsExternalId`

Get an AWS external id to use for streaming tailnet logs to S3 using role-based authentication,
creating a new one for this tailnet when necessary.

OAuth Scope: `log_streaming`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create or get AWS external id
description: |
    Get an AWS external id to use for streaming tailnet logs to S3 using role-based authentication,
    creating a new one for this tailnet when necessary.

    OAuth Scope: `log_streaming`.
operationId: getAwsExternalId
tags:
    - Logging
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    reusable:
                        type: boolean
                        description: If set to true, this same AWS external id will be returned on future calls to this endpoint, if and only if those calls also mark `reusable` as true, and the ID has not yet been linked with an AWS account.
            example:
                reusable: true
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/AwsExternalId'
    '403':
        description: User does not have sufficient access to obtain an AWS external id.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
```
