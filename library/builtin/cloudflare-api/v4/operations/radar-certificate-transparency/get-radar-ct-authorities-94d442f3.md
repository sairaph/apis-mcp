---
title: List certificate authorities
page_id: operation-get-radar-ct-authorities-5f7664e8
path: operations/radar-certificate-transparency
description: Retrieves a list of certificate authorities.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/ct/authorities
operation_ids:
    - radar-get-certificate-authorities
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List certificate authorities

`GET /radar/ct/authorities`

Operation ID: `radar-get-certificate-authorities`

Retrieves a list of certificate authorities.

## Definition

```yaml
{"operationId": "radar-get-certificate-authorities", "summary": "List certificate authorities", "description": "Retrieves a list of certificate authorities.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"certificateAuthorities": {"type": "array", "items": {"properties": {"certificateRecordType": {"description": "Specifies the type of certificate in the trust chain.", "type": "string", "enum": ["ROOT_CERTIFICATE", "INTERMEDIATE_CERTIFICATE"]}, "country": {"description": "The two-letter ISO country code where the CA organization is based.", "type": "string", "example": "PT"}, "countryName": {"description": "The full country name corresponding to the country code.", "type": "string", "example": "Portugal"}, "name": {"description": "The full name of the certificate authority (CA).", "type": "string", "example": "MULTICERT Advanced Certification Authority 005"}, "owner": {"description": "The organization that owns and operates the CA.", "type": "string", "example": "MULTICERT"}, "parentName": {"description": "The name of the parent/root certificate authority that issued this intermediate certificate.", "type": "string", "example": "MULTICERT Root Certification Authority 01"}, "parentSha256Fingerprint": {"description": "The SHA-256 fingerprint of the parent certificate.", "type": "string", "example": "24EDD4E503A8D3FDB5FFB4AF66C887359901CBE687A5A0760D10A08EED99A7C3"}, "revocationStatus": {"description": "The current revocation status of a Certificate Authority (CA) certificate.", "type": "string", "enum": ["NOT_REVOKED", "REVOKED", "PARENT_CERT_REVOKED"]}, "sha256Fingerprint": {"description": "The SHA-256 fingerprint of the intermediate certificate.", "type": "string", "example": "24EDD4E503A8D3FDB5FFB4AF66C887359901CBE687A5A0760D10A08EED99A7C3"}}, "required": ["sha256Fingerprint", "name", "owner", "parentName", "parentSha256Fingerprint", "certificateRecordType", "country", "countryName", "revocationStatus"], "type": "object"}}}, "required": ["certificateAuthorities"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Certificate Transparency"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.ct.authorities", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
