package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abbyck/husk/execcmd"
)

func Prompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		execcmd.ExecHost()

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Handle the execution of the input.
		if err = execcmd.ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
