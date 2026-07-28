---
title: Update an Account Capability
page_id: operation-post-v1-accounts-account-capabilities-capability-668a1932
path: operations/untagged
description: <p>Updates an existing Account Capability. Request or remove a capability by updating its <code>requested</code> parameter.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/accounts/{account}/capabilities/{capability}
operation_ids:
    - PostAccountsAccountCapabilitiesCapability
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update an Account Capability

`POST /v1/accounts/{account}/capabilities/{capability}`

Operation ID: `PostAccountsAccountCapabilitiesCapability`

<p>Updates an existing Account Capability. Request or remove a capability by updating its <code>requested</code> parameter.</p>

## Definition

```yaml
{"summary": "Update an Account Capability", "description": "<p>Updates an existing Account Capability. Request or remove a capability by updating its <code>requested</code> parameter.</p>", "operationId": "PostAccountsAccountCapabilitiesCapability", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "capability", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "requested": {"type": "boolean", "description": "To request a new capability for an account, pass true. There can be a delay before the requested capability becomes active. If the capability has any activation requirements, the response includes them in the `requirements` arrays.\n\nIf a capability isn't permanent, you can remove it from the account by passing false. Some capabilities are permanent after they've been requested. Attempting to remove a permanent capability returns an error."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/capability"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
