// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title SmartRouteProxyAdmin
 * @dev 智能路由代理管理员合约
 */
contract SmartRouteProxyAdmin is ProxyAdmin {
    constructor() ProxyAdmin() {}
} 