---
title: iam_scim_user
page_id: schema-iam-scim-user-4052c682
path: schemas
description: A SCIM 2.0 User resource representing an account member (RFC 7643 Section 4.1).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_user

A SCIM 2.0 User resource representing an account member (RFC 7643 Section 4.1).

```yaml
{"description": "A SCIM 2.0 User resource representing an account member (RFC 7643 Section 4.1).\n", "type": "object", "properties": {"active": {"description": "A Boolean value indicating the user's administrative status. Set to `false` to deprovision the user, removing their membership from the account.", "type": "boolean", "example": true, "x-auditable": true}, "displayName": {"description": "The display name shown for the user. Falls back to formatted name or userName if not set.", "type": "string", "example": "Jane Smith", "x-auditable": true}, "emails": {"description": "Always contains a single primary work email matching `userName`.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_user_email_object"}}, "externalId": {"description": "An identifier for the user as defined by the provisioning client (IdP). This value is stored and returned but not interpreted by Cloudflare.", "type": "string", "example": "idp-user-abc123", "x-auditable": true}, "groups": {"description": "A list of group identifiers to which the user belongs. Includes both system group tags (prefixed `cloudflare-v1-`) and custom user group tags.", "type": "array", "items": {"type": "string"}, "example": ["cloudflare-v1-023e105f4ecef8ad9ca31a8372d0c353"], "readOnly": true}, "id": {"description": "Unique identifier for the user, assigned by Cloudflare (user tag).", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "readOnly": true, "x-auditable": true}, "meta": {"description": "Resource metadata for a SCIM User.", "type": "object", "properties": {"resourceType": {"description": "The name of the resource type.", "type": "string", "example": "User"}}, "readOnly": true}, "name": {"$ref": "#/components/schemas/iam_scim_user_name_object"}, "schemas": {"description": "Must contain `urn:ietf:params:scim:schemas:core:2.0:User`.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:User"]}, "userName": {"description": "Unique identifier for the user, equal to the user's email address.", "type": "string", "format": "email", "example": "user@example.com", "x-auditable": true}}, "required": ["schemas", "id", "userName", "active"], "title": "SCIM User"}
```
