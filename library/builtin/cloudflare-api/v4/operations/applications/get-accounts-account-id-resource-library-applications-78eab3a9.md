---
title: List applications
page_id: operation-get-accounts-account-id-resource-library-applications-6c283896
path: operations/applications
description: List applications with different filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/resource-library/applications
operation_ids:
    - getApplications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List applications

`GET /accounts/{account_id}/resource-library/applications`

Operation ID: `getApplications`

List applications with different filters.

## Definition

```yaml
{"operationId": "getApplications", "summary": "List applications", "description": "List applications with different filters.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "filter", "in": "query", "description": "Filter applications using key:value format. Supported filter keys:\n- name: Filter by application name (e.g., name:HR)\n- id: Filter by application ID (e.g., id:0b63249c-95bf-4cc0-a7cc-d7faaaf1dac0)\n- human_id: Filter by human-readable ID (e.g., human_id:HR)\n- hostname: Filter by hostname or support domain (e.g., hostname:portal.example.com)\n- source: Filter by application source name (e.g., source:cloudflare)\n- ip_subnet: Filter by IP subnet using CIDR containment — returns applications where any stored subnet contains the search value (e.g., ip_subnet:10.0.1.5/32 matches apps with 10.0.0.0/16)\n- intel_id: Filter by Intel API ID (e.g., intel_id:498). also supports multiple values (e.g., intel_id:498,1001)\n- category_id: Filter by category ID (e.g., category_id:37f8ec03-8766-49d4-9a15-369b044c842c).\n- category_name: Filter by category name (e.g., category_name:HR).\n- supported: Filter by supported Cloudflare product (e.g., supported:ACCESS). Values: GATEWAY, ACCESS, CASB.\n.\n", "schema": {"type": "string"}, "example": "filter=name:HR&filter=source:cloudflare&filter=intel_id:498"}, {"name": "limit", "in": "query", "description": "Limit of number of results to return (max 250).", "schema": {"type": "integer", "default": 25}}, {"name": "offset", "in": "query", "description": "Offset of results to return.", "schema": {"type": "integer", "default": 0}}, {"name": "order_by", "in": "query", "description": "Order results by field name and direction (e.g., name:asc). Ignored when search is provided; results are ranked by relevance instead.", "schema": {"type": "string"}, "example": "name:asc"}, {"name": "search", "in": "query", "description": "Fuzzy search across application name and hostnames. Results are ranked by relevance. Must be between 2 and 200 characters. Can be combined with filter parameters.", "schema": {"type": "string", "maxLength": 200, "minLength": 2}, "example": "MyNewApp"}], "responses": {"200": {"description": "Get the application response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/alexandria_get_applications_response"}}}}, "4XX": {"description": "Get application response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/alexandria_get_applications_response"}, {"$ref": "#/components/schemas/alexandria_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.applications.get", "x-fern-sdk-method-name": "applications"}
```
