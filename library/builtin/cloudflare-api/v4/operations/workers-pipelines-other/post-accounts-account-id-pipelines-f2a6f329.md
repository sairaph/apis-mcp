---
title: '[DEPRECATED] Create Pipeline'
page_id: operation-post-accounts-account-id-pipelines-6926bba6
path: operations/workers-pipelines-other
description: '[DEPRECATED] Create a new pipeline. Use the new /pipelines/v1/pipelines endpoint instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pipelines
operation_ids:
    - postV4AccountsByAccount_idPipelines_deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# [DEPRECATED] Create Pipeline

`POST /accounts/{account_id}/pipelines`

Operation ID: `postV4AccountsByAccount_idPipelines_deprecated`

[DEPRECATED] Create a new pipeline. Use the new /pipelines/v1/pipelines endpoint instead.

## Definition

```yaml
{"operationId": "postV4AccountsByAccount_idPipelines_deprecated", "summary": "[DEPRECATED] Create Pipeline", "description": "[DEPRECATED] Create a new pipeline. Use the new /pipelines/v1/pipelines endpoint instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"destination": {"type": "object", "properties": {"batch": {"type": "object", "properties": {"max_bytes": {"description": "Specifies rough maximum size of files.", "type": "integer", "default": 100000000, "maximum": 100000000, "minimum": 1000}, "max_duration_s": {"description": "Specifies duration to wait to aggregate batches files.", "type": "number", "default": 300, "maximum": 300, "minimum": 0.25}, "max_rows": {"description": "Specifies rough maximum number of rows per file.", "type": "integer", "default": 10000000, "maximum": 10000000, "minimum": 100}}}, "compression": {"type": "object", "properties": {"type": {"description": "Specifies the desired compression algorithm and format.", "type": "string", "example": "gzip", "default": "gzip", "enum": ["none", "gzip", "deflate"]}}}, "credentials": {"type": "object", "properties": {"access_key_id": {"description": "Specifies the R2 Bucket Access Key Id.", "type": "string", "example": "<access key id>"}, "endpoint": {"description": "Specifies the R2 Endpoint.", "type": "string", "example": "https://123f8a8258064ed892a347f173372359.r2.cloudflarestorage.com"}, "secret_access_key": {"description": "Specifies the R2 Bucket Secret Access Key.", "type": "string", "example": "<secret key>"}}, "required": ["endpoint", "access_key_id", "secret_access_key"]}, "format": {"description": "Specifies the format of data to deliver.", "type": "string", "enum": ["json"]}, "path": {"type": "object", "example": {"bucket": "bucket", "prefix": "base"}, "properties": {"bucket": {"description": "Specifies the R2 Bucket to store files.", "type": "string", "example": "bucket"}, "filename": {"description": "Specifies the name pattern to for individual data files.", "example": "${slug}${extension}", "allOf": [{"type": "string"}, {"type": "string"}]}, "filepath": {"description": "Specifies the name pattern for directory.", "type": "string", "example": "${date}/${hour}"}, "prefix": {"description": "Specifies the base directory within the bucket.", "type": "string", "example": "base"}}, "required": ["bucket"]}, "type": {"description": "Specifies the type of destination.", "type": "string", "enum": ["r2"]}}, "required": ["type", "format", "batch", "compression", "path", "credentials"]}, "name": {"description": "Defines the name of the pipeline.", "type": "string", "example": "sample_pipeline", "maxLength": 128, "minLength": 1}, "source": {"type": "array", "items": {"discriminator": {"mapping": {"binding": "#/components/schemas/cloudflare-pipelines_workers_pipelines_binding_source", "http": "#/components/schemas/cloudflare-pipelines_workers_pipelines_http_source"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_workers_pipelines_http_source"}, {"$ref": "#/components/schemas/cloudflare-pipelines_workers_pipelines_binding_source"}]}, "minItems": 1}}, "required": ["name", "source", "destination"]}}}}, "responses": {"200": {"description": "[DEPRECATED] Indicates a successfully created pipeline. Use /pipelines/v1/pipelines instead.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline"}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in creating a pipeline.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}, "example": []}, "results": {"type": "object", "nullable": true, "x-stainless-empty-object": true}, "success": {"example": false, "allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}]}}, "required": ["success", "results", "errors"]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
