---
title: Cancel a crawl job.
page_id: operation-delete-accounts-account-id-browser-rendering-crawl-job-id-42b34268
path: operations/brapi
description: Cancels an ongoing crawl job by setting its status to cancelled and stopping all queued URLs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/browser-rendering/crawl/{job_id}
operation_ids:
    - brapi-delete_CancelCrawl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Cancel a crawl job.

`DELETE /accounts/{account_id}/browser-rendering/crawl/{job_id}`

Operation ID: `brapi-delete_CancelCrawl`

Cancels an ongoing crawl job by setting its status to cancelled and stopping all queued URLs.

## Definition

```yaml
{"operationId": "brapi-delete_CancelCrawl", "summary": "Cancel a crawl job.", "description": "Cancels an ongoing crawl job by setting its status to cancelled and stopping all queued URLs.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "job_id", "in": "path", "description": "The ID of the crawl job to cancel.", "required": true, "schema": {"description": "The ID of the crawl job to cancel.", "type": "string"}}], "responses": {"200": {"description": "Crawl job cancelled successfully.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "result": {"type": "object", "properties": {"job_id": {"description": "The ID of the cancelled job.", "type": "string"}, "message": {"description": "Cancellation confirmation message.", "type": "string"}}, "required": ["message", "job_id"]}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Job is already in final status and cannot be cancelled.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Crawl job not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.crawl", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
