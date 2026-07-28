---
title: POST /v1/accounts/{account}/external_accounts/{id}
page_id: operation-post-v1-accounts-account-external-accounts-id-47a25cd1
path: operations/untagged
description: |-
    <p>Updates the metadata, account holder name, account holder type of a bank account belonging to
    a connected account and optionally sets it as the default for its currency. Other bank account
    details are not editable by design.</p>

    <p>You can only update bank accounts when <a href="/api/accounts/object#account_object-controller-requirement_collection">account.controller.requirement_collection</a> is <code>application</code>, which includes <a href="/connect/custom-accounts">Custom accounts</a>.</p>

    <p>You can re-enable a disabled bank account by performing an update call without providing any
    arguments or changes.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/accounts/{account}/external_accounts/{id}
operation_ids:
    - PostAccountsAccountExternalAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/accounts/{account}/external_accounts/{id}

`POST /v1/accounts/{account}/external_accounts/{id}`

Operation ID: `PostAccountsAccountExternalAccountsId`

<p>Updates the metadata, account holder name, account holder type of a bank account belonging to
a connected account and optionally sets it as the default for its currency. Other bank account
details are not editable by design.</p>

<p>You can only update bank accounts when <a href="/api/accounts/object#account_object-controller-requirement_collection">account.controller.requirement_collection</a> is <code>application</code>, which includes <a href="/connect/custom-accounts">Custom accounts</a>.</p>

<p>You can re-enable a disabled bank account by performing an update call without providing any
arguments or changes.</p>

## Definition

```yaml
{"description": "<p>Updates the metadata, account holder name, account holder type of a bank account belonging to\na connected account and optionally sets it as the default for its currency. Other bank account\ndetails are not editable by design.</p>\n\n<p>You can only update bank accounts when <a href=\"/api/accounts/object#account_object-controller-requirement_collection\">account.controller.requirement_collection</a> is <code>application</code>, which includes <a href=\"/connect/custom-accounts\">Custom accounts</a>.</p>\n\n<p>You can re-enable a disabled bank account by performing an update call without providing any\narguments or changes.</p>", "operationId": "PostAccountsAccountExternalAccountsId", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name of the person or business that owns the bank account."}, "account_holder_type": {"maxLength": 5000, "type": "string", "description": "The type of entity that holds the account. This can be either `individual` or `company`.", "enum": ["", "company", "individual"]}, "account_type": {"maxLength": 5000, "type": "string", "description": "The bank account type. This can only be `checking` or `savings` in most countries. In Japan, this can only be `futsu` or `toza`.", "enum": ["checking", "futsu", "savings", "toza"]}, "address_city": {"maxLength": 5000, "type": "string", "description": "City/District/Suburb/Town/Village."}, "address_country": {"maxLength": 5000, "type": "string", "description": "Billing address country, if provided when creating card."}, "address_line1": {"maxLength": 5000, "type": "string", "description": "Address line 1 (Street address/PO Box/Company name)."}, "address_line2": {"maxLength": 5000, "type": "string", "description": "Address line 2 (Apartment/Suite/Unit/Building)."}, "address_state": {"maxLength": 5000, "type": "string", "description": "State/County/Province/Region."}, "address_zip": {"maxLength": 5000, "type": "string", "description": "ZIP or postal code."}, "default_for_currency": {"type": "boolean", "description": "When set to true, this becomes the default external account for its currency."}, "documents": {"title": "external_account_documents_param", "type": "object", "properties": {"bank_account_ownership_verification": {"title": "documents_param", "type": "object", "properties": {"files": {"type": "array", "items": {"maxLength": 500, "type": "string"}}}}}, "description": "Documents that may be submitted to satisfy various informational requests."}, "exp_month": {"maxLength": 5000, "type": "string", "description": "Two digit number representing the card’s expiration month."}, "exp_year": {"maxLength": 5000, "type": "string", "description": "Four digit number representing the card’s expiration year."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "name": {"maxLength": 5000, "type": "string", "description": "Cardholder name."}}, "additionalProperties": false}, "encoding": {"documents": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/external_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
