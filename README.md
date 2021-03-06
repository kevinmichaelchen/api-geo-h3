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

### What resolution do we use?

Brooklyn is 250 sq km (one cell at Resolution 5)...

Williamsburg is 5 sq km (one cell at Resolution 7)...

Each finer-resolution cell is 7 times smaller than its coarser parent.

H3 supports [multiple resolutions](https://h3geo.org/docs/core-library/restable):
| H3 Resolution | Average Hexagon Area (km 2) | Average Hexagon Edge Length (km) | Number of unique indexes |
|---------------|-----------------------------|----------------------------------|--------------------------|
| 0 | 4,250,546.8477000 | 1,107.712591000 | 122 |
| 1 | 607,220.9782429 | 418.676005500 | 842 |
| 2 | 86,745.8540347 | 158.244655800 | 5,882 |
| 3 | 12,392.2648621 | 59.810857940 | 41,162 |
| 4 | 1,770.3235517 | 22.606379400 | 288,122 |
| 5 | 252.9033645 | 8.544408276 | 2,016,842 |
| 6 | 36.1290521 |3.229482772 | 14,117,882 |
| 7 | 5.1612932 | 1.220629759 | 98,825,162 |
| 8 | 0.7373276 | 0.461354684 | 691,776,122 |
| 9 | 0.1053325 | 0.174375668 | 4,842,432,842 |
| 10 | 0.0150475 | 0.065907807 | 33,897,029,882 |
| 11 | 0.0021496 | 0.024910561 | 237,279,209,162 |
| 12 | 0.0003071 | 0.009415526 | 1,660,954,464,122 |
| 13 | 0.0000439 | 0.003559893 | 11,626,681,248,842 |
| 14 | 0.0000063 | 0.001348575 | 81,386,768,741,882 |
| 15 | 0.0000009 | 0.000509713 | 569,707,381,193,162 |