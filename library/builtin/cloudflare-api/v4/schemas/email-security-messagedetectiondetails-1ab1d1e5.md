---
title: email-security_MessageDetectionDetails
page_id: schema-email-security-messagedetectiondetails-1ab1d1e5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_MessageDetectionDetails

```yaml
{"type": "object", "properties": {"action": {"type": "string"}, "attachments": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Attachment"}}, "final_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "findings": {"type": "array", "items": {"properties": {"attachment": {"type": "string", "nullable": true}, "detail": {"type": "string", "nullable": true}, "detection": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "field": {"type": "string", "nullable": true}, "name": {"type": "string", "nullable": true}, "portion": {"type": "string", "nullable": true}, "reason": {"type": "string", "nullable": true}, "score": {"type": "number", "format": "double", "nullable": true}, "value": {"type": "string", "nullable": true}}, "type": "object"}, "nullable": true}, "headers": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_MessageHeader"}}, "links": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Link"}}, "sender_info": {"type": "object", "properties": {"as_name": {"description": "The name of the autonomous system.", "type": "string", "nullable": true}, "as_number": {"description": "The number of the autonomous system.", "type": "integer", "nullable": true}, "geo": {"type": "string", "nullable": true}, "ip": {"type": "string", "nullable": true}, "pld": {"type": "string", "nullable": true}}}, "threat_categories": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_ThreatCategory"}}, "validation": {"$ref": "#/components/schemas/email-security_Validation"}}, "required": ["validation", "headers", "threat_categories", "sender_info", "links", "action", "attachments", "findings"]}
```
