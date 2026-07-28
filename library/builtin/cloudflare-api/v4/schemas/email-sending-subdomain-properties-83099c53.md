---
title: email_sending_subdomain_properties
page_id: schema-email-sending-subdomain-properties-83099c53
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_sending_subdomain_properties

```yaml
{"type": "object", "properties": {"created": {"$ref": "#/components/schemas/email_created"}, "dkim_selector": {"description": "The DKIM selector used for email signing.", "type": "string", "example": "cf-bounce", "readOnly": true}, "enabled": {"description": "Whether Email Sending is enabled on this subdomain.", "type": "boolean", "readOnly": true}, "modified": {"$ref": "#/components/schemas/email_modified"}, "name": {"description": "The subdomain domain name.", "type": "string", "example": "sub.example.com", "x-auditable": true}, "preview_enabled": {"description": "Whether sent messages from this subdomain can be previewed in the activity log.", "type": "boolean", "example": true, "x-auditable": true}, "return_path_domain": {"description": "The return-path domain used for bounce handling.", "type": "string", "example": "cf-bounce.sub.example.com", "readOnly": true}, "tag": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, "required": ["tag", "name", "enabled"]}
```
