module src.creational

go 1.20

replace patterns/abstractfactory => ./src/creational/AbstractFactory

require (
	patterns/abstractfactory v0.0.0-00010101000000-000000000000
	patterns/builder v0.0.0-00010101000000-000000000000
)

replace patterns/builder => ./src/creational/Builder
