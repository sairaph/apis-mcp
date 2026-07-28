---
title: Revoke application tokens
page_id: operation-post-zones-zone-id-access-apps-app-id-revoke-tokens-7bc7da78
path: operations/zone-level-access-applications
description: Revokes all tokens issued for an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/revoke_tokens
operation_ids:
    - zone-level-access-applications-revoke-service-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke application tokens

`POST /zones/{zone_id}/access/apps/{app_id}/revoke_tokens`

Operation ID: `zone-level-access-applications-revoke-service-tokens`

Revokes all tokens issued for an application.

## Definition

```yaml
{"operationId": "zone-level-access-applications-revoke-service-tokens", "summary": "Revoke application tokens", "description": "Revokes all tokens issued for an application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Revoke application tokens response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_empty_response-2"}}}}, "4XX": {"description": "Revoke application tokens response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-api-token-group": ["Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.revoke.tokens", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation revokes all service tokens for the application."}
```
