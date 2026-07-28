---
title: Cancel a subscription
page_id: operation-delete-v1-subscriptions-subscription-exposed-id-86d8f323
path: operations/untagged
description: |-
    <p>Cancels a customer’s subscription immediately. The customer won’t be charged again for the subscription. After it’s canceled, the subscription is largely immutable. You can still update its <a href="/metadata">metadata</a> and <code>cancellation_details</code>.</p>

    <p>Any pending invoice items that you’ve created are still charged at the end of the period, unless manually <a href="/api/invoiceitems/delete">deleted</a>. If you’ve set the subscription to cancel at the end of the period, any pending prorations are also left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations are removed if <code>invoice_now</code> and <code>prorate</code> are both set to false.</p>

    <p>By default, upon subscription cancellation, Stripe stops automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/subscriptions/{subscription_exposed_id}
operation_ids:
    - DeleteSubscriptionsSubscriptionExposedId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel a subscription

`DELETE /v1/subscriptions/{subscription_exposed_id}`

Operation ID: `DeleteSubscriptionsSubscriptionExposedId`

<p>Cancels a customer’s subscription immediately. The customer won’t be charged again for the subscription. After it’s canceled, the subscription is largely immutable. You can still update its <a href="/metadata">metadata</a> and <code>cancellation_details</code>.</p>

<p>Any pending invoice items that you’ve created are still charged at the end of the period, unless manually <a href="/api/invoiceitems/delete">deleted</a>. If you’ve set the subscription to cancel at the end of the period, any pending prorations are also left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations are removed if <code>invoice_now</code> and <code>prorate</code> are both set to false.</p>

<p>By default, upon subscription cancellation, Stripe stops automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.</p>

## Definition

```yaml
{"summary": "Cancel a subscription", "description": "<p>Cancels a customer’s subscription immediately. The customer won’t be charged again for the subscription. After it’s canceled, the subscription is largely immutable. You can still update its <a href=\"/metadata\">metadata</a> and <code>cancellation_details</code>.</p>\n\n<p>Any pending invoice items that you’ve created are still charged at the end of the period, unless manually <a href=\"/api/invoiceitems/delete\">deleted</a>. If you’ve set the subscription to cancel at the end of the period, any pending prorations are also left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations are removed if <code>invoice_now</code> and <code>prorate</code> are both set to false.</p>\n\n<p>By default, upon subscription cancellation, Stripe stops automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.</p>", "operationId": "DeleteSubscriptionsSubscriptionExposedId", "parameters": [{"name": "subscription_exposed_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"cancellation_details": {"title": "cancellation_details_param", "type": "object", "properties": {"comment": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "feedback": {"type": "string", "enum": ["", "customer_service", "low_quality", "missing_features", "other", "switched_service", "too_complex", "too_expensive", "unused"]}}, "description": "Details about why this subscription was cancelled"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "invoice_now": {"type": "boolean", "description": "Will generate a final invoice that invoices for any un-invoiced metered usage and new/pending proration invoice items. Defaults to `false`."}, "prorate": {"type": "boolean", "description": "Will generate a proration invoice item that credits remaining unused time until the subscription period end. Defaults to `false`."}}, "additionalProperties": false}, "encoding": {"cancellation_details": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/subscription"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
