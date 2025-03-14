package story

import (
	"fmt"
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Wald///////////////////////////

// Spieler wird vom Bär angegriffen
// Gets playerstats, Kampfalgoritmus
// Returns playerstats nach Kampf
func Baerangriff(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 3)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Speiler heilt am Lagerfeuer
// Gets PlayerHp, PlayerMaxHp
// Returns PlayerHp mit addiertem ein fünftel vom PlayerMaxHp
func Bonfire(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du findest einen Lagerfeuer und ruhst dich aus")

	hp = hp + (maxHp / 5)

	return hp
}

// Spieler darf auswählen ob er einen Pilz isst
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach essen oder nicht essen vom Pilz
func GesunderPilz(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Du findest einen Pilz?")

	// Auswahl Pilz essen oder nicht
	fmt.Println("Möchtest du ihn essen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		hp = hp + (maxHp / 7)

		text.Print("Du hast den Pilz gegessen und etwas regeneriert")
	} else {
		text.Print("Du hast den Pilz nicht gegessen")
	}
	return hp
}

////////////////Funktionen für Burg///////////////////////////

////////////////Funktionen für Dorf///////////////////////////
