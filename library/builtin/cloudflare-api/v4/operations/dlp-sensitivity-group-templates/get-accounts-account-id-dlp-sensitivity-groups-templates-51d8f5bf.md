---
title: Retrieve all sensitivity group templates in an account
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-templates-a43a1f96
path: operations/dlp-sensitivity-group-templates
description: Lists available sensitivity group templates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/templates
operation_ids:
    - dlp-sensitivity-group-templates-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all sensitivity group templates in an account

`GET /accounts/{account_id}/dlp/sensitivity_groups/templates`

Operation ID: `dlp-sensitivity-group-templates-list`

Lists available sensitivity group templates.

## Definition

```yaml
{"operationId": "dlp-sensitivity-group-templates-list", "summary": "Retrieve all sensitivity group templates in an account", "description": "Lists available sensitivity group templates.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Sensitivity group template read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroupTemplateArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group template read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Group Templates"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
