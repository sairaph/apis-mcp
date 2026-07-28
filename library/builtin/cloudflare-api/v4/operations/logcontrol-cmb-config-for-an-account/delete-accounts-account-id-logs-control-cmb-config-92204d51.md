---
title: Delete CMB config
page_id: operation-delete-accounts-account-id-logs-control-cmb-config-b771b970
path: operations/logcontrol-cmb-config-for-an-account
description: Deletes CMB config.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/logs/control/cmb/config
operation_ids:
    - delete-accounts-account_id-logs-control-cmb-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete CMB config

`DELETE /accounts/{account_id}/logs/control/cmb/config`

Operation ID: `delete-accounts-account_id-logs-control-cmb-config`

Deletes CMB config.

## Definition

```yaml
{"operationId": "delete-accounts-account_id-logs-control-cmb-config", "summary": "Delete CMB config", "description": "Deletes CMB config.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logcontrol_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete CMB config response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common"}, {"properties": {"result": {"type": "object", "enum": [null], "nullable": true}}}]}}}}, "4XX": {"description": "Delete CMB config response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logcontrol CMB config for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}}
```
