---
title: Retrieve an Account Capability
page_id: operation-get-v1-accounts-account-capabilities-capability-e39cab8f
path: operations/untagged
description: <p>Retrieves information about the specified Account Capability.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/accounts/{account}/capabilities/{capability}
operation_ids:
    - GetAccountsAccountCapabilitiesCapability
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an Account Capability

`GET /v1/accounts/{account}/capabilities/{capability}`

Operation ID: `GetAccountsAccountCapabilitiesCapability`

<p>Retrieves information about the specified Account Capability.</p>

## Definition

```yaml
{"summary": "Retrieve an Account Capability", "description": "<p>Retrieves information about the specified Account Capability.</p>", "operationId": "GetAccountsAccountCapabilitiesCapability", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "capability", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/capability"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
