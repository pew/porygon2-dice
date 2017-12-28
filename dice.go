package bot

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"math/rand"
	"strconv"
	"strings"
	"log"
	"time"
)

func dice(command *bot.Cmd, matches []string) (msg string, err error) {
	seed := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(seed)

	var s []string

	for i := 0; i < 5; i++ {
		s = append(s, strconv.Itoa(rn.Intn(6)+1))
	}

    return fmt.Sprintf(strings.Join(s, " ")), nil
}

func init() {
	bot.RegisterCommand("^dice$", dice)
	log.Println("dice plugin loaded")
}
