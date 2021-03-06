package main

import (
	"fmt"
	"testing"

	"github.com/miekg/dns"
)

func TestBundDE(t *testing.T) {
	res := Result{}
	res.checkPath("bund.de")
	res.writeResult("outest.json")
}

func TestOutputANYwithDNSSECrrs(t *testing.T) {
	fqdn := "bsi.de"
	m := dnssecQuery(fqdn, dns.TypeA, "")
	fmt.Printf("\nAnswer Section: \n")
	for _, x := range m.Answer {
		fmt.Printf("%v\n", x)
	}
	fmt.Printf("\nNs Section: \n")
	for _, x := range m.Ns {
		fmt.Printf("%v\n", x)
	}
	fmt.Printf("\nExtra Section: \n")
	for _, x := range m.Extra {
		fmt.Printf("%v\n", x)
	}
	fmt.Printf("\n")
}

func TestCheckZSKverifiability(t *testing.T) {
	fqdn := "bund.de"
	x := checkZSKverifiability(fqdn)
	t.Log(x)
}

func TestCheckPath(t *testing.T) {
	r := Result{}
	r.checkPath("bsi.de")
}

func TestBsiDE(t *testing.T) {
	m := dnssecQuery("bsi.de", dns.TypeANY, "")
	if &m == nil {
		t.Error("No response from dnssecQuery()")
	}
	for _, x := range m.Answer {
		fmt.Printf("%v", x)
	}
	for _, x := range m.Ns {
		fmt.Printf("%v", x)
	}
	for _, x := range m.Extra {
		fmt.Printf("%v", x)
	}
}

func TestGetDS(t *testing.T) {
	m := dnssecQuery("bsi.de", dns.TypeDS, "")
	for _, x := range m.Answer {
		fmt.Printf("%v\n", x)
	}
	for _, x := range m.Ns {
		fmt.Printf("%v\n", x)
	}
	for _, x := range m.Extra {
		fmt.Printf("%v\n", x)
	}
}

/*func TestGetKSK(t *testing.T) {
	m := getDNSKEYs("bsi.de", 257)
	fmt.Printf("%v\n", m[0].KeyTag())
}*/

/*func TestMakeZones(t *testing.T) {
	y := makeZones("example.higher.tld")
	for _, x := range y {
		fmt.Printf(x)
		//	t.Log(x)
	}
}*/

func TestMakeQuery(t *testing.T) {
	m := dnssecQuery("bsi.de", dns.TypeDNSKEY, "")
	for _, x := range m.Answer {
		fmt.Printf("%v\n", x)
	}
}

func TestListOfDNSresolvers(t *testing.T) {
	servers := []string{"185.48.116.10", "185.48.118.6", "8.8.8.8", "8.8.4.4", "9.9.9.10", "4.2.2.1", "4.2.2.2", "4.2.2.3"}
	results := make([]Result, len(servers))
	for i := range servers {
		results[i].checkPath("bund.de")
	}
}
