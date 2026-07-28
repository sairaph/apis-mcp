---
title: Retrieve a SetupIntent
page_id: operation-get-v1-setup-intents-intent-649ee020
path: operations/untagged
description: |-
    <p>Retrieves the details of a SetupIntent that has previously been created. </p>

    <p>Client-side retrieval using a publishable key is allowed when the <code>client_secret</code> is provided in the query string. </p>

    <p>When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the <a href="#setup_intent_object">SetupIntent</a> object reference for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/setup_intents/{intent}
operation_ids:
    - GetSetupIntentsIntent
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a SetupIntent

`GET /v1/setup_intents/{intent}`

Operation ID: `GetSetupIntentsIntent`

<p>Retrieves the details of a SetupIntent that has previously been created. </p>

<p>Client-side retrieval using a publishable key is allowed when the <code>client_secret</code> is provided in the query string. </p>

<p>When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the <a href="#setup_intent_object">SetupIntent</a> object reference for more details.</p>

## Definition

```yaml
{"summary": "Retrieve a SetupIntent", "description": "<p>Retrieves the details of a SetupIntent that has previously been created. </p>\n\n<p>Client-side retrieval using a publishable key is allowed when the <code>client_secret</code> is provided in the query string. </p>\n\n<p>When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the <a href=\"#setup_intent_object\">SetupIntent</a> object reference for more details.</p>", "operationId": "GetSetupIntentsIntent", "parameters": [{"name": "client_secret", "in": "query", "description": "The client secret of the SetupIntent. We require this string if you use a publishable key to retrieve the SetupIntent.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/setup_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
