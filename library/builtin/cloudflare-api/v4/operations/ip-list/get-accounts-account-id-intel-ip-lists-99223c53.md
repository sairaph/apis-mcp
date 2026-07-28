---
title: Get Available IP Lists
page_id: operation-get-accounts-account-id-intel-ip-lists-1e5aca5b
path: operations/ip-list
description: Returns a list of available IP list categories (e.g., anonymizer, botnetcc, malware, tor, vpn, open_proxies). This endpoint provides metadata about which IP lists are available in the system.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/ip-lists
operation_ids:
    - ip-list-get-ip-lists
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Available IP Lists

`GET /accounts/{account_id}/intel/ip-lists`

Operation ID: `ip-list-get-ip-lists`

Returns a list of available IP list categories (e.g., anonymizer, botnetcc, malware, tor, vpn, open_proxies). This endpoint provides metadata about which IP lists are available in the system.

## Definition

```yaml
{"operationId": "ip-list-get-ip-lists", "summary": "Get Available IP Lists", "description": "Returns a list of available IP list categories (e.g., anonymizer, botnetcc, malware, tor, vpn, open_proxies). This endpoint provides metadata about which IP lists are available in the system.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}], "responses": {"200": {"description": "Get Available IP Lists response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_components-schemas-response"}}}}, "4XX": {"description": "Get Available IP Lists response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_components-schemas-response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["IP List"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
