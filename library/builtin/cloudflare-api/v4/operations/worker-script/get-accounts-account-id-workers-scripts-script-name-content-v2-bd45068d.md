---
title: Get script content
page_id: operation-get-accounts-account-id-workers-scripts-script-name-content-v2-e109ee51
path: operations/worker-script
description: Fetch script content only.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/content/v2
operation_ids:
    - worker-script-get-content
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get script content

`GET /accounts/{account_id}/workers/scripts/{script_name}/content/v2`

Operation ID: `worker-script-get-content`

Fetch script content only.

## Definition

```yaml
{"operationId": "worker-script-get-content", "summary": "Get script content", "description": "Fetch script content only.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Fetch script content.", "content": {"string": {"schema": {"type": "string", "example": "export default {\n  async fetch(request, env, ctx) {\n    return new Response(\"Hello, world!\");\n  }\n};\n"}}}}, "4XX": {"description": "Fetch script content failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.content", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
