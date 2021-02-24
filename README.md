# goa-snmp

Goa micro service for snmp

## Design

This Service will take an ip address or hostname, and snmp public/private, walk the tree and save
the mib/oid information for http retrieval.

`http://localhost/goa-snmp/device/{mame/ip}`
`http://localhost/goa-snmp/device/{name}/mibs`