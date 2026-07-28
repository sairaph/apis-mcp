---
title: Import DNS Records
page_id: operation-post-zones-zone-id-dns-records-import-3ae16940
path: operations/dns-records-for-a-zone
description: |-
    You can upload your [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this endpoint. It assumes that cURL is called from a location with bind_config.txt (valid BIND config) present.

    See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records") for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records/import
operation_ids:
    - dns-records-for-a-zone-import-dns-records
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Import DNS Records

`POST /zones/{zone_id}/dns_records/import`

Operation ID: `dns-records-for-a-zone-import-dns-records`

You can upload your [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this endpoint. It assumes that cURL is called from a location with bind_config.txt (valid BIND config) present.

See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records") for more information.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-import-dns-records", "summary": "Import DNS Records", "description": "You can upload your [BIND config](https://en.wikipedia.org/wiki/Zone_file \"Zone file\") through this endpoint. It assumes that cURL is called from a location with bind_config.txt (valid BIND config) present.\n\nSee [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ \"Import and export records\") for more information.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"description": "BIND config to import.\n\n**Tip:** When using cURL, a file can be uploaded using `--form 'file=@bind_config.txt'`.\n", "type": "string", "example": "www.example.com. 300 IN  A 127.0.0.1"}, "proxied": {"description": "Whether or not proxiable records should receive the performance and security benefits of Cloudflare.\n\nThe value should be either `true` or `false`.", "type": "string", "example": "true", "default": "false"}}, "required": ["file"]}}}}, "responses": {"200": {"description": "Import DNS Records response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_import_scan"}}}}, "4XX": {"description": "Import DNS Records response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_import_scan"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "import"}
```
