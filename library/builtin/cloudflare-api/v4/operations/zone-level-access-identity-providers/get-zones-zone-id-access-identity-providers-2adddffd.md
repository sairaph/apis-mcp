---
title: List Access identity providers
page_id: operation-get-zones-zone-id-access-identity-providers-51412a8b
path: operations/zone-level-access-identity-providers
description: Lists all configured identity providers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/identity_providers
operation_ids:
    - zone-level-access-identity-providers-list-access-identity-providers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access identity providers

`GET /zones/{zone_id}/access/identity_providers`

Operation ID: `zone-level-access-identity-providers-list-access-identity-providers`

Lists all configured identity providers.

## Definition

```yaml
{"operationId": "zone-level-access-identity-providers-list-access-identity-providers", "summary": "List Access identity providers", "description": "Lists all configured identity providers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List Access identity providers response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-17"}}}}, "4XX": {"description": "List Access identity providers response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-identity-providers", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
