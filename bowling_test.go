package bowling

import (
	"testing"
)

func (g *Game) rollMany(runs ,pin int){
	for i:=0;i<runs;i++{
		g.Roll(pin)
	}
}
func (g *Game) rollSpare() {
	g.Roll(5)
	g.Roll(5)
}

func (g *Game) rollStrike() {
	g.Roll(10)
}
func TestGutterGame(t *testing.T){
	game :=new(Game)
	game.rollMany(20,0)
	score := game.Score()
	if score != 0 {
		t.Error("Gutter game:", "Expected", 0, "was", score)
	}
}
func TestAllOne(t *testing.T){
	game := new(Game)
	game.rollMany(20,1)
	score := game.Score()
	if score != 20 {
		t.Error("All one game:", "Expected", 20, "was", score)
	}
}
func TestOneSpare(t *testing.T) {
	game := new(Game)

	game.rollSpare()
	game.Roll(3)
	game.rollMany(17, 0)

	score := game.Score()

	if score != 16 {
		t.Error("One spare:", "Expected", 16, "was", score)
	}
}
func TestOneStrike(t *testing.T){
	game:=new(Game)
	game.rollStrike()
	game.Roll(3)
	game.Roll(4)
	game.rollMany(16,0)
	score :=game.Score()
	if score !=24{
		t.Error("one strike:","expected",24,"was",score)
	}
}
func TestPerfectGame(t *testing.T){
	game:=new(Game)
	game.rollMany(12,10)
	score :=game.Score()

	if score !=300{
		t.Error("Perfect Game:","expected",300, "was",score)
	}
}