package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/FoxFurry/scholarlabs/virt/docker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	engine, err := docker.New(ctx)
	if err != nil {
		panic(err)
	}

	id, err := engine.Spin(ctx, "docker.io/library/alpine", "docker.io/library/alpine")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID: %s\n", id)

	fmt.Println("[STARTING TERMINAL]")

	con, err := engine.StartTerminal(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("[OPENING SCANNER]")

	scanner := bufio.NewScanner(con.GetReader())
	go func(sc *bufio.Scanner) {
		for sc.Scan() {
			fmt.Println("[OUT]:", sc.Text())
		}
	}(scanner)

	//var outBut, errBuf bytes.Buffer
	//go func(out, err *bytes.Buffer, reader io.Reader) {
	//	for {
	//		_, err := stdcopy.StdCopy(out, err, reader)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		fmt.Println("[OUT]:", out.String())
	//		fmt.Println("[ERR]:", errBuf.String())
	//	}
	//}(&outBut, &errBuf, con.Reader)

	fmt.Println("[STARTING INPUT ROUTINE]")
	bebra := bufio.NewScanner(os.Stdin)
	for bebra.Scan() {
		input := bebra.Text()
		if input == "exit" {
			break
		}

		fmt.Println("[IN]:", input)

		if _, err := con.GetConn().Write([]byte(input + "\n")); err != nil {
			panic(err)
		}
	}

	//reader := bufio.NewReader(os.Stdin)
	//
	//for {
	//	fmt.Print("[IN]: ")
	//	input, _ := reader.ReadString('\n')
	//
	//	if input == "exit" {
	//		break
	//	}
	//
	//	if _, err := con.Conn.Write([]byte(input + "\n")); err != nil {
	//		panic(err)
	//	}
	//}

	con.Close()
	cancel()
}
