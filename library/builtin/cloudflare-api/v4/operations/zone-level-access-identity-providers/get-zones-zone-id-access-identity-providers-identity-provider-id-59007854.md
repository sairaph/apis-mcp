---
title: Get an Access identity provider
page_id: operation-get-zones-zone-id-access-identity-providers-identity-provider-id-c81e1b81
path: operations/zone-level-access-identity-providers
description: Fetches a configured identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - zone-level-access-identity-providers-get-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access identity provider

`GET /zones/{zone_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `zone-level-access-identity-providers-get-an-access-identity-provider`

Fetches a configured identity provider.

## Definition

```yaml
{"operationId": "zone-level-access-identity-providers-get-an-access-identity-provider", "summary": "Get an Access identity provider", "description": "Fetches a configured identity provider.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-16"}}}}, "4XX": {"description": "Get an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-identity-providers", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
