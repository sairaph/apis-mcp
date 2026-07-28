---
title: access_proxy_endpoint_props
page_id: schema-access-proxy-endpoint-props-10d7ce19
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_proxy_endpoint_props

```yaml
{"allOf": [{"$ref": "#/components/schemas/access_feature_app_props"}, {"properties": {"domain": {"description": "The proxy endpoint domain in the format: 10 alphanumeric characters followed by .proxy.cloudflare-gateway.com", "example": "abcd123456.proxy.cloudflare-gateway.com", "pattern": "^[A-Za-z0-9]{10}\\.proxy\\.cloudflare-gateway\\.com$"}, "name": {"example": "Gateway Proxy", "default": "Gateway Proxy"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "proxy_endpoint"}]}}}]}
```
