---
title: access_failed_login_response
page_id: schema-access-failed-login-response-fc434ad3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_failed_login_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"expiration": {"type": "integer"}, "metadata": {"type": "object", "example": {"app_name": "Test App", "aud": "39691c1480a2352a18ece567debc2b32552686cbd38eec0887aa18d5d3f00c04", "datetime": "2022-02-02T21:54:34.914Z", "ray_id": "6d76a8a42ead4133", "user_email": "test@cloudflare.com", "user_uuid": "57171132-e453-4ee8-b2a5-8cbaad333207"}}}, "type": "object"}}}}]}
```
