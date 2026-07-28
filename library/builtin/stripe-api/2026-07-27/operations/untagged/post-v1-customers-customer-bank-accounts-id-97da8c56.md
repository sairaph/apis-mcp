---
title: POST /v1/customers/{customer}/bank_accounts/{id}
page_id: operation-post-v1-customers-customer-bank-accounts-id-dc7346f2
path: operations/untagged
description: <p>Update a specified source for a given customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/bank_accounts/{id}
operation_ids:
    - PostCustomersCustomerBankAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/customers/{customer}/bank_accounts/{id}

`POST /v1/customers/{customer}/bank_accounts/{id}`

Operation ID: `PostCustomersCustomerBankAccountsId`

<p>Update a specified source for a given customer.</p>

## Definition

```yaml
{"description": "<p>Update a specified source for a given customer.</p>", "operationId": "PostCustomersCustomerBankAccountsId", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string", "description": "The name of the person or business that owns the bank account."}, "account_holder_type": {"maxLength": 5000, "type": "string", "description": "The type of entity that holds the account. This can be either `individual` or `company`.", "enum": ["company", "individual"]}, "address_city": {"maxLength": 5000, "type": "string", "description": "City/District/Suburb/Town/Village."}, "address_country": {"maxLength": 5000, "type": "string", "description": "Billing address country, if provided when creating card."}, "address_line1": {"maxLength": 5000, "type": "string", "description": "Address line 1 (Street address/PO Box/Company name)."}, "address_line2": {"maxLength": 5000, "type": "string", "description": "Address line 2 (Apartment/Suite/Unit/Building)."}, "address_state": {"maxLength": 5000, "type": "string", "description": "State/County/Province/Region."}, "address_zip": {"maxLength": 5000, "type": "string", "description": "ZIP or postal code."}, "exp_month": {"maxLength": 5000, "type": "string", "description": "Two digit number representing the card’s expiration month."}, "exp_year": {"maxLength": 5000, "type": "string", "description": "Four digit number representing the card’s expiration year."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "name": {"maxLength": 5000, "type": "string", "description": "Cardholder name."}, "owner": {"title": "owner", "type": "object", "properties": {"address": {"title": "source_address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, "email": {"type": "string"}, "name": {"maxLength": 5000, "type": "string"}, "phone": {"maxLength": 5000, "type": "string"}}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "owner": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/card"}, {"$ref": "#/components/schemas/bank_account"}, {"$ref": "#/components/schemas/source"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
