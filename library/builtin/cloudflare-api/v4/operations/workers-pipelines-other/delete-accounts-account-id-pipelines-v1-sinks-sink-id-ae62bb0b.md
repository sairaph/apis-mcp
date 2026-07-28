---
title: Delete Sink
page_id: operation-delete-accounts-account-id-pipelines-v1-sinks-sink-id-b02091f3
path: operations/workers-pipelines-other
description: Delete Pipeline in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/sinks/{sink_id}
operation_ids:
    - deleteV4AccountsByAccount_idPipelinesV1SinksBySink_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Sink

`DELETE /accounts/{account_id}/pipelines/v1/sinks/{sink_id}`

Operation ID: `deleteV4AccountsByAccount_idPipelinesV1SinksBySink_id`

Delete Pipeline in Account.

## Definition

```yaml
{"operationId": "deleteV4AccountsByAccount_idPipelinesV1SinksBySink_id", "summary": "Delete Sink", "description": "Delete Pipeline in Account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "sink_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-sink-id"}}, {"name": "force", "in": "query", "schema": {"description": "Deprecated: Delete sink forcefully, including deleting any dependent pipelines.", "type": "string"}, "deprecated": true}], "responses": {"200": {"description": "Indicates a successfully deleted Sink.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object"}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in listing Sinks."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.sinks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
