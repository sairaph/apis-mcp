---
title: Search Resources
page_id: operation-get-accounts-account-id-load-balancers-search-74f713b3
path: operations/account-load-balancer-search
description: Search for Load Balancing resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/search
operation_ids:
    - account-load-balancer-search-search-resources
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search Resources

`GET /accounts/{account_id}/load_balancers/search`

Operation ID: `account-load-balancer-search-search-resources`

Search for Load Balancing resources.

## Definition

```yaml
{"operationId": "account-load-balancer-search-search-resources", "summary": "Search Resources", "description": "Search for Load Balancing resources.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "query", "in": "query", "schema": {"description": "Search query term.", "type": "string", "example": "primary", "default": ""}}, {"name": "references", "in": "query", "schema": {"description": "The type of references to include. \"*\" to include both referral and referrer references. \"\" to not include any reference information.", "type": "string", "example": "*", "default": "", "enum": ["", "*", "referral", "referrer"]}}, {"name": "page", "in": "query", "schema": {"type": "number", "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "number", "default": 25, "maximum": 1000, "minimum": 1}}], "responses": {"200": {"description": "Search Resources response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_api-paginated-response-collection"}, {"$ref": "#/components/schemas/load-balancing_search_result"}]}}}}, "4XX": {"description": "Search Resources response failure.", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-paginated-response-collection"}, {"$ref": "#/components/schemas/load-balancing_search_result"}]}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Search"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
