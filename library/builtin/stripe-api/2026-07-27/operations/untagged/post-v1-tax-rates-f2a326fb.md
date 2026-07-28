---
title: Create a tax rate
page_id: operation-post-v1-tax-rates-147d9baf
path: operations/untagged
description: <p>Creates a new tax rate.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/tax_rates
operation_ids:
    - PostTaxRates
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a tax rate

`POST /v1/tax_rates`

Operation ID: `PostTaxRates`

<p>Creates a new tax rate.</p>

## Definition

```yaml
{"summary": "Create a tax rate", "description": "<p>Creates a new tax rate.</p>", "operationId": "PostTaxRates", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["display_name", "inclusive", "percentage"], "type": "object", "properties": {"active": {"type": "boolean", "description": "Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set."}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2))."}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers."}, "display_name": {"maxLength": 50, "type": "string", "description": "The display name of the tax rate, which will be shown to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "inclusive": {"type": "boolean", "description": "This specifies if the tax rate is inclusive or exclusive."}, "jurisdiction": {"maxLength": 50, "type": "string", "description": "The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "percentage": {"type": "number", "description": "This represents the tax rate percent out of 100."}, "state": {"maxLength": 5000, "type": "string", "description": "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2), without country prefix. For example, \"NY\" for New York, United States."}, "tax_type": {"type": "string", "description": "The high-level tax type, such as `vat` or `sales_tax`.", "enum": ["amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat"], "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax_rate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
