---
title: List Scan Configs
page_id: operation-get-accounts-account-id-cloudforce-one-scans-config-60f1e7a2
path: operations/scans
description: List Scan Configs
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/scans/config
operation_ids:
    - get_ConfigFetch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Scan Configs

`GET /accounts/{account_id}/cloudforce-one/scans/config`

Operation ID: `get_ConfigFetch`

## Definition

```yaml
{"operationId": "get_ConfigFetch", "summary": "List Scan Configs", "parameters": [{"name": "account_id", "in": "path", "description": "Defines the Account ID.", "required": true, "schema": {"description": "Defines the Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns all Scan Configs.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_scan-config"}}}}]}}}}, "4XX": {"description": "List Scan Configs failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
