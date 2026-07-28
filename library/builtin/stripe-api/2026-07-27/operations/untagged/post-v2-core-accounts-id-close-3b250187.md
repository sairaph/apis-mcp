---
title: Close an account
page_id: operation-post-v2-core-accounts-id-close-f114b72e
path: operations/untagged
description: Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/accounts/{id}/close
operation_ids:
    - PostV2CoreAccountsIdClose
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Close an account

`POST /v2/core/accounts/{id}/close`

Operation ID: `PostV2CoreAccountsIdClose`

Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.

## Definition

```yaml
{"summary": "Close an account", "description": "Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.", "operationId": "PostV2CoreAccountsIdClose", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Account to close.", "required": true, "style": "simple", "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"applied_configurations": {"type": "array", "description": "Configurations on the Account to be closed. All configurations on the Account must be passed in for this request to succeed.", "items": {"type": "string", "enum": ["customer", "merchant", "recipient"]}}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.account"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_not_yet_compatible_with_v2"}, {"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.cannot_delete_account_with_balance"}, {"$ref": "#/components/schemas/v2.error.cannot_delete_customer_with_available_cash_balance"}, {"$ref": "#/components/schemas/v2.error.configs_must_match_to_close"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error.pending_transactions_cannot_be_deleted"}, {"$ref": "#/components/schemas/v2.error.platform_registration_required"}, {"$ref": "#/components/schemas/v2.error.stripe_loss_liable_cannot_be_deleted"}, {"$ref": "#/components/schemas/v2.error.v1_account_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error.v1_customer_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
