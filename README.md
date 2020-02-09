# go-netplan-types

Type definitions for netplan configuration YAML.

See also: https://netplan.io/reference

## Features

- Support full definitions for netplan configuration
- Support YAML marshalling and unmarshalling

## Example

```go
import (
	"fmt"

	"github.com/moznion/go-netplan-types/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	network := netplan.Network{
		Network: &netplan.NetworkConfig{
			Ethernets: netplan.Ethernets{
				"eno1": &netplan.Ethernet{
					Device: netplan.Device{
						DHCP4: netplan.NillableBoolOf(true),
					},
				},
			},
		},
	}
	yamlString, _ := yaml.Marshal(&network)
	fmt.Printf("%s\n", yamlString)
	// This should be the following:
	// network:
	//   version: 2
	//   ethernets:
	//     eno1:
	//       dhcp4: true
}
```

## Author

moznion (<moznion@gmail.com>)

## License

```
The MIT License (MIT)
Copyright © 2020 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
