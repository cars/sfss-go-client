module github.com/cars/sfss-go-client

go 1.22.2

require github.com/cars/sfss-go-client/sfssrest v0.0.0-20231002120000-abcdef123456
require github.com/cars/sfss-go-client/sfssapp v0.0.0-20231002120000-abcdef123456

replace github.com/cars/sfss-go-client/sfssapp => ./sfssapp
replace github.com/cars/sfss-go-client/sfssrest => ./sfssrest
