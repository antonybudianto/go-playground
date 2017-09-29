package main

import "fmt"

type IPAddr [4]byte
const iplen = 4

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
  str := ""
  for i := 0; i < iplen; i++ {
  	str += fmt.Sprintf("%d", ip[i])
	if i < iplen-1 {
	  str += "."
	}
  }
  return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
