---
title: Get an Access policy
page_id: operation-get-zones-zone-id-access-apps-app-id-policies-policy-id-524c30c9
path: operations/zone-level-access-policies
description: Fetches a single Access policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/policies/{policy_id}
operation_ids:
    - zone-level-access-policies-get-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access policy

`GET /zones/{zone_id}/access/apps/{app_id}/policies/{policy_id}`

Operation ID: `zone-level-access-policies-get-an-access-policy`

Fetches a single Access policy.

## Definition

```yaml
{"operationId": "zone-level-access-policies-get-an-access-policy", "summary": "Get an Access policy", "description": "Fetches a single Access policy.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-22"}}}}, "4XX": {"description": "Get an Access policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.policies", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
