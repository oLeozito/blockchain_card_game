// contracts/GameAssets.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GameAssets {
    // --- ESTRUTURAS DE DADOS ---

    struct Card {
        uint256 id;
        string rarity;
        address owner;
    }

    struct MatchResult {
        uint256 matchId;
        address winner;
        address loser;
        uint256 scoreWinner;
        uint256 scoreLoser;
        uint256 timestamp;
    }

    // --- VARIÁVEIS DE ESTADO ---

    mapping(uint256 => Card) public cards;
    mapping(address => uint256[]) public playerInventory;
    MatchResult[] public matches;

    // Endereço do servidor (Líder Raft) que tem permissão para gerenciar o jogo
    address public gameServer; 

    // --- EVENTOS ---

    event CardMinted(uint256 id, string rarity, address owner);
    event CardTransferred(uint256 id, address from, address to); 
    event MatchRecorded(uint256 matchId, address winner, address loser);

    // --- CONSTRUTOR ---

    constructor() {
        // Quem fizer o deploy do contrato (o servidor) será o "dono/juiz"
        gameServer = msg.sender; 
    }

    // --- FUNÇÕES DO JOGO ---

    // 1. MINT (Criar Carta)
    function mintCard(uint256 _id, string memory _rarity, address _owner) public {
        require(msg.sender == gameServer, "Apenas servidor pode criar cartas");
        require(cards[_id].id == 0, "Carta ja existe!");
        
        cards[_id] = Card(_id, _rarity, _owner);
        playerInventory[_owner].push(_id);
        
        emit CardMinted(_id, _rarity, _owner);
    }

    // 2. TRANSFER (Troca de Cartas)
    function transferCard(uint256 _cardId, address _to) public {
        // SEGURANÇA: Apenas o servidor pode autorizar a movimentação
        require(msg.sender == gameServer, "Apenas servidor pode mover cartas");
        
        require(cards[_cardId].id != 0, "Carta nao existe");
        
        address oldOwner = cards[_cardId].owner;
        require(oldOwner != _to, "Nao pode transferir para si mesmo");

        // 1. Atualiza o dono na struct principal
        cards[_cardId].owner = _to;

        // 2. Remove do inventário do antigo dono
        removeCardFromInventory(oldOwner, _cardId);

        // 3. Adiciona ao inventário do novo dono
        playerInventory[_to].push(_cardId);

        emit CardTransferred(_cardId, oldOwner, _to);
    }

    // 3. MATCH (Registrar Partida)
    function recordMatch(address _winner, address _loser, uint256 _sWin, uint256 _sLose) public {
        require(msg.sender == gameServer, "Apenas servidor pode registrar partidas");
        
        uint256 newId = matches.length + 1;
        
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
    
    // --- FUNÇÕES AUXILIARES ---

    // Remove um ID de carta de um array de inventário (swap and pop)
    function removeCardFromInventory(address _owner, uint256 _cardId) internal {
        uint256[] storage inv = playerInventory[_owner];
        for (uint256 i = 0; i < inv.length; i++) {
            if (inv[i] == _cardId) {
                // Substitui o elemento a ser removido pelo último e remove o último
                inv[i] = inv[inv.length - 1];
                inv.pop();
                break;
            }
        }
    }

    // Retorna quantas cartas um jogador tem
    function getBalance(address _owner) public view returns (uint256) {
        return playerInventory[_owner].length;
    }

    // Retorna o total de partidas registradas
    function getMatchCount() public view returns (uint256) {
        return matches.length;
    }
}