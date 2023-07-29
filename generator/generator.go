package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

type DishCategory struct {
	id   int
	name string
}

// TODO: ideally should also define amount of gas required to cook it, but that might be not that easy
// 	what makes more sense is an attempt to have groups of dishes, relations to drinks etc. to generate appropriate combinations.
// 	can we use categories for this matching?

type Dish struct {
	categories  []DishCategory
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
			breakfast: generateBreakfast(),
			lunch:     generateLunch(),
			dinner:    generateDinner(),
			snack:     generateSnacks(),
		}

		sumIngredients(ingrMap, day.breakfast, params)
		sumIngredients(ingrMap, day.lunch, params)
		sumIngredients(ingrMap, day.dinner, params)
		sumIngredients(ingrMap, day.snack, params)

		menu.days[i] = day
	}

	return menu
}

func generateBreakfast() []Dish {
	dish, err := getDish(BreakfastDish)
	if err != nil {
		fmt.Println("Failed to generate breakfast: " + err.Error())
		return nil
	}

	return []Dish{dish}
}

func generateLunch() []Dish {
	dish, err := getDish(LunchDish)
	if err != nil {
		fmt.Println("Failed to generate lunch: " + err.Error())
		return nil
	}

	return []Dish{dish}
}

func generateDinner() []Dish {
	dish, err := getDish(DinnerDish)
	if err != nil {
		fmt.Println("Failed to generate dinner: " + err.Error())
		return nil
	}

	return []Dish{dish}
}

func generateSnacks() []Dish {
	dish, err := getDish(SnackDish)
	if err != nil {
		fmt.Println("Failed to generate snacks: " + err.Error())
		return nil
	}

	return []Dish{dish}
}

func sumIngredients(totals map[string]DishIngredient, dishes []Dish, params InputParams) {
	for _, dish := range dishes {
		for _, ingr := range dish.ingredients {
			totalIngr, exists := totals[ingr.product.name]
			if !exists {
				totalIngr = DishIngredient{
					product: Product{name: ingr.product.name},
					amount:  0.0,
				}
			}

			totalIngr.amount += ingr.amount * float32(params.People)
		}
	}
}

// IDEA: what if we mark some ingredients as optional and let users exclude them? Is there a sense at all?

// getDish is supposed to get a random dish for the specified category
func getDish(cat DishCategory) (Dish, error) {
	dishes := FoodOfCategory(cat)
	foundCount := len(dishes)
	if foundCount == 0 {
		return Dish{}, errors.New("cannot find dish for the specified category")
	}

	rand.Seed(time.Now().Unix())

	return dishes[rand.Intn(foundCount)], nil
}
