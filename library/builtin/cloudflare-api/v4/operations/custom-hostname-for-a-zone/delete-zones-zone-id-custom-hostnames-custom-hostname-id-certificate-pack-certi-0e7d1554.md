---
title: Delete Single Certificate And Key For Custom Hostname
page_id: operation-delete-zones-zone-id-custom-hostnames-custom-hostname-id-certificate-pac-03a3ac53
path: operations/custom-hostname-for-a-zone
description: Delete a single custom certificate from a certificate pack that contains two bundled certificates. Deletion is subject to the following constraints. You cannot delete a certificate if it is the only remaining certificate in the pack. At least one certificate must remain in the pack.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}
operation_ids:
    - custom-hostname-for-a-zone-delete_single_certificate_and_key_in_a_custom_hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Single Certificate And Key For Custom Hostname

`DELETE /zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}`

Operation ID: `custom-hostname-for-a-zone-delete_single_certificate_and_key_in_a_custom_hostname`

Delete a single custom certificate from a certificate pack that contains two bundled certificates. Deletion is subject to the following constraints. You cannot delete a certificate if it is the only remaining certificate in the pack. At least one certificate must remain in the pack.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-delete_single_certificate_and_key_in_a_custom_hostname", "summary": "Delete Single Certificate And Key For Custom Hostname", "description": "Delete a single custom certificate from a certificate pack that contains two bundled certificates. Deletion is subject to the following constraints. You cannot delete a certificate if it is the only remaining certificate in the pack. At least one certificate must remain in the pack.", "parameters": [{"name": "custom_hostname_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"202": {"description": "Delete Single Certificate and Key In a Custom Hostname response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}}}}}, "4XX": {"description": "Delete Single Certificate and Key In a Custom Hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"properties": {"id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, "type": "object"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames.certificate-pack.certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
