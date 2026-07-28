---
title: Upload a Finetune Asset
page_id: operation-post-accounts-account-id-ai-finetunes-finetune-id-finetune-assets-542afad9
path: operations/workers-ai-finetune
description: Uploads training data assets for a Workers AI fine-tuning job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai/finetunes/{finetune_id}/finetune-assets
operation_ids:
    - workers-ai-upload-finetune-asset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a Finetune Asset

`POST /accounts/{account_id}/ai/finetunes/{finetune_id}/finetune-assets`

Operation ID: `workers-ai-upload-finetune-asset`

Uploads training data assets for a Workers AI fine-tuning job.

## Definition

```yaml
{"operationId": "workers-ai-upload-finetune-asset", "summary": "Upload a Finetune Asset", "description": "Uploads training data assets for a Workers AI fine-tuning job.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "finetune_id", "in": "path", "required": true, "schema": {"type": "string", "example": "bc451aef-f723-4b26-a6b2-901afd2e7a8a"}}], "requestBody": {"description": "Finetune asset file upload", "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"description": "File to upload", "type": "string", "format": "binary"}, "file_name": {"description": "Name of the file (adapter_config.json or adapter_model.safetensors)", "type": "string"}}, "required": ["file_name", "file"]}}}}, "responses": {"200": {"description": "Returns successfully if finetunes were uploaded", "content": {"application/json": {"schema": {"type": "object", "properties": {"success": {"type": "boolean"}}, "required": ["success"]}}}}, "400": {"description": "Finetune creation failed", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI Finetune"], "x-api-token-group": ["Workers AI Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
