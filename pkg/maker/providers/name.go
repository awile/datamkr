package providers

import (
	"math/rand"
	"time"

	"github.com/awile/datamkr/pkg/dataset"
)

type nameFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewNameWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &nameFieldMaker{definition: definition}
}

func (nfm *nameFieldMaker) MakeField() interface{} {
	const NUMBER_OF_NAMES = 50
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(NUMBER_OF_NAMES)
	return nfm.getName(i)
}

func (nfm *nameFieldMaker) getName(index int) string {
	var names []string = []string{
		"Mason Wyman",
		"Adria Marcinek",
		"Dorian Blomgren",
		"Thanh Losey",
		"Janyce Lucena",
		"Janean Rashid",
		"Tony Horak",
		"Laurel Fricke",
		"Kirsten Depaolo",
		"Lashay Kreidler",
		"Monika Weddell",
		"Lorrine Sanabria",
		"Apolonia Koons",
		"Agnus Curran",
		"Chun Brittingham",
		"Bree Hanover",
		"Keena Marsee",
		"Adan Fortuna",
		"Jarvis Richarson",
		"Fumiko Philpot",
		"Olivia Hora",
		"Denyse Cerda",
		"Tamela Tolleson",
		"Lesli Rowsey",
		"Ozie Schlachter",
		"Katharina Mcduffy",
		"Shantay Viviani",
		"Reanna Tharp",
		"Lola Tierney",
		"Carolynn Henriquez",
		"Levi Grell",
		"Lurlene Vansickle",
		"Laticia Westling",
		"Kylee Eastham",
		"Wyatt Ralston",
		"Calvin Siers",
		"Cicely Reily",
		"Latashia Pinion",
		"Inell Relyea",
		"Shanita Teaster",
		"Raisa Beyers",
		"Lulu Poovey",
		"Ollie Dumond",
		"Tammara Drinkard",
		"Jeanna Dove",
		"Virgen Manney",
		"Marlo Condrey",
		"Jonell Banfield",
		"Sanda Rahimi",
		"Roseann Caryl",
		"Chad Blais",
	}
	return names[index]
}
