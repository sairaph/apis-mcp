---
title: access_cloudflare
page_id: schema-access-cloudflare-29f09321
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_cloudflare

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"type": "object", "properties": {"redirect_url": {"type": "string", "readOnly": true}, "restrict_to_account_members": {"description": "When enabled, only users who are members of your Cloudflare account can authenticate through this identity provider. When disabled, any user with a Cloudflare account can authenticate, subject to your Access policies.", "type": "boolean", "default": false}}}, "type": {"type": "string", "enum": ["cloudflare"]}}}], "title": "Cloudflare"}
```
