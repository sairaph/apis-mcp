---
title: Get CMB config
page_id: operation-get-accounts-account-id-logs-control-cmb-config-71dc75b8
path: operations/logcontrol-cmb-config-for-an-account
description: Gets CMB config.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/control/cmb/config
operation_ids:
    - get-accounts-account_id-logs-control-cmb-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get CMB config

`GET /accounts/{account_id}/logs/control/cmb/config`

Operation ID: `get-accounts-account_id-logs-control-cmb-config`

Gets CMB config.

## Definition

```yaml
{"operationId": "get-accounts-account_id-logs-control-cmb-config", "summary": "Get CMB config", "description": "Gets CMB config.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logcontrol_identifier"}}], "responses": {"200": {"description": "Get CMB config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_cmb_config_response_single"}}}}, "4XX": {"description": "Get CMB config response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logcontrol CMB config for an account"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read", "#analytics:read"]}}
```
