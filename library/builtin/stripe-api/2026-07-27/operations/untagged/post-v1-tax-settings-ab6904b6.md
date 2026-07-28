---
title: Update settings
page_id: operation-post-v1-tax-settings-8f9a0b02
path: operations/untagged
description: <p>Updates Tax <code>Settings</code> parameters used in tax calculations. All parameters are editable but none can be removed once set.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/tax/settings
operation_ids:
    - PostTaxSettings
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update settings

`POST /v1/tax/settings`

Operation ID: `PostTaxSettings`

<p>Updates Tax <code>Settings</code> parameters used in tax calculations. All parameters are editable but none can be removed once set.</p>

## Definition

```yaml
{"summary": "Update settings", "description": "<p>Updates Tax <code>Settings</code> parameters used in tax calculations. All parameters are editable but none can be removed once set.</p>", "operationId": "PostTaxSettings", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"defaults": {"title": "defaults_param", "type": "object", "properties": {"tax_behavior": {"type": "string", "enum": ["exclusive", "inclusive", "inferred_by_currency"]}, "tax_code": {"type": "string"}}, "description": "Default configuration to be used on Stripe Tax calculations."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "head_office": {"title": "head_office_param", "required": ["address"], "type": "object", "properties": {"address": {"title": "validated_country_address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}}, "description": "The place where your business is located."}}, "additionalProperties": false}, "encoding": {"defaults": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "head_office": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.settings"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
