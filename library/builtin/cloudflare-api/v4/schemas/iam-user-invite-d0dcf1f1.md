---
title: iam_user_invite
page_id: schema-iam-user-invite-d0dcf1f1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_user_invite

```yaml
{"type": "object", "properties": {"expires_on": {"$ref": "#/components/schemas/iam_schemas-expires_on"}, "id": {"$ref": "#/components/schemas/iam_invite_components-schemas-identifier"}, "invited_by": {"$ref": "#/components/schemas/iam_invited_by"}, "invited_member_email": {"$ref": "#/components/schemas/iam_invited_member_email"}, "invited_member_id": {"description": "ID of the user to add to the organization.", "type": "string", "example": "5a7805061c76ada191ed06f989cc3dac", "maxLength": 32, "nullable": true, "readOnly": true, "x-auditable": true}, "invited_on": {"$ref": "#/components/schemas/iam_invited_on"}, "organization_id": {"description": "ID of the organization the user will be added to.", "type": "string", "example": "5a7805061c76ada191ed06f989cc3dac", "maxLength": 32, "readOnly": true, "x-auditable": true}, "organization_is_enforcing_twofactor": {"type": "boolean", "example": true, "x-auditable": true}, "organization_name": {"description": "Organization name.", "type": "string", "example": "Cloudflare, Inc.", "maxLength": 100, "readOnly": true, "x-auditable": true}, "roles": {"$ref": "#/components/schemas/iam_role_names"}, "status": {"description": "Current status of the invitation.", "example": "accepted", "enum": ["pending", "accepted", "rejected", "expired"], "x-auditable": true}}, "required": ["invited_member_id", "organization_id"]}
```
