---
title: ResponsesErrorField
page_id: schema-responseserrorfield-8dd0b192
path: schemas
description: Error information returned from the API
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ResponsesErrorField

Error information returned from the API

```yaml
{"description": "Error information returned from the API", "example": {"code": "rate_limit_exceeded", "message": "Rate limit exceeded. Please try again later."}, "properties": {"code": {"enum": ["server_error", "rate_limit_exceeded", "invalid_prompt", "vector_store_timeout", "invalid_image", "invalid_image_format", "invalid_base64_image", "invalid_image_url", "image_too_large", "image_too_small", "image_parse_error", "image_content_policy_violation", "invalid_image_mode", "image_file_too_large", "unsupported_image_media_type", "empty_image_file", "failed_to_download_image", "image_file_not_found", "bio_policy"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": ["object", "null"]}
```
