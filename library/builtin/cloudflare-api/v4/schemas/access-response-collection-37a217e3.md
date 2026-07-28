---
title: access_response_collection
page_id: schema-access-response-collection-37a217e3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_response_collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/access_azureAD"}, {"$ref": "#/components/schemas/access_centrify"}, {"$ref": "#/components/schemas/access_facebook"}, {"$ref": "#/components/schemas/access_github"}, {"$ref": "#/components/schemas/access_google"}, {"$ref": "#/components/schemas/access_google-apps"}, {"$ref": "#/components/schemas/access_linkedin"}, {"$ref": "#/components/schemas/access_oidc"}, {"$ref": "#/components/schemas/access_okta"}, {"$ref": "#/components/schemas/access_onelogin"}, {"$ref": "#/components/schemas/access_pingone"}, {"$ref": "#/components/schemas/access_saml"}, {"$ref": "#/components/schemas/access_yandex"}, {"$ref": "#/components/schemas/access_onetimepin"}, {"$ref": "#/components/schemas/access_cloudflare"}]}}}, "type": "object"}]}
```
