---
title: Get the Latest Scan Result
page_id: operation-get-accounts-account-id-cloudforce-one-scans-results-config-id-13537feb
path: operations/scans
description: Get the Latest Scan Result
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/scans/results/{config_id}
operation_ids:
    - get_GetOpenPorts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Latest Scan Result

`GET /accounts/{account_id}/cloudforce-one/scans/results/{config_id}`

Operation ID: `get_GetOpenPorts`

## Definition

```yaml
{"operationId": "get_GetOpenPorts", "summary": "Get the Latest Scan Result", "parameters": [{"name": "account_id", "in": "path", "description": "Defines the Account ID.", "required": true, "schema": {"description": "Defines the Account ID.", "type": "string"}}, {"name": "config_id", "in": "path", "description": "Defines the Config ID.", "required": true, "schema": {"description": "Defines the Config ID.", "type": "string"}}], "responses": {"200": {"description": "Returns Current Open Ports.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "string"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "properties": {"1.1.1.1": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_port"}}}, "required": ["1.1.1.1"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "messages", "errors"]}}}}, "4XX": {"description": "Get the Latest Scan Result failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
