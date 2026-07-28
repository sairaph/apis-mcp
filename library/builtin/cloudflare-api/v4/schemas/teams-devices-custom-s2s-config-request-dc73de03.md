---
title: teams-devices_custom_s2s_config_request
page_id: schema-teams-devices-custom-s2s-config-request-dc73de03
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_custom_s2s_config_request

```yaml
{"type": "object", "properties": {"access_client_id": {"description": "This id will be passed in the `CF-Access-Client-ID` header when hitting the `api_url`.", "type": "string", "example": "88bf3b6d86161464f6509f7219099e57.access"}, "access_client_secret": {"description": "This secret will be passed in the `CF-Access-Client-Secret` header when hitting the `api_url`.", "type": "string", "example": "bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5", "x-sensitive": true}, "api_url": {"description": "The Custom Device Posture Integration  API URL.", "type": "string", "example": "https://example.custom-s2s.com", "x-auditable": true}}, "required": ["api_url", "access_client_id", "access_client_secret"], "title": "Custom Device Posture Integration Config"}
```
