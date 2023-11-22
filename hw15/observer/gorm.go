package observer

import (
	"log"
	nativeMysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormState struct {
	db *gorm.DB
}

func NewGormConnection() *gorm.DB {
	cfg := nativeMysql.Config{
		User:   "root",
		Passwd: "1qazXSW@",
		Net:    "tcp",
		Addr:   "0.0.0.0:3306",
		DBName: "playersdb",
	}
	gorm, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return gorm
}

func (g *GormState) ListPlayers() []Player {
	players := []Player{}
	trx := g.db.Find(&players)

	if trx.Error != nil {
		log.Println(trx.Error)
		return nil
	}
	return players
}

func (g *GormState) SavePlayer(p Player) Player {
	saved := g.db.Save(&p)
	if saved.Error != nil {
		log.Println(saved.Error)
		return p
	}

	return p
}

func (g *GormState) AddGameRoom(r *GameRoom) {
	trx := g.db.Create(r)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
}

func (g *GormState) FindPlayers(playerName string) []Player {
	found := []Player{}
	trx := g.db.Where("Name LIKE ?", playerName).Find(&found)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
	return found
}

func (g *GormState) AddPlayers(players []*Player) {
	trx := g.db.Create(players)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
}

func NewGormState(db *gorm.DB) *GormState {
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&GameRoom{})
	return &GormState{db}
}

func (g *GormState) PrinlnAllPlayers() {
	players := g.ListPlayers()
	for _, p := range players {
		log.Println(p)
	}
}

func (g *GormState) PrinlnAllRooms() {
	room := []GameRoom{}
	trx := g.db.Find(&room)

	if trx.Error != nil {
		log.Println(trx.Error)
		return
	}

	for _, r := range room {
		log.Println(r)
	}
}

func (g *GormState) DeleteAllUsers() {
    trx := g.db.Where(gorm.Expr("1")).Delete(&Player{})
    if trx.Error != nil {
        log.Println(trx.Error)
    }
}

func (g *GormState) DeleteAllRooms() {
    trx := g.db.Where(gorm.Expr("1")).Delete(&GameRoom{})
    if trx.Error != nil {
        log.Println(trx.Error)
    }
}

