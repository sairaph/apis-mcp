---
title: Create a new Scan Config
page_id: operation-post-accounts-account-id-cloudforce-one-scans-config-3c033b46
path: operations/scans
description: Create a new Scan Config
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/scans/config
operation_ids:
    - post_ConfigCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Scan Config

`POST /accounts/{account_id}/cloudforce-one/scans/config`

Operation ID: `post_ConfigCreate`

## Definition

```yaml
{"operationId": "post_ConfigCreate", "summary": "Create a new Scan Config", "parameters": [{"name": "account_id", "in": "path", "description": "Defines the Account ID.", "required": true, "schema": {"description": "Defines the Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"frequency": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_frequency"}, "ips": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ips"}, "ports": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ports"}}, "required": ["ips"]}}}}, "responses": {"200": {"description": "Returns the created config.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_scan-config"}}}]}}}}, "4XX": {"description": "Create a new Scan Config failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"], "x-api-token-group": ["Cloudforce One Write"]}
```
