pragma solidity ^0.5.0;

contract SimpleStorage {

    function expmod(uint256 base, uint256 e, uint256 m) public returns (uint256 o) {
        bool success;
    
         assembly {
        
           let p := mload(0x40)
           // store data assembly-favouring ways
           mstore(p, 0x20)             // Length of Base
           mstore(add(p, 0x20), 0x20)  // Length of Exponent
           mstore(add(p, 0x40), 0x20)  // Length of Modulus
           mstore(add(p, 0x60), base)  // Base
           mstore(add(p, 0x80), e)     // Exponent
           mstore(add(p, 0xa0), m)     // Modulus
           // call modexp precompile
           success := call(sub(gas, 2000), 0x05, 0, p, 0xc0, p, 0x20)
           // gas fiddling
           
            o := mload(p)
           
        
            
        }
        require(success);
        return o;
    }
  function echo(uint256 ini) public returns (uint256){
      return ini;
  }
}
