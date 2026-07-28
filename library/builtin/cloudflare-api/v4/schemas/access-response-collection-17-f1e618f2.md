---
title: access_response_collection-17
page_id: schema-access-response-collection-17-f1e618f2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_response_collection-17

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/access_azureAD-2"}, {"$ref": "#/components/schemas/access_centrify-2"}, {"$ref": "#/components/schemas/access_facebook-2"}, {"$ref": "#/components/schemas/access_github-2"}, {"$ref": "#/components/schemas/access_google-2"}, {"$ref": "#/components/schemas/access_google-apps-2"}, {"$ref": "#/components/schemas/access_linkedin-2"}, {"$ref": "#/components/schemas/access_oidc-2"}, {"$ref": "#/components/schemas/access_okta-2"}, {"$ref": "#/components/schemas/access_onelogin-2"}, {"$ref": "#/components/schemas/access_pingone-2"}, {"$ref": "#/components/schemas/access_saml-2"}, {"$ref": "#/components/schemas/access_yandex-2"}, {"$ref": "#/components/schemas/access_onetimepin-2"}, {"$ref": "#/components/schemas/access_cloudflare-2"}]}}}, "type": "object"}]}
```
