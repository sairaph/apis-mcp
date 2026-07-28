---
title: User Details
page_id: operation-get-user-5184b908
path: operations/user
description: Retrieves detailed information about the currently authenticated user, including email, name, and account memberships.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user
operation_ids:
    - user-user-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# User Details

`GET /user`

Operation ID: `user-user-details`

Retrieves detailed information about the currently authenticated user, including email, name, and account memberships.

## Definition

```yaml
{"operationId": "user-user-details", "summary": "User Details", "description": "Retrieves detailed information about the currently authenticated user, including email, name, and account memberships.", "responses": {"200": {"description": "User Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_user_response"}}}}, "4XX": {"description": "User Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.user.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
