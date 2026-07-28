---
title: List fields
page_id: operation-get-accounts-account-id-logpush-datasets-dataset-id-fields-446f6524
path: operations/logpush-jobs-for-an-account
description: Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logpush/datasets/{dataset_id}/fields
operation_ids:
    - get-accounts-account_id-logpush-datasets-dataset_id-fields
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List fields

`GET /accounts/{account_id}/logpush/datasets/{dataset_id}/fields`

Operation ID: `get-accounts-account_id-logpush-datasets-dataset_id-fields`

Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.

## Definition

```yaml
{"operationId": "get-accounts-account_id-logpush-datasets-dataset_id-fields", "summary": "List fields", "description": "Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.", "parameters": [{"name": "dataset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_dataset"}, "example": "gateway_dns"}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "responses": {"200": {"description": "List fields response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_logpush_field_response_collection"}}}}, "4XX": {"description": "List fields response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-datasets.fields", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
