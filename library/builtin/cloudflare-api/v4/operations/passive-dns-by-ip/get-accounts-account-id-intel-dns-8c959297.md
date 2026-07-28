---
title: Get Passive DNS by IP
page_id: operation-get-accounts-account-id-intel-dns-7e1b0e6f
path: operations/passive-dns-by-ip
description: Gets a list of all the domains that have resolved to a specific IP address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/dns
operation_ids:
    - passive-dns-by-ip-get-passive-dns-by-ip
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Passive DNS by IP

`GET /accounts/{account_id}/intel/dns`

Operation ID: `passive-dns-by-ip-get-passive-dns-by-ip`

Gets a list of all the domains that have resolved to a specific IP address.

## Definition

```yaml
{"operationId": "passive-dns-by-ip-get-passive-dns-by-ip", "summary": "Get Passive DNS by IP", "description": "Gets a list of all the domains that have resolved to a specific IP address.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "start_end_params", "in": "query", "schema": {"$ref": "#/components/schemas/intel_start_end_params"}}, {"name": "ipv4", "in": "query", "schema": {"type": "string"}}, {"name": "page", "in": "query", "schema": {"description": "Requested page within paginated list of results.", "type": "number", "example": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results requested.", "type": "number", "example": 20}}], "responses": {"200": {"description": "Get Passive DNS by IP response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_components-schemas-single_response"}}}}, "4XX": {"description": "Get Passive DNS by IP response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_components-schemas-single_response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Passive DNS by IP"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
