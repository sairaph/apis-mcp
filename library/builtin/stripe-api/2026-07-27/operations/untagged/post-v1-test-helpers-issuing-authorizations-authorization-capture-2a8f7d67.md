---
title: Capture a test-mode authorization
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-capture-257d3c13
path: operations/untagged
description: <p>Capture a test-mode authorization.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/capture
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationCapture
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Capture a test-mode authorization

`POST /v1/test_helpers/issuing/authorizations/{authorization}/capture`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationCapture`

<p>Capture a test-mode authorization.</p>

## Definition

```yaml
{"summary": "Capture a test-mode authorization", "description": "<p>Capture a test-mode authorization.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationCapture", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"capture_amount": {"type": "integer", "description": "The amount to capture from the authorization. If not provided, the full amount of the authorization will be captured. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "close_authorization": {"type": "boolean", "description": "Whether to close the authorization after capture. Defaults to true. Set to false to enable multi-capture flows."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "purchase_details": {"title": "purchase_details_specs", "type": "object", "properties": {"fleet": {"title": "fleet_specs", "type": "object", "properties": {"cardholder_prompt_data": {"title": "fleet_cardholder_prompt_data_specs", "type": "object", "properties": {"driver_id": {"maxLength": 5000, "type": "string"}, "odometer": {"type": "integer"}, "unspecified_id": {"maxLength": 5000, "type": "string"}, "user_id": {"maxLength": 5000, "type": "string"}, "vehicle_number": {"maxLength": 5000, "type": "string"}}}, "purchase_type": {"maxLength": 5000, "type": "string", "enum": ["fuel_and_non_fuel_purchase", "fuel_purchase", "non_fuel_purchase"]}, "reported_breakdown": {"title": "fleet_reported_breakdown_specs", "type": "object", "properties": {"fuel": {"title": "fleet_reported_breakdown_fuel_specs", "type": "object", "properties": {"gross_amount_decimal": {"type": "string", "format": "decimal"}}}, "non_fuel": {"title": "fleet_reported_breakdown_non_fuel_specs", "type": "object", "properties": {"gross_amount_decimal": {"type": "string", "format": "decimal"}}}, "tax": {"title": "fleet_reported_breakdown_tax_specs", "type": "object", "properties": {"local_amount_decimal": {"type": "string", "format": "decimal"}, "national_amount_decimal": {"type": "string", "format": "decimal"}}}}}, "service_type": {"maxLength": 5000, "type": "string", "enum": ["full_service", "non_fuel_transaction", "self_service"]}}}, "flight": {"title": "flight_specs", "type": "object", "properties": {"departure_at": {"type": "integer", "format": "unix-time"}, "passenger_name": {"maxLength": 5000, "type": "string"}, "refundable": {"type": "boolean"}, "segments": {"type": "array", "items": {"title": "flight_segment_specs", "type": "object", "properties": {"arrival_airport_code": {"maxLength": 3, "type": "string"}, "carrier": {"maxLength": 5000, "type": "string"}, "departure_airport_code": {"maxLength": 3, "type": "string"}, "flight_number": {"maxLength": 5000, "type": "string"}, "service_class": {"maxLength": 5000, "type": "string"}, "stopover_allowed": {"type": "boolean"}}}}, "travel_agency": {"maxLength": 5000, "type": "string"}}}, "fuel": {"title": "fuel_specs", "type": "object", "properties": {"industry_product_code": {"maxLength": 5000, "type": "string"}, "quantity_decimal": {"type": "string", "format": "decimal"}, "type": {"maxLength": 5000, "type": "string", "enum": ["diesel", "other", "unleaded_plus", "unleaded_regular", "unleaded_super"]}, "unit": {"maxLength": 5000, "type": "string", "enum": ["charging_minute", "imperial_gallon", "kilogram", "kilowatt_hour", "liter", "other", "pound", "us_gallon"]}, "unit_cost_decimal": {"type": "string", "format": "decimal"}}}, "lodging": {"title": "lodging_specs", "type": "object", "properties": {"check_in_at": {"type": "integer", "format": "unix-time"}, "nights": {"type": "integer"}}}, "receipt": {"type": "array", "items": {"title": "receipt_specs", "type": "object", "properties": {"description": {"maxLength": 26, "type": "string"}, "quantity": {"type": "string", "format": "decimal"}, "total": {"type": "integer"}, "unit_cost": {"type": "integer"}}}}, "reference": {"maxLength": 5000, "type": "string"}}, "description": "Additional purchase information that is optionally provided by the merchant."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "purchase_details": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
