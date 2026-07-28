---
title: Update a charge
page_id: operation-post-v1-charges-charge-1535ea0f
path: operations/untagged
description: <p>Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}
operation_ids:
    - PostChargesCharge
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a charge

`POST /v1/charges/{charge}`

Operation ID: `PostChargesCharge`

<p>Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>

## Definition

```yaml
{"summary": "Update a charge", "description": "<p>Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>", "operationId": "PostChargesCharge", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "The ID of an existing customer that will be associated with this request. This field may only be updated if there is no existing associated customer with this charge."}, "description": {"maxLength": 40000, "type": "string", "description": "An arbitrary string which you can attach to a charge object. It is displayed when in the web interface alongside the charge. Note that if you use Stripe to send automatic email receipts to your customers, your receipt emails will include the `description` of the charge(s) that they are describing."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "fraud_details": {"title": "fraud_details", "required": ["user_report"], "type": "object", "properties": {"user_report": {"maxLength": 5000, "type": "string", "enum": ["", "fraudulent", "safe"]}}, "description": "A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms."}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "receipt_email": {"maxLength": 5000, "type": "string", "description": "This is the email address that the receipt for this charge will be sent to. If this field is updated, then a new email receipt will be sent to the updated address."}, "shipping": {"title": "optional_fields_shipping", "required": ["address", "name"], "type": "object", "properties": {"address": {"title": "optional_fields_address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "carrier": {"maxLength": 5000, "type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"maxLength": 5000, "type": "string"}, "tracking_number": {"maxLength": 5000, "type": "string"}}, "description": "Shipping information for the charge. Helps prevent fraud on charges for physical goods."}, "transfer_group": {"type": "string", "description": "A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "fraud_details": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "shipping": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/charge"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
