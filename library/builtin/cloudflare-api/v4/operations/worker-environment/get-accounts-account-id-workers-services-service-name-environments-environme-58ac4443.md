---
title: Get script content
page_id: operation-get-accounts-account-id-workers-services-service-name-environments-envir-29840f09
path: operations/worker-environment
description: Get script content from a worker with an environment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/content
operation_ids:
    - worker-environment-get-script-content
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get script content

`GET /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/content`

Operation ID: `worker-environment-get-script-content`

Get script content from a worker with an environment.

## Definition

```yaml
{"operationId": "worker-environment-get-script-content", "summary": "Get script content", "description": "Get script content from a worker with an environment.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "service_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_service"}}, {"name": "environment_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_environment"}}], "responses": {"200": {"description": "Get script content.", "content": {"string": {"schema": {"type": "string", "example": "export default {\n  async fetch(request, env, ctx) {\n    return new Response(\"Hello, world!\");\n  }\n};\n"}}}}, "4XX": {"description": "Get script content failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Environment"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.services.environments.content", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
