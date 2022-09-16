package main

import (
	"sync"

	application "github.com/lulingua/hexa-arch/application"
	"github.com/lulingua/hexa-arch/domain"
	handler "github.com/lulingua/hexa-arch/infrastructure/http"
	slice "github.com/lulingua/hexa-arch/infrastructure/repositories/slice"
)

func main() {

	p1 := domain.Person{
		Name:     "Isaac",
		LastName: "Asimov",
	}

	p2 := domain.Person{
		Name:     "stanislaw",
		LastName: "lem",
	}

	p3 := domain.Person{
		Name:     "arthur C.",
		LastName: "Clarck",
	}

	memoryDB := slice.NewInMemoryDB([]domain.Person{
		p1,
		p2,
		p3,
	})

	personService := application.NewPersonaService(memoryDB)

	muxHandler := handler.NewGorillaHandler(personService)

	//handler.NewRouter()

	wg := sync.WaitGroup{}
	wg.Add(2)

	muxHandler.SetupRoutes()

	go func() {
		muxHandler.RunServer("default")
		wg.Done()
	}()

	go func() {
		muxHandler.RunServer(":9000")
		wg.Done()
	}()

	wg.Wait()
}
