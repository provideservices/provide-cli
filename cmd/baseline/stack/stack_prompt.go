package stack

import (
	"os"
	"strconv"

	"github.com/provideservices/provide-cli/cmd/common"
	"github.com/spf13/cobra"
)

const promptStepRun = "Run"
const promptStepStop = "Stop"
const promptStepLogs = "Logs"

var emptyPromptArgs = []string{promptStepRun, promptStepStop, promptStepLogs}
var emptyPromptLabel = "What would you like to do"

var boolPromptArgs = []string{"No", "Yes"}
var tunnelPromptLabel = "Would you like to set up a tunnel"
var tunnelAPIPromptLabel = "Would you like to set up a tunnel"
var tunnelMessagingPromptLabel = "Would you like to set up a tunnel"
var autoRemovePromptLabel = "Would you like to automatically remove"
var localVaultPromptLabel = "Would you like to set up vault locally"
var localIdentPromptLabel = "Would you like to set up ident locally"
var localNchainPromptLabel = "Would you like to set up nachain locally"
var localPrivacyPromptLabel = "Would you like to set up privacy locally"

// General Endpoints
func generalPrompt(cmd *cobra.Command, args []string, currentStep string) {
	switch step := currentStep; step {
	case promptStepRun:
		common.RequireOrganization()
		if optional {
			if name == "" {
				name = common.FreeInput("Name")
			}
			if common.APIEndpoint == "" {
				common.APIEndpoint = common.FreeInput("API endpoint")
			}
			if common.MessagingEndpoint == "" {
				common.MessagingEndpoint = common.FreeInput("Messaging endpoint")
			}
			if common.Tunnel == false {
				common.Tunnel = common.SelectInput(boolPromptArgs, tunnelPromptLabel) == "Yes"
			}
			if common.ExposeAPITunnel == false {
				common.ExposeAPITunnel = common.SelectInput(boolPromptArgs, tunnelAPIPromptLabel) == "Yes"
			}
			if common.ExposeMessagingTunnel == false {
				common.ExposeMessagingTunnel = common.SelectInput(boolPromptArgs, tunnelMessagingPromptLabel) == "Yes"
			}
			if sorID == "" {
				sorID = common.FreeInput("SoR")
			}
			if sorURL == "" {
				sorURL = common.FreeInput("SoR URL")
			}
			if apiHostname == "" {
				apiHostname = common.FreeInput("API Hostname")
			}
			if port == 8080 {
				port, _ = strconv.Atoi(common.FreeInput("Port"))
			}
			if consumerHostname == name+"-consumer" {
				consumerHostname = common.FreeInput("Consumer Hostname")
			}
			if natsHostname == name+"-nats" {
				natsHostname = common.FreeInput("Nats Hostname")
			}
			if natsPort == 4222 {
				natsPort, _ = strconv.Atoi(common.FreeInput("Nats Port"))
			}
			if natsWebsocketPort == 4221 {
				natsWebsocketPort, _ = strconv.Atoi(common.FreeInput("Nats Websocket Port"))
			}
			if natsAuthToken == "testtoken" {
				natsAuthToken = common.FreeInput("Nats Auth Token")
			}
			if natsStreamingHostname == name+"-nats-streaming" {
				natsStreamingHostname = common.FreeInput("Nats Streaming Token")
			}
			if natsStreamingPort == 4220 {
				natsStreamingPort, _ = strconv.Atoi(common.FreeInput("Nats Streaming Port"))
			}
			if redisHostname == name+"-reddis" {
				redisHostname = common.FreeInput("Reddis Host Name")
			}
			if redisPort == 6379 {
				redisPort, _ = strconv.Atoi(common.FreeInput("Reddis Port"))
			}
			// if redisHosts == redisHostname+":"+strconv.Atoi(redisContainerPort) {
			// 	redisPort, _ = strconv.Atoi(common.FreeInput("Reddis Port"))
			// }
			if autoRemove == false {
				autoRemove = common.SelectInput(boolPromptArgs, tunnelMessagingPromptLabel) == "Yes"
			}
			if logLevel == "DEBUG" {
				logLevel = common.FreeInput("Reddis Host Name")
			}
			if jwtSignerPublicKey == "" {
				jwtSignerPublicKey = common.FreeInput("JWT Signer Public Key")
			}
			if identAPIHost == "ident.provide.services" {
				nchainAPIHost = common.FreeInput("Ident API Host")
			}
			if identAPIScheme == "https" {
				nchainAPIScheme = common.FreeInput("Ident API Scheme")
			}
			if nchainAPIHost == "nchain.provide.services" {
				nchainAPIHost = common.FreeInput("Nchain API Host")
			}
			if nchainAPIScheme == "https" {
				nchainAPIScheme = common.FreeInput("Nchain API Scheme")
			}
			if privacyAPIHost == "privacy.provide.services" {
				privacyAPIHost = common.FreeInput("Privacy API Scheme")
			}
			if privacyAPIScheme == "https" {
				privacyAPIScheme = common.FreeInput("Privacy API Scheme")
			}
			if vaultAPIHost == "vault.provide.services" {
				vaultAPIHost = common.FreeInput("Vault API Scheme")
			}
			if vaultAPIScheme == "https" {
				vaultAPIScheme = common.FreeInput("Vault API Scheme")
			}
			if vaultRefreshToken == os.Getenv("VAULT_REFRESH_TOKEN") {
				vaultRefreshToken = common.FreeInput("Vault API Scheme")
			}
			if vaultSealUnsealKey == os.Getenv("VAULT_SEAL_UNSEAL_KEY") {
				vaultSealUnsealKey = common.FreeInput("Vault API Scheme")
			}
			if withLocalVault == false {
				withLocalVault = common.SelectInput(boolPromptArgs, localVaultPromptLabel) == "Yes"
			}
			if withLocalIdent == false {
				withLocalIdent = common.SelectInput(boolPromptArgs, localIdentPromptLabel) == "Yes"
			}
			if withLocalNchain == false {
				withLocalNchain = common.SelectInput(boolPromptArgs, localNchainPromptLabel) == "Yes"
			}
			if withLocalPrivacy == false {
				withLocalPrivacy = common.SelectInput(boolPromptArgs, localPrivacyPromptLabel) == "Yes"
			}
			if organizationRefreshToken == os.Getenv("PROVIDE_ORGANIZATION_REFRESH_TOKEN") {
				organizationRefreshToken = common.FreeInput("Organization Refresh Token")
			}
			if baselineOrganizationAddress == "0x" {
				baselineOrganizationAddress = common.FreeInput("Baseline Organization Address")
			}
			if baselineRegistryContractAddress == "0x" {
				baselineOrganizationAddress = common.FreeInput("Baseline Registry Contract Address")
			}
			if baselineWorkgroupID == "" {
				baselineOrganizationAddress = common.FreeInput("Baseline Workgroup ID")
			}
			if nchainBaselineNetworkID == "0x" {
				baselineOrganizationAddress = common.FreeInput("Nchain Baseline Network ID")
			}
			runProxy(cmd, args)
		}
	// case promptStepStop:
	// 	if optional {
	// 		fmt.Println("Optional Flags:")
	// 		if !nonCustodial {
	// 			nonCustodial = common.SelectInput(custodyPromptArgs, custodyPromptLabel) == "Yes"
	// 		}
	// 		if walletName == "" {
	// 			walletName = common.FreeInput("Wallet Name")
	// 		}
	// 		if purpose == 44 {
	// 			purpose, _ = strconv.Atoi(common.FreeInput("Wallet Purpose"))
	// 		}
	// 	}
	// 	CreateWalletRun(cmd, args)
	// case promptStepLogs:
	// 	if optional {
	// 		fmt.Println("Optional Flags:")
	// 		if common.ApplicationID == "" {
	// 			common.RequireApplication()
	// 		}
	// 	}
	// 	listWalletsRun(cmd, args)
	case "":
		result := common.SelectInput(emptyPromptArgs, emptyPromptLabel)
		generalPrompt(cmd, args, result)
	}
}
