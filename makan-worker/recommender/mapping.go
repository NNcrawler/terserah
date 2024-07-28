package recommender

var temperatureToFoodChoiceMap = map[string][]string{
	"cold":      {"hot red velvet latte", "coffee", "americano", "machiato", "matcha", "hot chocolate", "soup", "chicken rice carbonara", "chicken butter rice", "pisang goreng", "bakso", "mie"},
	"cool":      {"latte series", "chocolate", "croissant", "tiramisu", "kentang goreng", "roti bakar", "bakso"},
	"normal":    {"ice latte", "matcha", "cookies", "sandwich", "fruit salad", "gado-gado", "mie"},
	"hot":       {"ice latte", "red velvet", "chocolate", "taro", "cold brew coffee", "salad", "fruit", "es teler", "mie"},
	"super_hot": {"es campur", "ice cream", "iced tea", "es buah", "shaved ice", "sorbet"},
}

var timeOfDayToFoodChoiceMap = map[string][]string{
	"morning":   {"coffee", "bubur", "porridge", "roti", "bread", "croissant", "pisang goreng", "nasi uduk"},
	"afternoon": {"ice latte", "latte series", "machiato", "red velvet", "chocolate", "cookies", "sandwich", "chicken rice carbonara", "chicken butter rice", "taro", "kentang goreng", "gado-gado", "soto ayam", "bakso", "mie"},
	"night":     {"hot red velvet latte", "coffee", "americano", "machiato", "matcha", "tiramisu", "pisang goreng", "light dinner", "fruit", "nasi goreng", "sate", "sop buntut", "bakso", "mie"},
}

var conditionToFoodChoiceMap = map[string][]string{
	"clear":  {"ice latte", "matcha", "croissant", "salad", "es teler", "fruit salad"},
	"cloudy": {"americano", "cookies", "kentang goreng", "hot chocolate", "soup", "mie ayam", "bakso"},
	"rain":   {"hot red velvet latte", "soup", "chicken butter rice", "pisang goreng", "sop buntut", "mie goreng", "bakso"},
}
