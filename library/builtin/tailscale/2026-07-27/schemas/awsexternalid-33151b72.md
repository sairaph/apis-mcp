---
title: AwsExternalId
page_id: schema-awsexternalid-33151b72
path: schemas
description: An external ID for use in authenticating to AWS using role-based authentication.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# AwsExternalId

An external ID for use in authenticating to AWS using role-based authentication.

```yaml
type: object
description: An external ID for use in authenticating to AWS using role-based authentication.
example:
    externalId: 60fe9ce7-7791-4ab3-ab34-4294f5972725
    tailscaleAwsAccountId: '001234567890'
properties:
    externalId:
        type: string
        description: The external id that Tailscale will supply to AWS when authenticating using role-based authentication.
        example: 60fe9ce7-7791-4ab3-ab34-4294f5972725
    tailscaleAwsAccountId:
        type: string
        description: The AWS account id that Tailscale will supply to AWS when authenticating using role-based authentication.
        example: '001234567890'
```
