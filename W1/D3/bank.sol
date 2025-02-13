// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Bank {
    address public admin;
    mapping(address => uint256) public balances;
    address[3] public topUsers;

    constructor() {
        admin = msg.sender; // 部署合约的地址为管理员
    }

    // 接收以太币
    receive() external payable {
        balances[msg.sender] += msg.value;
        updateTopUsers(msg.sender);
    }

    // 仅管理员可以提取资金
    function withdraw(uint256 amount) public {
        require(msg.sender == admin, "Only admin can withdraw");
        require(address(this).balance >= amount, "Insufficient balance in contract");
        payable(admin).transfer(amount);
    }

    // 更新存款金额前 3 名用户
    function updateTopUsers(address user) internal {
        for (uint256 i = 0; i < topUsers.length; i++) {
            if (topUsers[i] == user) {
                break;
            }
            if (balances[user] > balances[topUsers[i]]) {
                for (uint256 j = topUsers.length - 1; j > i; j--) {
                    topUsers[j] = topUsers[j - 1];
                }
                topUsers[i] = user;
                break;
            }
        }
    }

    // 获取前 3 名存款用户
    function getTopUsers() public view returns (address[3] memory) {
        return topUsers;
    }
}