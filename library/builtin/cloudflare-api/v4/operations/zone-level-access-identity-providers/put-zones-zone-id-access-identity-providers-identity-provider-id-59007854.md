---
title: Update an Access identity provider
page_id: operation-put-zones-zone-id-access-identity-providers-identity-provider-id-e6b30225
path: operations/zone-level-access-identity-providers
description: Updates a configured identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - zone-level-access-identity-providers-update-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access identity provider

`PUT /zones/{zone_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `zone-level-access-identity-providers-update-an-access-identity-provider`

Updates a configured identity provider.

## Definition

```yaml
{"operationId": "zone-level-access-identity-providers-update-an-access-identity-provider", "summary": "Update an Access identity provider", "description": "Updates a configured identity provider.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_identity-providers-2"}}}}, "responses": {"200": {"description": "Update an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-16"}}}}, "4XX": {"description": "Update an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-identity-providers", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stability": "beta"}
```
