// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

/// @custom:security-contact r10922188@csie.ntu.edu.tw
/// Owner, Minter
contract SoulBoundToken is ERC721, ERC721Enumerable, ERC721URIStorage, AccessControl {
    using Counters for Counters.Counter;
   // event UpdateGuardLog(uint256 indexed tokenId, address indexed newGuard, address oldGuard, uint64 expires);
    enum BurnAuth{
        IssuerOnly,
        OwnerOnly,
        Both,
        Neither
    }
    event Issued(
        address indexed from,
        address indexed to,
        uint256 indexed  tokenId,
        BurnAuth burnAuth
    );

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    
    Counters.Counter private _tokenIdCounter;


    constructor() ERC721("SoulBoundToken", "SBT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function safeMint(address to, string memory uri) public onlyRole(MINTER_ROLE) {
        
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        emit Issued(msg.sender,to,tokenId,BurnAuth.Both);
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }
    function setTokenURI(uint256 tokenId,string memory uri) public onlyRole(MINTER_ROLE){
        _setTokenURI(tokenId,uri);
    }


    // The following functions are overrides required by Solidity.

    function _beforeTokenTransfer(address from, address to, uint256 tokenId, uint256 batchSize)
        internal
        override(ERC721, ERC721Enumerable)
    {
        if(from != address(0) && to!=address(0)){
            revert();
        }
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }

    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
    function  burnAuth(uint256 tokenId) public view  returns (BurnAuth){
        require(tokenId < totalSupply());
        return BurnAuth.Both;
    }

    function burn(uint256 tokenId) public {
        if(hasRole(MINTER_ROLE, msg.sender) || ownerOf(tokenId)==msg.sender){
            _burn(tokenId);
        }else{
            revert();
        }
    }
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

}