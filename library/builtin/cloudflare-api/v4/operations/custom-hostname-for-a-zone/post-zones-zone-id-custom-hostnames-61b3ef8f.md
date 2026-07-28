---
title: Create Custom Hostname
page_id: operation-post-zones-zone-id-custom-hostnames-7ebc287e
path: operations/custom-hostname-for-a-zone
description: Add a new custom hostname and request that an SSL certificate be issued for it. One of three validation methods—http, txt, email—should be used, with 'http' recommended if the CNAME is already in place (or will be soon). Specifying 'email' will send an email to the WHOIS contacts on file for the base domain plus hostmaster, postmaster, webmaster, admin, administrator. If http is used and the domain is not already pointing to the Managed CNAME host, the PATCH method must be used once it is (to complete validation).  Enable bundling of certificates using the custom_cert_bundle field. The bundling process requires the following condition One certificate in the bundle must use an RSA, and the other must use an ECDSA.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/custom_hostnames
operation_ids:
    - custom-hostname-for-a-zone-create-custom-hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Custom Hostname

`POST /zones/{zone_id}/custom_hostnames`

Operation ID: `custom-hostname-for-a-zone-create-custom-hostname`

Add a new custom hostname and request that an SSL certificate be issued for it. One of three validation methods—http, txt, email—should be used, with 'http' recommended if the CNAME is already in place (or will be soon). Specifying 'email' will send an email to the WHOIS contacts on file for the base domain plus hostmaster, postmaster, webmaster, admin, administrator. If http is used and the domain is not already pointing to the Managed CNAME host, the PATCH method must be used once it is (to complete validation).  Enable bundling of certificates using the custom_cert_bundle field. The bundling process requires the following condition One certificate in the bundle must use an RSA, and the other must use an ECDSA.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-create-custom-hostname", "summary": "Create Custom Hostname", "description": "Add a new custom hostname and request that an SSL certificate be issued for it. One of three validation methods—http, txt, email—should be used, with 'http' recommended if the CNAME is already in place (or will be soon). Specifying 'email' will send an email to the WHOIS contacts on file for the base domain plus hostmaster, postmaster, webmaster, admin, administrator. If http is used and the domain is not already pointing to the Managed CNAME host, the PATCH method must be used once it is (to complete validation).  Enable bundling of certificates using the custom_cert_bundle field. The bundling process requires the following condition One certificate in the bundle must use an RSA, and the other must use an ECDSA.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"custom_metadata": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_metadata"}, "custom_origin_server": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_origin_server"}, "custom_origin_sni": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_origin_sni"}, "hostname": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_post"}, "ssl": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_sslpost"}}, "required": ["hostname"]}}}}, "responses": {"200": {"description": "Create Custom Hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}}}}, "4XX": {"description": "Create Custom Hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
