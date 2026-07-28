---
title: Validate SQL
page_id: operation-post-accounts-account-id-pipelines-v1-validate-sql-bbafcb77
path: operations/workers-pipelines-other
description: Validate Arroyo SQL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/validate_sql
operation_ids:
    - postV4AccountsByAccount_idPipelinesV1Validate_sql
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate SQL

`POST /accounts/{account_id}/pipelines/v1/validate_sql`

Operation ID: `postV4AccountsByAccount_idPipelinesV1Validate_sql`

Validate Arroyo SQL.

## Definition

```yaml
{"operationId": "postV4AccountsByAccount_idPipelinesV1Validate_sql", "summary": "Validate SQL", "description": "Validate Arroyo SQL.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"sql": {"description": "Specifies SQL to validate.", "type": "string", "example": "insert into sink select * from source;"}}, "required": ["sql"]}}}}, "responses": {"200": {"description": "Indicates SQL validation success.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"graph": {"$ref": "#/components/schemas/cloudflare-pipelines_PipelineGraph"}, "tables": {"description": "Indicates tables involved in the processing.", "type": "object", "additionalProperties": {"properties": {"id": {"type": "string"}, "name": {"type": "string"}, "type": {"type": "string"}, "version": {"type": "number"}}, "required": ["id", "version", "type", "name"], "type": "object"}}}, "required": ["tables"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates SQL validation failed."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "validate-sql", "x-forge-hidden": true}
```
