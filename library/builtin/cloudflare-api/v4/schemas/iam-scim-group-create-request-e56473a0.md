---
title: iam_scim_group_create_request
page_id: schema-iam-scim-group-create-request-e56473a0
path: schemas
description: Request body for creating a SCIM Group. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_group_create_request

Request body for creating a SCIM Group. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).

```yaml
{"description": "Request body for creating a SCIM Group. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).\n", "type": "object", "properties": {"displayName": {"description": "A human-readable name for the Group. REQUIRED. Must not start with `CF` (reserved prefix for Cloudflare-managed virtual groups).\n", "type": "string", "example": "My IdP Group", "x-auditable": true}, "externalId": {"description": "Identifier for the Group as defined by the provisioning client (IdP).", "type": "string", "example": "idp-group-456", "x-auditable": true}}, "required": ["displayName"], "title": "SCIM Group Create Request"}
```
