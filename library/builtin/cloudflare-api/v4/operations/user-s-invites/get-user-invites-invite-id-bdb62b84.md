---
title: Invitation Details
page_id: operation-get-user-invites-invite-id-cd045017
path: operations/user-s-invites
description: Gets the details of an invitation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/invites/{invite_id}
operation_ids:
    - user'-s-invites-invitation-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Invitation Details

`GET /user/invites/{invite_id}`

Operation ID: `user'-s-invites-invitation-details`

Gets the details of an invitation.

## Definition

```yaml
{"operationId": "user'-s-invites-invitation-details", "summary": "Invitation Details", "description": "Gets the details of an invitation.", "parameters": [{"name": "invite_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_invite_components-schemas-identifier"}}], "responses": {"200": {"description": "Invitation Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_invite_response"}}}}, "4XX": {"description": "Invitation Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Invites"], "x-api-token-group": ["Memberships Write", "Memberships Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
