---
title: Update Membership
page_id: operation-put-memberships-membership-id-772c6f03
path: operations/user-s-account-memberships
description: Accept or reject this account invitation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /memberships/{membership_id}
operation_ids:
    - user'-s-account-memberships-update-membership
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Membership

`PUT /memberships/{membership_id}`

Operation ID: `user'-s-account-memberships-update-membership`

Accept or reject this account invitation.

## Definition

```yaml
{"operationId": "user'-s-account-memberships-update-membership", "summary": "Update Membership", "description": "Accept or reject this account invitation.", "parameters": [{"name": "membership_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"status": {"description": "Whether to accept or reject this account invitation.", "example": "accepted", "enum": ["accepted", "rejected"]}}, "required": ["status"]}}}}, "responses": {"200": {"description": "Update Membership response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_membership_response_with_policies"}}}}, "4XX": {"description": "Update Membership response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Account Memberships"], "x-api-token-group": ["Memberships Write"]}
```
