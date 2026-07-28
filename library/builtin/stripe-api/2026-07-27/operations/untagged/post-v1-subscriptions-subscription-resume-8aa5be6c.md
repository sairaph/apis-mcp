---
title: Resume a subscription
page_id: operation-post-v1-subscriptions-subscription-resume-b6768de2
path: operations/untagged
description: <p>Initiates resumption of a paused subscription, optionally resetting the billing cycle anchor and creating prorations. Resume is only available for subscriptions that use <code>charge_automatically</code> collection. If Stripe doesn’t generate a resumption invoice, the subscription becomes <code>active</code> immediately. When a resumption invoice is generated, Stripe finalizes it immediately. If the invoice is paid or marked uncollectible, the subscription becomes <code>active</code>. If the invoice is manually voided, the subscription stays <code>paused</code>. If there is no payment attempt within 23 hours, Stripe voids the invoice and the subscription stays <code>paused</code>. Learn more about <a href="/docs/billing/subscriptions/pause#resume-subscriptions">resuming subscriptions</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/subscriptions/{subscription}/resume
operation_ids:
    - PostSubscriptionsSubscriptionResume
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Resume a subscription

`POST /v1/subscriptions/{subscription}/resume`

Operation ID: `PostSubscriptionsSubscriptionResume`

<p>Initiates resumption of a paused subscription, optionally resetting the billing cycle anchor and creating prorations. Resume is only available for subscriptions that use <code>charge_automatically</code> collection. If Stripe doesn’t generate a resumption invoice, the subscription becomes <code>active</code> immediately. When a resumption invoice is generated, Stripe finalizes it immediately. If the invoice is paid or marked uncollectible, the subscription becomes <code>active</code>. If the invoice is manually voided, the subscription stays <code>paused</code>. If there is no payment attempt within 23 hours, Stripe voids the invoice and the subscription stays <code>paused</code>. Learn more about <a href="/docs/billing/subscriptions/pause#resume-subscriptions">resuming subscriptions</a>.</p>

## Definition

```yaml
{"summary": "Resume a subscription", "description": "<p>Initiates resumption of a paused subscription, optionally resetting the billing cycle anchor and creating prorations. Resume is only available for subscriptions that use <code>charge_automatically</code> collection. If Stripe doesn’t generate a resumption invoice, the subscription becomes <code>active</code> immediately. When a resumption invoice is generated, Stripe finalizes it immediately. If the invoice is paid or marked uncollectible, the subscription becomes <code>active</code>. If the invoice is manually voided, the subscription stays <code>paused</code>. If there is no payment attempt within 23 hours, Stripe voids the invoice and the subscription stays <code>paused</code>. Learn more about <a href=\"/docs/billing/subscriptions/pause#resume-subscriptions\">resuming subscriptions</a>.</p>", "operationId": "PostSubscriptionsSubscriptionResume", "parameters": [{"name": "subscription", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"billing_cycle_anchor": {"type": "string", "description": "The billing cycle anchor that applies when the subscription is resumed. Either `now` or `unchanged`. The default is `now`. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).", "enum": ["now", "unchanged"]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "proration_behavior": {"type": "string", "description": "Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) resulting from the `billing_cycle_anchor` being `unchanged`. When the `billing_cycle_anchor` is set to `now` (default value), no prorations are generated. If no value is passed, the default is `create_prorations`.", "enum": ["always_invoice", "create_prorations", "none"]}, "proration_date": {"type": "integer", "description": "If set, prorations will be calculated as though the subscription was resumed at the given time. This can be used to apply exactly the same prorations that were previewed with the [create preview](https://stripe.com/docs/api/invoices/create_preview) endpoint.", "format": "unix-time"}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/subscription"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
