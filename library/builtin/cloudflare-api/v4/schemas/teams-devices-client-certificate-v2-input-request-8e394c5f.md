---
title: teams-devices_client_certificate_v2_input_request
page_id: schema-teams-devices-client-certificate-v2-input-request-8e394c5f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_client_certificate_v2_input_request

```yaml
{"type": "object", "properties": {"certificate_id": {"description": "UUID of Cloudflare managed certificate.", "type": "string", "example": "b14ddcc4-bcd2-4df4-bd4f-eb27d5a50c30", "maxLength": 36}, "check_private_key": {"description": "Confirm the certificate was not imported from another device. We recommend keeping this enabled unless the certificate was deployed without a private key.", "type": "boolean", "example": true}, "cn": {"description": "Certificate Common Name. This may include one or more variables in the ${ } notation. Only ${serial_number} and ${hostname} are valid variables.", "type": "string", "example": "${hostname}.com.${serial_number}"}, "extended_key_usage": {"description": "List of values indicating purposes for which the certificate public key can be used.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_extended_key_usage_enum"}, "example": ["clientAuth", "emailProtection"]}, "locations": {"type": "object", "properties": {"paths": {"$ref": "#/components/schemas/teams-devices_paths"}, "trust_stores": {"$ref": "#/components/schemas/teams-devices_trust_stores"}}}, "operating_system": {"description": "Operating System.", "type": "string", "example": "windows", "enum": ["windows", "mac", "linux"]}, "subject_alternative_names": {"description": "List of certificate Subject Alternative Names.", "type": "array", "items": {"type": "string"}, "example": ["example.com", "sample.com"]}}, "required": ["certificate_id", "check_private_key", "operating_system"], "title": "Client Certificate V2"}
```
