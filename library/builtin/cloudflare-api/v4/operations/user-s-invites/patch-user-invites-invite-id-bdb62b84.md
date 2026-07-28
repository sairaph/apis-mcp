---
title: Respond to Invitation
page_id: operation-patch-user-invites-invite-id-3805565f
path: operations/user-s-invites
description: Responds to an invitation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /user/invites/{invite_id}
operation_ids:
    - user'-s-invites-respond-to-invitation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Respond to Invitation

`PATCH /user/invites/{invite_id}`

Operation ID: `user'-s-invites-respond-to-invitation`

Responds to an invitation.

## Definition

```yaml
{"operationId": "user'-s-invites-respond-to-invitation", "summary": "Respond to Invitation", "description": "Responds to an invitation.", "parameters": [{"name": "invite_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_invite_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"status": {"description": "Status of your response to the invitation (rejected or accepted).", "example": "accepted", "enum": ["accepted", "rejected"]}}, "required": ["status"]}}}}, "responses": {"200": {"description": "Respond to Invitation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_invite_response"}}}}, "4XX": {"description": "Respond to Invitation response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Invites"], "x-api-token-group": ["Memberships Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
