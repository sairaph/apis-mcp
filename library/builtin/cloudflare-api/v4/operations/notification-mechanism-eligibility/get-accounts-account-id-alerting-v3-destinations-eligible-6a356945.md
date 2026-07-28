---
title: Get delivery mechanism eligibility
page_id: operation-get-accounts-account-id-alerting-v3-destinations-eligible-e5741e03
path: operations/notification-mechanism-eligibility
description: Get a list of all delivery mechanism types for which an account is eligible.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/eligible
operation_ids:
    - notification-mechanism-eligibility-get-delivery-mechanism-eligibility
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get delivery mechanism eligibility

`GET /accounts/{account_id}/alerting/v3/destinations/eligible`

Operation ID: `notification-mechanism-eligibility-get-delivery-mechanism-eligibility`

Get a list of all delivery mechanism types for which an account is eligible.

## Definition

```yaml
{"operationId": "notification-mechanism-eligibility-get-delivery-mechanism-eligibility", "summary": "Get delivery mechanism eligibility", "description": "Get a list of all delivery mechanism types for which an account is eligible.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "Get delivery mechanism eligibility response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_schemas-response_collection"}}}}, "4XX": {"description": "Get delivery mechanism eligibility response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Mechanism Eligibility"]}
```
