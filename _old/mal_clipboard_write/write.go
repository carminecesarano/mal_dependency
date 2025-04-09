package mal_clipboard_read

import "github.com/atotto/clipboard"

func ClipboardWrite() {

	text := "Hello, World!"
	clipboard.WriteAll(text)
}
