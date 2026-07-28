---
title: Patch Latest Version
page_id: operation-patch-accounts-account-id-workers-workers-worker-id-versions-latest-19550733
path: operations/versions
description: Only `/versions/latest` is supported. Creates a new version by applying a JSON Merge Patch (RFC 7396) to the latest version. Patching a specific version ID is not supported. Omitted fields are inherited from the latest version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}/versions/latest
operation_ids:
    - patchLatestWorkerVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Latest Version

`PATCH /accounts/{account_id}/workers/workers/{worker_id}/versions/latest`

Operation ID: `patchLatestWorkerVersion`

Only `/versions/latest` is supported. Creates a new version by applying a JSON Merge Patch (RFC 7396) to the latest version. Patching a specific version ID is not supported. Omitted fields are inherited from the latest version.

## Definition

```yaml
{"operationId": "patchLatestWorkerVersion", "summary": "Patch Latest Version", "description": "Only `/versions/latest` is supported. Creates a new version by applying a JSON Merge Patch (RFC 7396) to the latest version. Patching a specific version ID is not supported. Omitted fields are inherited from the latest version.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}, {"name": "deploy", "in": "query", "schema": {"description": "If true, a deployment will be created that sends 100% of traffic to the new version.", "type": "boolean"}}], "requestBody": {"description": "JSON Merge Patch to apply to the latest version. An empty object creates a new version with no changes.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Version"}, {"type": "object"}]}}, "application/merge-patch+json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Version"}, {"type": "object"}]}}}}, "responses": {"200": {"description": "Patch version success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Version"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Patch version failure. When the patch uses the declarative\n`exports` field and one or more entries fail reconciliation,\nthe response is the exports reconciliation error envelope\n(error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Versions"]}
```
