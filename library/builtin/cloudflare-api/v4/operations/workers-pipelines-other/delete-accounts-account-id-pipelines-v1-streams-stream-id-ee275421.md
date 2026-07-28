---
title: Delete Stream
page_id: operation-delete-accounts-account-id-pipelines-v1-streams-stream-id-3393dd4d
path: operations/workers-pipelines-other
description: Delete Stream in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/streams/{stream_id}
operation_ids:
    - deleteV4AccountsByAccount_idPipelinesV1StreamsByStream_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Stream

`DELETE /accounts/{account_id}/pipelines/v1/streams/{stream_id}`

Operation ID: `deleteV4AccountsByAccount_idPipelinesV1StreamsByStream_id`

Delete Stream in Account.

## Definition

```yaml
{"operationId": "deleteV4AccountsByAccount_idPipelinesV1StreamsByStream_id", "summary": "Delete Stream", "description": "Delete Stream in Account.", "parameters": [{"name": "force", "in": "query", "schema": {"description": "Deprecated: Delete stream forcefully, including deleting any dependent pipelines.", "type": "string"}, "deprecated": true}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "stream_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-stream-id"}}], "responses": {"200": {"description": "Indicates a successfully deleted Stream.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object"}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in listing Streams."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.streams", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
