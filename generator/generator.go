package generator

type InputParams struct {
	Days   int
	People int
	// TODO: there may be also more details & specifics, like first\last day being partial, some preferences etc.
}

type Product struct {
	name string
}

type DishIngredient struct {
	product Product
	amount  float32 // The idea is to measure in kg/person
}

// TODO: ideally should also define amount of gas required to cook it, but that might be not that easy
// 	what makes more sense is an attempt to have groups of dishes, relations to drinks etc. to generate appropriate combinations.

type Dish struct {
	name        string
	ingredients []DishIngredient
}

type AdventureDay struct {
	number    int
	breakfast []Dish // Can include main dish and some additions, like a chocolate bar, cheese etc.
	lunch     []Dish
	dinner    []Dish
	snack     []Dish
}

type Menu struct {
	days    []AdventureDay
	summary []DishIngredient // TODO: is it wise to use the same type everywhere?
}

func GenerateMenu(params InputParams) Menu {
	menu := Menu{
		days: make([]AdventureDay, params.Days),
	}

	ingrMap := make(map[string]DishIngredient)

	for i := 0; i < params.Days; i++ {
		day := AdventureDay{
			number:    i + 1,
			breakfast: []Dish{getDish()},
			lunch:     []Dish{getDish()},
			dinner:    []Dish{getDish()},
		}

		// TODO: can be many
		for _, ingr := range day.breakfast[0].ingredients {
			// TODO: add if present, create otherwise
			ingrMap[ingr.product.name] = DishIngredient{
				product: Product{name: ingr.product.name},
				amount:  ingr.amount * float32(params.People),
			}
		}

		menu.days[i] = day
	}

	return menu
}

// getDish is supposed to get a random (not really) dish
func getDish() Dish {
	return Dish{
		name: "Yummy",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Onion"},
				amount:  0.1,
			},
			{
				product: Product{name: "Garlic"},
				amount:  0.1,
			},
		},
	}
}
