---
title: Post Worker subdomain
page_id: operation-post-accounts-account-id-workers-scripts-script-name-subdomain-567c9177
path: operations/worker-script
description: Enable or disable the Worker on the workers.dev subdomain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/subdomain
operation_ids:
    - worker-script-post-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Post Worker subdomain

`POST /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Operation ID: `worker-script-post-subdomain`

Enable or disable the Worker on the workers.dev subdomain.

## Definition

```yaml
{"operationId": "worker-script-post-subdomain", "summary": "Post Worker subdomain", "description": "Enable or disable the Worker on the workers.dev subdomain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"enabled": {"description": "Whether the Worker should be available on the workers.dev subdomain.", "type": "boolean", "example": true, "x-auditable": true}, "previews_enabled": {"description": "Whether the Worker's Preview URLs should be available on the workers.dev subdomain.", "type": "boolean", "example": false, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["enabled"]}}}}, "responses": {"200": {"description": "Post subdomain response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_subdomain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Post subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.subdomain", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
