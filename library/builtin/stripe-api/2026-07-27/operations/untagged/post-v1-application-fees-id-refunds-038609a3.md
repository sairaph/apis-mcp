---
title: Create an application fee refund
page_id: operation-post-v1-application-fees-id-refunds-9d06c768
path: operations/untagged
description: |-
    <p>Refunds an application fee that has previously been collected but not yet refunded.
    Funds will be refunded to the Stripe account from which the fee was originally collected.</p>

    <p>You can optionally refund only part of an application fee.
    You can do so multiple times, until the entire fee has been refunded.</p>

    <p>Once entirely refunded, an application fee can’t be refunded again.
    This method will raise an error when called on an already-refunded application fee,
    or when trying to refund more money than is left on an application fee.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/application_fees/{id}/refunds
operation_ids:
    - PostApplicationFeesIdRefunds
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an application fee refund

`POST /v1/application_fees/{id}/refunds`

Operation ID: `PostApplicationFeesIdRefunds`

<p>Refunds an application fee that has previously been collected but not yet refunded.
Funds will be refunded to the Stripe account from which the fee was originally collected.</p>

<p>You can optionally refund only part of an application fee.
You can do so multiple times, until the entire fee has been refunded.</p>

<p>Once entirely refunded, an application fee can’t be refunded again.
This method will raise an error when called on an already-refunded application fee,
or when trying to refund more money than is left on an application fee.</p>

## Definition

```yaml
{"summary": "Create an application fee refund", "description": "<p>Refunds an application fee that has previously been collected but not yet refunded.\nFunds will be refunded to the Stripe account from which the fee was originally collected.</p>\n\n<p>You can optionally refund only part of an application fee.\nYou can do so multiple times, until the entire fee has been refunded.</p>\n\n<p>Once entirely refunded, an application fee can’t be refunded again.\nThis method will raise an error when called on an already-refunded application fee,\nor when trying to refund more money than is left on an application fee.</p>", "operationId": "PostApplicationFeesIdRefunds", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer, in _cents (or local equivalent)_, representing how much of this fee to refund. Can refund only up to the remaining unrefunded amount of the fee."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/fee_refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
