---
title: Update an existing Scan Config
page_id: operation-patch-accounts-account-id-cloudforce-one-scans-config-config-id-c510a161
path: operations/scans
description: Update an existing Scan Config
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/scans/config/{config_id}
operation_ids:
    - post_ConfigUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an existing Scan Config

`PATCH /accounts/{account_id}/cloudforce-one/scans/config/{config_id}`

Operation ID: `post_ConfigUpdate`

## Definition

```yaml
{"operationId": "post_ConfigUpdate", "summary": "Update an existing Scan Config", "parameters": [{"name": "account_id", "in": "path", "description": "Defines the Account ID.", "required": true, "schema": {"description": "Defines the Account ID.", "type": "string"}}, {"name": "config_id", "in": "path", "description": "Defines the Config ID.", "required": true, "schema": {"description": "Defines the Config ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"frequency": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_frequency"}, "ips": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ips"}, "ports": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_ports"}}}}}}, "responses": {"200": {"description": "Returns the updated config.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_scan-config"}}}]}}}}, "4XX": {"description": "Update an Existing Scan Config failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"], "x-api-token-group": ["Cloudforce One Write"]}
```
