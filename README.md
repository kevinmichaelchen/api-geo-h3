# api-geo-h3
This is a proof-of-concept component that wraps [H3](https://h3geo.org/),
a hexagonal hierarchical geospatial indexing system.

<img src="./dispatch.png" />

For dispatch, our geospatial needs can be satisfied in the simplest way by 
asking: given a trip pickup location, which (drivers') coordinates are the
closest.

This can be done with [`GeoToH3`](https://h3geo.org/docs/api/indexing/#geotoh3)
(for converting geographic coordinates to a hex cell index) and with
[`H3Distance`](https://h3geo.org/docs/api/traversal#h3distance) (which returns 
the distance in grid cells between two indexes).
