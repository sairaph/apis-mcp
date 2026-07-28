---
title: cc_DNSConfiguration
page_id: schema-cc-dnsconfiguration-9667758a
path: schemas
description: |-
    Represents the /etc/resolv.conf that will appear in the deployment.
    If the 'dns' property is specified, even if empty object, will override the default resolv.conf of the container.
    The default resolv.conf of a container is 'servers = ["1.1.1.1", "9.9.9.9", "2606:4700:4700::1111"]', only if an IPv4 is assigned.
    The default for a non IPv4 deployment is 'servers = ["2606:4700:4700::1111", "2620:fe::fe"]'.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_DNSConfiguration

Represents the /etc/resolv.conf that will appear in the deployment.
If the 'dns' property is specified, even if empty object, will override the default resolv.conf of the container.
The default resolv.conf of a container is 'servers = ["1.1.1.1", "9.9.9.9", "2606:4700:4700::1111"]', only if an IPv4 is assigned.
The default for a non IPv4 deployment is 'servers = ["2606:4700:4700::1111", "2620:fe::fe"]'.

```yaml
{"description": "Represents the /etc/resolv.conf that will appear in the deployment.\nIf the 'dns' property is specified, even if empty object, will override the default resolv.conf of the container.\nThe default resolv.conf of a container is 'servers = [\"1.1.1.1\", \"9.9.9.9\", \"2606:4700:4700::1111\"]', only if an IPv4 is assigned.\nThe default for a non IPv4 deployment is 'servers = [\"2606:4700:4700::1111\", \"2620:fe::fe\"]'.\n", "type": "object", "properties": {"searches": {"description": "The container resolver will append these domains to every resolve query. For example, if you have 'google.com',\nand your deployment queries 'web', it will append 'google.com' to 'web' in the search query before trying 'web'.\nLimited to 6 domains.\n", "type": "array", "items": {"$ref": "#/components/schemas/cc_Domain"}}, "servers": {"description": "List of DNS servers that the deployment will use to resolve domain names. You can only specify a maximum of 3.", "type": "array", "items": {"$ref": "#/components/schemas/cc_IP"}}}}
```
