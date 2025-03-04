// Basic gui to capture user input and display output from rediscli
package gui

import (
	"fmt"
	"strings"

	"github.com/aria-afk/rediscli/redis"
	"github.com/pkg/term"
	goredis "github.com/redis/go-redis/v9"
	"golang.design/x/clipboard"
)

type GUI struct {
	db       *redis.Redis
	address  string
	usertext []string
}

func NewGUI(db *redis.Redis) GUI {
	g := GUI{
		db: db,
	}
	parsed, _ := goredis.ParseURL(g.db.Opts.URI)
	g.address = fmt.Sprintf("%s/%d", parsed.Addr, parsed.DB)
	return g
}

// Start render loop
func (g *GUI) Run() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	g.initialRender()
	g.renderUserLine()
	for {
		userInput := getUserInput()

		// exit conditions
		if userInput == ctrl_c {
			return
		}
		// paste NOTE: this enforces ctrl_v across all users
		// eventually can refactor getUserInput to accept larger byte buffers than 3
		if userInput == ctrl_v {
			clipboardText := clipboard.Read(clipboard.FmtText)
			g.usertext = append(g.usertext, string(clipboardText))
			g.renderUserLine()
			continue
		}
		// delete a character
		if userInput == backspace {
			if len(g.usertext) > 0 {
				g.usertext = g.usertext[0 : len(g.usertext)-1]
				g.renderUserLine()
				continue
			}
		}
		// dispatch a command
		if userInput == enter {
			fmt.Println()
			g.usertext = make([]string, 0)
			g.renderUserLine()
			continue
		}

		_, reserved := reservedKeys[userInput]
		if !reserved {
			g.usertext = append(g.usertext, string(userInput))
		}
		g.renderUserLine()
	}
}

// Renders the user line with the users given input
func (g *GUI) renderUserLine() {
	fmt.Print("\033[2K")
	fmt.Printf("\r[%s] > %s", g.address, strings.Join(g.usertext, ""))
}

// Renders initial state of GUI
func (g *GUI) initialRender() {
	// get the uri(s) connected
	fmt.Printf("Successfully connected to redis :: %s\n", g.address)
}

// map for key codes in bytes
var (
	ctrl_c    byte = 3
	enter     byte = 13
	ctrl_v    byte = 22
	backspace byte = 127
)

var reservedKeys = map[byte]bool{
	3:   true,
	13:  true,
	22:  true,
	127: true,
}

func getUserInput() byte {
	// TODO: Spec for all distros/os
	t, err := term.Open("/dev/tty")
	if err != nil {
		panic(err)
	}
	term.RawMode(t)

	var action int
	bytes := make([]byte, 3)
	action, _ = t.Read(bytes)

	t.Restore()
	t.Close()

	if action == 3 {
		return bytes[2]
	} else {
		return bytes[0]
	}
}
