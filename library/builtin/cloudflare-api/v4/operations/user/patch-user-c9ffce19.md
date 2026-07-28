---
title: Edit User
page_id: operation-patch-user-59d74f16
path: operations/user
description: Edit part of your user details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /user
operation_ids:
    - user-edit-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit User

`PATCH /user`

Operation ID: `user-edit-user`

Edit part of your user details.

## Definition

```yaml
{"operationId": "user-edit-user", "summary": "Edit User", "description": "Edit part of your user details.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"country": {"$ref": "#/components/schemas/iam_country"}, "first_name": {"$ref": "#/components/schemas/iam_first_name"}, "last_name": {"$ref": "#/components/schemas/iam_last_name"}, "telephone": {"$ref": "#/components/schemas/iam_telephone"}, "zipcode": {"$ref": "#/components/schemas/iam_zipcode"}}}}}}, "responses": {"200": {"description": "Edit User response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_user_response"}}}}, "4XX": {"description": "Edit User response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["User"], "x-api-token-group": ["User Details Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.user.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
