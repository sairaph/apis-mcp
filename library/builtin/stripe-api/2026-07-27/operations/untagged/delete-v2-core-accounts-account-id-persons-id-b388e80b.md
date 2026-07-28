---
title: Delete a person
page_id: operation-delete-v2-core-accounts-account-id-persons-id-562d5faf
path: operations/untagged
description: Delete a Person associated with an Account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v2/core/accounts/{account_id}/persons/{id}
operation_ids:
    - DeleteV2CoreAccountsAccountIdPersonsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a person

`DELETE /v2/core/accounts/{account_id}/persons/{id}`

Operation ID: `DeleteV2CoreAccountsAccountIdPersonsId`

Delete a Person associated with an Account.

## Definition

```yaml
{"summary": "Delete a person", "description": "Delete a Person associated with an Account.", "operationId": "DeleteV2CoreAccountsAccountIdPersonsId", "parameters": [{"name": "account_id", "in": "path", "description": "The Account the Person is associated with.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "id", "in": "path", "description": "The ID of the Person to delete.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.deleted_object"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_not_yet_compatible_with_v2"}, {"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error.v1_account_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error.v1_customer_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
