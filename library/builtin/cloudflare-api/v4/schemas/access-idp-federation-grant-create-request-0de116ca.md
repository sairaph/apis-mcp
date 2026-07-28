---
title: access_idp_federation_grant_create_request
page_id: schema-access-idp-federation-grant-create-request-0de116ca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_idp_federation_grant_create_request

```yaml
{"type": "object", "properties": {"idp_id": {"description": "UID of the identity provider to federate. Must be an existing identity provider in this account. One-time pin and Cloudflare-managed identity providers cannot be federated.", "type": "string", "format": "uuid", "example": "a79de439-0e7f-4ebb-8a02-222222222222", "x-auditable": true}}, "required": ["idp_id"]}
```
