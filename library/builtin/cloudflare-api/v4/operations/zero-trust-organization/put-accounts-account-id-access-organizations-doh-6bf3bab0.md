---
title: Update your Zero Trust organization DoH settings
page_id: operation-put-accounts-account-id-access-organizations-doh-805c2b5d
path: operations/zero-trust-organization
description: Updates the DoH settings for your Zero Trust organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/organizations/doh
operation_ids:
    - zero-trust-organization-update-your-zero-trust-organization-doh-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update your Zero Trust organization DoH settings

`PUT /accounts/{account_id}/access/organizations/doh`

Operation ID: `zero-trust-organization-update-your-zero-trust-organization-doh-settings`

Updates the DoH settings for your Zero Trust organization.

## Definition

```yaml
{"operationId": "zero-trust-organization-update-your-zero-trust-organization-doh-settings", "summary": "Update your Zero Trust organization DoH settings", "description": "Updates the DoH settings for your Zero Trust organization.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"properties": {"doh_jwt_duration": {"$ref": "#/components/schemas/access_doh_jwt_duration"}, "service_token_id": {"description": "The uuid of the service token you want to use for DoH authentication", "type": "string", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}}}}}, "responses": {"201": {"description": "Update your Zero Trust organization DoH settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-2"}, {"properties": {"result": {"type": "object", "properties": {"doh_jwt_duration": {"$ref": "#/components/schemas/access_doh_jwt_duration"}}}}}]}}}}, "4XX": {"description": "Update your Zero Trust organization DoH settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.organizations.doh", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
