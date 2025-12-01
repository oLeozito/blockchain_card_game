// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GameAssets {
    struct Card {
        uint256 id;
        string rarity;
        address owner;
    }

    // CORREÇÃO: Adicionados scoreWinner e scoreLoser na struct principal
    struct MatchResult {
        uint256 matchId;
        address winner;
        address loser;
        uint256 scoreWinner; // Novo campo
        uint256 scoreLoser;  // Novo campo
        uint256 timestamp;
    }

    // Variáveis de Estado
    mapping(uint256 => Card) public cards;
    mapping(address => uint256[]) public playerInventory;
    MatchResult[] public matches; // Agora usa a struct completa

    address public gameServer; 

    constructor() {
        gameServer = msg.sender; 
    }

    event CardMinted(uint256 id, string rarity, address owner);
    event CardTransferred(uint256 id, address from, address to); 
    event MatchRecorded(uint256 matchId, address winner, address loser);

    // 1. MINT (Criar Carta)
    function mintCard(uint256 _id, string memory _rarity, address _owner) public {
        require(msg.sender == gameServer, "Apenas servidor");
        require(cards[_id].id == 0, "Carta ja existe!");
        
        cards[_id] = Card(_id, _rarity, _owner);
        playerInventory[_owner].push(_id);
        
        emit CardMinted(_id, _rarity, _owner);
    }

    // 2. TRANSFER (Troca de Carta)
    function transferCard(uint256 _cardId, address _from, address _to) public {
        require(msg.sender == gameServer, "Apenas servidor");
        require(cards[_cardId].owner == _from, "Remetente nao e dono");

        cards[_cardId].owner = _to;
        emit CardTransferred(_cardId, _from, _to);
    }

    // 3. MATCH (Registrar Partida)
    function recordMatch(address _winner, address _loser, uint256 _sWin, uint256 _sLose) public {
        require(msg.sender == gameServer, "Apenas servidor");
        
        uint256 newId = matches.length + 1; // Corrigido para usar 'matches.length'
        
        // Agora a struct tem 6 campos, então esse push vai funcionar
        matches.push(MatchResult({
            matchId: newId,
            winner: _winner,
            loser: _loser,
            scoreWinner: _sWin,
            scoreLoser: _sLose,
            timestamp: block.timestamp
        }));

        emit MatchRecorded(newId, _winner, _loser);
    }
    
    // Função auxiliar para ver quantas partidas foram jogadas
    function getMatchCount() public view returns (uint256) {
        return matches.length;
    }
}