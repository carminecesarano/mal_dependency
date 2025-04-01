package clipboard_read

import (
	"fmt"

	"golang.design/x/clipboard"
)

func ClipboardRead() {

	err_clip := clipboard.Init()
	if err_clip != nil {
		fmt.Println("Errore nell'inizializzare la clipboard:", err_clip)
	} else {
		data := clipboard.Read(clipboard.FmtText)
		fmt.Println("Testo letto dalla clipboard:", string(data))
	}

}
