package templates

import (
	"math/rand"
	"time"

	"github.com/luisfcofv/indexter/models"
)

func GetEventTemplates(world models.World) []models.Event {
	return []models.Event{
		getFirstTemplate(world),
		getSecondTemplate(world),
		getThirdTemplate(world),
		getFourthTemplate(world),
		getFifthTemplate(world),
	}
}

func getFirstTemplate(world models.World) models.Event {
	time := time.Now().Unix()
	rand.Seed(time)

	location := world.Locations[rand.Intn(len(world.Locations))]
	agent := world.Agents[rand.Intn(len(world.Agents))]
	goal := world.Goals[rand.Intn(len(world.Goals))]
	causation := world.Goals[rand.Intn(len(world.Goals))]

	protagonist := false
	if rand.Intn(2) == 1 {
		protagonist = true
	}

	event := models.Event{
		Name:        "Something is happening",
		Description: "Nobody knows what is going on, be ready for an adventure",
		Protagonist: protagonist,
		Agents:      []models.Agent{agent},
		Location:    location,
		Goal:        goal,
		Propagation: rand.Intn(2) + 1,
		Cause:       causation,
	}

	return event
}

func getSecondTemplate(world models.World) models.Event {
	event := models.Event{
		Name:        "The orcs are gathering",
		Description: "This unusal situation means they are looking for something valuable",
		Protagonist: false,
		Agents:      []models.Agent{world.Agents[4]}, // Traveler
		Location:    world.Locations[3],
		Goal:        world.Goals[1], // Find the treasure
		Propagation: 1,
		Cause:       world.Goals[2], // Collect coins
	}

	return event
}

func getThirdTemplate(world models.World) models.Event {
	event := models.Event{
		Name:        "The king has gathered important information",
		Description: "This information can solve an imporantant mistery",
		Protagonist: false,
		Agents:      []models.Agent{world.Agents[0], world.Agents[1]}, // The king & Witness
		Location:    world.Locations[1],                               // King
		Goal:        world.Goals[0],                                   // Talk to the king
		Propagation: 1,
		Cause:       world.Goals[4], // Rescue the princess
	}

	return event
}

func getFourthTemplate(world models.World) models.Event {
	event := models.Event{
		Name:        "Kill the dragon",
		Description: "The most feared creature in this world, this dragon has a broad, elongated body with a very long tail",
		Protagonist: true,
		Agents:      []models.Agent{world.Agents[3]}, // Wizard
		Location:    world.Locations[4],
		Goal:        world.Goals[3], // Fight the dragon
		Propagation: 1,
	}

	return event
}

func getFifthTemplate(world models.World) models.Event {
	time := time.Now().Unix()
	rand.Seed(time)

	possibleCities := []int{3, 4}
	cityID := rand.Intn(len(possibleCities))

	event := models.Event{
		Name:        "The dragon has taken the princess",
		Description: "The princess disappeared at midnight, some people saw the dragon close to the city last night",
		Agents:      []models.Agent{world.Agents[4]}, // Princess
		Location:    world.Locations[possibleCities[cityID]],
		Goal:        world.Goals[4], // Rescue the princess
		Propagation: 2,
		Cause:       world.Goals[3], // Fight the dragon
	}

	return event
}
