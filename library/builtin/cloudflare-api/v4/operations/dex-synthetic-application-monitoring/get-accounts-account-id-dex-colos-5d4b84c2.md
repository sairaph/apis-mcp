---
title: List Cloudflare colos
page_id: operation-get-accounts-account-id-dex-colos-4782f558
path: operations/dex-synthetic-application-monitoring
description: List Cloudflare colos that account's devices were connected to during a time period, sorted by usage starting from the most used colo. Colos without traffic are also returned and sorted alphabetically.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/colos
operation_ids:
    - dex-endpoints-list-colos
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Cloudflare colos

`GET /accounts/{account_id}/dex/colos`

Operation ID: `dex-endpoints-list-colos`

List Cloudflare colos that account's devices were connected to during a time period, sorted by usage starting from the most used colo. Colos without traffic are also returned and sorted alphabetically.

## Definition

```yaml
{"operationId": "dex-endpoints-list-colos", "summary": "List Cloudflare colos", "description": "List Cloudflare colos that account's devices were connected to during a time period, sorted by usage starting from the most used colo. Colos without traffic are also returned and sorted alphabetically.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "from", "in": "query", "description": "Start time for connection period in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-08-20T20:45:00Z"}}, {"name": "to", "in": "query", "description": "End time for connection period in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-08-24T20:45:00Z"}}, {"name": "sortBy", "in": "query", "description": "Type of usage that colos should be sorted by. If unspecified, returns all Cloudflare colos sorted alphabetically.", "schema": {"type": "string", "enum": ["fleet-status-usage", "application-tests-usage"]}}], "responses": {"200": {"description": "List colos response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_colos_response"}}}]}}}}, "4XX": {"description": "List colos failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.colos", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
