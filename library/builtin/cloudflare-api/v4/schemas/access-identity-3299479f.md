---
title: access_identity
page_id: schema-access-identity-3299479f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_identity

```yaml
{"type": "object", "properties": {"account_id": {"type": "string", "example": "1234567890", "x-auditable": true}, "auth_status": {"type": "string", "example": "NONE", "x-auditable": true}, "common_name": {"type": "string", "example": ""}, "devicePosture": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/access_device_posture_rule-2"}}, "device_id": {"type": "string", "example": ""}, "device_sessions": {"$ref": "#/components/schemas/access_string_key_map_device_session"}, "email": {"type": "string", "example": "test@cloudflare.com"}, "geo": {"$ref": "#/components/schemas/access_geo"}, "iat": {"type": "number", "example": 1694791905}, "idp": {"type": "object", "properties": {"id": {"type": "string"}, "type": {"type": "string"}}}, "ip": {"type": "string", "example": "127.0.0.0"}, "is_gateway": {"type": "boolean", "example": false}, "is_warp": {"type": "boolean", "example": false}, "mtls_auth": {"type": "object", "properties": {"auth_status": {"type": "string"}, "cert_issuer_dn": {"type": "string"}, "cert_issuer_ski": {"type": "string"}, "cert_presented": {"type": "boolean"}, "cert_serial": {"type": "string"}}}, "service_token_id": {"type": "string", "example": ""}, "service_token_status": {"type": "boolean", "example": false}, "user_uuid": {"type": "string", "example": "57cf8cf2-f55a-4588-9ac9-f5e41e9f09b4"}, "version": {"type": "number", "example": 2}}}
```
