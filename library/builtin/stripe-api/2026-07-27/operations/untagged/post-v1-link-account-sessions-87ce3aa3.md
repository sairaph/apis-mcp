---
title: Create a Session
page_id: operation-post-v1-link-account-sessions-cd44d74f
path: operations/untagged
description: <p>To launch the Financial Connections authorization flow, create a <code>Session</code>. The session’s <code>client_secret</code> can be used to launch the flow using Stripe.js.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/link_account_sessions
operation_ids:
    - PostLinkAccountSessions
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Session

`POST /v1/link_account_sessions`

Operation ID: `PostLinkAccountSessions`

<p>To launch the Financial Connections authorization flow, create a <code>Session</code>. The session’s <code>client_secret</code> can be used to launch the flow using Stripe.js.</p>

## Definition

```yaml
{"summary": "Create a Session", "description": "<p>To launch the Financial Connections authorization flow, create a <code>Session</code>. The session’s <code>client_secret</code> can be used to launch the flow using Stripe.js.</p>", "operationId": "PostLinkAccountSessions", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["account_holder", "permissions"], "type": "object", "properties": {"account_holder": {"title": "accountholder_params", "required": ["type"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string"}, "customer": {"maxLength": 5000, "type": "string"}, "customer_account": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["account", "customer"]}}, "description": "The account holder to link accounts for."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "filters": {"title": "filters_params", "type": "object", "properties": {"account_subcategories": {"type": "array", "items": {"type": "string", "enum": ["checking", "credit_card", "line_of_credit", "mortgage", "savings"]}}, "countries": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, "description": "Filters to restrict the kinds of accounts to collect."}, "permissions": {"type": "array", "description": "List of data features that you would like to request access to.\n\nPossible values are `balances`, `transactions`, `ownership`, and `payment_method`.", "items": {"maxLength": 5000, "type": "string", "enum": ["balances", "ownership", "payment_method", "transactions"], "x-stripeBypassValidation": true}}, "prefetch": {"type": "array", "description": "List of data features that you would like to retrieve upon account creation.", "items": {"type": "string", "enum": ["balances", "ownership", "transactions"], "x-stripeBypassValidation": true}}, "return_url": {"maxLength": 5000, "type": "string", "description": "For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app."}}, "additionalProperties": false}, "encoding": {"account_holder": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "filters": {"style": "deepObject", "explode": true}, "permissions": {"style": "deepObject", "explode": true}, "prefetch": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
