---
title: tls-certificates-and-hostnames_custom_cert_and_key
page_id: schema-tls-certificates-and-hostnames-custom-cert-and-key-562fb83c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_cert_and_key

```yaml
{"type": "object", "properties": {"custom_certificate": {"description": "If a custom uploaded certificate is used.", "type": "string", "example": "-----BEGIN CERTIFICATE-----\nMIIDdjCCAl6gAwIBAgIJAPnMg0Fs+/B0MA0GCSqGSIb3DQEBCwUAMFsx...\n-----END CERTIFICATE-----\n"}, "custom_key": {"description": "The key for a custom uploaded certificate.", "type": "string", "example": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/SCB5...\n-----END PRIVATE KEY-----\n", "x-sensitive": true}}, "required": ["custom_certificate", "custom_key"]}
```
