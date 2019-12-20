// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package constants

import (
	"fmt"
	"math"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var initialized = false

func InitConfig() error {
	if initialized {
		time.Sleep(time.Millisecond)
		return nil
	}
	initialized = true

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		return fmt.Errorf("could not find environment file: %v", err)
	}
	return nil
}

func getEnvironmentString(key string) string {
	err := InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("No environment configuration variable set for key: \"%v\"", key)
	}
	return value
}

/* Returns all contract config names from transformer configuration:
[exporter.vow_file]
	path = "transformers/events/vow_file/initializer"
	type = "eth_event"
	repository = "github.com/makerdao/vdb-mcd-transformers"
	migrations = "db/migrations"
	contracts = ["MCD_VOW"]   <----
	rank = "0"
*/
func GetTransformerContractNames(transformerLabel string) []string {
	err := InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}
	configKey := "exporter." + transformerLabel + ".contracts"
	contracts := viper.GetStringSlice(configKey)
	if len(contracts) == 0 {
		log.Fatalf("No contracts configured for transformer: \"%v\"", transformerLabel)
	}
	return contracts
}

// Get the ABI for multiple contracts from config
// Makes sure the ABI matches for all, since a single transformer may run against many contracts.
func GetContractsABI(contractNames []string) (string, error) {
	err := InitConfig()
	if err != nil {
		return "", fmt.Errorf("failed to init config: %s", err.Error())
	}
	if len(contractNames) < 1 {
		return "", fmt.Errorf("No contracts to get ABI for")
	}
	abi := getContractABI(contractNames[0])
	if len(contractNames) > 1 {
		for _, contractName := range contractNames[:1] {
			if abi != getContractABI(contractName) {
				log.WithField("contracts", contractNames).Fatalf("ABIs not consistent between contracts")
			}
		}
	}
	return abi, nil
}

func getContractABI(contractName string) string {
	configKey := "contract." + contractName + ".abi"
	abi := viper.GetString(configKey)
	if abi == "" {
		log.Fatalf("No ABI configured for contract: \"%v\"", contractName)
	}
	return abi
}

// Get the minimum deployment block for multiple contracts from config
func GetMinDeploymentBlock(contractNames []string) int64 {
	if len(contractNames) < 1 {
		log.Fatalf("No contracts supplied")
	}
	err := InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}
	minBlock := int64(math.MaxInt64)
	for _, c := range contractNames {
		deployed := getDeploymentBlock(c)
		if deployed < minBlock {
			minBlock = deployed
		}
	}
	return minBlock
}

func getDeploymentBlock(contractName string) int64 {
	configKey := "contract." + contractName + ".deployed"
	value := viper.GetInt64(configKey)
	if value == -1 {
		log.Infof("No deployment block configured for contract \"%v\", defaulting to 0.", contractName)
		return 0
	}
	return value
}

// Get the addresses for multiple contracts from config
func GetContractAddresses(contractNames []string) (addresses []string) {
	if len(contractNames) < 1 {
		log.Fatalf("No contracts supplied")
	}
	err := InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}
	for _, contractName := range contractNames {
		addresses = append(addresses, GetContractAddress(contractName))
	}
	return
}

func GetContractAddress(contract string) string {
	return getEnvironmentString("contract." + contract + ".address")
}
