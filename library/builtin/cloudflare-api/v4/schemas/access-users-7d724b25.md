---
title: access_users
page_id: schema-access-users-7d724b25
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_users

```yaml
{"type": "object", "properties": {"active": {"description": "Determines the status of the SCIM User resource.", "type": "boolean", "example": true}, "displayName": {"description": "The name of the SCIM User resource.", "type": "string", "example": "John Smith"}, "emails": {"type": "array", "items": {"properties": {"primary": {"description": "Indicates if the email address is the primary email belonging to the SCIM User resource.", "type": "boolean", "example": true}, "type": {"description": "Indicates the type of the email address.", "type": "string", "example": "work"}, "value": {"description": "The email address of the SCIM User resource.", "type": "string", "format": "email", "example": "john.smith@example.com"}}, "type": "object"}}, "externalId": {"description": "The IdP-generated Id of the SCIM resource.", "type": "string", "example": "john_smith"}, "id": {"$ref": "#/components/schemas/access_id"}, "meta": {"$ref": "#/components/schemas/access_meta"}, "schemas": {"description": "The list of URIs which indicate the attributes contained within a SCIM resource.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:User"]}}}
```
