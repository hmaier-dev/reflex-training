package main

import (
	"fmt"
	"math/rand"
	"time"
    "flag"
)

type Aktion struct {
    Name string
    Verteidigungen []string
}

var aktionen = []Aktion{
    { 
        Name: "Jab",
        Verteidigungen: []string{"Doppeldeckung", "Winkekatze"},
    },
    { 
        Name: "Punch",
        Verteidigungen: []string{"Doppeldeckung", "Block"},
    },
    { 
        Name: "hinterer Push-Kick",
        Verteidigungen: []string{"Parieren", "Wegwischen", "Ausweichen"},
    },
    { 
        Name: "vorderer Push-Kick",
        Verteidigungen: []string{"Parieren", "Wegwischen", "Ausweichen"},
    },
    { 
        Name: "vorderer Haken",
        Verteidigungen: []string{"Block"},
    },
    { 
        Name: "hinterer Haken",
        Verteidigungen: []string{"Block"},
    },
    { 
        Name: "vorderer Ribkick",
        Verteidigungen: []string{"Dreipunkt-Block", "Schienbein"},
    },
    { 
        Name: "hinterer Ribkick",
        Verteidigungen: []string{"Dreipunkt-Block", "Schienbein"},
    },
}

var maximum int

func init(){
    flag.IntVar(&maximum ,"random-range", 5, "display action in range from 1 to n seconds")
    flag.Parse()
}

func main(){

    for {
        sleeping := rand.Intn(maximum - 1)+1 
        fmt.Println(sleeping)
        time.Sleep(time.Second * time.Duration(sleeping))
        aktion := aktionen[rand.Intn(len(aktionen))]
        fmt.Println("Verteidige ",  aktion.Name, " mit ", aktion.Verteidigungen[rand.Intn(len(aktion.Verteidigungen))])
    }

}
