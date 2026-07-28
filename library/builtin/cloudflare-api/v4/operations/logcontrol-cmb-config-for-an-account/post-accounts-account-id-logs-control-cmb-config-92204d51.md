---
title: Update CMB config
page_id: operation-post-accounts-account-id-logs-control-cmb-config-98d01fe7
path: operations/logcontrol-cmb-config-for-an-account
description: Updates CMB config.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/logs/control/cmb/config
operation_ids:
    - post-accounts-account_id-logs-control-cmb-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update CMB config

`POST /accounts/{account_id}/logs/control/cmb/config`

Operation ID: `post-accounts-account_id-logs-control-cmb-config`

Updates CMB config.

## Definition

```yaml
{"operationId": "post-accounts-account_id-logs-control-cmb-config", "summary": "Update CMB config", "description": "Updates CMB config.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logcontrol_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_cmb_config"}}}}, "responses": {"200": {"description": "Update CMB config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_cmb_config_response_single"}}}}, "4XX": {"description": "Update CMB config response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logcontrol CMB config for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}}
```
