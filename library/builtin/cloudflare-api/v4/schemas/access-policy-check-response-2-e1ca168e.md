---
title: access_policy_check_response-2
page_id: schema-access-policy-check-response-2-e1ca168e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_policy_check_response-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"app_state": {"type": "object", "properties": {"app_uid": {"$ref": "#/components/schemas/access_uuid"}, "aud": {"type": "string", "example": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe389"}, "hostname": {"type": "string", "example": "test.com"}, "name": {"type": "string", "example": "Test App"}, "policies": {"type": "array", "items": {}, "example": [{"decision": "allow", "exclude": [], "include": [{"_type": "email", "email": "testuser@gmail.com"}], "precedence": 0, "require": [], "status": "Success"}]}, "status": {"type": "string", "example": "Success"}}}, "user_identity": {"type": "object", "properties": {"account_id": {"type": "string", "example": "41ecfbb341f033e52b46742756aabb8b"}, "device_sessions": {"type": "object", "example": {}}, "email": {"type": "string", "example": "testuser@gmail.com"}, "geo": {"type": "object", "properties": {"country": {"type": "string", "example": "US"}}}, "iat": {"type": "integer"}, "id": {"type": "string", "example": "1164449231815010287495"}, "is_gateway": {"type": "boolean", "example": false}, "is_warp": {"type": "boolean", "example": false}, "name": {"type": "string", "example": "Test User"}, "user_uuid": {"$ref": "#/components/schemas/access_uuid"}, "version": {"type": "integer"}}}}}}}]}
```
