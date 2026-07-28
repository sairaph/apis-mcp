---
title: Revoke all Access tokens for a user
page_id: operation-post-zones-zone-id-access-organizations-revoke-user-21cf9cf2
path: operations/zone-level-zero-trust-organization
description: Revokes a user's access across all applications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/organizations/revoke_user
operation_ids:
    - zone-level-zero-trust-organization-revoke-all-access-tokens-for-a-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke all Access tokens for a user

`POST /zones/{zone_id}/access/organizations/revoke_user`

Operation ID: `zone-level-zero-trust-organization-revoke-all-access-tokens-for-a-user`

Revokes a user's access across all applications.

## Definition

```yaml
{"operationId": "zone-level-zero-trust-organization-revoke-all-access-tokens-for-a-user", "summary": "Revoke all Access tokens for a user", "description": "Revokes a user's access across all applications.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-4"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"email": {"description": "The email of the user to revoke.", "type": "string", "example": "test@example.com"}}, "required": ["email"]}}}}, "responses": {"200": {"description": "Revoke all Access tokens for a user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_empty_response"}}}}, "4xx": {"description": "Revoke all Access tokens for a user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-organizations", "x-fern-sdk-method-name": "revoke-user-create", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation revokes all access tokens for a user, affecting their access to all applications within the zone."}
```
