---
title: Leave Organization
page_id: operation-delete-user-organizations-organization-id-260c3070
path: operations/user-s-organizations
description: Removes association to an organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /user/organizations/{organization_id}
operation_ids:
    - user'-s-organizations-leave-organization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Leave Organization

`DELETE /user/organizations/{organization_id}`

Operation ID: `user'-s-organizations-leave-organization`

Removes association to an organization.

## Definition

```yaml
{"operationId": "user'-s-organizations-leave-organization", "summary": "Leave Organization", "description": "Removes association to an organization.", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Leave Organization response", "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}}}}}}, "4XX": {"description": "Leave Organization response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Organizations"], "x-cfDeprecation": {"description": "This endpoint and its related APIs are deprecated in favor of the `/accounts` equivalent APIs, which have a broader range of features and are backwards compatible with these API.", "display": true, "eol": "2020-02-04", "id": "org_deprecation"}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
