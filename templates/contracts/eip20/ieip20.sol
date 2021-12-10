// SPDX-License-Identifier: GPL-3.0
// Copyright 2021 The expo-go Authors
// This file is part of expo-go.
//
// expo-go is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

pragma solidity 0.8.10;

// Interface eip20 standard
interface IEIP20 {
	// Returns the name of the token - e.g. "MyToken".
	function name() external view returns (string memory);
	// Returns the symbol of the token. E.g. “HIX”.
	function symbol() external view returns (string memory);
	// Returns the number of decimals the token uses - e.g. 8
	// means to divide the token amount by 100000000 to get its user representation.
	function decimals() external view returns (uint8);
	// Returns the total token supply.
	function totalSupply() external view returns (uint256);
	// Returns the account balance of another account with address _owner.
	function balanceOf(address _owner) external view returns (uint256);
	// Transfers _value amount of tokens to address _to
	// and MUST fire the Transfer event. The function SHOULD throw if the message caller’s account balance does not have enough tokens to spend.
	function transfer(address _to, uint256 _value) external returns (bool);
	// Transfers _value amount of tokens from address _from to address _to
	// and MUST fire the Transfer event.
	// The transferFrom method is used for a withdraw workflow
	// allowing contracts to transfer tokens on your behalf
	// This can be used for example to allow a contract to transfer tokens on your behalf and/or to charge fees in sub-currencies
	// The function SHOULD throw unless the _from account has deliberately authorized the sender of the message via some mechanism.
	function transferFrom(address _from, address _to, uint256 _value) external returns (bool);
	// Allows _spender to withdraw from your account multiple times, up to the _value amount. If this function is called again it overwrites the current allowance with _value.
	function approve(address _spender, uint256 _value) external returns (bool);
	// Returns the amount which _spender is still allowed to withdraw from _owner.
	function allowance(address _owner, address _spender) external view returns (uint256);
	// MUST trigger when tokens are transferred, including zero value transfers.
	// A token contract which creates new tokens SHOULD trigger a Transfer event with the _from address set to 0x0 when tokens are created.
	event Transfer(address indexed _from, address indexed _to, uint256 _value);
	// MUST trigger on any successful call to approve(address _spender, uint256 _value).
	event Approval(address indexed _from, address indexed _to, uint256 _value);
}
