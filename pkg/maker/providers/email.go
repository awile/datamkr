package providers

import (
	"fmt"
	"math/rand"

	"github.com/awile/datamkr/pkg/dataset"
)

type emailFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewEmailWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &emailFieldMaker{definition: definition}
}

func (efm *emailFieldMaker) MakeField() ProviderField {
	emailValue := fmt.Sprintf("%s@%s", efm.getName(), efm.getDomain())
	return &emailField{value: emailValue}
}

func (efm *emailFieldMaker) getName() string {
	var firstNames []string = []string{
		"Mason", "Adria", "Dorian", "Thanh", "Janyce",
		"Janean", "Tony", "Laurel", "Kirsten", "Lashay",
		"Monika", "Lorrine", "Apolonia", "Agnus", "Chun",
		"Bree", "Keena", "Adan", "Jarvis", "Fumiko",
		"Olivia", "Denyse", "Tamela", "Lesli", "Ozie",
		"Katharina", "Shantay", "Reanna", "Lola", "Carolynn",
		"Levi", "Lurlene", "Laticia", "Kylee", "Wyatt",
		"Calvin", "Cicely", "Latashia", "Inell", "Shanita",
		"Raisa", "Lulu", "Ollie", "Tammara", "Jeanna",
		"Virgen", "Marlo", "Jonell", "Sanda", "Roseann", "Chad",
	}
	var lastNames []string = []string{
		"Wyman", "Marcinek", "Blomgren", "Losey", "Lucena", "Rashid",
		"Horak", "Fricke", "Depaolo", "Kreidler", "Weddell",
		"Sanabria", "Koons", "Curran", "Brittingham", "Hanover",
		"Marsee", "Fortuna", "Richarson", "Philpot", "Hora",
		"Cerda", "Tolleson", "Rowsey", "Schlachter", "Mcduffy",
		"Viviani", "Tharp", "Tierney", "Henriquez", "Grell",
		"Vansickle", "Westling", "Eastham", "Ralston", "Siers",
		"Reily", "Pinion", "Relyea", "Teaster", "Beyers",
		"Poovey", "Dumond", "Drinkard", "Dove", "Manney",
		"Condrey", "Banfield", "Rahimi", "Caryl", "Blais",
	}
	var separators []string = []string{
		"-", ".", "", "_",
	}
	return fmt.Sprintf("%s%s%s",
		firstNames[rand.Intn(len(firstNames))],
		separators[rand.Intn(len(separators))],
		lastNames[rand.Intn(len(lastNames))],
	)
}

func (efm *emailFieldMaker) getDomain() string {
	var domains []string = []string{
		"hotmail.com",
		"fakedomain.org",
		"uni.edu",
		"fake.com",
		"gmail.com",
	}
	return domains[rand.Intn(len(domains))]
}

type emailField struct {
	value string
}

func (email *emailField) String() string {
	return email.value
}

func (email *emailField) Value() interface{} {
	return email.value
}
