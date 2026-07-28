---
title: Retrieve an application fee
page_id: operation-get-v1-application-fees-id-0253af62
path: operations/untagged
description: <p>Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/application_fees/{id}
operation_ids:
    - GetApplicationFeesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an application fee

`GET /v1/application_fees/{id}`

Operation ID: `GetApplicationFeesId`

<p>Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.</p>

## Definition

```yaml
{"summary": "Retrieve an application fee", "description": "<p>Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.</p>", "operationId": "GetApplicationFeesId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/application_fee"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
