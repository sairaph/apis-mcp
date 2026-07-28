---
title: Finalize a test-mode authorization's amount
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-finalize-amoun-836e0536
path: operations/untagged
description: <p>Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/finalize_amount
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationFinalizeAmount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Finalize a test-mode authorization's amount

`POST /v1/test_helpers/issuing/authorizations/{authorization}/finalize_amount`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationFinalizeAmount`

<p>Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.</p>

## Definition

```yaml
{"summary": "Finalize a test-mode authorization's amount", "description": "<p>Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationFinalizeAmount", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["final_amount"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "final_amount": {"type": "integer", "description": "The final authorization amount that will be captured by the merchant. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "fleet": {"title": "fleet_specs", "type": "object", "properties": {"cardholder_prompt_data": {"title": "fleet_cardholder_prompt_data_specs", "type": "object", "properties": {"driver_id": {"maxLength": 5000, "type": "string"}, "odometer": {"type": "integer"}, "unspecified_id": {"maxLength": 5000, "type": "string"}, "user_id": {"maxLength": 5000, "type": "string"}, "vehicle_number": {"maxLength": 5000, "type": "string"}}}, "purchase_type": {"maxLength": 5000, "type": "string", "enum": ["fuel_and_non_fuel_purchase", "fuel_purchase", "non_fuel_purchase"]}, "reported_breakdown": {"title": "fleet_reported_breakdown_specs", "type": "object", "properties": {"fuel": {"title": "fleet_reported_breakdown_fuel_specs", "type": "object", "properties": {"gross_amount_decimal": {"type": "string", "format": "decimal"}}}, "non_fuel": {"title": "fleet_reported_breakdown_non_fuel_specs", "type": "object", "properties": {"gross_amount_decimal": {"type": "string", "format": "decimal"}}}, "tax": {"title": "fleet_reported_breakdown_tax_specs", "type": "object", "properties": {"local_amount_decimal": {"type": "string", "format": "decimal"}, "national_amount_decimal": {"type": "string", "format": "decimal"}}}}}, "service_type": {"maxLength": 5000, "type": "string", "enum": ["full_service", "non_fuel_transaction", "self_service"]}}, "description": "Fleet-specific information for authorizations using Fleet cards."}, "fuel": {"title": "fuel_specs", "type": "object", "properties": {"industry_product_code": {"maxLength": 5000, "type": "string"}, "quantity_decimal": {"type": "string", "format": "decimal"}, "type": {"maxLength": 5000, "type": "string", "enum": ["diesel", "other", "unleaded_plus", "unleaded_regular", "unleaded_super"]}, "unit": {"maxLength": 5000, "type": "string", "enum": ["charging_minute", "imperial_gallon", "kilogram", "kilowatt_hour", "liter", "other", "pound", "us_gallon"]}, "unit_cost_decimal": {"type": "string", "format": "decimal"}}, "description": "Information about fuel that was purchased with this transaction."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "fleet": {"style": "deepObject", "explode": true}, "fuel": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
