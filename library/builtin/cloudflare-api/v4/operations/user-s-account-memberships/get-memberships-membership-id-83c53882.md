---
title: Membership Details
page_id: operation-get-memberships-membership-id-28dcd798
path: operations/user-s-account-memberships
description: Get a specific membership.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /memberships/{membership_id}
operation_ids:
    - user'-s-account-memberships-membership-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Membership Details

`GET /memberships/{membership_id}`

Operation ID: `user'-s-account-memberships-membership-details`

Get a specific membership.

## Definition

```yaml
{"operationId": "user'-s-account-memberships-membership-details", "summary": "Membership Details", "description": "Get a specific membership.", "parameters": [{"name": "membership_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}], "responses": {"200": {"description": "Membership Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_membership_response_with_policies"}}}}, "4XX": {"description": "Membership Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Account Memberships"], "x-api-token-group": ["Memberships Write", "Memberships Read"]}
```
