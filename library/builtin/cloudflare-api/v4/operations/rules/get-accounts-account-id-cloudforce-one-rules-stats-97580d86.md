---
title: Get dashboard stats
page_id: operation-get-accounts-account-id-cloudforce-one-rules-stats-0da6a7fc
path: operations/rules
description: Get statistics about rules for the dashboard.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/stats
operation_ids:
    - cloudforce-one-get-rule-stats
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get dashboard stats

`GET /accounts/{account_id}/cloudforce-one/rules/stats`

Operation ID: `cloudforce-one-get-rule-stats`

Get statistics about rules for the dashboard.

## Definition

```yaml
{"operationId": "cloudforce-one-get-rule-stats", "summary": "Get dashboard stats", "description": "Get statistics about rules for the dashboard.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "responses": {"200": {"description": "Dashboard statistics.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_StatsResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
