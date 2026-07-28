---
title: magic-visibility-pcaps_pcaps_ownership_response
page_id: schema-magic-visibility-pcaps-pcaps-ownership-response-e98b3330
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_ownership_response

```yaml
{"type": "object", "properties": {"destination_conf": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_destination_conf"}, "filename": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_challenge"}, "id": {"description": "The bucket ID associated with the packet captures API.", "type": "string", "example": "9883874ecac311ec8475433579a6bf5f", "maxLength": 32, "minLength": 32}, "status": {"description": "The status of the ownership challenge. Can be pending, success or failed.", "type": "string", "example": "success", "enum": ["pending", "success", "failed"]}, "submitted": {"description": "The RFC 3339 timestamp when the bucket was added to packet captures API.", "type": "string", "example": "2020-01-01T08:00:00Z"}, "validated": {"description": "The RFC 3339 timestamp when the bucket was validated.", "type": "string", "example": "2020-01-01T08:00:00Z"}}, "required": ["id", "status", "submitted", "destination_conf", "filename"]}
```
