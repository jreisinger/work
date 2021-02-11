package main

import (
	"errors"

	"github.com/miekg/dns"
)

func lookup(fqdn, serverAddr string) (ipAddrs []string) {
	var cfqdn = fqdn // Don't modify the original.
	for {
		cnames, err := lookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue // We have to process the next CNAME.
		}
		ips, err := lookupA(cfqdn, serverAddr)
		if err != nil {
			break // There are no A records for this hostname.
		}
		for _, ip := range ips {
			ipAddrs = append(ipAddrs, ip)
		}
		break // We have processed all the results.
	}
	return
}

func lookupA(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	r, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(r.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, answer := range r.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func lookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return fqdns, err
	}
	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}
	return fqdns, nil
}
