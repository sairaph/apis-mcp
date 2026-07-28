---
title: Retrieve a specific sensitivity group template.
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-templates-template-id-3c7c15f7
path: operations/dlp-sensitivity-group-templates
description: Gets an available sensitivity group template.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/templates/{template_id}
operation_ids:
    - dlp-sensitivity-group-template-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific sensitivity group template.

`GET /accounts/{account_id}/dlp/sensitivity_groups/templates/{template_id}`

Operation ID: `dlp-sensitivity-group-template-read`

Gets an available sensitivity group template.

## Definition

```yaml
{"operationId": "dlp-sensitivity-group-template-read", "summary": "Retrieve a specific sensitivity group template.", "description": "Gets an available sensitivity group template.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "template_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Sensitivity group template read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroupTemplate"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group template read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Group Templates"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
