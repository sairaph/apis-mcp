---
title: Create an account link
page_id: operation-post-v1-account-links-d95f4fa8
path: operations/untagged
description: <p>Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/account_links
operation_ids:
    - PostAccountLinks
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an account link

`POST /v1/account_links`

Operation ID: `PostAccountLinks`

<p>Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.</p>

## Definition

```yaml
{"summary": "Create an account link", "description": "<p>Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.</p>", "operationId": "PostAccountLinks", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["account", "type"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string", "description": "The identifier of the account to create an account link for."}, "collect": {"type": "string", "description": "The collect parameter is deprecated. Use `collection_options` instead.", "enum": ["currently_due", "eventually_due"]}, "collection_options": {"title": "collection_options_params", "type": "object", "properties": {"fields": {"type": "string", "enum": ["currently_due", "eventually_due"]}, "future_requirements": {"type": "string", "enum": ["include", "omit"]}}, "description": "Specifies the requirements that Stripe collects from connected accounts in the Connect Onboarding flow."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "refresh_url": {"type": "string", "description": "The URL the user will be redirected to if the account link is expired, has been previously-visited, or is otherwise invalid. The URL you specify should attempt to generate a new account link with the same parameters used to create the original account link, then redirect the user to the new account link's URL so they can continue with Connect Onboarding. If a new account link cannot be generated or the redirect fails you should display a useful error to the user."}, "return_url": {"type": "string", "description": "The URL that the user will be redirected to upon leaving or completing the linked flow."}, "type": {"type": "string", "description": "The type of account link the user is requesting.\n\nYou can create Account Links of type `account_update` only for connected accounts where your platform is responsible for collecting requirements, including Custom accounts. You can't create them for accounts that have access to a Stripe-hosted Dashboard. If you use [Connect embedded components](/connect/get-started-connect-embedded-components), you can include components that allow your connected accounts to update their own information. For an account without Stripe-hosted Dashboard access where Stripe is liable for negative balances, you must use embedded components.", "enum": ["account_onboarding", "account_update"], "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"collection_options": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/account_link"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
