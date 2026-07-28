---
title: Delete Membership
page_id: operation-delete-memberships-membership-id-62623860
path: operations/user-s-account-memberships
description: Remove the associated member from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /memberships/{membership_id}
operation_ids:
    - user'-s-account-memberships-delete-membership
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Membership

`DELETE /memberships/{membership_id}`

Operation ID: `user'-s-account-memberships-delete-membership`

Remove the associated member from an account.

## Definition

```yaml
{"operationId": "user'-s-account-memberships-delete-membership", "summary": "Delete Membership", "description": "Remove the associated member from an account.", "parameters": [{"name": "membership_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Membership response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"properties": {"id": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Membership response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Account Memberships"], "x-api-token-group": ["Memberships Write"]}
```
