/*
 * Copyright 2022 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package appconfig

import (
	"errors"
	"fmt"
	"github.com/citrix/adc-nitro-go/service"
	"log"
)

type Environment struct {
	Name        string             `json:"name" yaml:"name"`               // Target environment name, such as "Production"
	Type        EnvironmentType    `json:"type" yaml:"type"`               // Target type: "StandAlone", "HighAvailabilityPair", "Cluster"
	SNIP        Node               `json:"snip" yaml:"snip"`               // Connection details for the shared SNIP (SNIP) of the environment
	Nodes       []Node             `json:"nodes" yaml:"nodes"`             // Connection details for the individual NSIP of each node
	Credentials Credentials        `json:"credentials" yaml:"credentials"` // Credentials
	Settings    ConnectionSettings `json:"settings" yaml:"settings"`       // Connections settings
}

//GetAllNitroClients Get a map of NitroClient for every node in the environment (NSIP/SNIP)
func (e *Environment) GetAllNitroClients() (map[string]service.NitroClient, error) {
	clients := make(map[string]service.NitroClient)
	if len(e.Nodes) != 0 {
		for _, n := range e.Nodes {
			client, err := service.NewNitroClientFromParams(
				service.NitroParams{
					Url:           n.GetNodeUrl(e.Settings.UrlScheme),
					Username:      e.Credentials.Username,
					Password:      e.Credentials.Password,
					ProxiedNs:     "",
					SslVerify:     e.Settings.ValidateServerCertificate,
					Timeout:       e.Settings.Timeout,
					RootCAPath:    "",
					ServerName:    "",
					Headers:       nil,
					LogLevel:      "",
					JSONLogFormat: false,
				})

			if err != nil {
				log.Printf("Could not create client for environment %s, node %s", e.Name, n.Name)
				return clients, err
			}

			clients[n.Name] = *client
		}
	}

	emptyNode := Node{}
	if e.SNIP != emptyNode {
		client, err := service.NewNitroClientFromParams(
			service.NitroParams{
				Url:           e.SNIP.GetNodeUrl(e.Settings.UrlScheme),
				Username:      e.Credentials.Username,
				Password:      e.Credentials.Password,
				ProxiedNs:     "",
				SslVerify:     e.Settings.ValidateServerCertificate,
				Timeout:       e.Settings.Timeout,
				RootCAPath:    "",
				ServerName:    "",
				Headers:       nil,
				LogLevel:      "",
				JSONLogFormat: false,
			})

		if err != nil {
			log.Printf("Could not create client for environment %s, SNIP %s", e.Name, e.SNIP.Name)
			return clients, err
		}

		clients["SNIP"] = *client
	}

	return clients, nil
}

//GetPrimaryNodeName Get the client name of the primary node in the environment
func (e *Environment) GetPrimaryNodeName() (string, error) {
	clients, err := e.GetAllNitroClients()
	if err != nil {
		return "", err
	}

	// Return client for SNIP if defined, as it always points to the primary node
	if _, exists := clients["SNIP"]; exists {
		return "SNIP", nil
	}

	// Return error if there are no individual nodes defined
	if len(e.Nodes) == 0 {
		errText := fmt.Sprintf("invalid number of nodes defined for the environment %s (%d)", e.Name, len(e.Nodes))
		return "", errors.New(errText)
	}

	// Return client for NSIP of the only node in a Standalone NetScaler environment
	if e.Type == Standalone {
		if len(e.Nodes) == 1 {
			return e.Nodes[0].Name, nil
		} else {
			errText := fmt.Sprintf("invalid number of nodes defined for the environment %s (%d)", e.Name, len(e.Nodes))
			return "", errors.New(errText)
		}
	}

	// Return client for the primary node by checking the HANODE state
	for _, n := range e.Nodes {
		if _, err := CheckNodeIsPrimary(clients[n.Name]); err == nil {
			return n.Name, nil
		}
	}

	// Not able to select a client for the primary node
	errText := fmt.Sprintf("invalid number of nodes defined for the environment %s", e.Name)
	return "", errors.New(errText)
}

//CheckNodeIsPrimary Check if the provided NitroClient is acting as a primary node
func CheckNodeIsPrimary(client service.NitroClient) (bool, error) {
	response, err := client.FindResource(service.Hanode.Type(), "0")
	if err == nil {
		if response["state"] == "Primary" {
			return true, nil
		}
		return false, nil
	}
	return false, err
}
