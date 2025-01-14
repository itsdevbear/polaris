// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package solidity

//go:generate abigen --pkg generated --abi ./out/staking.sol/IStakingModule.abi.json --bin ./out/staking.sol/IStakingModule.bin --out ./generated/i_staking_module.abigen.go --type StakingModule
//go:generate abigen --pkg generated --abi ./out/bank.sol/IBankModule.abi.json --bin ./out/bank.sol/IbankModule.bin --out ./generated/i_bank_module.abigen.go --type BankModule
//go:generate abigen --pkg generated --abi ./out/address.sol/IAddress.abi.json --bin ./out/address.sol/IAddress.bin --out ./generated/i_address.abigen.go --type Address
//go:generate abigen --pkg generated --abi ./out/distribution.sol/IDistributionModule.abi.json --bin ./out/distribution.sol/IDistributionModule.bin --out ./generated/i_distribution_module.abigen.go --type DistributionModule
