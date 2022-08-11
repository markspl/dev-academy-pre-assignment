# Helsinki city bike app

This is pre-assignment for Solita Dev Academy Finland (fall 2022).

> Journey dataset:
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-05.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-06.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-07.csv

### To-Do

##### Recommended
- [ ] Import data from the CSV files to a database or in-memory storage
- [ ] Validate data before importing
- [x] Don't import journeys that lasted for less than ten seconds
- [x] Don't import journeys that covered distances shorter than 10 meters

After filtering short (time and distance) journeys, `2021-05.csv` input file's data dropped from `814676` to `784794` rows.