---
title: tls-certificates-and-hostnames_setting_id
page_id: schema-tls-certificates-and-hostnames-setting-id-ea75326a
path: schemas
description: |-
    The TLS Setting name.
    The value type depends on the setting:
    - `ciphers`: value is an array of cipher suite strings (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`).
    - `min_tls_version`: value is a TLS version string (`"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"`).
    - `http2`: value is `"on"` or `"off"`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_setting_id

The TLS Setting name.
The value type depends on the setting:
- `ciphers`: value is an array of cipher suite strings (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`).
- `min_tls_version`: value is a TLS version string (`"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"`).
- `http2`: value is `"on"` or `"off"`.

```yaml
{"description": "The TLS Setting name.\nThe value type depends on the setting:\n- `ciphers`: value is an array of cipher suite strings (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`).\n- `min_tls_version`: value is a TLS version string (`\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"`).\n- `http2`: value is `\"on\"` or `\"off\"`.", "type": "string", "enum": ["ciphers", "min_tls_version", "http2"], "x-auditable": true}
```
