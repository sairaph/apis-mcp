---
title: iam_scim_user_replace_request
page_id: schema-iam-scim-user-replace-request-640dea32
path: schemas
description: Request body for replacing a SCIM User (PUT). Only `schemas` and `userName` are required; other fields are optional.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_user_replace_request

Request body for replacing a SCIM User (PUT). Only `schemas` and `userName` are required; other fields are optional.

```yaml
{"description": "Request body for replacing a SCIM User (PUT). Only `schemas` and `userName` are required; other fields are optional.\n", "type": "object", "properties": {"active": {"description": "A Boolean value indicating the user's administrative status. Optional; if omitted, the current value is preserved. Set to `false` to deprovision the user, removing their membership from the account.", "type": "boolean", "example": true, "x-auditable": true}, "displayName": {"description": "The name of the user, suitable for display to end-users. If not explicitly set, falls back to the formatted name or userName.", "type": "string", "example": "Jane Smith", "x-auditable": true}, "emails": {"description": "Email addresses for the user. If a primary email is provided, it must match `userName`.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_user_email_object"}}, "externalId": {"description": "An identifier for the user as defined by the provisioning client (IdP). This value is stored and returned but not interpreted by Cloudflare.", "type": "string", "example": "idp-user-abc123", "x-auditable": true}, "name": {"$ref": "#/components/schemas/iam_scim_user_name_object"}, "schemas": {"description": "Must contain `urn:ietf:params:scim:schemas:core:2.0:User`.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:User"]}, "userName": {"description": "Unique identifier for the user, equal to the user's email address.", "type": "string", "format": "email", "example": "user@example.com", "x-auditable": true}}, "required": ["schemas", "userName"], "title": "SCIM User Replace Request"}
```
