---
title: Create a credit note
page_id: operation-post-v1-credit-notes-98c43b44
path: operations/untagged
description: |-
    <p>Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice’s <code>amount_remaining</code> (and <code>amount_due</code>), but not below zero.
    This amount is indicated by the credit note’s <code>pre_payment_amount</code>. The excess amount is indicated by <code>post_payment_amount</code>, and it can result in any combination of the following:</p>

    <ul>
    <li>Refunds: create a new refund (using <code>refund_amount</code>) or link existing refunds (using <code>refunds</code>).</li>
    <li>Customer balance credit: credit the customer’s balance (using <code>credit_amount</code>) which will be automatically applied to their next invoice when it’s finalized.</li>
    <li>Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using <code>out_of_band_amount</code>).</li>
    </ul>

    <p>The sum of refunds, customer balance credits, and outside of Stripe credits must equal the <code>post_payment_amount</code>.</p>

    <p>You may issue multiple credit notes for an invoice. Each credit note may increment the invoice’s <code>pre_payment_credit_notes_amount</code>,
    <code>post_payment_credit_notes_amount</code>, or both, depending on the invoice’s <code>amount_remaining</code> at the time of credit note creation.</p>

    <p>For invoices that also have refunds created through the <a href="/docs/api/refunds">Refund API</a>, the credit note API subtracts those refund amounts from the maximum creditable amount. This prevents the combined credit notes and refunds from exceeding the invoice amount. If you use both, ensure the combined total does not exceed the invoice’s paid amount.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/credit_notes
operation_ids:
    - PostCreditNotes
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a credit note

`POST /v1/credit_notes`

Operation ID: `PostCreditNotes`

<p>Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice’s <code>amount_remaining</code> (and <code>amount_due</code>), but not below zero.
This amount is indicated by the credit note’s <code>pre_payment_amount</code>. The excess amount is indicated by <code>post_payment_amount</code>, and it can result in any combination of the following:</p>

<ul>
<li>Refunds: create a new refund (using <code>refund_amount</code>) or link existing refunds (using <code>refunds</code>).</li>
<li>Customer balance credit: credit the customer’s balance (using <code>credit_amount</code>) which will be automatically applied to their next invoice when it’s finalized.</li>
<li>Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using <code>out_of_band_amount</code>).</li>
</ul>

<p>The sum of refunds, customer balance credits, and outside of Stripe credits must equal the <code>post_payment_amount</code>.</p>

<p>You may issue multiple credit notes for an invoice. Each credit note may increment the invoice’s <code>pre_payment_credit_notes_amount</code>,
<code>post_payment_credit_notes_amount</code>, or both, depending on the invoice’s <code>amount_remaining</code> at the time of credit note creation.</p>

<p>For invoices that also have refunds created through the <a href="/docs/api/refunds">Refund API</a>, the credit note API subtracts those refund amounts from the maximum creditable amount. This prevents the combined credit notes and refunds from exceeding the invoice amount. If you use both, ensure the combined total does not exceed the invoice’s paid amount.</p>

## Definition

```yaml
{"summary": "Create a credit note", "description": "<p>Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice’s <code>amount_remaining</code> (and <code>amount_due</code>), but not below zero.\nThis amount is indicated by the credit note’s <code>pre_payment_amount</code>. The excess amount is indicated by <code>post_payment_amount</code>, and it can result in any combination of the following:</p>\n\n<ul>\n<li>Refunds: create a new refund (using <code>refund_amount</code>) or link existing refunds (using <code>refunds</code>).</li>\n<li>Customer balance credit: credit the customer’s balance (using <code>credit_amount</code>) which will be automatically applied to their next invoice when it’s finalized.</li>\n<li>Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using <code>out_of_band_amount</code>).</li>\n</ul>\n\n<p>The sum of refunds, customer balance credits, and outside of Stripe credits must equal the <code>post_payment_amount</code>.</p>\n\n<p>You may issue multiple credit notes for an invoice. Each credit note may increment the invoice’s <code>pre_payment_credit_notes_amount</code>,\n<code>post_payment_credit_notes_amount</code>, or both, depending on the invoice’s <code>amount_remaining</code> at the time of credit note creation.</p>\n\n<p>For invoices that also have refunds created through the <a href=\"/docs/api/refunds\">Refund API</a>, the credit note API subtracts those refund amounts from the maximum creditable amount. This prevents the combined credit notes and refunds from exceeding the invoice amount. If you use both, ensure the combined total does not exceed the invoice’s paid amount.</p>", "operationId": "PostCreditNotes", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["invoice"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The integer amount in cents (or local equivalent) representing the total amount of the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided."}, "credit_amount": {"type": "integer", "description": "The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice."}, "effective_at": {"type": "integer", "description": "The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.", "format": "unix-time"}, "email_type": {"type": "string", "description": "Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.", "enum": ["credit_note", "none"]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "invoice": {"maxLength": 5000, "type": "string", "description": "ID of the invoice."}, "lines": {"type": "array", "description": "Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.", "items": {"title": "credit_note_line_item_params", "required": ["type"], "type": "object", "properties": {"amount": {"type": "integer"}, "description": {"maxLength": 5000, "type": "string"}, "invoice_line_item": {"maxLength": 5000, "type": "string"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}}, "quantity": {"type": "integer"}, "tax_amounts": {"anyOf": [{"type": "array", "items": {"title": "tax_amount_with_tax_rate_param", "required": ["amount", "tax_rate", "taxable_amount"], "type": "object", "properties": {"amount": {"type": "integer"}, "tax_rate": {"maxLength": 5000, "type": "string"}, "taxable_amount": {"type": "integer"}}}}, {"type": "string", "enum": [""]}]}, "tax_rates": {"anyOf": [{"type": "array", "items": {"maxLength": 5000, "type": "string"}}, {"type": "string", "enum": [""]}]}, "type": {"type": "string", "enum": ["custom_line_item", "invoice_line_item"]}, "unit_amount": {"type": "integer"}, "unit_amount_decimal": {"type": "string", "format": "decimal"}}}}, "memo": {"maxLength": 5000, "type": "string", "description": "The credit note's memo appears on the credit note PDF."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "out_of_band_amount": {"type": "integer", "description": "The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe."}, "reason": {"type": "string", "description": "Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`", "enum": ["duplicate", "fraudulent", "order_change", "product_unsatisfactory"], "x-stripeBypassValidation": true}, "refund_amount": {"type": "integer", "description": "The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice."}, "refunds": {"type": "array", "description": "Refunds to link to this credit note.", "items": {"title": "credit_note_refund_params", "type": "object", "properties": {"amount_refunded": {"type": "integer"}, "payment_record_refund": {"title": "payment_record_refund_params", "required": ["payment_record", "refund_group"], "type": "object", "properties": {"payment_record": {"maxLength": 5000, "type": "string"}, "refund_group": {"maxLength": 5000, "type": "string"}}}, "refund": {"type": "string"}, "type": {"type": "string", "enum": ["payment_record_refund", "refund"]}}}}, "shipping_cost": {"title": "credit_note_shipping_cost", "type": "object", "properties": {"shipping_rate": {"maxLength": 5000, "type": "string"}}, "description": "When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "lines": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "refunds": {"style": "deepObject", "explode": true}, "shipping_cost": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/credit_note"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
