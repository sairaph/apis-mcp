---
title: Organization Details
page_id: operation-get-user-organizations-organization-id-71b8574f
path: operations/user-s-organizations
description: Gets a specific organization the user is associated with.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/organizations/{organization_id}
operation_ids:
    - user'-s-organizations-organization-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Organization Details

`GET /user/organizations/{organization_id}`

Operation ID: `user'-s-organizations-organization-details`

Gets a specific organization the user is associated with.

## Definition

```yaml
{"operationId": "user'-s-organizations-organization-details", "summary": "Organization Details", "description": "Gets a specific organization the user is associated with.", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}}], "responses": {"200": {"description": "Organization Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_organization_response"}}}}, "4XX": {"description": "Organization Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Organizations"], "x-cfDeprecation": {"description": "This endpoint and its related APIs are deprecated in favor of the `/accounts` equivalent APIs, which have a broader range of features and are backwards compatible with these API.", "display": true, "eol": "2020-02-04", "id": "org_deprecation"}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
