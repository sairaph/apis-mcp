---
title: iam_scim_group_summary
page_id: schema-iam-scim-group-summary-fbfb5067
path: schemas
description: A SCIM 2.0 Group resource as returned in list responses. Does not include members; use the individual GET endpoint to retrieve members.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_group_summary

A SCIM 2.0 Group resource as returned in list responses. Does not include members; use the individual GET endpoint to retrieve members.

```yaml
{"description": "A SCIM 2.0 Group resource as returned in list responses. Does not include members; use the individual GET endpoint to retrieve members.", "type": "object", "properties": {"displayName": {"description": "A human-readable name for the Group.", "type": "string", "example": "Administrators", "x-auditable": true}, "externalId": {"description": "Identifier for the Group as defined by the provisioning client (IdP).", "type": "string", "example": "idp-group-456", "x-auditable": true}, "id": {"description": "Unique identifier for the Group, assigned by Cloudflare. System groups are prefixed `cloudflare-v1-<permissionGroupTag>`; custom groups use a UUID-style tag.", "type": "string", "example": "cloudflare-v1-023e105f4ecef8ad9ca31a8372d0c353", "readOnly": true, "x-auditable": true}, "meta": {"$ref": "#/components/schemas/iam_scim_group_meta"}, "schemas": {"description": "Must contain `urn:ietf:params:scim:schemas:core:2.0:Group`.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:Group"]}}, "required": ["schemas", "id", "displayName"], "title": "SCIM Group Summary"}
```
