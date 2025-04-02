package main

import (
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"fmt"
	"os/exec"
	"log"
)

func main() {

	offuscatedData := []byte{120, 156, 18, 230, 97, 16, 72, 77, 206, 200, 87, 240, 72, 205, 201, 201, 87, 8, 207, 47, 202, 73, 1, 4, 0, 0, 255, 255, 52, 116, 6, 11}

	r, err := zlib.NewReader(bytes.NewReader(offuscatedData))
	if err != nil {
		log.Fatal("Errore durante la decompressione con zlib:", err)
	}
	var decodedBuf bytes.Buffer
	_, err = decodedBuf.ReadFrom(r)
	if err != nil {
		log.Fatal("Errore durante la lettura dei dati decompressi:", err)
	}

	var decodedCommand string
	dec := gob.NewDecoder(&decodedBuf)
	err = dec.Decode(&decodedCommand)
	if err != nil {
		log.Fatal("Errore durante la decodifica con Gob:", err)
	}

	cmd := exec.Command("sh", "-c", decodedCommand)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Errore nell'esecuzione del comando:", err)
	}

	fmt.Println(string(output)) 
}
