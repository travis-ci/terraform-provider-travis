package travis

import (
	"fmt"
	"net"
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTravisExpandedCIDR() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTravisExpandedCIDRRead,
		Schema: map[string]*schema.Schema{
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addrs": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceTravisExpandedCIDRRead(d *schema.ResourceData, meta interface{}) error {
	cidr := d.Get("cidr").(string)
	ips, err := netIPs(cidr)
	if err != nil {
		return fmt.Errorf("error parsing cidr %q: %s", cidr, err)
	}

	sort.Strings(ips)
	d.Set("addrs", ips)
	d.SetId(cidr)

	return nil
}

func netIPs(cidr string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	ips := []string{}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); nextIp(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func nextIp(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
