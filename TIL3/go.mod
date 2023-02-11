module main

go 1.20

replace module => ./module

replace myConst => ./myConst

require (
	myConst v0.0.0-00010101000000-000000000000 // indirect
	module v0.0.0-00010101000000-000000000000 // indirect
)
