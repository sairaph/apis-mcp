---
title: Delete an Access identity provider
page_id: operation-delete-zones-zone-id-access-identity-providers-identity-provider-id-1f63665d
path: operations/zone-level-access-identity-providers
description: Deletes an identity provider from Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - zone-level-access-identity-providers-delete-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access identity provider

`DELETE /zones/{zone_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `zone-level-access-identity-providers-delete-an-access-identity-provider`

Deletes an identity provider from Access.

## Definition

```yaml
{"operationId": "zone-level-access-identity-providers-delete-an-access-identity-provider", "summary": "Delete an Access identity provider", "description": "Deletes an identity provider from Access.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-identity-providers", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
