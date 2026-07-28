---
title: List all ReceivedCredits
page_id: operation-get-v1-treasury-received-credits-2b7b4fa1
path: operations/untagged
description: <p>Returns a list of ReceivedCredits.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/treasury/received_credits
operation_ids:
    - GetTreasuryReceivedCredits
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all ReceivedCredits

`GET /v1/treasury/received_credits`

Operation ID: `GetTreasuryReceivedCredits`

<p>Returns a list of ReceivedCredits.</p>

## Definition

```yaml
{"summary": "List all ReceivedCredits", "description": "<p>Returns a list of ReceivedCredits.</p>", "operationId": "GetTreasuryReceivedCredits", "parameters": [{"name": "ending_before", "in": "query", "description": "A cursor for use in pagination. `ending_before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with `obj_bar`, your subsequent call can include `ending_before=obj_bar` in order to fetch the previous page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "financial_account", "in": "query", "description": "The FinancialAccount that received the funds.", "required": true, "style": "form", "explode": true, "schema": {"type": "string"}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "linked_flows", "in": "query", "description": "Only return ReceivedCredits described by the flow.", "required": false, "style": "deepObject", "explode": true, "schema": {"title": "linked_flows_param", "required": ["source_flow_type"], "type": "object", "properties": {"source_flow_type": {"type": "string", "enum": ["credit_reversal", "other", "outbound_payment", "outbound_transfer", "payout"]}}}}, {"name": "starting_after", "in": "query", "description": "A cursor for use in pagination. `starting_after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with `obj_foo`, your subsequent call can include `starting_after=obj_foo` in order to fetch the next page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "status", "in": "query", "description": "Only return ReceivedCredits that have the given status: `succeeded` or `failed`.", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["failed", "succeeded"], "x-stripeBypassValidation": true}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "TreasuryReceivedCreditsResourceReceivedCreditList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/treasury.received_credit"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
