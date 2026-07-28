---
title: Create an external account
page_id: operation-post-v1-accounts-account-external-accounts-56d78837
path: operations/untagged
description: <p>Create an external account for a given account.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/accounts/{account}/external_accounts
operation_ids:
    - PostAccountsAccountExternalAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an external account

`POST /v1/accounts/{account}/external_accounts`

Operation ID: `PostAccountsAccountExternalAccounts`

<p>Create an external account for a given account.</p>

## Definition

```yaml
{"summary": "Create an external account", "description": "<p>Create an external account for a given account.</p>", "operationId": "PostAccountsAccountExternalAccounts", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"bank_account": {"description": "Either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary containing a user's bank account details.", "anyOf": [{"title": "external_account_payout_bank_account", "required": ["account_number", "country"], "type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string"}, "account_holder_type": {"maxLength": 5000, "type": "string", "enum": ["company", "individual"]}, "account_number": {"maxLength": 5000, "type": "string"}, "account_type": {"maxLength": 5000, "type": "string", "enum": ["checking", "futsu", "savings", "toza"]}, "country": {"maxLength": 5000, "type": "string"}, "currency": {"type": "string", "format": "currency"}, "documents": {"title": "external_account_documents_param", "type": "object", "properties": {"bank_account_ownership_verification": {"title": "documents_param", "type": "object", "properties": {"files": {"type": "array", "items": {"maxLength": 500, "type": "string"}}}}}}, "object": {"maxLength": 5000, "type": "string", "enum": ["bank_account"]}, "routing_number": {"maxLength": 5000, "type": "string"}}}, {"maxLength": 5000, "type": "string"}]}, "default_for_currency": {"type": "boolean", "description": "When set to true, or if this is the first external account added in this currency, this account becomes the default external account for its currency."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "external_account": {"maxLength": 5000, "type": "string", "description": "A token, like the ones returned by [Stripe.js](https://docs.stripe.com/js) or a dictionary containing a user's external account details (with the options shown below). Please refer to full [documentation](https://stripe.com/docs/api/external_accounts) instead.", "x-stripeBypassValidation": true}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}}, "additionalProperties": false}, "encoding": {"bank_account": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/external_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
