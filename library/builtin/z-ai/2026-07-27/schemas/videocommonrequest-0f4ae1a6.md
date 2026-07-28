---
title: VideoCommonRequest
page_id: schema-videocommonrequest-0f4ae1a6
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# VideoCommonRequest

```yaml
{"type": "object", "properties": {"request_id": {"type": "string", "description": "Passed by the user side, needs to be unique; used to distinguish each request, 6–64 characters. If not provided by the user side, the platform will generate one by default.", "minLength": 6, "maxLength": 64}, "user_id": {"type": "string", "description": "Unique ID of the end-user, assists the platform in intervening in end-user violations, generating illegal or inappropriate information, or other abusive behaviors. ID length requirement: minimum `6` characters, maximum `128` characters."}}}
```
