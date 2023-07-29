package generator

var BreakfastDish = DishCategory{id: 1, name: "Breakfast"}
var SweetDish = DishCategory{id: 2, name: "Sweet"}
var SoupDish = DishCategory{id: 3, name: "Soup"}
var DishWithChicken = DishCategory{id: 4, name: "Chicken"}
var DinnerDish = DishCategory{id: 7, name: "Dinner"}
var LunchDish = DishCategory{id: 8, name: "Lunch"}
var PastaDish = DishCategory{id: 5, name: "Pasta"}
var NoMeatDish = DishCategory{id: 6, name: "No meat"}
var SnackDish = DishCategory{id: 9, name: "Snack"}

var Food = []Dish{
	{
		categories: []DishCategory{
			BreakfastDish,
			SweetDish,
		},
		name: "Oat porridge",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Oat flakes"},
				amount:  0.05,
			},
			{
				product: Product{name: "Raisin"},
				amount:  0.01,
			},
			{
				product: Product{name: "Milk powder"},
				amount:  0.005,
			},
			{
				product: Product{name: "Sugar"},
				amount:  0.005,
			},
		},
	},
	{
		categories: []DishCategory{
			SoupDish,
			DishWithChicken,
			DinnerDish,
			LunchDish,
		},
		name: "Chicken noodles",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Dried chicken"},
				amount:  0.01,
			},
			{
				product: Product{name: "Noodles"},
				amount:  0.03,
			},
			{
				product: Product{"Fried onion"},
				amount:  0.001,
			},
		},
	},
	{
		categories: []DishCategory{
			PastaDish,
			NoMeatDish,
			DinnerDish,
			LunchDish,
		},
		name: "Pasta with tomatoes and cheese",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Pasta"},
				amount:  0.05,
			},
			{
				product: Product{name: "Dried tomatoes"},
				amount:  0.01,
			},
			{
				product: Product{name: "Cheese"},
				amount:  0.02,
			},
		},
	},
	{
		categories: []DishCategory{
			SweetDish,
			SnackDish,
		},
		name: "Dried fruits",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Dried fruits of choice"},
				amount:  0.05,
			},
		},
	},
	{
		categories: []DishCategory{
			SnackDish,
		},
		name: "Nuts",
		ingredients: []DishIngredient{
			{
				product: Product{name: "Nuts of choice"},
				amount:  0.05,
			},
		},
	},
}

func FoodOfCategory(cat DishCategory) []Dish {
	var result []Dish

	for _, dish := range Food {
		for _, category := range dish.categories {
			if category.id == cat.id {
				result = append(result, dish)

				break
			}
		}
	}

	return result
}
