---
title: Delete Custom CSR
page_id: operation-delete-zones-zone-id-custom-csrs-custom-csr-id-02f2cfa3
path: operations/custom-csrs-for-a-zone
description: Delete a custom Certificate Signing Request (CSR) and its associated private key. A CSR cannot be deleted if any custom certificates still reference it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/custom_csrs/{custom_csr_id}
operation_ids:
    - custom-csrs-for-a-zone-delete-custom-csr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Custom CSR

`DELETE /zones/{zone_id}/custom_csrs/{custom_csr_id}`

Operation ID: `custom-csrs-for-a-zone-delete-custom-csr`

Delete a custom Certificate Signing Request (CSR) and its associated private key. A CSR cannot be deleted if any custom certificates still reference it.

## Definition

```yaml
{"operationId": "custom-csrs-for-a-zone-delete-custom-csr", "summary": "Delete Custom CSR", "description": "Delete a custom Certificate Signing Request (CSR) and its associated private key. A CSR cannot be deleted if any custom certificates still reference it.", "parameters": [{"name": "custom_csr_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Custom CSR response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_id_only"}}}}, "4XX": {"description": "Delete Custom CSR response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_id_only"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
