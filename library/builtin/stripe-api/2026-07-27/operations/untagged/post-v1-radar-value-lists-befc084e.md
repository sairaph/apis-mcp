---
title: Create a value list
page_id: operation-post-v1-radar-value-lists-41ab6f08
path: operations/untagged
description: <p>Creates a new <code>ValueList</code> object, which can then be referenced in rules.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/radar/value_lists
operation_ids:
    - PostRadarValueLists
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a value list

`POST /v1/radar/value_lists`

Operation ID: `PostRadarValueLists`

<p>Creates a new <code>ValueList</code> object, which can then be referenced in rules.</p>

## Definition

```yaml
{"summary": "Create a value list", "description": "<p>Creates a new <code>ValueList</code> object, which can then be referenced in rules.</p>", "operationId": "PostRadarValueLists", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["alias", "name"], "type": "object", "properties": {"alias": {"maxLength": 100, "type": "string", "description": "The name of the value list for use in rules."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "item_type": {"maxLength": 5000, "type": "string", "description": "Type of the items in the value list. One of `card_fingerprint`, `card_bin`, `crypto_fingerprint`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, `customer_id`, `account`, `sepa_debit_fingerprint`, or `us_bank_account_fingerprint`. Use `string` if the item type is unknown or mixed.", "enum": ["account", "card_bin", "card_fingerprint", "case_sensitive_string", "country", "crypto_fingerprint", "customer_id", "email", "ip_address", "sepa_debit_fingerprint", "string", "us_bank_account_fingerprint"]}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "name": {"maxLength": 100, "type": "string", "description": "The human-readable name of the value list."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.value_list"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
