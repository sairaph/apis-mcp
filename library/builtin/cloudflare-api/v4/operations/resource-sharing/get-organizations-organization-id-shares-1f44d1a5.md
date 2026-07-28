---
title: List organization shares
page_id: operation-get-organizations-organization-id-shares-3a352b66
path: operations/resource-sharing
description: Lists all organization shares.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/shares
operation_ids:
    - organization-shares-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List organization shares

`GET /organizations/{organization_id}/shares`

Operation ID: `organization-shares-list`

Lists all organization shares.

## Definition

```yaml
{"operationId": "organization-shares-list", "summary": "List organization shares", "description": "Lists all organization shares.", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_organization_id"}}, {"$ref": "#/components/parameters/resource-sharing_status"}, {"$ref": "#/components/parameters/resource-sharing_kind"}, {"$ref": "#/components/parameters/resource-sharing_target_type"}, {"$ref": "#/components/parameters/resource-sharing_resource_types"}, {"$ref": "#/components/parameters/resource-sharing_order"}, {"$ref": "#/components/parameters/resource-sharing_direction"}, {"$ref": "#/components/parameters/resource-sharing_page"}, {"$ref": "#/components/parameters/resource-sharing_per_page"}], "responses": {"200": {"description": "List organization shares response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_response_collection"}}}}, "4XX": {"description": "List organization shares response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "List organization shares response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.shares", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
