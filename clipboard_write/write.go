package clipboard_read

import (
	"fmt"

	"golang.design/x/clipboard"
)

func ClipboardWrite() {

	clipboard.Init()

	text := "prova"

	clipboard.Write(clipboard.FmtText, []byte(text))
	fmt.Println("Testo copiato sulla clipboard:", text)
}
