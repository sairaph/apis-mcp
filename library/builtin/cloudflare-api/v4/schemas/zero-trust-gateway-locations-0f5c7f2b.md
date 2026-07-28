---
title: zero-trust-gateway_locations
page_id: schema-zero-trust-gateway-locations-0f5c7f2b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_locations

```yaml
{"type": "object", "properties": {"client_default": {"$ref": "#/components/schemas/zero-trust-gateway_client-default"}, "created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "dns_destination_ips_id": {"$ref": "#/components/schemas/zero-trust-gateway_dns-destination-ips-id-read"}, "dns_destination_ipv6_block_id": {"$ref": "#/components/schemas/zero-trust-gateway_dns_destination_ipv6_block_id"}, "doh_subdomain": {"$ref": "#/components/schemas/zero-trust-gateway_subdomain"}, "ecs_support": {"$ref": "#/components/schemas/zero-trust-gateway_ecs-support"}, "endpoints": {"$ref": "#/components/schemas/zero-trust-gateway_endpoints"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}, "ip": {"$ref": "#/components/schemas/zero-trust-gateway_ip"}, "ipv4_destination": {"description": "Show the primary destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.", "type": "string", "example": "172.64.36.1", "readOnly": true, "x-stainless-terraform-configurability": "computed"}, "ipv4_destination_backup": {"description": "Show the backup destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.", "type": "string", "example": "172.64.36.2", "readOnly": true, "x-stainless-terraform-configurability": "computed"}, "max_ttl": {"$ref": "#/components/schemas/zero-trust-gateway_max-ttl"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-2"}, "networks": {"$ref": "#/components/schemas/zero-trust-gateway_ipv4_networks"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}}}
```
