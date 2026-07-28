---
title: List Invitations
page_id: operation-get-user-invites-da9fa742
path: operations/user-s-invites
description: Lists all invitations associated with my user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/invites
operation_ids:
    - user'-s-invites-list-invitations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Invitations

`GET /user/invites`

Operation ID: `user'-s-invites-list-invitations`

Lists all invitations associated with my user.

## Definition

```yaml
{"operationId": "user'-s-invites-list-invitations", "summary": "List Invitations", "description": "Lists all invitations associated with my user.", "responses": {"200": {"description": "List Invitations response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_schemas-collection_invite_response"}}}}, "4XX": {"description": "List Invitations response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Invites"], "x-api-token-group": ["Memberships Write", "Memberships Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
