---
title: Validate external ID integration with IAM role trust policy
page_id: operation-post-tailnet-tailnet-aws-external-id-id-validate-aws-trust-policy-fdfda372
path: operations/logging
description: |-
    Validate that Tailscale can assume your IAM role with (and only with)
    this external ID.

    OAuth Scope: `log_streaming`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/aws-external-id/{id}/validate-aws-trust-policy
operation_ids:
    - validateAwsExternalId
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Validate external ID integration with IAM role trust policy

`POST /tailnet/{tailnet}/aws-external-id/{id}/validate-aws-trust-policy`

Operation ID: `validateAwsExternalId`

Validate that Tailscale can assume your IAM role with (and only with)
this external ID.

OAuth Scope: `log_streaming`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- name: id
  in: path
  description: The AWS external ID to validate.
  required: true
  example: 60fe9ce7-7791-4ab3-ab34-4294f5972725
  schema:
    type: string
```

## Definition

```yaml
summary: Validate external ID integration with IAM role trust policy
description: |
    Validate that Tailscale can assume your IAM role with (and only with)
    this external ID.

    OAuth Scope: `log_streaming`.
operationId: validateAwsExternalId
tags:
    - Logging
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    roleArn:
                        type: string
                        description: ARN of the AWS IAM role to validate with this external ID.
                example:
                    roleArn: arn:aws:iam::000000000000:role/tailscale-log-writer
responses:
    '200':
        description: Validation succeeded for this external ID and IAM role.
    '403':
        description: User does not have sufficient access for this tailnet.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet or external ID not found.
        $ref: '#/components/responses/404'
    '422':
        description: Validation failed for this external ID and IAM role.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        message:
                            type: string
                            description: The reason for validation failure.
```
