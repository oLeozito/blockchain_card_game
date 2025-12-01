package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"text/tabwriter"

	"pbl_redes/contracts" // O caminho deve bater com o go.mod

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. Conecta na Blockchain (Hardhat)
	clientUrl := "http://localhost:8545"
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar na blockchain: %v. (O 'make up' está rodando?)", err)
	}
	fmt.Println("--- CONECTADO À BLOCKCHAIN (HARDHAT) ---")

	// 2. Endereço do Contrato (Padrão Hardhat)
	contractAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// 3. Carrega o Contrato
	instance, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		log.Fatalf("Erro ao carregar contrato: %v", err)
	}

	// =========================================================================
	// 1. AUDITORIA DE CRIAÇÃO (MINT)
	// =========================================================================
	
	iterCards, err := instance.FilterCardMinted(&bind.FilterOpts{Start: 0})
	if err != nil {
		log.Printf("Aviso: Não foi possível ler histórico de cartas: %v", err)
	} else {
		fmt.Println("\n=== LIVRO RAZÃO: CRIAÇÃO DE ATIVOS (MINT) ===")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "BLOCO\t| ID (HASH)\t| RARIDADE\t| DONO INICIAL")
		fmt.Fprintln(w, "-----\t| ---------\t| --------\t| -------------")

		countCards := 0
		for iterCards.Next() {
			event := iterCards.Event
			idShort := event.Id.String()
			if len(idShort) > 12 { idShort = idShort[:12] + "..." }
			
			fmt.Fprintf(w, "#%d\t| %s\t| %s\t| %s\n",
				event.Raw.BlockNumber, idShort, event.Rarity, event.Owner.Hex())
			countCards++
		}
		w.Flush()
		fmt.Printf("Total de Cartas Criadas: %d\n", countCards)
	}

	// =========================================================================
	// 2. AUDITORIA DE TRANSFERÊNCIAS (TROCAS) - NOVO!
	// =========================================================================

	iterTransfers, err := instance.FilterCardTransferred(&bind.FilterOpts{Start: 0})
	if err != nil {
		// Se der erro aqui, é porque precisa rodar 'abigen' de novo!
		log.Printf("\nAviso: Não foi possível ler transferências (Verifique o abigen): %v", err)
	} else {
		fmt.Println("\n=== LIVRO RAZÃO: HISTÓRICO DE TROCAS (TRANSFERÊNCIAS) ===")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "BLOCO\t| ID CARTA\t| DE (REMETENTE)\t| PARA (DESTINO)")
		fmt.Fprintln(w, "-----\t| --------\t| --------------\t| --------------")

		countTransfers := 0
		for iterTransfers.Next() {
			evt := iterTransfers.Event
			
			idShort := evt.Id.String()
			if len(idShort) > 10 { idShort = idShort[:10] + "..." }

			fromShort := evt.From.Hex()
			if len(fromShort) > 10 { fromShort = fromShort[:10] + "..." }

			toShort := evt.To.Hex()
			if len(toShort) > 10 { toShort = toShort[:10] + "..." }

			fmt.Fprintf(w, "#%d\t| %s\t| %s\t| %s\n",
				evt.Raw.BlockNumber, idShort, fromShort, toShort)
			countTransfers++
		}
		w.Flush()
		fmt.Printf("Total de Trocas Realizadas: %d\n", countTransfers)
	}

	// =========================================================================
	// 3. AUDITORIA DE PARTIDAS (MATCHES)
	// =========================================================================

	iterMatches, err := instance.FilterMatchRecorded(&bind.FilterOpts{Start: 0})
	if err != nil {
		log.Printf("\nAviso: Não foi possível ler histórico de partidas: %v", err)
	} else {
		fmt.Println("\n=== LIVRO RAZÃO: HISTÓRICO DE PARTIDAS ===")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\t| VENCEDOR\t| PERDEDOR\t| PLACAR")
		fmt.Fprintln(w, "--\t| --------\t| --------\t| ------")

		countMatches := 0
		for iterMatches.Next() {
			evt := iterMatches.Event
			
			matchIndex := new(big.Int).Sub(evt.MatchId, big.NewInt(1))
			placarStr := "? x ?"
			
			matchData, err := instance.Matches(nil, matchIndex)
			if err == nil {
				placarStr = fmt.Sprintf("%d x %d", matchData.ScoreWinner, matchData.ScoreLoser)
			}

			winShort := evt.Winner.Hex()
			if len(winShort) > 10 { winShort = winShort[:10] + "..." }
			
			loseShort := evt.Loser.Hex()
			if len(loseShort) > 10 { loseShort = loseShort[:10] + "..." }

			fmt.Fprintf(w, "%s\t| %s\t| %s\t| %s\n",
				evt.MatchId.String(), winShort, loseShort, placarStr)
			countMatches++
		}
		w.Flush()
		fmt.Printf("Total de Partidas: %d\n", countMatches)
	}

	fmt.Println("================================================")
}