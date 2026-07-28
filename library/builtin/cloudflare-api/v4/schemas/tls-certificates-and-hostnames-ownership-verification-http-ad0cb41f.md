---
title: tls-certificates-and-hostnames_ownership_verification_http
page_id: schema-tls-certificates-and-hostnames-ownership-verification-http-ad0cb41f
path: schemas
description: This presents the token to be served by the given http url to activate a hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_ownership_verification_http

This presents the token to be served by the given http url to activate a hostname.

```yaml
{"description": "This presents the token to be served by the given http url to activate a hostname.", "type": "object", "oneOf": [{"properties": {"http_body": {"description": "Token to be served.", "type": "string", "example": "5cc07c04-ea62-4a5a-95f0-419334a875a4"}, "http_url": {"description": "The HTTP URL that will be checked during custom hostname verification and where the customer should host the token.", "type": "string", "example": "http://custom.test.com/.well-known/cf-custom-hostname-challenge/0d89c70d-ad9f-4843-b99f-6cc0252067e9"}}, "type": "object"}]}
```
