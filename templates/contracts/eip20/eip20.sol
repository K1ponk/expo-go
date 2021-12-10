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

// import and use the eip20 interface
import "./ieip20.sol";

// implementation of eip20 ./interface.sol
contract EIP20 is IEIP20 {
	mapping (address => uint256) private _balances;
	mapping (address => mapping (address => uint256)) private _alowances;
	uint256 private _totalSupply;
	string private _name;
	string private _symbol;

	constructor(string memory name_, string memory symbol_) {
		_name = name_;
		_symbol = symbol_;
	}

	function name() public view virtual override returns (string memory) {
		return _name;
	}

	function symbol() public view virtual override returns (string memory) {
		return _symbol;
	}

	function decimals() public view virtual override returns (uint8) {
		return 18;
	}

	function totalSupply() public view virtual override returns (uint256) {
		return _totalSupply;
	}

	function balanceOf(address _owner) public view virtual returns (uint256) {
		return _balances[_owner];
	}

	function transfer(address _to, uint256 _value) public virtual override returns (bool) {
		return _safeTransfer(msg.sender, _to, _value);
	}

	function transferFrom(address _from, address _to, uint256 _value) public virtual override returns (bool) {
		return _safeTransfer(_from, _to, _value);
	}

	function approve(address _spender, uint256 _value) public virtual override returns (bool) {
		require(_spender != address(0), "approval to zero address");
		_alowances[msg.sender][_spender] = _value;

		emit Approval(msg.sender, _spender, _value);

		return true;
	}

	function allowance(address _owner, address _spender) public view virtual override returns (uint256) {
		return allowance[_owner][_spender];
	}

	// _safeTransfer for handling all transfer in save way
	function _safeTransfer(address _from, address _to, uint256 _value) internal pure returns (bool) {
		require(_from != address(0), "transfer from zero address");
		require(_to != address(0), "transfer to zero address");
		require(balanceOf(_from) >= _value, "not enough balance");

		if (_from != msg.sender) {
			require(allowance(_from, msg.sender) >= _value, "not enough allowances");

			_alowances[_from][msg.sender] -= _value;
		}

		_balances[_from] -= _value;
		_balances[_to] += _value;

		emit Transfer(_from, _to, _value);

		return true;
	}

	// token minting function
	function _mint(address _to, uint256 _value) external returns (bool) {
		require(_to != address(0), "mint to zero address");
		_balances[_to] += _value;
		_totalSupply += _value;

		emit Transfer(address(0), _to, _value);

		return true;
	}

	// token burning function
	function _burn(address _owner, uint256 _value) external view returns (bool) {
		require(_owner != address(0), "burn from zero address");
		require(balanceOf(_owner) >= _value, "not enough balances");
		if (_owner != msg.sender) {
			require(allowance(_owner, msg.sender) >= _value, "not enough allowances");

			_alowances[_owner][msg.sender] -= _value;
		}

		_balances[_owner] -= _value;
		_totalSupply -= _value;

		emit Transfer(_owner, address(0), _value);

		return true;
	}
}
