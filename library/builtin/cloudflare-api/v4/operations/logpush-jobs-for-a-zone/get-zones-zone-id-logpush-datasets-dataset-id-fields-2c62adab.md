---
title: List fields
page_id: operation-get-zones-zone-id-logpush-datasets-dataset-id-fields-cd599b9c
path: operations/logpush-jobs-for-a-zone
description: Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logpush/datasets/{dataset_id}/fields
operation_ids:
    - get-zones-zone_id-logpush-datasets-dataset_id-fields
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List fields

`GET /zones/{zone_id}/logpush/datasets/{dataset_id}/fields`

Operation ID: `get-zones-zone_id-logpush-datasets-dataset_id-fields`

Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logpush-datasets-dataset_id-fields", "summary": "List fields", "description": "Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.", "parameters": [{"name": "dataset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_dataset"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "responses": {"200": {"description": "List fields response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_logpush_field_response_collection"}}}}, "4XX": {"description": "List fields response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for a zone"], "x-api-token-group": ["Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.datasets.fields", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
