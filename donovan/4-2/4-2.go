package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
  shaFlag := flag.String("sha", "256", "sha type")

  flag.Parse()

  if len(os.Args) >= 1 && flag.Arg(0) != "" {
    switch *shaFlag {
    case "384":
      fmt.Printf("%x 384\n", sha512.Sum384([]byte(flag.Arg(0))))
    case "512":
      fmt.Printf("%x 512\n", sha512.Sum512([]byte(flag.Arg(0))))
    case "256":
      fmt.Printf("%x 256\n", sha256.Sum256([]byte(flag.Arg(0))))
    }
  } else {
    fmt.Println("No input data")
  }
}
