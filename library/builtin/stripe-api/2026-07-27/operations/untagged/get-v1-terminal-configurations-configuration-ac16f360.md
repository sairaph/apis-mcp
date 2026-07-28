---
title: Retrieve a Configuration
page_id: operation-get-v1-terminal-configurations-configuration-14ef7b9c
path: operations/untagged
description: <p>Retrieves a <code>Configuration</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/terminal/configurations/{configuration}
operation_ids:
    - GetTerminalConfigurationsConfiguration
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Configuration

`GET /v1/terminal/configurations/{configuration}`

Operation ID: `GetTerminalConfigurationsConfiguration`

<p>Retrieves a <code>Configuration</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a Configuration", "description": "<p>Retrieves a <code>Configuration</code> object.</p>", "operationId": "GetTerminalConfigurationsConfiguration", "parameters": [{"name": "configuration", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/terminal.configuration"}, {"$ref": "#/components/schemas/deleted_terminal.configuration"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
