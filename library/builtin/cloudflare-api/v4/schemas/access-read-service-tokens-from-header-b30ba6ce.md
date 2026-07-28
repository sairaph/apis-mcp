---
title: access_read_service_tokens_from_header
page_id: schema-access-read-service-tokens-from-header-b30ba6ce
path: schemas
description: |-
    Allows matching Access Service Tokens passed HTTP in a single header with this name.
    This works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.
    The header value will be interpreted as a json object similar to:
      {
        "cf-access-client-id": "88bf3b6d86161464f6509f7219099e57.access.example.com",
        "cf-access-client-secret": "bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5"
      }
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_read_service_tokens_from_header

Allows matching Access Service Tokens passed HTTP in a single header with this name.
This works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.
The header value will be interpreted as a json object similar to:
  {
    "cf-access-client-id": "88bf3b6d86161464f6509f7219099e57.access.example.com",
    "cf-access-client-secret": "bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5"
  }

```yaml
{"description": "Allows matching Access Service Tokens passed HTTP in a single header with this name.\nThis works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.\nThe header value will be interpreted as a json object similar to:\n  {\n    \"cf-access-client-id\": \"88bf3b6d86161464f6509f7219099e57.access.example.com\",\n    \"cf-access-client-secret\": \"bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5\"\n  }\n", "type": "string", "example": "Authorization"}
```
