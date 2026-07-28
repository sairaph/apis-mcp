---
title: List OAuth Scopes
page_id: operation-get-oauth-scopes-7e1e5f63
path: operations/oauth-clients
description: List all available OAuth scopes. This endpoint requires authentication but has no authorization role requirements.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /oauth/scopes
operation_ids:
    - oauth-scopes-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List OAuth Scopes

`GET /oauth/scopes`

Operation ID: `oauth-scopes-list`

List all available OAuth scopes. This endpoint requires authentication but has no authorization role requirements.

## Definition

```yaml
{"operationId": "oauth-scopes-list", "summary": "List OAuth Scopes", "description": "List all available OAuth scopes. This endpoint requires authentication but has no authorization role requirements.", "responses": {"200": {"description": "List OAuth Scopes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_oauth_scopes_response"}}}}, "4XX": {"description": "List OAuth Scopes response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
