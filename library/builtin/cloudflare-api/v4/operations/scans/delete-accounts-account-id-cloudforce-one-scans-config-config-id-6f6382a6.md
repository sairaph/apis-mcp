---
title: Delete a Scan Config
page_id: operation-delete-accounts-account-id-cloudforce-one-scans-config-config-id-24452ccf
path: operations/scans
description: Delete a Scan Config
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/scans/config/{config_id}
operation_ids:
    - delete_DeleteScans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Scan Config

`DELETE /accounts/{account_id}/cloudforce-one/scans/config/{config_id}`

Operation ID: `delete_DeleteScans`

## Definition

```yaml
{"operationId": "delete_DeleteScans", "summary": "Delete a Scan Config", "parameters": [{"name": "account_id", "in": "path", "description": "Defines the Account ID.", "required": true, "schema": {"description": "Defines the Account ID.", "type": "string"}}, {"name": "config_id", "in": "path", "description": "Defines the Config ID.", "required": true, "schema": {"description": "Defines the Config ID.", "type": "string"}}], "responses": {"200": {"description": "Delete a Scan Config.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "string"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["success", "result", "messages", "errors"]}}}}, "4XX": {"description": "Delete a Scan Config failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-port-scan-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"], "x-api-token-group": ["Cloudforce One Write"]}
```
