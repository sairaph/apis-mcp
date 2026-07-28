---
title: iam_scim_group
page_id: schema-iam-scim-group-bc19fb93
path: schemas
description: A SCIM 2.0 Group resource representing an account user group (RFC 7643 Section 4.2). May be a system group (backed by a Cloudflare permission group, read-only except for member management) or a custom user group (full CRUD).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_group

A SCIM 2.0 Group resource representing an account user group (RFC 7643 Section 4.2). May be a system group (backed by a Cloudflare permission group, read-only except for member management) or a custom user group (full CRUD).

```yaml
{"description": "A SCIM 2.0 Group resource representing an account user group (RFC 7643 Section 4.2). May be a system group (backed by a Cloudflare permission group, read-only except for member management) or a custom user group (full CRUD).", "type": "object", "properties": {"displayName": {"description": "A human-readable name for the Group.", "type": "string", "example": "Administrators", "x-auditable": true}, "externalId": {"description": "Identifier for the Group as defined by the provisioning client (IdP).", "type": "string", "example": "idp-group-456", "x-auditable": true}, "id": {"description": "Unique identifier for the Group, assigned by Cloudflare. System groups are prefixed `cloudflare-v1-<permissionGroupTag>`; custom groups use a UUID-style tag.", "type": "string", "example": "cloudflare-v1-023e105f4ecef8ad9ca31a8372d0c353", "readOnly": true, "x-auditable": true}, "members": {"description": "A list of members of the Group. Only populated for custom (Phase 2) groups on individual GET requests. Each member object contains a `value` (user tag) and optional `display` (email).", "type": "array", "items": {"properties": {"display": {"description": "The display name (email) of the group member.", "type": "string", "example": "user@example.com"}, "value": {"description": "The user tag identifier of the group member.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, "type": "object"}, "example": [{"display": "user@example.com", "value": "023e105f4ecef8ad9ca31a8372d0c353"}], "readOnly": true}, "meta": {"$ref": "#/components/schemas/iam_scim_group_meta"}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:Group"]}}, "required": ["schemas", "id", "displayName"], "title": "SCIM Group"}
```
