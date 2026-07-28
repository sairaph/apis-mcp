---
title: Retrieve a Country Spec
page_id: operation-get-v1-country-specs-country-b3fc853f
path: operations/untagged
description: <p>Returns a Country Spec for a given Country code.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/country_specs/{country}
operation_ids:
    - GetCountrySpecsCountry
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Country Spec

`GET /v1/country_specs/{country}`

Operation ID: `GetCountrySpecsCountry`

<p>Returns a Country Spec for a given Country code.</p>

## Definition

```yaml
{"summary": "Retrieve a Country Spec", "description": "<p>Returns a Country Spec for a given Country code.</p>", "operationId": "GetCountrySpecsCountry", "parameters": [{"name": "country", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/country_spec"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
