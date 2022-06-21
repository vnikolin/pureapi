package pureapi

import "encoding/json"

// FetchArrayInfo
func (c *StorageClient) FetchArrayInfo() (ArrayInfo, error) {

	url := c.formatPath("array")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return ArrayInfo{}, err
	}

	var array ArrayInfo

	json.Unmarshal(resp, &array)

	return array, nil
}

// FetchPhoneHomeInfo
func (c *StorageClient) FetchPhoneHomeInfo() (PhoneHomeInfo, error) {

	url := c.formatPath("array?phonehome=true")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return PhoneHomeInfo{}, err
	}

	var phonehome PhoneHomeInfo

	json.Unmarshal(resp, &phonehome)

	return phonehome, nil
}

// FetchNTPInfo
func (c *StorageClient) FetchNTPInfo() (NTPServerInfo, error) {

	url := c.formatPath("array?ntpserver=true")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return NTPServerInfo{}, err
	}

	var ntp NTPServerInfo

	json.Unmarshal(resp, &ntp)

	return ntp, nil
}

//FetchHardwareInfo
func (c *StorageClient) FetchHardwareInfo() ([]HardwareInfo, error) {

	url := c.formatPath("hardware")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var array []HardwareInfo

	json.Unmarshal(resp, &array)

	return array, nil
}

// FetchNetworkInfo
func (c *StorageClient) FetchNetworkInfo() ([]NetworkInfo, error) {

	url := c.formatPath("network")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var network []NetworkInfo

	json.Unmarshal(resp, &network)

	return network, nil
}

//FetchDNSInfo
func (c *StorageClient) FetchDNSInfo() (DNSInfo, error) {

	url := c.formatPath("dns")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return DNSInfo{}, err
	}

	var dns DNSInfo

	json.Unmarshal(resp, &dns)

	return dns, nil
}

//FetchSNMPInfo
func (c *StorageClient) FetchSNMPInfo() ([]SNMPInfo, error) {

	url := c.formatPath("snmp")

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var snmp []SNMPInfo

	json.Unmarshal(resp, &snmp)

	return snmp, nil
}
