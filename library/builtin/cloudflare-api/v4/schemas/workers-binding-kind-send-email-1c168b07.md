---
title: workers_binding_kind_send_email
page_id: schema-workers-binding-kind-send-email-1c168b07
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_send_email

```yaml
{"type": "object", "properties": {"allowed_destination_addresses": {"description": "List of allowed destination addresses.", "type": "array", "items": {"format": "email", "type": "string"}, "example": ["user1@example.com", "user2@example.com"], "x-auditable": true}, "allowed_sender_addresses": {"description": "List of allowed sender addresses.", "type": "array", "items": {"format": "email", "type": "string"}, "example": ["user1@example.com", "user2@example.com"], "x-auditable": true}, "destination_address": {"description": "Destination address for the email.", "type": "string", "format": "email", "example": "user@example.com", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["send_email"], "x-auditable": true}}, "required": ["name", "type"]}
```
